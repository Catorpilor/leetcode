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

// useOnePass time complexity O(N), space complexity O(1)
func useOnePass(nums []int) {
	redIdx, whiteIdx, blueIdx := 0, 0, len(nums)-1
	for whiteIdx <= blueIdx {
		if nums[whiteIdx] == 0 {
			// it's red
			nums[redIdx], nums[whiteIdx] = nums[whiteIdx], nums[redIdx]
			redIdx++
			whiteIdx++
		} else if nums[whiteIdx] == 1 {
			// it's white just increase whiteIdx
			whiteIdx++
		} else {
			nums[whiteIdx], nums[blueIdx] = nums[blueIdx], nums[whiteIdx]
			blueIdx--
		}
	}
}
