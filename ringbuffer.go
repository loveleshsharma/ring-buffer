package main

import "errors"

/*
Features of RingBuffer
	1. Generics
	2. Thread Safe
*/

var (
	ErrBufferFull = errors.New("buffer is full")
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

func (b *RingBuffer) isBufferFull() bool {
	if (b.write == b.read) && b.isEmpty == false {
		return true
	}
	return false
}
