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
	list *linkedlist.LinkedList[*Tab] // Tab: shared resource, so store it by reference
	trie *trie.Trie
}

func NewTab(title, url string) *Tab {
	return &Tab{
		title: title,
		url: url,
	}
}

func NewTabManager() *TabManager {
	return &TabManager{
		list: linkedlist.NewLinkedList[*Tab](),
		trie: trie.NewTrie(),
	}
}

// handle the insertion part in linkedlist
func (tabManager *TabManager) Insert(tab *Tab) {
	// first insert it in linkedlist
	nodeL := linkedlist.NewNode(tab)
	tabManager.list.Insert(nodeL)

	// now insert the tab in trie
	tabManager.trie.Insert(tab.title)
}

func dfs(node *trie.Node, soFar string, searchResult *[]string) {
	f := false
	for i, neigh := range node.Children {
		if neigh != nil {
			f = true
			dfs(neigh, soFar + string(rune('a' + i)), searchResult)
		}
	}

	if !f {
		*searchResult = append(*searchResult, soFar)
	}
}

func (tabManager *TabManager) Search(name string) []string {
	// 1. get to the node which matches the prefix the most
	// 2. do dfs/bfs to get the full name of tabs in the trie and return that list

	//1. node
	node := tabManager.trie.ReturnPrefixNode(name)
	
	//2. doing dfs - no need of keeping visited array (no cycles, indegree of all nodes =1)
	searchResult := make([]string, 0)

	dfs(node, name, &searchResult)

	return searchResult
}