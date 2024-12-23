package main

// snip ------------------------------------------------------------------------

type Numeric interface {
	int8 | int16 | int32 | int64 |
		uint8 | uint16 | uint32 | uint64 |
		int | uint | uintptr |
		float32 | float64 |
		complex64 | complex128
}

type Integer interface {
	int8 | int16 | int32 | int64 |
		uint8 | uint16 | uint32 | uint64 |
		int | uint | uintptr
}

type OrderedNumeric interface {
	int8 | int16 | int32 | int64 |
		uint8 | uint16 | uint32 | uint64 |
		int | uint | uintptr |
		float32 | float64
}

type Ordered interface {
	int8 | int16 | int32 | int64 |
		uint8 | uint16 | uint32 | uint64 |
		int | uint | uintptr |
		float32 | float64 |
		string
}

type SignedNumeric interface {
	int8 | int16 | int32 | int64 |
		int | float32 | float64
}

type SignedInteger interface {
	int8 | int16 | int32 | int64 |
		int
}
