package main

import "fmt"

/*
Given four integer arrays nums1, nums2, nums3, and nums4 all of length n, return the number of tuples (i, j, k, l) such that:

0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0


Example 1:

Input: nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
Output: 2
Explanation:
The two tuples are:
1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0

Example 2:

Input: nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
Output: 1


Constraints:

n == nums1.length
n == nums2.length
n == nums3.length
n == nums4.length
1 <= n <= 200
-2^28 <= nums1[i], nums2[i], nums3[i], nums4[i] <= 2^28
*/
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {

	ans := 0
	tmp_map := make(map[int]int)
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			tmp_map[num1+num2]++
		}
	}
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			ans += tmp_map[-(num3 + num4)]
		}
	}
	return ans
}

/*
func fourSumCount(A []int, B []int, C []int, D []int) int {
    ans, tmp := 0, make(map[int]int)
    for _, a := range A {
        for _, b := range B {
            tmp[a + b]++
        }
    }
    for _, c := range C {
        for _, d := range D {
            ans += tmp[-(c + d)]
        }
    }
    return ans
}

func fourSumCount(A []int, B []int, C []int, D []int) int {

    dict := make(map[int]int)
    l := len(A)
    res := 0

    for i := 0; i < l; i++ {
        for j := 0; j < l; j++ {
            dict[A[i] + B[j]]++
        }
    }

    for i := 0; i < l; i++ {
        for j := 0; j < l; j++ {
            res += dict[-C[i]-D[j]]
        }
    }

    return res
}

*/
func main() {
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}

	ans := fourSumCount(nums1, nums2, nums3, nums4)
	fmt.Println(ans)
}
