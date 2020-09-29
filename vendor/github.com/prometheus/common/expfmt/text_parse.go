// Copyright 2014 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package expfmt

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	dto "github.com/prometheus/client_model/go"

	"github.com/golang/protobuf/proto"
	"github.com/prometheus/common/model"
)

// A stateFn is a function that represents a state in a state machine. By
// executing it, the state is progressed to the next state. The stateFn returns
// another stateFn, which represents the new state. The end state is represented
// by nil.
type stateFn func() stateFn

// ParseError signals errors while parsing the simple and flat text-based
// exchange format.
type ParseError struct {
	Line int
	Msg  string
}

// Error implements the error interface.
func (e ParseError) Error() string {
	return fmt.Sprintf("text format parsing error in line %d: %s", e.Line, e.Msg)
}

// TextParser is used to parse the simple and flat text-based exchange format. Its
// zero value is ready to use.
type TextParser struct {
	metricFamiliesByName map[string]*dto.MetricFamily
	buf                  *bufio.Reader // Where the parsed input is read through.
	err                  error         // Most recent error.
	lineCount            int           // Tracks the line count for error messages.
	currentByte          byte          // The most recent byte read.
	currentToken         bytes.Buffer  // Re-used each time a token has to be gathered from multiple bytes.
	currentMF            *dto.MetricFamily
	currentMetric        *dto.Metric
	currentLabelPair     *dto.LabelPair

	// The remaining member variables are only used for summaries/histograms.
	currentLabels map[string]string // All labels including '__name__' but excluding 'quantile'/'le'
	// Summary specific.
	summaries       map[uint64]*dto.Metric // Key is created with LabelsToSignature.
	currentQuantile float64
	// Histogram specific.
	histograms    map[uint64]*dto.Metric // Key is created with LabelsToSignature.
	currentBucket float64
	// These tell us if the currently processed line ends on '_count' or
	// '_sum' respectively and belong to a summary/histogram, representing the sample
	// count and sum of that summary/histogram.
	currentIsSummaryCount, currentIsSummarySum     bool
	currentIsHistogramCount, currentIsHistogramSum bool
}

// TextToMetricFamilies reads 'in' as the simple and flat text-based exchange
// format and creates MetricFamily proto messages. It returns the MetricFamily
// proto messages in a map where the metric names are the keys, along with any
// error encountered.
//
// If the input contains duplicate metrics (i.e. lines with the same metric name
// and exactly the same label set), the resulting MetricFamily will contain
// duplicate Metric proto messages. Similar is true for duplicate label
// names. Checks for duplicates have to be performed separately, if required.
// Also note that neither the metrics within each MetricFamily are sorted nor
// the label pairs within each Metric. Sorting is not required for the most
// frequent use of this method, which is sample ingestion in the Prometheus
// server. However, for presentation purposes, you might want to sort the
// metrics, and in some cases, you must sort the labels, e.g. for consumption by
// the metric family injection hook of the Prometheus registry.
//
// Summaries and histograms are rather special beasts. You would probably not
// use them in the simple text format anyway. This method can deal with
// summaries and histograms if they are presented in exactly the way the
// text.Create function creates them.
//
// This method must not be called concurrently. If you want to parse different
// input concurrently, instantiate a separate Parser for each goroutine.
func (p *TextParser) TextToMetricFamilies(in io.Reader) (map[string]*dto.MetricFamily, error) {
	p.reset(in)
	for nextState := p.startOfLine; nextState != nil; nextState = nextState() {
		// Magic happens here...
	}
	// Get rid of empty metric families.
	for k, mf := range p.metricFamiliesByName {
		if len(mf.GetMetric()) == 0 {
			delete(p.metricFamiliesByName, k)
		}
	}
	// If p.err is io.EOF now, we have run into a premature end of the input
	// stream. Turn this error into something nicer and more
	// meaningful. (io.EOF is often used as a signal for the legitimate end
	// of an input stream.)
	if p.err == io.EOF {
		p.parseError("unexpected end of input stream")
	}
	return p.metricFamiliesByName, p.err
}

