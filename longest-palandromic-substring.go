
func longestPalindrome(s string) string {
    longest := ""
    temp := ""

    for i := 0; i < len(s); i++ {
	
		temp = expandAndCount(s, i, i)
		
		if len(temp) > len(longest) {
            longest = temp
        }
        
        temp = expandAndCount(s, i, i+1)
                
        if len(temp) > len(longest) {
            longest = temp
        }
    }
    
    return longest
}

func expandAndCount(s string, start, end int) string {
    for start >= 0 && end <= len(s) - 1 && s[start] == s[end] {
        start--
        end++
    }
    return s[start+1: end]
} 