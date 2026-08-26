package strto

// StrTo provides object-style access to the conversion functions in this package.
var StrTo = facade{}

type facade struct{}

func (facade) StrToBool(value string) bool             { return StrToBool(value) }
func (facade) StrToBytes(value string) []byte          { return StrToBytes(value) }
func (facade) StrToInt(value string) int               { return StrToInt(value) }
func (facade) StrToInt8(value string) int8             { return StrToInt8(value) }
func (facade) StrToInt16(value string) int16           { return StrToInt16(value) }
func (facade) StrToInt32(value string) int32           { return StrToInt32(value) }
func (facade) StrToInt64(value string) int64           { return StrToInt64(value) }
func (facade) StrToUint(value string) uint             { return StrToUint(value) }
func (facade) StrToUint8(value string) uint8           { return StrToUint8(value) }
func (facade) StrToUint16(value string) uint16         { return StrToUint16(value) }
func (facade) StrToUint32(value string) uint32         { return StrToUint32(value) }
func (facade) StrToUint64(value string) uint64         { return StrToUint64(value) }
func (facade) ArrStrToBool(values []string) []bool     { return ArrStrToBool(values) }
func (facade) ArrStrToInt(values []string) []int       { return ArrStrToInt(values) }
func (facade) ArrStrToInt8(values []string) []int8     { return ArrStrToInt8(values) }
func (facade) ArrStrToInt16(values []string) []int16   { return ArrStrToInt16(values) }
func (facade) ArrStrToInt32(values []string) []int32   { return ArrStrToInt32(values) }
func (facade) ArrStrToInt64(values []string) []int64   { return ArrStrToInt64(values) }
func (facade) ArrStrToUint(values []string) []uint     { return ArrStrToUint(values) }
func (facade) ArrStrToUint8(values []string) []uint8   { return ArrStrToUint8(values) }
func (facade) ArrStrToUint16(values []string) []uint16 { return ArrStrToUint16(values) }
func (facade) ArrStrToUint32(values []string) []uint32 { return ArrStrToUint32(values) }
func (facade) ArrStrToUint64(values []string) []uint64 { return ArrStrToUint64(values) }
