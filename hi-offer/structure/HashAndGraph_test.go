package structure

import "testing"

func TestLinkedHashMap(t *testing.T) {
	mp := &LinkedHashMap{}
	mp.Put(1, "1,")
	mp.Put(5, "5,")
	mp.Put(4, "4,")
	mp.Put(3, "3,")

	cur := mp.Next()
	for nil != cur {
		print(mp.Get(cur.Value).(string))
		cur = cur.Next
	}

	mp.Remove(3)

	println()
	cur = mp.Next()
	for nil != cur {
		print(mp.Get(cur.Value).(string))
		cur = cur.Next
	}
}

func TestFirstNotRepeatingChar(t *testing.T) {
	src := "google"
	idx := FirstNotRepeatingChar(src)
	println(idx)
}

func TestHasPath(t *testing.T) {
	/**
	 * a b c e
	 * s f c s
	 * a d e e
	 */
	matrix := [][]rune{{'a', 'b', 'c', 'e'}, {'s', 'f', 'c', 's'}, {'a', 'd', 'e', 'e'}}
	target := []rune{'b', 'c', 'c', 'e', 'd'}
	res := HasPath(matrix, target)
	println(res)
}

func TestCountMove(t *testing.T) {
	res := CountMove(18, 1000, 1000)
	println(res)
}
