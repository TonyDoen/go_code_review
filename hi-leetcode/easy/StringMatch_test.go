package easy

import (
	"fmt"
	"testing"
)

func TestViolentMatch(t *testing.T) {
	s, p := "BBC ABCDAB ABCDABCDABDE", "ABCDABD"
	res := ViolentMatch(s, p)
	fmt.Printf("%d\n", res)
}

/**
匹配过程如下:
1. S[0]为B，P[0]为A，不匹配，执行第②条指令：“如果失配（即S[i]! = P[j]），令i = i - (j - 1)，j = 0”，S[1]跟P[0]匹配，相当于模式串要往右移动一位（i=1，j=0）
|B|BC ABCDAB ABCDABCDABDE
|A|BCDABD

2. S[1]跟P[0]还是不匹配，继续执行第②条指令：“如果失配（即S[i]! = P[j]），令i = i - (j - 1)，j = 0”，S[2]跟P[0]匹配（i=2，j=0），从而模式串不断的向右移动一位（不断的执行“令i = i - (j - 1)，j = 0”，i从2变到4，j一直为0）
B|B|C ABCDAB ABCDABCDABDE
 |A|BCDABD

......

3. 直到S[4]跟P[0]匹配成功（i=4，j=0），此时按照上面的暴力匹配算法的思路，转而执行第①条指令：“如果当前字符匹配成功（即S[i] == P[j]），则i++，j++”，可得S[i]为S[5]，P[j]为P[1]，即接下来S[5]跟P[1]匹配（i=5，j=1）
BBC |A|BCDAB ABCDABCDABDE
    |A|BCDABD

4. S[5]跟P[1]匹配成功，继续执行第①条指令：“如果当前字符匹配成功（即S[i] == P[j]），则i++，j++”，得到S[6]跟P[2]匹配（i=6，j=2），如此进行下去
BBC A|B|CDAB ABCDABCDABDE
    A|B|CDABD

......

5. 直到S[10]为空格字符，P[6]为字符D（i=10，j=6），因为不匹配，重新执行第②条指令：“如果失配（即S[i]! = P[j]），令i = i - (j - 1)，j = 0”，相当于S[5]跟P[0]匹配（i=5，j=0）
BBC ABCDAB| |ABCDABCDABDE
    ABCDAB|D|

6. 至此，我们可以看到，如果按照暴力匹配算法的思路，尽管之前文本串和模式串已经分别匹配到了S[9]、P[5]，但因为S[10]跟P[6]不匹配，所以文本串回溯到S[5]，模式串回溯到P[0]，从而让S[5]跟P[0]匹配。
BBC A|B|CDAB ABCDABCDABDE
     |A|BCDABD

而S[5]肯定跟P[0]失配。为什么呢？因为在之前第4步匹配中，我们已经得知S[5] = P[1] = B，而P[0] = A，即P[1] != P[0]，故S[5]必定不等于P[0]，所以回溯过去必然会导致失配。那有没有一种算法，让i 不往回退，只需要移动j 即可呢？
答案是肯定的。这种算法就是本文的主旨KMP算法，它利用之前已经部分匹配这个有效信息，保持i 不回溯，通过修改j 的位置，让模式串尽量地移动到有效的位置。
*/

func TestKMPMatch(t *testing.T) {
	s, p := "BBC ABCDAB ABCDABCDABDE", "ABCDABD"
	res := KMPMatch(s, p)
	fmt.Printf("%d\n", res)
}

func TestSunDayMatch(t *testing.T) {
	//s, p := "BBC ABCDAB ABCDABCDABDE", "ABCDABD"
	//res := SundayMatch(s, p)
	res := SundayMatch("substring searching algorithm", "search")
	fmt.Printf("%d\n", res)
}
