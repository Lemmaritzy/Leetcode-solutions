package main

func main() {
	nums := [5]int{3, 2, 4}
	
	Calc(nums[0:5], 9)
}

func Calc(nums []int, target int) (result []int) {
	for i1, v1 := range nums {

		for i2, v2 := range nums {
			sum := v1 + v2
			if sum == target && i1 != i2 {
				return []int{i1, i2}
			}
		}
	}

	return
}
