package jianzhi

import (
	"strings"
)

//func main() {
//	dictionary := []string{"cat", "bat", "rat"}
//	sentence := "the cattle was rattled by the battery"
//	fmt.Println(ReplaceWords(dictionary, sentence))
//	// 输出：the cat was rat by the bat
//}

type TrieNode struct {
	childrens [26]*TrieNode
	isWordEnd bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{},
	}
}

func (t Trie) insert(word string) {
	curNode := t.root
	for _, c := range word {
		index := c - 'a'
		if curNode.childrens[index] == nil {
			curNode.childrens[index] = &TrieNode{}
		}
		curNode = curNode.childrens[index]
	}
	curNode.isWordEnd = true
}

func (t Trie) find(word string) (bool, int) {
	curNode := t.root
	for i, c := range word {
		index := c - 'a'
		if curNode.childrens[index] == nil {
			return false, -1
		}
		curNode = curNode.childrens[index]
		if curNode.isWordEnd {
			return true, i
		}
	}
	return false, -1
}

func ReplaceWords(dictionary []string, sentence string) string {
	// 构建字典树（前缀树）
	myTrie := NewTrie()
	for _, str := range dictionary {
		myTrie.insert(str)
	}
	// 遍历字典树，查找最短前缀
	words := strings.Split(sentence, " ")
	for i, word := range words {
		if exist, index := myTrie.find(word); exist {
			words[i] = word[:index+1]
		}
	}
	return strings.Join(words, " ")
}
