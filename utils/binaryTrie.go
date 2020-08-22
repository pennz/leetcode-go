package utils

const (
	//BINARY_SIZE total characters in english alphabet
	BINARY_SIZE = 2
	LEN         = 32
)

type BinaryTrieNode struct {
	Children  [BINARY_SIZE]*BinaryTrieNode
	IsWordEnd bool
	Value     int
}

type BinaryTrie struct {
	Root *BinaryTrieNode
}

func InitBinaryTrie() *BinaryTrie {
	return &BinaryTrie{
		Root: &BinaryTrieNode{},
	}
}

func (t *BinaryTrie) Insert(number int) {
	current := t.Root
	mask := 1 << (LEN - 1)
	for i := 0; i < LEN; i++ {
		index := 1
		if number&mask == 0 {
			index = 0
		}

		// log.Print("mask:", mask)
		// log.Print("index:", index)
		if current.Children[index] == nil {
			current.Children[index] = &BinaryTrieNode{}
		}
		current = current.Children[index]
		mask = mask >> 1
	}
	current.IsWordEnd = true
	current.Value = number
}

//func (t *BinaryTrie) find(word string) bool {
//	wordLength := len(word)
//	current := t.root
//	for i := 0; i < wordLength; i++ {
//		index := word[i] - 'a'
//		if current.childrens[index] == nil {
//			return false
//		}
//		current = current.childrens[index]
//	}
//	if current.isWordEnd {
//		return true
//	}
//	return false
//}
