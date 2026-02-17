package trie

type node struct {
	children []*node
	isLeaf bool
}

type Trie struct {
	root *node
}

func NewTrieNode() *node {
	return &node{
		children: make([]*node, 26),
		isLeaf: false,
	}
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (trie *Trie) Insert(s string) {
	temp := trie.root

	for _,c := range s {
		if temp.children[c - 'a'] == nil {
			// create new node
			temp.children[c-'a'] = NewTrieNode() 
		} 
		temp = temp.children[c-'a']
	}

	temp.isLeaf = true
}

func (trie *Trie) Search(s string) bool {
	temp := trie.root

	for _,c := range s {
		if temp.children[c-'a'] == nil {
			return false
		}
		temp = temp.children[c-'a']
	}
	return temp.isLeaf
}

func (trie *Trie) IsPrefix(s string) bool {
	temp := trie.root

	for _,c := range s {
		if temp.children[c-'a'] == nil {
			return false
		}
		temp = temp.children[c-'a']
	}
	return true
}