package jianzhi

import (
	"fmt"
	"strings"
)

func main() {
	dictionary := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	fmt.Println(replaceWords(dictionary, sentence))
	// 输出：the cat was rat by the bat
}

func replaceWords(dictionary []string, sentence string) string {
	// 构建字典树（前缀树）
	type trie map[rune]trie
	root := trie{}
	for _, s := range dictionary {
		cur := root
		for _, c := range s {
			if cur[c] == nil {
				cur[c] = trie{}
			}
			cur = cur[c]
		}
		cur['#'] = trie{}
	}

	// 遍历字典树，查找最短前缀
	words := strings.Split(sentence, " ")
	for i, word := range words {
		cur := root
		for j, w := range word {
			if cur['#'] != nil {
				// 用词根替换单词
				words[i] = word[:j]
				break
			}
			if cur[w] == nil {
				break
			}
			cur = cur[w]
		}
	}
	return strings.Join(words, " ")
}
