package sort

// MergeSort sorts a slice of any ordered type using the merge sort algorithm.
// It returns a new sorted slice without modifying the original
func MergeSort[T Ordered](items []T) []T {
	if len(items) <= 1 {
		result := make([]T, len(items))
		copy(result, items)
		return result
	}

	middle := len(items) / 2

	left := MergeSort(items[:middle])
	right := MergeSort(items[middle:])

	return merge(left, right)
}

// merge combines two sorted slices into a single sorted xlice using the comparison function
func merge[T Ordered](left, right []T) []T {
	result := make([]T, len(left)+len(right))
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	return result
}

// MergeSortFuncInPlace sorts slice of any type using a comparison function
// It modifies the slice in-palce
func MergeSortInPlace[T Ordered](items []T) {
	temp := make([]T, len(items))

	mergeSortImpl(items, 0, len(items)-1, temp)
}

// mergeSortImpl is a helper function for MergeSortInPlace
func mergeSortImpl[T Ordered](items []T, left, right int, temp []T) {
	if left < right {
		mid := left + (right-left)/2

		mergeSortImpl(items, left, mid, temp)
		mergeSortImpl(items, mid+1, right, temp)

		mergeInPlace(items, left, mid, right, temp)
	}
}

// mergeInPlace merges two sorted subarrays into a single sorted subarray in-place
func mergeInPlace[T Ordered](items []T, left, mid, right int, temp []T) {
	for i := left; i <= right; i++ {
		temp[i] = items[i]
	}

	i, j, k := left, mid+1, left

	for i <= mid && j <= right {
		if temp[i] <= temp[j] {
			items[k] = temp[i]
			i++
		} else {
			items[k] = temp[j]
			j++
		}
		k++
	}

	for i <= mid {
		items[k] = temp[i]
		i++
		k++
	}
}

// MergeSortFunc sorts a slice of any type using a comparsion func
// It returns a new sorted slice without modifying the original
func MergeSortFunc[T any](items []T, less func(a, b T) bool) []T {
	if len(items) <= 1 {
		result := make([]T, len(items))
		copy(result, items)
		return result
	}

	middle := len(items) / 2

	left := MergeSortFunc(items[:middle], less)
	right := MergeSortFunc(items[middle:], less)

	return mergeFunc(left, right, less)
}

// mergeFunc combines two sorted slices into a single int oa single sorted slice using a comparison function
func mergeFunc[T any](left, right []T, less func(a, b T) bool) []T {
	result := make([]T, len(left)+len(right))
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if less(left[i], right[j]) || !less(right[j], left[i]) {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}
	return result
}

// MergeSortFuncInPlace sorts a slice of any type using a compasrison func
// It modifies the slice in-place
func MergeSortFuncInPlace[T any](items []T, less func(a, b T) bool) {
	temp := make([]T, len(items))

	mergeSortFuncImpl(items, 0, len(items)-1, temp, less)
}

// mergeSortFuncImpl is a helper func for MergeSortFuncInPlace
func mergeSortFuncImpl[T any](items []T, left, right int, temp []T, less func(a, b T) bool) {
	if left < right {
		mid := left + (right-left)/2

		mergeSortFuncImpl(items, left, mid, temp, less)
		mergeSortFuncImpl(items, mid+1, right, temp, less)

		mergeFuncInPlace(items, left, mid, right, temp, less)
	}
}

// mergeFuncInPlace merges two sorted subarrays into a single sorted subarray in-place using a comparison func
func mergeFuncInPlace[T any](items []T, left, mid, right int, temp []T, less func(a, b T) bool) {
	for i := left; i <= right; i++ {
		temp[i] = items[i]
	}

	i, j, k := left, mid+1, left

	for i <= mid && j <= right {
		if less(temp[i], temp[j]) || !less(temp[j], temp[i]) {
			items[k] = temp[i]
			i++
		} else {
			items[k] = temp[j]
			j++
		}
		k++
	}

	for i <= mid {
		items[k] = temp[i]
		i++
		k++
	}
}
