package main

import "fmt"

type packet struct {
	id int
	payload string
}

type videoStreamingBuffer struct {
	buffer []packet
	capacity int
}

func NewVideoBuffer(cap int) *videoStreamingBuffer {
	return &videoStreamingBuffer{
		buffer: make([]packet, cap),
		capacity: cap,
	}
}

func (v *videoStreamingBuffer) addPacket(x packet) {
	// checking to which index does this packet belong
	idx := x.id
	
	for v.capacity <= idx {
		v.resize()
	}
	v.buffer[idx] = x
}

func (v *videoStreamingBuffer) resize() {
	newArray := make([]packet, v.capacity * 2)

	copy(newArray, v.buffer)

	v.buffer = newArray
	v.capacity = cap(newArray)

	fmt.Println("Resized to", v.capacity)
}

// the above approach of using id of packet as its index in the buffer is memory costly, time complexity has been traded off with space complexity

func main() {
	v := NewVideoBuffer(8)

	for i:=1;i<=10;i++ {
		v.addPacket(packet{
			id: i,
			payload: "source ip: 192.168.0.1 | destination ip: 52.0.131.132",
		})
	}
	v.addPacket(packet{
		id: 15,
		payload: "the end of series",
	})
	fmt.Println(v)
}