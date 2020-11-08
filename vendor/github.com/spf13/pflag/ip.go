package pflag

import (
	"fmt"
	"net"
	"strings"
)

// -- net.IP value
type ipValue net.IP

func newIPValue(val net.IP, p *net.IP) *ipValue {
	*p = val
	return (*ipValue)(p)
}

func (i *ipValue) String() string { return net.IP(*i).S