package easy

import "bytes"

/**
  url: https://leetcode.com/problems/unique-morse-code-words/solution/
  url: http://www.cnblogs.com/grandyang/p/9338988.html

  International Morse Code defines a standard encoding where each letter is mapped to a series of dots and dashes, as follows: "a" maps to ".-", "b" maps to "-...", "c" maps to "-.-.", and so on.
  For convenience, the full table for the 26 letters of the English alphabet is given below:
    a    b      c      d     e   f      g     h      i    j      k     l      m    n    o     p      q      r     s     t   u     v      w     x      y      z
  [".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]
  Now, given a list of words, each word can be written as a concatenation of the Morse code of each letter. For example, "cab" can be written as "-.-.-....-", (which is the concatenation "-.-." + "-..." + ".-"). We'll call such a concatenation, the transformation of a word.
  Return the number of different transformations among all words we have.

  Example:
  Input: words = ["gin", "zen", "gig", "msg"]
  Output: 2
  Explanation:
  The transformation of each word is:
  "gin" -> "--...-."
  "zen" -> "--...-."
  "gig" -> "--...--."
  "msg" -> "--...--."
  There are 2 different transformations, "--...-." and "--...--.".

  Note:
  The length of words will be at most 100.
  Each words[i] will have length in range [1, 12].
  words[i] will only consist of lowercase letters.
*/

/**
  题意：独特的摩斯码单词
  给了我们所有字母的摩斯码的写法，然后给了我们一个单词数组，问我们表示这些单词的摩斯码有多少种。因为某些单词的摩斯码表示是相同的，比如gin和zen就是相同的。最简单直接的方法就是我们求出每一个单词的摩斯码，然后将其放入一个HashSet中，利用其去重复的特性，从而实现题目的要求，最终HashSet中元素的个数即为所求
*/

var morse = []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}

func UniqueMorseRepresentations1(words []string) int {
	//morse := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	mp := make(map[string]int)
	for _, wd := range words {
		buff := new(bytes.Buffer)
		n := len(wd)
		for i := 0; i < n; i++ {
			buff.WriteString(morse[wd[i] - 'a'])
		}
		mp[buff.String()] = 1
	}
	return len(mp)
}

