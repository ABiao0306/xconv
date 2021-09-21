package xconv

import (
	"encoding/json"
	"strconv"

	"gopkg.in/yaml.v2"
)

var _dconv Convertor = Convertor{}

// Set default value
func SetDftInt8(v int8) {
	_dconv.SetDftInt8(v)
}

func SetDftInt16(v int16) {
	_dconv.SetDftInt16(v)
}

func SetDftInt32(v int32) {
	_dconv.SetDftInt32(v)
}

func SetDftInt64(v int64) {
	_dconv.SetDftInt64(v)
}

func SetDftUint8(v uint8) {
	_dconv.SetDftUint8(v)
}

func SetDftUint16(v uint16) {
	_dconv.SetDftUint16(v)
}

func SetDftUint32(v uint32) {
	_dconv.SetDftUint32(v)
}

func SetDftUint64(v uint64) {
	_dconv.SetDftUint64(v)
}

func SetDftBool(v bool) {
	_dconv.SetDftBool(v)
}

func SetDftFloat32(v float32) {
	_dconv.SetDftFloat32(v)
}

func SetDftFloat64(v float64) {
	_dconv.SetDftFloat64(v)
}

func ToStr4I8(trgt int8) string {
	return strconv.FormatInt(int64(trgt), 10)
}

func ToStr4I16(trgt int16) string {
	return strconv.FormatInt(int64(trgt), 10)
}

func ToStr4I32(trgt int32) string {
	return strconv.FormatInt(int64(trgt), 10)
}

func ToStr4I64(trgt int64) string {
	return strconv.FormatInt(int64(trgt), 10)
}

func ToStr4Ui8(trgt uint8) string {
	return strconv.FormatUint(uint64(trgt), 10)
}

func ToStr4Ui16(trgt uint16) string {
	return strconv.FormatUint(uint64(trgt), 10)
}

func ToStr4Ui32(trgt uint32) string {
	return strconv.FormatUint(uint64(trgt), 10)
}

func ToStr4Ui64(trgt uint64) string {
	return strconv.FormatUint(uint64(trgt), 10)
}

func ToStr4Bool(trgt bool) string {
	return strconv.FormatBool(trgt)
}

func ToStr4F32(trgt float32) string {
	return strconv.FormatFloat(float64(trgt), 'f', 10, 32)
}

func ToStr4F64(trgt float64) string {
	return strconv.FormatFloat(float64(trgt), 'f', 10, 64)
}

func ToI8(trgt string) int8 {
	return _dconv.ToI8(trgt)
}

func ToI16(trgt string) int16 {
	return _dconv.ToI16(trgt)
}

func ToI32(trgt string) int32 {
	return _dconv.ToI32(trgt)
}

func ToI64(trgt string) int64 {
	return _dconv.ToI64(trgt)
}

func ToUi8(trgt string) uint8 {
	return _dconv.ToUi8(trgt)
}

func ToUi16(trgt string) uint16 {
	return _dconv.ToUi16(trgt)
}

func ToUi32(trgt string) uint32 {
	return _dconv.ToUi32(trgt)
}

func ToUi64(trgt string) uint64 {
	return _dconv.ToUi64(trgt)
}

func ToBool(trgt string) bool {
	return _dconv.ToBool(trgt)
}

func ToF32(trgt string) float32 {
	return _dconv.ToF32(trgt)
}

func ToF64(trgt string) float64 {
	return _dconv.ToF64(trgt)
}

func ToJsonByte(trgt interface{}) []byte {
	data, err := json.Marshal(trgt)
	if err != nil {
		return []byte("")
	}
	return data
}

func ToJsonStr(trgt interface{}) string {
	return string(ToJsonByte(trgt))
}

func Parse4JsonStr(data string, trgt interface{}) {
	json.Unmarshal([]byte(data), trgt)
}

func Parse4JsonByte(data []byte, trgt interface{}) {
	json.Unmarshal(data, trgt)
}

func ToYamlByte(trgt interface{}) []byte {
	data, err := yaml.Marshal(trgt)
	if err != nil {
		return []byte("")
	}
	return data
}

func ToYamlStr(trgt interface{}) string {
	return string(ToJsonByte(trgt))
}

func Parse4YamlStr(str string, trgt interface{}) {
	yaml.Unmarshal([]byte(str), trgt)
}

func Parse4YamlByte(data []byte, trgt interface{}) {
	yaml.Unmarshal(data, trgt)
}
