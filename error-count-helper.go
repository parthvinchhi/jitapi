package jitapi

import (
	"fmt"
	"strings"
)

func subFuncIR(str, substr, substr2 string) int {
	if strings.Contains(str, substr) && strings.Contains(str, substr2) {
		return 1
	}
	return 0
}

func subFunc(data Data, substr string) int {
	if strings.Contains(data.Info.RaspiIssueMsg, substr) {
		return 1
	}
	return 0
}

func returnData(no int, msg1, msg2 string, count int) []string {
	result := make([]string, 0)
	result = append(result, fmt.Sprintf("%d", no), msg1, msg2, fmt.Sprintf("%d", count))
	return result
}

// func writeCountToCsv(data [][]string, filename string) error {
// 	csvFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Fatal("error opening file", err)
// 	}
// 	defer csvFile.Close()

// 	w := csv.NewWriter(csvFile)
// 	defer w.Flush()

// 	for _, item := range data {
// 		err := w.Write(item)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
