package math

import (
	"fmt"
)

func _check_bit(_idx uint32) bool {
	return _idx >= 0 && _idx < 32
}
func _check_byte(_idx uint32) bool {
	return _idx >= 0 && _idx < 4

}
func _check_word(_idx uint32) bool {
	return _idx >= 0 && _idx < 2
}

func GetBit(_val uint32, _idx uint32) (uint32, error) {
	if _check_bit(_idx) {
		r := (_val & (1 << _idx)) >> _idx
		return r, nil
	}
	return 0, fmt.Errorf("GetBit(%d,%d) bit index not in range", _val, _idx)
}
func GetByte(_val uint32, _idx uint32) (uint32, error) {
	if _check_byte(_idx) {
		bits := _idx * 8
		r := (_val & (0xFF << bits)) >> bits
		return r, nil
	}
	return 0, fmt.Errorf("GetByte(%d,%d) byte index not in range", _val, _idx)
}
func GetWord(_val uint32, _idx uint32) (uint32, error) {
	if _check_word(_idx) {
		bits := _idx * 16
		r := (_val & (0xFFFF << bits)) >> bits
		return r, nil
	}
	return 0, fmt.Errorf("GetWord(%d,%d) _word index not in range", _val, _idx)
}
