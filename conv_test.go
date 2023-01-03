package conv

import (
	"errors"
	"testing"
	"time"
)

func Test_Bool(t *testing.T) {
	b := true
	if *Bool(b) != b {
		t.Errorf("Bool(%v) = %v, want %v", b, *Bool(b), b)
	}

	b = false
	if *Bool(b) != b {
		t.Errorf("Bool(%v) = %v, want %v", b, *Bool(b), b)
	}
}

func Test_Int(t *testing.T) {
	i := int(1)
	if *Int(1) != i {
		t.Errorf("Int(%v) = %v, want %v", i, *Int(i), i)
	}
}

func Test_Int8(t *testing.T) {
	i := int8(1)
	if *Int8(1) != i {
		t.Errorf("Int8(%v) = %v, want %v", i, *Int8(i), i)
	}
}

func Test_Int16(t *testing.T) {
	i := int16(1)
	if *Int16(1) != i {
		t.Errorf("Int16(%v) = %v, want %v", i, *Int16(i), i)
	}
}

func Test_Int32(t *testing.T) {
	i := int32(1)
	if *Int32(1) != i {
		t.Errorf("Int32(%v) = %v, want %v", i, *Int32(i), i)
	}
}

func Test_Int64(t *testing.T) {
	i := int64(1)
	if *Int64(1) != i {
		t.Errorf("Int64(%v) = %v, want %v", i, *Int64(i), i)
	}
}

func Test_Uint(t *testing.T) {
	i := uint(1)
	if *Uint(1) != i {
		t.Errorf("Uint(%v) = %v, want %v", i, *Uint(i), i)
	}
}

func Test_Uint8(t *testing.T) {
	i := uint8(1)
	if *Uint8(1) != i {
		t.Errorf("Uint8(%v) = %v, want %v", i, *Uint8(i), i)
	}
}

func Test_Uint16(t *testing.T) {
	i := uint16(1)
	if *Uint16(1) != i {
		t.Errorf("Uint16(%v) = %v, want %v", i, *Uint16(i), i)
	}
}

func Test_Uint32(t *testing.T) {
	i := uint32(1)
	if *Uint32(1) != i {
		t.Errorf("Uint32(%v) = %v, want %v", i, *Uint32(i), i)
	}
}

func Test_Uint64(t *testing.T) {
	i := uint64(1)
	if *Uint64(1) != i {
		t.Errorf("Uint64(%v) = %v, want %v", i, *Uint64(i), i)
	}
}

func Test_Float32(t *testing.T) {
	f := float32(1)
	if *Float32(1) != f {
		t.Errorf("Float32(%v) = %v, want %v", f, *Float32(f), f)
	}
}

func Test_Float64(t *testing.T) {
	f := float64(1)
	if *Float64(1) != f {
		t.Errorf("Float64(%v) = %v, want %v", f, *Float64(f), f)
	}
}

func Test_Complex64(t *testing.T) {
	c := complex64(1)
	if *Complex64(1) != c {
		t.Errorf("Complex64(%v) = %v, want %v", c, *Complex64(c), c)
	}
}

func Test_Complex128(t *testing.T) {
	c := complex128(1)
	if *Complex128(1) != c {
		t.Errorf("Complex128(%v) = %v, want %v", c, *Complex128(c), c)
	}
}

func Test_String(t *testing.T) {
	s := "hello"
	if *String(s) != s {
		t.Errorf("String(%v) = %v, want %v", s, *String(s), s)
	}
}

func Test_Rune(t *testing.T) {
	r := rune('a')
	if *Rune(r) != r {
		t.Errorf("Rune(%v) = %v, want %v", r, *Rune(r), r)
	}
}

func Test_Byte(t *testing.T) {
	b := byte('a')
	if *Byte(b) != b {
		t.Errorf("Byte(%v) = %v, want %v", b, *Byte(b), b)
	}
}

func Test_Time(t *testing.T) {
	tm := time.Now()
	if *Time(tm) != tm {
		t.Errorf("Time(%v) = %v, want %v", tm, *Time(tm), tm)
	}
}

func Test_Error(t *testing.T) {
	err := errors.New("error")
	if *Error(err) != err {
		t.Errorf("Error(%v) = %v, want %v", err, *Error(err), err)
	}
}
