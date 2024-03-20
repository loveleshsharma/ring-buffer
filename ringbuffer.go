package main

import "errors"

/*
Features of RingBuffer
	1. Generics
	2. Thread Safe
*/

var (
	ErrBufferFull  = errors.New("buffer is full")
	ErrBufferEmpty = errors.New("buffer is empty")
)

type RingBuffer struct {
	arr []int

	read  int
	write int

	isEmpty bool
}

func NewRingBuffer(size int) RingBuffer {
	return RingBuffer{
		arr:     make([]int, size),
		read:    0,
		write:   0,
		isEmpty: true,
	}
}

func (b *RingBuffer) Write(data int) error {
	if b.isBufferFull() {
		return ErrBufferFull
	}

	b.arr[b.write] = data
	b.isEmpty = false
	b.incrementWritePointer()

	return nil
}

func (b *RingBuffer) Read() (int, error) {
	if b.isEmpty == true {
		return 0, ErrBufferEmpty
	}

	value := b.arr[b.read]
	b.incrementReadPointer()

	if b.read == b.write { // all values from the buffer is read
		b.isEmpty = true
	}

	return value, nil
}

func (b *RingBuffer) isBufferFull() bool {
	if (b.write == b.read) && b.isEmpty == false {
		return true
	}
	return false
}

func (b *RingBuffer) incrementWritePointer() {
	b.write = incrementPointer(b.write, len(b.arr))
}

func (b *RingBuffer) incrementReadPointer() {
	b.read = incrementPointer(b.read, len(b.arr))
}

func incrementPointer(currPtr int, arrLen int) int {
	if currPtr == arrLen-1 {
		currPtr = 0
		return currPtr
	}

	currPtr++
	return currPtr
}
