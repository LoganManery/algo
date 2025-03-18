package sort

func SelectionSort[T Ordered](items []T) {
	n := len(items)
	for i := 0; i < n-1; i++ {
		minIdx := i

		for j := i + 1; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}

		if minIdx != i {
			items[i], items[minIdx] = items[minIdx], items[i]
		}
	}
}

func SelectionSortFunc[T any](items []T, less func(a, b T) bool) {
	n := len(items)

	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if less(items[j], items[minIdx]) {
				minIdx = j
			}
		}
		if minIdx != i {
			items[i], items[minIdx] = items[minIdx], items[i]
		}
	}
}
