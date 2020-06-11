package sort

func sortColors(nums []int) {
	useCountSort(nums)
}

// useCountSort time complexity O(N), space complexity O(1)
func useCountSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	counts := [3]int{}
	for _, num := range nums {
		counts[num]++
	}
	prev := 0
	for i, count := range counts {
		if count != 0 {
			for j := prev; j < prev+count; j++ {
				nums[j] = i
			}
		}
		prev += count
	}
}
