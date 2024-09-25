package jitapi

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func (h *Helper) customErrorData(data []Data, issueMsg string, reason string) {

	countMap := make(map[string]int)
	dataMap := make(map[string]CsvData)

	for _, item := range data {
		if strings.Contains(item.Info.RaspiIssueMsg, issueMsg) {
			key := item.Info.BoardRole
			countMap[key]++
			dataMap[key] = CsvData{
				RaspiIssueMsg: issueMsg,
				Reason:        reason,
				BoardRole:     item.Info.BoardRole,
			}
		}
	}

	// Write the filtered data to the CSV file
	for key, count := range countMap {
		item := dataMap[key]
		row := []string{item.RaspiIssueMsg, reason, item.BoardRole, fmt.Sprintf("%d", count)}
		h.Result = append(h.Result, row)
	}
}

func (h *Helper) WriteCustomError(data []Data) {
	h.Result = append(h.Result, []string{"Message", "Reason", "Board Role", "Count"})
	h.customErrorData(data, "Exception when starting client session", "Network Issue")
	h.customErrorData(data, "Session is inactive.", "Socket Connection Issue")
	h.customErrorData(data, "read or base64 encode failed", "Checking...")
	h.customErrorData(data, "video upload failed!", "Checking...")
	h.Result = append(h.Result, []string{})
}

func (h *Helper) VideoSavedFilter(data []Data) {
	h.MasterCount = 0
	h.SlaveCount = 0

	for _, item := range data {
		if item.Source == "video_info" {
			if item.Info.BoardRole == "master" {
				// h.MasterData = append(h.MasterData, item)
				h.MasterCount++
			} else {
				// h.SlaveData = append(h.SlaveData, item)
				h.SlaveCount++
			}
		}
	}

	// var result [][]string
	h.Result = append(h.Result, returnData(1, "Video saved count", "Master", h.MasterCount))
	h.Result = append(h.Result, returnData(2, "Video saved count", "Slave", h.SlaveCount))
	h.Result = append(h.Result, []string{})

}

func (h *Helper) ZeroFramesFilter(data []Data) {
	h.MasterCount = 0
	h.SlaveCount = 0

	for _, item := range data {
		if item.Info.BoardRole == "master" {
			if strings.Contains(item.Info.RaspiIssueMsg, "Error! Captured 0 frames") {
				if item.Info.BoardRole == "master" {
					h.MasterCount++
				} else {
					h.SlaveCount++
				}
			}
		}
	}
	h.Result = append(h.Result, returnData(3, "Zero frames captured count", "master", h.MasterCount))
	h.Result = append(h.Result, returnData(4, "Zero frames captured count", "slave", h.SlaveCount))
	h.Result = append(h.Result, []string{})

}

func (h *Helper) GetMissedVIDs(data []Data) {
	h.Result = append(h.Result, []string{"xcall_id", "board_role", "missed_vid", "count"})

	// Map to store VIDs by xcall_id and board_role
	vidMap := make(map[string]map[string][]int)

	// Populate the vidMap with data from the JSON
	for _, entry := range data {
		xcallID := entry.Info.XcallId
		boardRole := entry.Info.BoardRole
		vid := extractVID(entry.Info.RaspiIssueMsg)

		if vid == -1 {
			continue
		}

		if _, ok := vidMap[xcallID]; !ok {
			vidMap[xcallID] = map[string][]int{
				"master": {},
				"slave":  {},
			}
		}

		vidMap[xcallID][boardRole] = append(vidMap[xcallID][boardRole], vid)
	}

	// Compare VIDs and identify discrepancies
	for xcallID, roles := range vidMap {
		masterVIDs := roles["master"]
		slaveVIDs := roles["slave"]

		sort.Ints(masterVIDs)
		sort.Ints(slaveVIDs)

		// Find VIDs in slave but not in master
		missedInMaster := difference(slaveVIDs, masterVIDs)
		// Find VIDs in master but not in slave
		missedInSlave := difference(masterVIDs, slaveVIDs)

		// Write missed VIDs for master
		if len(missedInMaster) > 0 {
			h.Result = append(h.Result, []string{xcallID, "master", fmt.Sprint(missedInMaster), strconv.Itoa(len(missedInMaster))})
			// writer.Write([]string{xcallID, "master", fmt.Sprint(missedInMaster), strconv.Itoa(len(missedInMaster))})
		}

		// Write missed VIDs for slave
		if len(missedInSlave) > 0 {
			h.Result = append(h.Result, []string{xcallID, "slave", fmt.Sprint(missedInSlave), strconv.Itoa(len(missedInSlave))})
			// writer.Write([]string{xcallID, "slave", fmt.Sprint(missedInSlave), strconv.Itoa(len(missedInSlave))})
		}
	}
}
