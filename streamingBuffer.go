package main

import (
	"fmt"
	"strconv"
)

type packet struct {
	seqID int
	payload string
}

type videoStreamingBuffer struct {
	buffer []*packet // sliding window
	size int
	header int // header gives the index of the expected id in the buffer
	expectedID int
}

func NewVideoBuffer(size int) *videoStreamingBuffer {
	v := videoStreamingBuffer{
		buffer: make([]*packet, size),
		size: size,
		header: 0,
		expectedID: 0,
	}

	for i := range v.buffer {
		v.buffer[i] = nil
	}

	return &v
}

func (v *videoStreamingBuffer) addPacket(x *packet) {
	// out of bounds check
	header := v.header
	size := v.size
	packetID := x.seqID

	if packetID < v.expectedID || packetID >= v.expectedID + size {
		fmt.Printf("Packet %d dropped\n", packetID)
		return
	} 
	
	diffID := packetID - v.expectedID
	idx := (header + diffID) % size

	v.buffer[idx] = x

	v.play()
}

func (v *videoStreamingBuffer) play() {
	for v.buffer[v.header] != nil {
		fmt.Printf("Packet %d: %s\n", v.expectedID, v.buffer[v.header].payload)
		v.buffer[v.header] = nil
		v.header++
		v.header %= v.size
		v.expectedID++
	}
}

func callStreamingBuffer() {
	v := NewVideoBuffer(4)

	for i:=1;i<10;i+=2 {
		v.addPacket(&packet{
			seqID: i,
			payload: "hello from " + strconv.Itoa(i),
		})
		v.addPacket(&packet{
			seqID: i-1,
			payload: "hello from " + strconv.Itoa(i-1),
		})
	}
}