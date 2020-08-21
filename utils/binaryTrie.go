package utils

const (
	//BINARY_SIZE total characters in english alphabet
	BINARY_SIZE = 2
	LEN         = 32
)

type binaryTrieNode struct {
	childrens [BINARY_SIZE]*binaryTrieNode
	isWordEnd bool
}

type binaryTrie struct {
	root *binaryTrieNode
}

func InitBinaryTrie() *binaryTrie {
	return &binaryTrie{
		root: &binaryTrieNode{},
	}
}

func (t *binaryTrie) Insert(number int) {
	current := t.root
	mask := 1 << (LEN - 1)
	for i := 0; i < LEN; i++ {
		index := 1
		if number&mask == 0 {
			index = 0
		}

		// log.Print("mask:", mask)
		// log.Print("index:", index)
		if current.childrens[index] == nil {
			current.childrens[index] = &binaryTrieNode{}
		}
		current = current.childrens[index]
		mask = mask >> 1
	}
	current.isWordEnd = true
}

//func (t *binaryTrie) find(word string) bool {
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
