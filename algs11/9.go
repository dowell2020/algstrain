package algs11

func MaxArea(height []int) int {
	area := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			if height[i] > height[j] {
				if area < (j-i)*height[j] {
					area = (j - i) * height[j]
				}
			} else {
				if area < (j-i)*height[i] {
					area = (j - i) * height[i]
				}
			}
		}
	}
	return area
}
