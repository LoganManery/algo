package sort

func InsertionSort[T Ordered](items []T) {
	n := len(items)

	for i := 1; i < n; i++ {
		current := items[i]
		j := i - 1
		for j >= 0 && items[j] > current {
			items[j+1] = items[j]
			j--
		}
		items[j+1] = current
	}
}

func InsertionSortFunc[T any](items []T, less func(a, b T) bool) {
	n := len(items)
	for i := 1; i < n; i++ {
		current := items[i]
		j := i - 1
		for j >= 0 && !less(items[j], current) {
			items[j+1] = items[j]
			j--
		}
		items[j+1] = current
	}
}
