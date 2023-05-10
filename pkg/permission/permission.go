package permission

const (
	readMask    = 0x01
	writeMask   = 0x02
	deleteMask  = 0x04
	custom1Mask = 0x08
	custom2Mask = 0x10
	custom3Mask = 0x20
	custom4Mask = 0x40
	custom5Mask = 0x80
)

// SetRead function
func SetRead(value int) int {

	return value | readMask
}

// ClearRead function
func ClearRead(value int) int {

	return value & (readMask ^ 0xFF)
}

// CanRead function
func CanRead(value int) bool {

	flag := value & readMask
	if flag == readMask {
		return true
	}

	return false
}

// SetWrite function
func SetWrite(value int) int {

	return value | writeMask
}

// ClearWrite function
func ClearWrite(value int) int {

	return value & (writeMask ^ 0xFF)
}

// CanWrite function
func CanWrite(value int) bool {

	flag := value & writeMask
	if flag == writeMask {
		return true
	}

	return false
}

// SetDelete function
func SetDelete(value int) int {

	return value | deleteMask
}

// ClearDelete function
func ClearDelete(value int) int {

	return value & (deleteMask ^ 0xFF)
}

// CanDelete function
func CanDelete(value int) bool {

	flag := value & deleteMask
	if flag == deleteMask {
		return true
	}

	return false
}

// SetCustom1 function
func SetCustom1(value int) int {

	return value | custom1Mask
}

// ClearCustom1 function
func ClearCustom1(value int) int {

	return value & (custom1Mask ^ 0xFF)
}

// CanCustom1 function
func CanCustom1(value int) bool {

	flag := value & custom1Mask
	if flag == custom1Mask {
		return true
	}

	return false
}

// SetCustom2 function
func SetCustom2(value int) int {

	return value | custom2Mask
}

// ClearCustom2 function
func ClearCustom2(value int) int {

	return value & (custom2Mask ^ 0xFF)
}

// CanCustom2 function
func CanCustom2(value int) bool {

	flag := value & custom2Mask
	if flag == custom2Mask {
		return true
	}

	return false
}

// SetCustom3 function
func SetCustom3(value int) int {

	return value | custom3Mask
}

// ClearCustom3 function
func ClearCustom3(value int) int {

	return value & (custom3Mask ^ 0xFF)
}

// CanCustom3 function
func CanCustom3(value int) bool {

	flag := value & custom3Mask
	if flag == custom3Mask {
		return true
	}

	return false
}

// SetCustom4 function
func SetCustom4(value int) int {

	return value | custom4Mask
}

// ClearCustom4 function
func ClearCustom4(value int) int {

	return value & (custom4Mask ^ 0xFF)
}

// CanCustom4 function
func CanCustom4(value int) bool {

	flag := value & custom4Mask
	if flag == custom4Mask {
		return true
	}

	return false
}

// SetCustom5 function
func SetCustom5(value int) int {
	return value | custom5Mask
}

// ClearCustom5 function
func ClearCustom5(value int) int {
	return value & (custom5Mask ^ 0xFF)
}

// CanCustom5 function
func CanCustom5(value int) bool {

	flag := value & custom5Mask
	if flag == custom5Mask {
		return true
	}

	return false
}
