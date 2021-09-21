package xconv

import "strconv"

type Convertor struct {
	dftInt8    int8
	dftInt16   int16
	dftInt32   int32
	dftInt64   int64
	dftUint8   uint8
	dftUint16  uint16
	dftUint32  uint32
	dftUint64  uint64
	dftBool    bool
	dftFloat32 float32
	dftFloat64 float64
}

func (conv Convertor) SetDftInt8(v int8) {
	conv.dftInt8 = v
}

func (conv Convertor) SetDftInt16(v int16) {
	conv.dftInt16 = v
}

func (conv Convertor) SetDftInt32(v int32) {
	conv.dftInt32 = v
}

func (conv Convertor) SetDftInt64(v int64) {
	conv.dftInt64 = v
}

func (conv Convertor) SetDftUint8(v uint8) {
	conv.dftUint8 = v
}

func (conv Convertor) SetDftUint16(v uint16) {
	conv.dftUint16 = v
}

func (conv Convertor) SetDftUint32(v uint32) {
	conv.dftUint32 = v
}

func (conv Convertor) SetDftUint64(v uint64) {
	conv.dftUint64 = v
}

func (conv Convertor) SetDftBool(v bool) {
	conv.dftBool = v
}

func (conv Convertor) SetDftFloat32(v float32) {
	conv.dftFloat32 = v
}

func (conv Convertor) SetDftFloat64(v float64) {
	conv.dftFloat64 = v
}

func (conv Convertor) ToStr4I8(trgt int8) string {
	return strconv.FormatInt(int64(trgt), 10)
}

func (conv Convertor) ToStr4I16(trgt int16) string {
	return strconv.FormatInt(int64(trgt), 10)
}

func (conv Convertor) ToStr4I32(trgt int32) string {
	return strconv.FormatInt(int64(trgt), 10)
}

func (conv Convertor) ToStr4I64(trgt int64) string {
	return strconv.FormatInt(int64(trgt), 10)
}

func (conv Convertor) ToStr4Ui8(trgt uint8) string {
	return strconv.FormatUint(uint64(trgt), 10)
}

func (conv Convertor) ToStr4Ui16(trgt uint16) string {
	return strconv.FormatUint(uint64(trgt), 10)
}

func (conv Convertor) ToStr4Ui32(trgt uint32) string {
	return strconv.FormatUint(uint64(trgt), 10)
}

func (conv Convertor) ToStr4Ui64(trgt uint64) string {
	return strconv.FormatUint(uint64(trgt), 10)
}

func (conv Convertor) ToStr4Bool(trgt bool) string {
	return strconv.FormatBool(trgt)
}

func (conv Convertor) ToStr4F32(trgt float32) string {
	return strconv.FormatFloat(float64(trgt), 'f', 10, 32)
}

func (conv Convertor) ToStr4F64(trgt float64) string {
	return strconv.FormatFloat(float64(trgt), 'f', 10, 64)
}

func (conv Convertor) ToI8(trgt string) int8 {
	i, err := strconv.ParseInt(trgt, 10, 8)
	if err != nil {
		return conv.dftInt8
	}
	return int8(i)
}

func (conv Convertor) ToI16(trgt string) int16 {
	i, err := strconv.ParseInt(trgt, 10, 16)
	if err != nil {
		return conv.dftInt16
	}
	return int16(i)
}

func (conv Convertor) ToI32(trgt string) int32 {
	i, err := strconv.ParseInt(trgt, 10, 32)
	if err != nil {
		return conv.dftInt32
	}
	return int32(i)
}

func (conv Convertor) ToI64(trgt string) int64 {
	i, err := strconv.ParseInt(trgt, 10, 64)
	if err != nil {
		return conv.dftInt64
	}
	return i
}

func (conv Convertor) ToUi8(trgt string) uint8 {
	i, err := strconv.ParseUint(trgt, 10, 8)
	if err != nil {
		return conv.dftUint8
	}
	return uint8(i)
}

func (conv Convertor) ToUi16(trgt string) uint16 {
	i, err := strconv.ParseUint(trgt, 10, 16)
	if err != nil {
		return conv.dftUint16
	}
	return uint16(i)
}

func (conv Convertor) ToUi32(trgt string) uint32 {
	i, err := strconv.ParseUint(trgt, 10, 32)
	if err != nil {
		return conv.dftUint32
	}
	return uint32(i)
}

func (conv Convertor) ToUi64(trgt string) uint64 {
	i, err := strconv.ParseUint(trgt, 10, 64)
	if err != nil {
		return conv.dftUint64
	}
	return i
}

func (conv Convertor) ToBool(trgt string) bool {
	b, err := strconv.ParseBool(trgt)
	if err != nil {
		return conv.dftBool
	}
	return b
}

func (conv Convertor) ToF32(trgt string) float32 {
	f, err := strconv.ParseFloat(trgt, 32)
	if err != nil {
		return conv.dftFloat32
	}
	return float32(f)
}

func (conv Convertor) ToF64(trgt string) float64 {
	f, err := strconv.ParseFloat(trgt, 64)
	if err != nil {
		return conv.dftFloat64
	}
	return f
}
