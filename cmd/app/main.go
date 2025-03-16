package main

import (
	"fmt"

	"github.com/loganmanery/go/internal/algo"
)

func main() {
	testRnkg()
}

func testMKSLL() {
	create := algo.CreateLinkedList
	list1 := create([]int{1, 4, 5})
	list2 := create([]int{1, 3, 4})
	list3 := create([]int{2, 6})

	merge := algo.MergeKLists
	result := merge([]*algo.ListNode{list1, list2, list3})

	slice := algo.LinkedListToSlice
	fmt.Println("Example 1 Output: ", slice(result))

	result = merge([]*algo.ListNode{})
	fmt.Println("Example 2 Output: ", slice(result))

	result = merge([]*algo.ListNode{nil})
	fmt.Println("Example 3 Output: ", slice(result))

	result = merge([]*algo.ListNode{nil, nil, nil})
	fmt.Println("Example 4 Output: ", slice(result))
}

func testRem() {
	match := algo.IsMatch

	s1 := "aa"
	p1 := "a"
	fmt.Printf("Example 1: s = %q, p = %q, IsMatch = %t\n", s1, p1, match(s1, p1))

	s2 := "aa"
	p2 := "a"
	fmt.Printf("Example 2: s = %q, p = %q, IsMatch = %t\n", s1, p1, match(s2, p2))

	s3 := "aa"
	p3 := "a"
	fmt.Printf("Example 3: s = %q, p = %q, IsMatch = %t\n", s1, p1, match(s3, p3))

	testCases := []struct {
		s        string
		p        string
		expected bool
	}{
		{"mississippi", "mis*is*p", true},
		{"aab", "a*a", true},
		{"aaa", "ab*a*c*a", true},
		{"aaaa", "a*", true},
		{"", ".*", true},
		{"", "a*", true},
		{"a", "", false},
		{"aaaa", "a{4}", false},
		{"abcd", "abcd", false},
		{"ab", ".*c", false},
	}

	for i, tc := range testCases {
		result := match(tc.s, tc.p)
		fmt.Printf("Test case %d: s = %q, p = %q, expected = %t, got = %t\n", i+1, tc.s, tc.p, tc.expected, result)
	}
}

func testRnkg() {
	reverse := algo.ReverseKGroup
	example1 := createLinkedList([]int{1, 2, 3, 4, 5})
	result1 := reverse(example1, 2)
	fmt.Println("Example 1 Output:", linkedListToSlice(result1))

	example2 := createLinkedList([]int{1, 2, 3, 4, 5})
	result2 := reverse(example2, 3)
	fmt.Println("Example 2 Output:", linkedListToSlice(result2))

	testCases := []struct {
		values []int
		k      int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 4},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{1, 2, 3, 4, 5, 6}, 6},
		{[]int{1}, 1},
		{[]int{1, 2}, 3},
	}

	for i, tc := range testCases {
		list := createLinkedList(tc.values)
		result := reverse(list, tc.k)
		fmt.Printf("Test case %d: values = %v, k = %d, result = %v\n", i+1, tc.values, tc.k, linkedListToSlice(result))
	}

}

func createLinkedList(values []int) *algo.ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &algo.ListNode{Val: values[0]}
	curr := head

	for i := 1; i < len(values); i++ {
		curr.Next = &algo.ListNode{Val: values[i]}
		curr = curr.Next
	}
	return head
}

func linkedListToSlice(head *algo.ListNode) []int {
	var result []int

	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func LVPar() {
	testCases := []string{
		"(()",
		")()())",
		"",
	}

	for _, s := range testCases {
		result := algo.LongestValidParentheses(s)
		fmt.Printf("Input: \"%s\"\nOutput: %d\n\n", s, result)
	}
}
