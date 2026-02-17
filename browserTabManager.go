package main

import (
	"github.com/VivekPatel3/go-dsa-systems/ds/linkedlist"
	"github.com/VivekPatel3/go-dsa-systems/ds/trie"
)

type Tab struct {
	title string
	url string
}

type TabManager struct {
	list *linkedlist.LinkedList[Tab]
	trie trie.Trie
}

func NewTab() *Tab {
	return &Tab{}
}

func NewTabManager() *TabManager {
	return &TabManager{
		list: linkedlist.NewLinkedList[Tab](),
	}
}

// handle the insertion part in linkedlist
func (tabManager *TabManager) Insert(tab Tab) {
	// first insert it in linkedlist
	nodeL := linkedlist.NewNode(tab)
	tabManager.list.Insert(nodeL)

	// now insert the tab in trie
	tabManager.trie.Insert(tab.title)
}

func (tabManager *TabManager) Search(name string) {
	// need a list of tabs matching this name 
	// therefore creating a wrapper around trie node to also store a list of tabs with the given title
	
}