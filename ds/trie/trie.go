package trie

type Node struct {
	Children []*Node
	IsLeaf bool
}

type Trie struct {
	Root *Node
}

func NewTrieNode() *Node {
	return &Node{
		Children: make([]*Node, 26),
		IsLeaf: false,
	}
}

func NewTrie() *Trie {
	return &Trie{
		Root: NewTrieNode(),
	}
}

func (trie *Trie) Insert(s string) {
	temp := trie.Root

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
	temp := trie.Root

	for _,c := range s {
		if temp.Children[c-'a'] == nil {
			return false
		}
		temp = temp.Children[c-'a']
	}
	return temp.IsLeaf
}

func (trie *Trie) IsPrefix(s string) bool {
	temp := trie.Root

	for _,c := range s {
		if temp.Children[c-'a'] == nil {
			return false
		}
		temp = temp.Children[c-'a']
	}
	return true
}

func (trie *Trie) ReturnPrefixNode(s string) *Node {
	temp := trie.Root

	for _,c := range s {
		if temp.Children[c-'a'] == nil {
			break
		}
		temp = temp.Children[c-'a']
	}

	return temp
}

func (trie *Trie) Delete(node *Node, title string, index int) bool {
	next := node.Children[title[index] - 'a']

	// check if this was the last index
	if len(title) == index + 1 {
		// delete this entire node
		node = nil
		return true
	}

	if trie.Delete(next, title, index + 1) {
		// now check whether this node has multiple children
		ct := 0

		for _, temp := range node.Children {
			if temp != nil {
				ct++
			}
		}

		if ct == 1 {
			// delete this node as well
			node = nil
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}