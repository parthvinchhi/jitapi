package jitapi

import "strings"

type Methods Helper

func (m *Methods) Source(data []Data, source string) ([]Data, int) {
	for _, item := range data {
		if item.Source == source {
			m.FilteredData = append(m.FilteredData, item)
			m.Count++
		}
	}
	return m.FilteredData, m.Count
}

func (m *Methods) StoreId(data []Data, storeId string) ([]Data, int) {
	for _, item := range data {
		if item.StoreId == storeId {
			m.FilteredData = append(m.FilteredData, item)
			m.Count++
		}
	}

	return m.FilteredData, m.Count
}

func (m *Methods) StoreIdwithCartId(data []Data, storeId, cartId string) ([]Data, int) {
	for _, item := range data {
		if item.StoreId == storeId && item.CartId == cartId {
			m.FilteredData = append(m.FilteredData, item)
			m.Count++
		}
	}

	return m.FilteredData, m.Count
}

func (m *Methods) XcallId(data []Data, xcallId string) ([]Data, int) {
	for _, item := range data {
		if item.Info.XcallId == xcallId {
			m.FilteredData = append(m.FilteredData, item)
			m.Count++
		}
	}

	return m.FilteredData, m.Count
}

func (m *Methods) ZeroFrames(data []Data) ([]Data, []Data, int, int) {
	for _, item := range data {
		if strings.Contains(item.Info.RaspiIssueMsg, "Error! Captured 0 frames") {
			if item.Info.BoardRole == "master" {
				m.MasterData = append(m.MasterData, item)
				m.MasterCount++
			} else if item.Info.BoardRole == "slave" {
				m.SlaveData = append(m.SlaveData, item)
				m.SlaveCount++
			}
		}
	}

	return m.MasterData, m.SlaveData, m.MasterCount, m.SlaveCount
}

func (m *Methods) ZeroFramesWithStoreId(data []Data, storeId string) ([]Data, []Data, int, int) {
	for _, item := range data {
		if strings.Contains(item.Info.RaspiIssueMsg, "Error! Captured 0 frames") && item.StoreId == storeId {
			if item.Info.BoardRole == "master" {
				m.MasterData = append(m.MasterData, item)
				m.MasterCount++
			} else if item.Info.BoardRole == "slave" {
				m.SlaveData = append(m.SlaveData, item)
				m.SlaveCount++
			}
		}
	}

	return m.MasterData, m.SlaveData, m.MasterCount, m.SlaveCount
}

func (m *Methods) VideoSaved(data []Data) ([]Data, []Data, int, int) {
	for _, item := range data {
		if item.Source == "video_info" {
			if item.Info.BoardRole == "master" {
				m.MasterData = append(m.MasterData, item)
				m.MasterCount++
			} else if item.Source == "slave" {
				m.SlaveData = append(m.SlaveData, item)
				m.SlaveCount++
			}
		}
	}

	return m.MasterData, m.SlaveData, m.MasterCount, m.SlaveCount
}
