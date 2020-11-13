package pflag

import "strconv"

// -- uint8 Value
type uint8Value uint8

func newUint8Value(val uint8, p *uint8) *uint8Value {
	*p = val
	return (*uint8Value)(p)
}

func (i *uint8Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 8)
	*i = uint8Value(v)
	return err
}

func (i *uint8Value) Type() string {
	return "uint8"
}

func (i *uint8Value) String() string { return strconv.FormatUint(uint64(*i), 10) }

func uint8Conv(sval string) (interface{}, error) {
	v, err := strconv.ParseUint(sval, 0, 8)
	if err != nil {
		return 0, err
	}
	return uint8(v), nil
}

// GetUint8 return the uint8 value of a flag with the given name
func (f *FlagSet) GetUint8(name string) (uint8, error) {
	val, err := f.getFlagType(name, "uint8", uint8Conv)
	if 