package tools

import (
	"encoding/binary"
	"math"
	"sort"
)

func Uint16ToFloat64(r map[uint16]uint16) float64 {
	return float64(binary.BigEndian.Uint16(getBytes(r)))
}

func Uint32ToFloat64(r map[uint16]uint16) float64 {
	return float64(binary.BigEndian.Uint32(getBytes(r)))
}

func Uint64ToFloat64(r map[uint16]uint16) float64 {
	return float64(binary.BigEndian.Uint64(getBytes(r)))
}

func Int16ToFloat64(r map[uint16]uint16) float64 {
	return float64(int16(Uint16ToFloat64(r)))
}

func Int32ToFloat64(r map[uint16]uint16) float64 {
	return float64(int32(Uint32ToFloat64(r)))
}

func Int64ToFloat64(r map[uint16]uint16) float64 {
	return float64(int64(Uint64ToFloat64(r)))
}

func BitsToFloat64(r map[uint16]uint16) float64 {
	return float64(math.Float32frombits(binary.BigEndian.Uint32(getBytes(r))))
}

func getBytes(r map[uint16]uint16) []byte {
	b := make([]byte, len(r)*2)

	keys := make([]int, 0, len(r))
	for k := range r {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	for i, k := range keys {
		binary.BigEndian.PutUint16(b[i*2:i*2+2], r[uint16(k)])
	}

	return b
}
