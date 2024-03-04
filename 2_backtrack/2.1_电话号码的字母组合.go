package backtrack

/**
Tag: 回溯算法

Description:
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

Analysis:
时间复杂度：O(3^m * 4^n)
空间复杂度：O(m+n)
其中 m是输入中对应 3个字母的数字个数（包括数字 2、3、4、5、6、8），n 是输入中对应 4 个字母的数字个数（包括数字 7、9）
*/

var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}
var ans []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	ans = []string{}
	// 递归函数
	phoneBackTrack(digits, 0, "")
	return ans
}

func phoneBackTrack(digits string, index int, combination string) {
	if index == len(digits) {
		ans = append(ans, combination)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		for i := 0; i < len(letters); i++ {
			phoneBackTrack(digits, index+1, combination+string(letters[i]))
		}
	}
}
