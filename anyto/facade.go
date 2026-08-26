package anyto

// AnyTo provides object-style access to the conversion functions in this package.
var AnyTo = facade{}

type facade struct{}

func (facade) AnyToBool(value any) bool                { return AnyToBool(value) }
func (facade) AnyToFloat32(value any) (float32, error) { return AnyToFloat32(value) }
func (facade) AnyToFloat64(value any) (float64, error) { return AnyToFloat64(value) }
func (facade) AnyToInt(value any) (int, error)         { return AnyToInt(value) }
func (facade) AnyToInt8(value any) (int8, error)       { return AnyToInt8(value) }
func (facade) AnyToInt16(value any) (int16, error)     { return AnyToInt16(value) }
func (facade) AnyToInt32(value any) (int32, error)     { return AnyToInt32(value) }
func (facade) AnyToInt64(value any) (int64, error)     { return AnyToInt64(value) }
func (facade) AnyToStr(value any) string               { return AnyToStr(value) }
func (facade) AnyToUint(value any) (uint, error)       { return AnyToUint(value) }
func (facade) AnyToUint8(value any) (uint8, error)     { return AnyToUint8(value) }
func (facade) AnyToUint16(value any) (uint16, error)   { return AnyToUint16(value) }
func (facade) AnyToUint32(value any) (uint32, error)   { return AnyToUint32(value) }
func (facade) AnyToUint64(value any) (uint64, error)   { return AnyToUint64(value) }
func (facade) ArrAnyToStr(values []any) []string       { return ArrAnyToStr(values) }
func (facade) ArrIntToStr(values []int) []string       { return ArrIntToStr(values) }
func (facade) ArrInt8ToStr(values []int8) []string     { return ArrInt8ToStr(values) }
func (facade) ArrInt16ToStr(values []int16) []string   { return ArrInt16ToStr(values) }
func (facade) ArrInt32ToStr(values []int32) []string   { return ArrInt32ToStr(values) }
func (facade) ArrInt64ToStr(values []int64) []string   { return ArrInt64ToStr(values) }
func (facade) ArrUintToStr(values []uint) []string     { return ArrUintToStr(values) }
func (facade) ArrUint8ToStr(values []uint8) []string   { return ArrUint8ToStr(values) }
func (facade) ArrUint16ToStr(values []uint16) []string { return ArrUint16ToStr(values) }
func (facade) ArrUint32ToStr(values []uint32) []string { return ArrUint32ToStr(values) }
func (facade) ArrUint64ToStr(values []uint64) []string { return ArrUint64ToStr(values) }
