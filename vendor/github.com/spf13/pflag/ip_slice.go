package pflag

import (
	"fmt"
	"io"
	"net"
	"strings"
)

// -- ipSlice Value
type ipSliceValue struct {
	value   *[]net.IP
	changed bool
}

func newIPSliceValue(val []net.IP, p *[]net.IP) *ipSliceValue {
	ipsv := new(ipSliceValue)
	ipsv.value = p
	*ipsv.value = val
	return ipsv
}

// Set converts, and assigns, the comma-separated IP argument string representation as the []net.IP value of this flag.
// If Set is called on a flag that already has a []net.IP assigned, the newly converted values will be appended.
func (s *ipSliceValue) Set(val string) error {

	// remove all quote characters
	rmQuote := strings.NewReplacer(`"`, "", `'`, "", "`", "")

	// read flag arguments with CSV parser
	ipStrSlice, err := readAsCSV(rmQuote.Replace(val))
	if err != nil 