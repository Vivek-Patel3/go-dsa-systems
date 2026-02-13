package main

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
	songItem := NewSongNode(song)
	 
	if playlist.header == nil {
		playlist.header = songItem
		playlist.tail = playlist.header
		return
	}

	lastSong := playlist.tail
	lastSong.next = songItem

	playlist.tail = songItem
}

