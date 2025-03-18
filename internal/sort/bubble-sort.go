package sort

import (
	"cmp"
)

type Ordered interface {
	cmp.Ordered
}

func BubbleSort[T Ordered](items []T) {
	n := len(items)
	swapped := true

	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				// Swap the item
				items[i], items[i+1] = items[i+1], items[i]
				swapped = true
			}
		}
		n--
	}
}

func BubbleSortFunc[T any](items []T, less func(a, b T) bool) {
	n := len(items)
	swapped := true

	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			// If items[i] should come after items[i+1], swap them
			if !less(items[i], items[i+1]) {
				// Swap the items
				items[i], items[i+1] = items[i+1], items[i]
				swapped = true
			}
		}
		n--
	}
}
