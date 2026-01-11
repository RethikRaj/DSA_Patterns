package two_pointers

import (
	"slices"
)

// func util_two_sum_sorted(nums []int, leftPtr int, rightPtr int, target int, triplets *[][]int) {
// 	for leftPtr < rightPtr {
// 		sum := nums[leftPtr] + nums[rightPtr]
// 		if sum > target {
// 			// skip duplicates and shift rightPtr
// 			for leftPtr < rightPtr && nums[rightPtr-1] == nums[rightPtr] {
// 				rightPtr -= 1
// 			}
// 			rightPtr -= 1
// 		} else if sum < target {
// 			// skip duplicates and shift leftPtr
// 			for leftPtr < rightPtr && nums[leftPtr] == nums[leftPtr+1] {
// 				leftPtr += 1
// 			}
// 			leftPtr += 1
// 		} else {
// 			triplet := make([]int, 3)
// 			triplet[0] = nums[leftPtr]
// 			triplet[1] = nums[rightPtr]
// 			triplet[2] = -target
// 			*triplets = append(*triplets, triplet)
// 			// Update leftPtr and rightPtr
// 			// skip duplicates and shift rightPtr
// 			for leftPtr < rightPtr && nums[rightPtr-1] == nums[rightPtr] {
// 				rightPtr -= 1
// 			}
// 			rightPtr -= 1
// 			// skip duplicates and shift leftPtr
// 			for leftPtr < rightPtr && nums[leftPtr] == nums[leftPtr+1] {
// 				leftPtr += 1
// 			}
// 			leftPtr += 1
// 		}
// 	}
// }

// func threeSum(nums []int) [][]int {
// 	// 1. Sort the nums
// 	slices.Sort(nums)

// 	triplets := make([][]int, 0)

// 	// 2. Now do two_sum_sorted
// 	index := 0
// 	for index < len(nums)-2 {
// 		target := -nums[index]

// 		util_two_sum_sorted(nums, index+1, len(nums)-1, target, &triplets)

// 		// Update index -> skip duplicates
// 		for index < len(nums)-1 && nums[index] == nums[index+1] {
// 			index += 1
// 		}
// 		index += 1
// 	}
// 	return triplets
// }
//

// ----------- The below is same as above , the only difference is the way i skipped duplicates -------------.
func util_two_sum_sorted(nums []int, leftPtr int, rightPtr int, target int, triplets *[][]int) {
	for leftPtr < rightPtr {
		sum := nums[leftPtr] + nums[rightPtr]

		if sum == target {
			*triplets = append(*triplets, []int{-target, nums[leftPtr], nums[rightPtr]})

			// Move both pointers and skip duplicates
			for leftPtr < rightPtr && nums[leftPtr] == nums[leftPtr+1] {
				leftPtr++
			}
			for leftPtr < rightPtr && nums[rightPtr] == nums[rightPtr-1] {
				rightPtr--
			}
			leftPtr++
			rightPtr--
		} else if sum < target {
			leftPtr++
		} else {
			rightPtr--
		}
	}
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	triplets := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicates for the first element
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		util_two_sum_sorted(nums, i+1, len(nums)-1, -nums[i], &triplets)

	}
	return triplets
}
