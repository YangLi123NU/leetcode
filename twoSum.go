package leetcode

import "fmt"

func TwoSum(nums []int, target int) []int {
	res := make([]int, 0)
	a := make(map[int]int)
	for idx, val := range nums {
		fmt.Println(idx, val)
		if _, ok := a[target - val]; ok {
			res = append(res, idx)
			res = append(res, a[target - val])
		}
		a[val] = idx
		fmt.Println(a)
	}
	return res
}


