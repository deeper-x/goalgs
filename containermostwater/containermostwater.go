package containermostwater

// Solution to container problem
func Solution(height []int) int {
	maxArea := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			h := min(height[i], height[j])
			b := j - i

			currArea := h * b
			if currArea > maxArea {
				maxArea = currArea
			}
		}
	}

	return maxArea
}
