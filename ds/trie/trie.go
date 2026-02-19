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
	if node == nil {
		return false // this case is because next might be nil
	}

	if index == len(title) {
		// end of word
		// now check whether this node is the leaf node (if yes, then it has to be deleted by its parent)
		if node.IsLeaf {
			return true
		} else {
			return false
		}
	}

	next := node.Children[title[index] - 'a']

	if trie.Delete(next, title, index + 1) {
		// true, so need to delete this child pointed by index
		node.Children[title[index] - 'a'] = nil

		// now check whether this node has to be deleted or not
		ct := 0

		for _, child := range node.Children {
			if child != nil {
				ct++
			}
		}

		if ct > 0 {
			// has more children apart from the deleted
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}