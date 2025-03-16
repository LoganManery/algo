package algo

func LongestValidParentheses(s string) int {
	// Initialize variables to track the longest valid substring
	maxLength := 0
	stack := []int{-1} // Start with -1 as a base index

	// Iterate through each cahracter in the string
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			// Push the index of opening
			stack = append(stack, i)
		} else {
			// Pop the top element (corresponding opening parenthesis or base index)
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				// If stack is empty, push current index as new base
				stack = append(stack, i)
			} else {
				// Calculate the length of valid substring ending at current position
				// Current position - position before the start of valid substring
				currentLength := i - stack[len(stack)-1]

				if currentLength > maxLength {
					maxLength = currentLength
				}
			}
		}
	}
	return maxLength
}
