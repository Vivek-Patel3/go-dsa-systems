package main

import "fmt"

type song struct {
	title string
	lyrics string
	duration uint64 // conversion to minutes:second is not too costly, so we will perform that operation on demand when the song is played
	artist string
	next *song
	previous *song
}

type songNode struct {
	node *song // cutting memory cost - by value would create a copy of song to be inserted as a node in playlist
	next *songNode
	prev *songNode
}

type songPlaylist struct {
	header *songNode
	tail *songNode // to allow insertion of song at the end
}

func NewSong(title, lyrics, artist string, duration uint64) *song {
	return &song {
		title: title,
		lyrics: lyrics,
		artist: artist,
		duration: duration,
	}
}
	
func NewSongNode(s *song) *songNode {
	return &songNode{
		node: s,
	}
}

func NewSongPlaylist() *songPlaylist {
	return &songPlaylist{
		header: nil,
		tail: nil,
	}
}

func (playlist *songPlaylist) AddToPlaylist(song *song) *songNode {
	songNode := NewSongNode(song)

	if playlist.header == nil {
		playlist.header = songNode
		playlist.tail = playlist.header
		return songNode
	}

	lastSong := playlist.tail
	lastSong.next = songNode

	songNode.prev = lastSong
	playlist.tail = songNode
	return songNode // returning songNode so that it can be stored to be reused during shuffle logic
}

// here we passed songNode instead of song, because at the deletion time (from the playlist), we have songNodes ready with us. Adding the song phase did not have songNode ready
func (playlist *songPlaylist) DeleteFromPlaylist(songNode *songNode) {
	removeNode, _ := playlist.searchNode(songNode)

	if removeNode == nil {
		fmt.Println("Song doesn't exist in playlist")
	}

	removeNode.prev = removeNode.next

	removeNode = nil // garbage collector will free the songNode memory since now it has no pointers pointing to it
}

func (playlist *songPlaylist) searchNode(songNd *songNode) (*songNode, int) {
	temp := playlist.header
	i := 0

	check := func(node1, node2 *songNode) bool {
		if node1.node.title == node2.node.title && node1.node.artist == node2.node.artist {
			return true
		} 
		return false
	}

	for temp != nil {
		if check(temp, songNd) {
			return temp, i
		}
		temp = temp.next
		i++
	}
	return nil, -1
}

func (playlist *songPlaylist) len() int {
	temp := playlist.header
	ct := 0

	for temp != nil {
		temp = temp.next
		ct++
	}

	return ct
}

func (playlist *songPlaylist) shuffleNode(songNode *songNode, destinationPos int) {
	// temp points to the location where songNode needs to be
	temp := playlist.header
	toRight := false

	for i:=0;i<destinationPos;i++ {
		if temp == songNode {
			toRight = true
		}

		temp = temp.next
	}
	
	// changing the neighbours at the original location
	if songNode.prev != nil {
		songNode.prev.next = songNode.next
	} else {
		playlist.header = songNode.next
	} 

	if songNode.next != nil {
		songNode.next.prev = songNode.prev
	} else {
		playlist.tail = songNode.prev
	}

	// now change the neighbours at the destination location

	if !toRight {
		songNode.prev = temp.prev
		songNode.next = temp
	
		if temp.prev != nil {
			temp.prev.next = songNode
			temp.prev = songNode
		} else {
			playlist.header.prev = songNode
			playlist.header = songNode
		}
	} else {
		songNode.prev = temp
		songNode.next = temp.next

		if temp.next != nil {
			temp.next.prev = songNode
			temp.next = songNode
		} else {
			temp.next = songNode
		}
	}
}

func (playlist *songPlaylist) printPlaylist() {
	if playlist.header == nil {
		fmt.Println("Empty playlist")
		return
	}

	temp := playlist.header

	for temp!=nil {
		fmt.Printf("Song: %s\n", temp.node.title)
		
		temp = temp.next
	}

	fmt.Println("--------")
	fmt.Println("End of Playlist")
}