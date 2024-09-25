package jitapi

import "fmt"

func GetQuery() (string, string) {
	var queryDate string
	fmt.Print("Enter date in YYYY-MM-DD format : ")
	fmt.Scanln(&queryDate)

	query := "select * from HARDWARE_ISSUES_INFOS WHERE created_at > '" + queryDate + " 03:10:18' ORDER BY created_at DESC;"

	return query, queryDate
}
