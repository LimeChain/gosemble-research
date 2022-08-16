package utils

import (
	"unsafe"
)

func BytesToPointerAndSize(data []byte) int64 {
	dataPtr := uintptr(unsafe.Pointer(&data))
	dataLen := len(data)
	return PointerAndSizeToInt64(int32(dataPtr), int32(dataLen))
}

// func PointerAndSizeToString(ptr uint32, size uint32) string {
// 	// We use SliceHeader, not StringHeader as it allows us to fix the capacity to what was allocated.
// 	// Tinygo requires these as uintptrs even if they are int fields.
// 	// https://github.com/tinygo-org/tinygo/issues/1284
// 	return *(*string)(unsafe.Pointer(&reflect.SliceHeader{
// 		Data: uintptr(ptr),
// 		Len:  uintptr(size),
// 		Cap:  uintptr(size),
// 	}))
// }

// func StringToPointerAndSize(s string) (uint32, uint32) {
// 	buf := []byte(s)
// 	ptr := &buf[0]
// 	unsafePtr := uintptr(unsafe.Pointer(ptr))
// 	return uint32(unsafePtr), uint32(len(buf))
// }

func Int64ToPointerAndSize(pointerAndSize int64) (dataPtr int32, dataLen int32) {
	return int32(pointerAndSize), int32(pointerAndSize >> 32)
}

func PointerAndSizeToInt64(dataPtr int32, dataLen int32) int64 {
	return int64(dataPtr) | (int64(dataLen) << 32)
}

func WriteMemory(offset int32, size int32, data [11]byte) {
	bs := unsafe.Slice((*byte)(unsafe.Pointer(uintptr(offset))), uintptr(size))

	for i := int32(0); i < size; i++ {
		limit := int32(0)
		if int32(len(data)) > size {
			limit = size
		} else {
			limit = int32(len(data))
		}
		if i >= limit {
			break
		}

		bs[i] = byte(data[i])

		// ptr := (*byte)(unsafe.Pointer(uintptr(offset) + uintptr(i)))
		// *ptr = byte(data[i])

		// ptr := unsafe.Pointer(uintptr(offset))
		// bs := (*[MAX_ARRAY_SIZE]byte)(ptr)[:]
	}
}
