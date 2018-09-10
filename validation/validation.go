package validation

// ValidRegister returns true if the input points to a valid register.
// Only 2 general purpose registers at this time (0x00 reserved for program counter)
func ValidRegister(r uint) bool {
	return r == 0x01 || r == 0x02
}

// ValidInputAddr returns true if the address is one of 2 specificed inputs
func ValidInputAddr(a uint) bool {
	return a == 0x10 || a == 0x12
}

// ValidOutputAddr returns true if the address is the specified output
func ValidOutputAddr(a uint) bool {
	return a == 0x0e
}
