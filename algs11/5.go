package algs11

func Convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	newstring := ""
	n := len(s)
	cycleLen := 2*numRows - 2
	i := 0
	for i < numRows {
		for j := 0; j+i < n; j += cycleLen {
			newstring = newstring + string(s[j+i])
			if i != 0 && i != numRows-1 && j+cycleLen-i < n {
				newstring = newstring + string(s[j+cycleLen-i])
			}
		}
		i++
	}
	return newstring
}
