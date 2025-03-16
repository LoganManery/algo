package algo

// Regular Expression Matching

// IsMatch checls if string s matches pattern p with support for '.' and '*'
func IsMatch(s string, p string) bool {
	// Create a 2D DP table: dp[i][j] represents if s[0:1] matches p[0:j]
	// dp[i][j] = true means s[0...i-1] matches p[0...j-1]
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	// Empty pattern matches empty string
	dp[0][0] = true

	// Handle patterns like a*, a*b*, a*b*c* ... which can match empty string
	for j := 1; j <= len(p); j++ {
		if p[j-1] == '*' && j >= 2 {
			dp[0][j] = dp[0][j-2]
		}
	}

	// Fill the dp table
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				// Current characters match directly
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// * can match zero of the preceding elements
				dp[i][j] = dp[i][j-2]

				// * can match one or more of the preceding element
				// if the preceding character in pattern matches current character in string
				// or if the preceding character is '.'
				if j >= 2 && (p[j-2] == '.' || p[j-2] == s[i-1]) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
			// else dp[i][j] remains false (default value)
		}
	}
	return dp[len(s)][len(p)]
}
