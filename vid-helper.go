package jitapi

import (
	"regexp"
	"strconv"
)

func extractVID(msg string) int {
	re := regexp.MustCompile(`VID=(\d+)`)
	match := re.FindStringSubmatch(msg)
	if len(match) > 1 {
		vid, _ := strconv.Atoi(match[1])
		return vid
	}
	return -1
}

func difference(slice1, slice2 []int) []int {
	diff := []int{}
	m := map[int]bool{}

	for _, v := range slice2 {
		m[v] = true
	}

	for _, v := range slice1 {
		if !m[v] {
			diff = append(diff, v)
		}
	}

	return diff
}
