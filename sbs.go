package jitapi

func SideBySideData(leftData, rightData [][]string) [][]string {
	maxLength := len(leftData)
	if len(rightData) > maxLength {
		maxLength = len(rightData)
	}

	combinedData := make([][]string, maxLength)

	for i := 0; i < maxLength; i++ {
		var row []string

		if i < len(leftData) {
			row = append(row, leftData[i]...)
		} else {
			row = append(row, make([]string, len(leftData[0]))...)
		}

		row = append(row, "", "")

		if i < len(rightData) {
			row = append(row, rightData[i]...)
		}

		combinedData[i] = row
	}

	return combinedData
}
