package hard

/**
 * There are N children standing in a line. Each child is assigned a rating value.
 * You are giving candies to these children subjected to the following requirements:
 * 1. Each child must have at least one candy.
 * 2. Children with a higher rating get more candies than their neighbors.
 * What is the minimum candies you must give?
 */
/**
 * 题意：
 * 先来看看两遍遍历的解法，首先初始化每个人一个糖果，然后这个算法需要遍历两遍，第一遍从左向右遍历，如果右边的小盆友的等级高，等加一个糖果，这样保证了一个方向上高等级的糖果多。然后再从右向左遍历一遍，如果相邻两个左边的等级高，而左边的糖果又少的话，则左边糖果数为右边糖果数加一。最后再把所有小盆友的糖果数都加起来返回即可。
 *
 * 思路：
 * 1、先全部分一个糖果
 * 2、从左到右判断是否满足规则，不满足的增加糖果
 * 3、从右往左判断是否满足规则，不满足的增加糖果
 */
func GiveCandy(ratings []int) int {
	length := len(ratings)
	if length < 1 {
		return 0
	}
	candies := make([]int, length)
	for i := 0; i < length; i++ { // 构建每个孩子拥有糖果数量的数组,初始化为1
		candies[i] = 1
	}
	for i := 1; i < length; i++ { // 从左到右处理数组
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	sum := 0
	for i := length - 1; i > 0; i-- { // 从右到左处理数组，并对数组中的元素求和，但是要记得最后加上第一个元素
		if ratings[i-1] > ratings[i] { // 如果左侧元素不大于右侧元素，则更新左侧元素为右侧元素+1
			if candies[i-1] <= candies[i] {
				candies[i-1] = candies[i] + 1
			}
		}
		sum += candies[i]
	}
	return sum + candies[0]
}
