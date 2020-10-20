
// Copyright 2013 The Prometheus Authors
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

package model

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

var (
	// ZeroSamplePair is the pseudo zero-value of SamplePair used to signal a
	// non-existing sample pair. It is a SamplePair with timestamp Earliest and
	// value 0.0. Note that the natural zero value of SamplePair has a timestamp
	// of 0, which is possible to appear in a real SamplePair and thus not
	// suitable to signal a non-existing SamplePair.
	ZeroSamplePair = SamplePair{Timestamp: Earliest}

	// ZeroSample is the pseudo zero-value of Sample used to signal a
	// non-existing sample. It is a Sample with timestamp Earliest, value 0.0,
	// and metric nil. Note that the natural zero value of Sample has a timestamp
	// of 0, which is possible to appear in a real Sample and thus not suitable
	// to signal a non-existing Sample.
	ZeroSample = Sample{Timestamp: Earliest}
)