

func lengthOfLongestSubstring(s string) int {
	finalCount, temp := 0, 0
	var m map[string]int
	m = make(map[string]int)

	for i:=0; i< len(s); i++{	
		for j:=i; j<len(s); j++{
			if _, ok := m[string(s[j])]; !ok{
				m[string(s[j])] = 1
				temp++
				if temp > finalCount{ finalCount = temp }
			}else{
				break
			}			
		}
		temp = 0
		m = make(map[string]int)
	}

    return finalCount 
}