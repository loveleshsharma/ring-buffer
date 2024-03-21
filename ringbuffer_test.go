package ringbuffer

import (
	"errors"
	"testing"
)

func TestNewRingBuffer_isBufferFull_ShouldReturnTrueIfBufferIsFull(t *testing.T) {
	testBuffer := RingBuffer{
		arr:     make([]int, 5),
		read:    0,
		write:   0,
		isEmpty: false,
	}

	actual := testBuffer.isBufferFull()

	if actual != true {
		t.Errorf("isBufferFull should return true, actual value: %v", actual)
	}
}

func TestNewRingBuffer_isBufferFull_ShouldReturnTrueIfBufferIsFullAndReadAndWritePointersAreDifferent(t *testing.T) {
	testBuffer := RingBuffer{
		arr:     make([]int, 5),
		read:    3,
		write:   3,
		isEmpty: false,
	}

	actual := testBuffer.isBufferFull()

	if actual != true {
		t.Errorf("isBufferFull should return true, actual value: %v", actual)
	}
}

func TestRingBuffer_Write_ShouldReturnNilIfWriteIsSuccessful(t *testing.T) {
	testBuffer := RingBuffer{
		arr:     make([]int, 5),
		read:    0,
		write:   0,
		isEmpty: true,
	}

	actual := testBuffer.Write(10)

	if actual != nil {
		t.Errorf("Write should return nil, actual value: %v", actual)
	}
}

func TestRingBuffer_Write_ShouldReturnErrorIfBufferIsFull(t *testing.T) {
	testBuffer := RingBuffer{
		arr:     make([]int, 5),
		read:    0,
		write:   0,
		isEmpty: false,
	}

	actual := testBuffer.Write(10)

	if !errors.Is(actual, ErrBufferFull) {
		t.Errorf("Write should return error, actual value: %v", actual)
	}
}

func TestRingBuffer_Write_ShouldSetIsEmptyFlagToFalseAfterWrite(t *testing.T) {
	testBuffer := RingBuffer{
		arr:     make([]int, 5),
		read:    0,
		write:   0,
		isEmpty: true,
	}

	_ = testBuffer.Write(10)

	if testBuffer.isEmpty != false {
		t.Errorf("isEmpty flag should be set to false, actual value: %v", testBuffer.isEmpty)
	}
}

func TestRingBuffer_Write_ShouldInsertValueInRingBuffer(t *testing.T) {
	testBuffer := RingBuffer{
		arr:     make([]int, 5),
		read:    0,
		write:   0,
		isEmpty: true,
	}

	_ = testBuffer.Write(10)

	if testBuffer.write != 1 {
		t.Errorf("write pointer should set as 1, actual value: %v", testBuffer.write)
	}
}

func TestRingBuffer_incrementPointerShouldReturnNxtPtrIfCurrPtrIsLessThanLastPtr(t *testing.T) {
	actual := incrementPointer(3, 5)

	if actual != 4 {
		t.Errorf("current pointer should increment to 4, actual value: %v", actual)
	}
}

func TestRingBuffer_incrementPointerShouldReturnZeroIfCurrPtrIsTheLastIndex(t *testing.T) {
	actual := incrementPointer(4, 5)

	if actual != 0 {
		t.Errorf("current pointer should increment to 0, actual value: %v", actual)
	}
}
