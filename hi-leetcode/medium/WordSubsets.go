package medium

import "container/list"

/**
 * Word Subsets
 * We are given two arrays A and B of words.  Each word is a string of lowercase letters.
 * Now, say that word b is a subset of word a if every letter in b occurs in a, including multiplicity.  For example, "wrr" is a subset of "warrior", but is not a subset of "world".
 * Now say a word a from A is universal if for every b in B, b is a subset of a.
 * Return a list of all universal words in A.  You can return the words in any order.
 *
 * Example 1:
 * Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["e","o"]
 * Output: ["facebook","google","leetcode"]
 *
 * Example 2:
 * Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["l","e"]
 * Output: ["apple","google","leetcode"]
 *
 * Example 3:
 * Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["e","oo"]
 * Output: ["facebook","google"]
 *
 * Example 4:
 * Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["lo","eo"]
 * Output: ["google","leetcode"]
 *
 * Example 5:
 * Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["ec","oc","ceo"]
 * Output: ["facebook","leetcode"]
 *
 * Note:
 * 1 <= A.length, B.length <= 10000
 * 1 <= A[i].length, B[i].length <= 10
 * A[i] and B[i] consist only of lowercase letters.
 * All words in A[i] are unique: there isn't i != j with A[i] == A[j].
 */

/**
 * 题意：单词子集合
 * 思路：这道题定义了两个单词之间的一种子集合关系，就是说假如单词b中的每个字母都在单词a中出现了（包括重复字母），就说单词b是单词a的子集合。现在给了两个单词集合A和B，让找出集合A中的所有满足要求的单词，使得集合B中的所有单词都是其子集合。
 * 配合上题目中给的一堆例子，意思并不难理解，根据子集合的定义关系，其实就是说若单词a中的每个字母的出现次数都大于等于单词b中每个字母的出现次数，单词b就一定是a的子集合。现在由于集合B中的所有单词都必须是A中某个单词的子集合，那么其实只要对于每个字母，都统计出集合B中某个单词中出现的最大次数，
 * 比如对于这个例子，B=["eo","oo"]，其中e最多出现1次，而o最多出现2次，那么只要集合A中有单词的e出现不少1次，o出现不少于2次，则集合B中的所有单词一定都是其子集合。
 * 这就是本题的解题思路，这里使用一个大小为 26 的一维数组 charCnt 来统计集合B中每个字母的最大出现次数，而将统计每个单词的字母次数的操作放到一个子函数 helper 中，当 charCnt 数组更新完毕后，下面就开始检验集合A中的所有单词了。对于每个遍历到的单词，还是要先统计其每个字母的出现次数，然后跟 charCnt 中每个位置上的数字比较，只要均大于等于 charCnt 中的数字，就可以加入到结果 res 中了
 */
func WordSubset1(a, b []string) *list.List {
	res := &list.List{}
	charCnt := make([]int, 26)

	bLength := len(b)
	for i := 0; i < bLength; i++ {
		t := count(b[i])
		for i := 0; i < 26; i++ {
			charCnt[i] = max(charCnt[i], t[i])
		}
	}
	aLength := len(a)
	for i := 0; i < aLength; i++ {
		as := a[i]
		t := count(as)
		j := 0
		for ; j < 26; j++ {
			if t[j] < charCnt[j] {
				break
			}
		}
		if 26 == j {
			res.PushBack(as)
		}
	}
	return res
}

func count(c string) []int {
	res := make([]int, 26)
	length := len(c)
	for i := 0; i < length; i++ {
		index := c[i] - 'a'
		res[index]++
	}
	return res
}