func (p *TextParser) reset(in io.Reader) {
	p.metricFamiliesByName = map[string]*dto.MetricFamily{}
	if p.buf == nil {
		p.buf = bufio.NewReader(in)
	} else {
		p.buf.Reset(in)
	}
	p.err = nil
	p.lineCount = 0
	if p.summaries == nil || len(p.summaries) > 0 {
		p.summaries = map[uint64]*dto.Metric{}
	}
	if p.histograms == nil || len(p.histograms) > 0 {
		p.histograms = map[uint64]*dto.Metric{}
	}
	p.currentQuantile = math.NaN()
	p.currentBucket = math.NaN()
}

// startOfLine represents the state where the next byte read from p.buf is the
// start of a line (or whitespace leading up to it).
func (p *TextParser) startOfLine() stateFn {
	p.lineCount++
	if p.skipBlankTab(); p.err != nil {
		// End of input reached. This is the only case where
		// that is not an error but a signal that we are done.
		p.err = nil
		return nil
	}
	switch p.currentByte {
	case '#':
		return p.startComment
	case '\n':
		return p.startOfLine // Empty line, start the next one.
	}
	return p.readingMetricName
}

// startComment represents the state where the next byte read from p.buf is the
// start of a comment (or whitespace leading up to it).
func (p *TextParser) startComment() stateFn {
	if p.skipBlankTab(); p.err != nil {
		return nil // Unexpected end of input.
	}
	if p.currentByte == '\n' {
		return p.startOfLine
	}
	if p.readTokenUntilWhitespace(); p.err != nil {
		return nil // Unexpected end of input.
	}
	// If we have hit the end of line already, there is nothing left
	// to do. This is not considered a syntax error.
	if p.currentByte == '\n' {
		return p.startOfLine
	}
	keyword := p.currentToken.String()
	if keyword != "HELP" && keyword != "TYPE" {
		// Generic comment, ignore by fast forwarding to end of line.
		for p.currentByte != '\n' {
			if p.currentByte, p.err = p.buf.ReadByte(); p.err != nil {
				return nil // Unexpected end of input.
			}
		}
		return p.startOfLine
	}
	// There is something. Next has to be a metric name.
	if p.skipBlankTab(); p.err != nil {
		return nil // Unexpected end of input.
	}
	if p.readTokenAsMetricName(); p.err != nil {
		return nil // Unexpected end of input.
	}
	if p.currentByte == '\n' {
		// At the end of the line already.
		// Again, this is not considered a syntax error.
		return p.startOfLine
	}
	if !isBlankOrTab(p.currentByte) {
		p.parseError("invalid metric name in comment")
		return nil
	}
	p.setOrCreateCurrentMF()
	if p.skipBlankTab(); p.err != nil {
		return nil // Unexpected end of input.
	}
	if p.currentByte == '\n' {
		// At the end of the line already.
		// Again, this is not considered a syntax error.
		return p.startOfLine
	}
	switch keyword {
	case "HELP":
		return p.readingHelp
	case "TYPE":
		return p.readingType
	}
	panic(fmt.Sprintf("code error: unexpected keyword %q", keyword))
}

// readingMetricName represents the state where the last byte read (now in
// p.currentByte) is the first byte of a metric name.
func (p *TextParser) readingMetricName() stateFn {
	if p.readTokenAsMetricName(); p.err != nil {
		return nil
	}
	if p.currentToken.Len() == 0 {
		p.parseError("invalid metric name")
		return nil
	}
	p.setOrCreateCurrentMF()
	// Now is the time to fix the type if it hasn't happened yet.
	if p.currentMF.Type == nil {
		p.currentMF.Type = dto.MetricType_UNTYPED.Enum()
	}
	p.currentMetric = &dto.Metric{}
	// Do not append the newly created currentMetric to
	// currentMF.Metric right now. First wait if this is a summary,
	// and the metric exists already, which we can only know after
	// having read all the labels.
	if p.skipBlankTabIfCurrentBlankTab(); p.err != nil {
		return nil // Unexpected end of input.
	}
	return p.readingLabels
}

// readingLabels represents the state where the last byte read (now in
// p.currentByte) is either the first byte of the label set (i.e. a '{'), or the
// first byte of the value (otherwise).
func (p *TextParser) readingLabels() stateFn {
	// Summaries/histograms are special. We have to reset the
	// currentLabels map, currentQuantile and currentBucket before starting to
	// read labels.
	if p.currentMF.GetType() == dto.MetricType_SUMMARY || p.currentMF.GetType() == dto.MetricType_HISTOGRAM {
		p.currentLabels = map[string]string{}
		p.currentLabels[string(model.MetricNameLabel)] = p.currentMF.GetName()
		