package medium

/**
 * Given a non-empty array of integers, every element appears twice except for one. Find that single one.
 *
 * Note:
 * Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
 *
 * Example 1:
 * Input: [2,2,1]
 * Output: 1
 *
 * Example 2:
 * Input: [4,1,2,1,2]
 * Output: 4
 *
 *
 * 题意：
 * 这道题给了我们一个非空的整数数组，说是除了一个数字之外所有的数字都正好出现了两次，让我们找出这个只出现一次的数字。题目中让我们在线性的时间复杂度内求解，那么一个非常直接的思路就是使用 HashSet，利用其常数级的查找速度。
 *
 *
 * 思路：
 * 两个相同的数异或运算结果为0,所以数组中的元素,
 * 依次进行异或运算，最后剩下的就是只存在一次的数字
 *
 * 数字在计算机是以二进制存储的，每位上都是0或1，
 * 如果我们把两个相同的数字异或，0与0 '异或' 是0，1与1 '异或' 也是0，那么我们会得到0。
 */
func SingleNumber00(arr []int) int {
	if nil == arr {
		return -1
	}

	result := 0
	for _, num := range arr {
		result ^= num
	}
	return result
}
func SingleNumber01(arr []int) int {
	if nil == arr {
		return -1
	}

	st := NewSet()
	for _, num := range arr {
		if st.Contains(num) {
			_ = st.Remove(num)
		} else {
			_ = st.Add(num)
		}
	}
	for key := range st.m {
		return key.(int)
	}
	return -1
}

/**
 * single number ii
 * Given an array of integers, every element appears three times except for one.
 * Find that single one.
 *
 * Note:
 * Your algorithm should have a linear runtime complexity.
 * Could you implement it without using extra memory?
 *
 *
 * 题意：
 * 这道题给了我们一个非空的整数数组，说是除了一个数字之外所有的数字都正好出现了3次，
 */
func SingleNumber10(arr []int) int {
	if nil == arr {
		return -1
	}
	result := 0
	for i := 0; i < 32; i++ {
		curBitCnt := 0
		for _, num := range arr {
			bit := (num >> uint(i)) & 1
			curBitCnt += bit
		}
		curBitCnt %= 3
		result |= curBitCnt << uint(i)
	}
	return result
}

func SingleNumber11(arr []int) int {
	if nil == arr {
		return -1
	}
	ones, twos, threes := 0, 0, 0
	for _, num := range arr {
		twos |= ones & num
		ones ^= num
		threes = ones & twos
		ones &= ^threes
		twos &= ^threes
	}
	return ones
}

func SingleNumber12(arr []int) int {
	if nil == arr {
		return -1
	}
	low, high := 0, 0
	for _, num := range arr {
		low = (low ^ num) &^ high
		high = (high ^ num) &^ low
	}
	return low
}
