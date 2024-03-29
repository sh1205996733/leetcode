package backtrack

import (
	"fmt"
	"testing"
)

// 全排列
func TestPermute(t *testing.T) {
	ret := permute([]int{1, 2})
	for _, r := range ret {
		fmt.Println(r)
	}
	fmt.Println(len(ret))
}

// 全排列II
func TestPermuteUnique(t *testing.T) {
	ret := permuteUnique([]int{1, 1, 2})
	for _, r := range ret {
		fmt.Println(r)
	}
	fmt.Println(len(ret))
}

// 组合总和
func TestCombinationSum(t *testing.T) {
	ret := combinationSum([]int{2, 3, 6, 7}, 7)
	for _, r := range ret {
		fmt.Println(r)
	}
}

// 括号生成
func TestGenerateParenthesis(t *testing.T) {
	ret := generateParenthesis(2)
	fmt.Println(ret)
}

// 电话号码的字母组合
func TestLetterCombinations(t *testing.T) {
	ret := letterCombinations("9")
	fmt.Println(ret)
}

// 子集
func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	ret := subsets(nums)
	fmt.Println(ret)
}
