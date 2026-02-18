package trie

type Node struct {
	Children []*Node
	IsLeaf bool
}

type Trie struct {
	root *Node
}

func NewTrieNode() *Node {
	return &Node{
		Children: make([]*Node, 26),
		IsLeaf: false,
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
		if temp.Children[c - 'a'] == nil {
			// create new Node
			temp.Children[c-'a'] = NewTrieNode() 
		} 
		temp = temp.Children[c-'a']
	}

	temp.IsLeaf = true
}

func (trie *Trie) Search(s string) bool {
	temp := trie.root

	for _,c := range s {
		if temp.Children[c-'a'] == nil {
			return false
		}
		temp = temp.Children[c-'a']
	}
	return temp.IsLeaf
}

func (trie *Trie) IsPrefix(s string) bool {
	temp := trie.root

	for _,c := range s {
		if temp.Children[c-'a'] == nil {
			return false
		}
		temp = temp.Children[c-'a']
	}
	return true
}

func (trie *Trie) ReturnPrefixNode(s string) *Node {
	temp := trie.root

	for _,c := range s {
		if temp.Children[c-'a'] == nil {
			break
		}
		temp = temp.Children[c-'a']
	}

	return temp
}