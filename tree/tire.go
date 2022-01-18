package tree

type TireNode struct {
	char  string
	isEnd bool
	node  map[rune]*TireNode
}
type TireTree struct {
	root *TireNode
}

// 初始化树

func NewTireTree() *TireTree {
	// 初始化根节点
	trieNode := NewTrieNode("/")
	return &TireTree{trieNode}
}

// 初始化节点
func NewTrieNode(char string) *TireNode {
	return &TireNode{
		char:  char,
		isEnd: false,
		node:  make(map[rune]*TireNode),
	}
}

// 往 Trie 树中插入一个单词
func (t *TireTree) InsertTireTree(word string) {
	tree := t.root
	for _, code := range word {
		value, ok := tree.node[code]
		if !ok {
			value = NewTrieNode(string(code))
			tree.node[code] = value
		}
		tree = value
	}
	tree.isEnd = true
}

// 在 Trie 树中查找一个单词
func (t *TireTree) FindWord(word string) bool {
	tree := t.root
	for _, code := range word {
		value, ok := tree.node[code]
		if !ok {
			return false
		}
		tree = value
	}
	if tree.isEnd == false {
		return false
	}
	return true
}
