func maxArea(height []int) int {

	maxArea := 0
	temp := 0

	start := 0
	end := len(height)-1

	for start < end {
		if height[start] <= height[end] {
			temp = height[start] * (end - start)
			if temp > maxArea {
				maxArea = temp
			}
			start++
		} else {
			temp = height[end] * (end - start)
			if temp > maxArea {
				maxArea = temp
			}
			end--
		}

	}

	return maxArea
}