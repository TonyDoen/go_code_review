package hard

import (
	"math"
	"strconv"
)

/**
 * We have a sorted set of digits D, a non-empty subset of {'1','2','3','4','5','6','7','8','9'}.  (Note that '0' is not included.)
 * Now, we write numbers using these digits, using each digit as many times as we want.  For example, if D = {'1','3','5'}, we may write numbers such as '13', '551', '1351315'.
 * Return the number of positive integers that can be written (using the digits of D) that are less than or equal to N.
 *
 * Example 1:
 * Input: D = ["1","3","5","7"], N = 100
 * Output: 20
 * Explanation:
 * The 20 numbers that can be written are:
 * 1, 3, 5, 7, 11, 13, 15, 17, 31, 33, 35, 37, 51, 53, 55, 57, 71, 73, 75, 77.
 *
 * Example 2:
 * Input: D = ["1","4","9"], N = 1000000000
 * Output: 29523
 * Explanation:
 * We can write 3 one digit numbers, 9 two digit numbers, 27 three digit numbers,
 * 81 four digit numbers, 243 five digit numbers, 729 six digit numbers,
 * 2187 seven digit numbers, 6561 eight digit numbers, and 19683 nine digit numbers.
 * In total, this is 29523 integers that can be written using the digits of D.
 *
 * Note:
 * D is a subset of digits '1'-'9' in sorted order.
 * 1 <= N <= 10^9
 */
/**
 * 题意：最大为 N 的数字组合
 *
 * 这道题给了我们一个有序字符串数组，里面是0到9之间的数（这里博主就纳闷了，既然只有一位数字，为啥不用 char 型，而要用 string 型），然后又给了一个整型数字N，问无限制次数使用D中的任意个数字，能组成多个不同的小于等于D的数字。先来分析例子1，当N为 100 时，所有的一位数和两位数都是可以的，既然可以重复使用数字，假设总共有n个数字，那么对于两位数来说，十位上和个位上分别都有n种可能，总共就是 n^2 种可能，对于一位数来说，总共n种可能。那么看到这里就可以归纳出当N总共有 len 位的话，我们就可以快速的求出不超过 len-1 位的所有情况综合，用个 for 循环，累加n的指数即可。然后就要来分析和数字N位数相等的组合，题目中的两个的例子的N都是1开始的，实际上N可以是任何数字，举个例子来说吧，假如 D={"1","3","5","7"}，N=365，那么根据前面的分析，我们可以很快的算出所有的两位数和一位数的组合情况总数 4 + 4^2 = 20 个。现在要来分析三位数都有哪些组合，由于D数组是有序的，所以我们从开头遍历的话先取到的就是最小的，这时候有三种情况，小于，等于，和大于，每种的处理情况都有些许不同，这里就拿上面提到的例子进行一步一步的分析：
 *
 * 对于N的百位数字3来说，D中的1小于N中的百位上的3，那么此时百位上固定为1，十位和个位上就可以是任意值了，即 1xx，共有 4^2 = 16 个。
 * 对于N的百位数字3来说，D中的3等于N中的百位上的3，那么此时百位上固定为3，十位和个位的值还是不确定，此时就不能再继续遍历D中的数字了，因为之后的数字肯定大于3，但是我们可以继续尝试N的下一位。
 * 对于N的十位数字6来说，D中的1小于N中的十位上的6，那么百位和十位分别固定为3和1，个位上就可以是任意值了，即 31x，共有 4 个。
 * 对于N的十位数字6来说，D中的3小于N中的十位上的6，那么百位和十位分别固定为3和3，个位上就可以是任意值了，即 33x，共有 4 个。
 * 对于N的十位数字6来说，D中的5小于N中的十位上的6，那么百位和十位分别固定为3和5，个位上就可以是任意值了，即 35x，共有 4 个。
 * 对于N的十位数字6来说，D中的7大于N中的十位上的6，此时再也组不成小于N的数字了，直接返回最终的 20+16+4+4+4=48 个。
 * 代码如下：
 */
type Set struct {
	Map map[interface{}]struct{} // struct为结构体类型的变量
}
func (s *Set) Add(items ...interface{}) error {
	for _, item := range items {
		s.Map[item] = struct{}{} // 空结构体变量的内存占用大小为0
	}
	return nil
}
func (s *Set) Size() int {
	return len(s.Map)
}
func (s *Set) Clear() {
	s.Map = make(map[interface{}]struct{})
}
func (s *Set) Contains(item interface{}) bool {
	_, ok := s.Map[item]
	return ok
}
func (s *Set) Equal(other *Set) bool {
	if s.Size() != other.Size() { // 如果两者Size不相等，就不用比较了
		return false
	}

	for key := range s.Map { // 迭代查询遍历
		if !other.Contains(key) { // 只要有一个不存在就返回false
			return false
		}
	}
	return true
}
func NewSet() *Set {
	return &Set{Map: make(map[interface{}]struct{})} // 获取Set的地址 // 声明map类型的数据结构
}
func NewSetWithItem(items ...interface{}) *Set {
	s := NewSet()
	_ = s.Add(items...)
	return s
}

func AtMostNGivenDigitSet(d *Set, n int) int {
	nStr := strconv.Itoa(n)
	length := len(nStr)
	res := 0
	dSize := d.Size()
	for i := 1; i < length; i++ {
		res += int(math.Pow(float64(dSize), float64(i)))
	}
	for i := 0; i < length; i++ {
		hasSameNum := false
		for key :=range d.Map {
			s := key.(string)
			if s[0] < nStr[i] {
				res += int(math.Pow(float64(dSize), float64(length - 1 - i)))
			} else if s[0] == nStr[i] {
				hasSameNum = true
			}
		}
		if !hasSameNum {
			return res
		}
	}
	return res + 1
}
