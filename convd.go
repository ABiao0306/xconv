package xconv

import "strconv"

func ToInt8Str(trgt string, dft int8) int8 {
	i, err := strconv.ParseInt(trgt, 10, 8)
	if err != nil {
		return dft
	}
	return int8(i)
}

func ToInt16Str(trgt string, dft int16) int16 {
	i, err := strconv.ParseInt(trgt, 10, 16)
	if err != nil {
		return dft
	}
	return int16(i)
}

func ToInt32Str(trgt string, dft int32) int32 {
	i, err := strconv.ParseInt(trgt, 10, 32)
	if err != nil {
		return dft
	}
	return int32(i)
}

func ToInt64Str(trgt string, dft int64) int64 {
	i, err := strconv.ParseInt(trgt, 10, 64)
	if err != nil {
		return dft
	}
	return i
}

func ToUint8Str(trgt string, dft uint8) uint8       {
    i, err := strconv.ParseUint(trgt, 10, 8)
	if err != nil {
        return dft
	}
	return uint8(i)
}

func ToUint16Str(trgt string, dft uint16) uint16    {
    i, err := strconv.ParseUint(trgt, 10, 16)
	if err != nil {
        return dft
	}
	return uint16(i)
}

func ToUint32Str(trgt string, dft uint32) uint32    {
    i, err := strconv.ParseUint(trgt, 10, 32)
	if err != nil {
        return dft
	}
	return uint32(i)
}

func ToUint64Str(trgt string, dft uint64) uint64    {
    i, err := strconv.ParseUint(trgt, 10, 64)
	if err != nil {
        return dft
	}
	return i
}

func ToBoolStr(trgt string, dft bool) bool          {
    b, err := strconv.ParseBool(trgt)
	if err != nil {
        return dft
	}
	return b
}

func ToFloat32Str(trgt string, dft float32) float32 {
    f, err := strconv.ParseFloat(trgt, 32)
	if err != nil {
        return dft
	}
	return float32(f)
}

func ToFloat64Str(trgt string, dft float64) float64 {
    f, err := strconv.ParseFloat(trgt, 64)
	if err != nil {
		return dft
	}
	return f
}
