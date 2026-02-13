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

func (playlist *songPlaylist) AddToPlaylist(song *song) {
	songNode := NewSongNode(song)

	if playlist.header == nil {
		playlist.header = songNode
		playlist.tail = playlist.header
		return
	}

	lastSong := playlist.tail
	lastSong.next = songNode

	playlist.tail = songNode
}

// here we passed songNode instead of song, because at the deletion time (from the playlist), we have songNodes ready with us. Adding the song phase did not have songNode ready
func (playlist *songPlaylist) DeleteFromPlaylist(songNode *songNode) {
	removeNode := playlist.searchNode(songNode)

	if removeNode == nil {
		fmt.Println("Song doesn't exist in playlist")
	}

	removeNode.prev = removeNode.next

	removeNode = nil // garbage collector will free the songNode memory since now it has no pointers pointing to it
}

func (playlist *songPlaylist) searchNode(songNode *songNode) *songNode {
	temp := playlist.header

	for temp != nil {
		if temp == songNode {
			return temp
		}
		temp = temp.next
	}
	return nil
}