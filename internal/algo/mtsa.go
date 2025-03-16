package algo

// Median Two Sorted Arrays

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the smaller array for simplicity
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	x, y := len(nums1), len(nums2)
	low, high := 0, x

	for low <= high {
		// Partition for smaller array
		partitionX := (low + high) / 2
		// Calculate the partition position for the larger array
		partitionY := (x+y+1)/2 - partitionX

		// Find the four boundary elements
		// maxLeftX: maximum element on the left side
		// maxRightX: maximum element on the right side
		// maxLeftY: maximum element on the left side
		// maxRightY: maximum element on the right side
		maxLeftX := getMax(nums1, partitionX)
		minRightX := getMin(nums1, partitionX, x)
		maxLeftY := getMax(nums2, partitionY)
		minRightY := getMin(nums2, partitionY, y)

		// Found the correct partition
		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			// If the total length is odd
			if (x+y)%2 != 0 {
				return float64(max(maxLeftX, maxLeftY))
			}
			// If the total length is even
			return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2.0
		} else if maxLeftX > minRightY {
			// move towards left in nums1
			high = partitionX - 1
		} else {
			// move towards right in nums1
			low = partitionX + 1
		}
	}
	// should never reach with valid inputs
	return 0.0
}

// Helper, gets the minimum element
func getMax(nums []int, partition int) int {
	if partition == 0 {
		return -1000001
	}
	return nums[partition-1]
}

// Helper, gets the maximum element
func getMin(nums []int, partition, length int) int {
	if partition == length {
		return 1000001
	}
	return nums[partition]
}

// Helper, max of two ints
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper, min of two ints
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
