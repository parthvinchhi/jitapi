package jitapi

type Variables Count

func (v *Variables) CountDataByString(data []Data) {
	for _, val := range data {
		v.IrBottomCountFixed += subFuncIR(val.Info.RaspiIssueMsg, "Bottom IR sensor", "is fixed")
		v.IrBottomCountON += subFuncIR(val.Info.RaspiIssueMsg, "Bottom IR sensor", "is in ON state for too long")
		v.IrTopCountFixed += subFuncIR(val.Info.RaspiIssueMsg, "Top IR sensor", "is fixed")
		v.IrTopCountON += subFuncIR(val.Info.RaspiIssueMsg, "Top IR sensor", "is in ON state for too long")
		v.NetworkOnlineCount += subFunc(val, "Network connection is online")
		v.NetworkOfflineCount += subFunc(val, "Network connection is offline")
		v.FailedToUploadAisleImageCount += subFunc(val, "Failed to upload aisle image")
		v.UploadFailedCount += subFunc(val, "upload failed with status -1: Skipping uploading videos, PANIC is triggered, shopping is done")
		v.FailedToConnectToSlaveCount += subFunc(val, "Failed to connect to slave device!")
		v.HttpsCounnectionPoolCount += subFunc(val, "catch exception: HTTPSConnectionPool")
		v.CsvUploadFailedCount += subFunc(val, "Failed to upload csv file, catch exception")
		v.NoResponseFromSlaveCount += subFunc(val, "There's been no response from slave")
		v.ZeroFramesCapturedCount += subFunc(val, "Error! Captured 0 frames")
		v.FailedToInitializeOneCamCount += subFunc(val, "STOP. Failed to initialize 1 cameras")
		v.TopCamErrorCount += subFunc(val, "Device not found for `top` camera")
		v.BottomCamErrorCount += subFunc(val, "Device not found for `bottom` camera")
		v.WeightModuleStartingCount += subFunc(val, "Reboot of WeightModule: Starting...")
		v.WeightModuleCompleteCount += subFunc(val, "Reboot of WeightModule: Startup is complete")
		v.FailedToConnectToUARTCount += subFunc(val, "Application failed to connect to scale's UART device")
		v.SocketConnectionLostCount += subFunc(val, "Socket connection lost,reconnecting..")
		v.XcallIdMismatchCount += subFunc(val, "Xcall ID mismatch")
	}
}

func (v *Variables) WriteCountToCsv() [][]string {
	var tempData [][]string

	tempData = append(tempData, returnData(1, "Network", "Network connection is online", v.NetworkOnlineCount))
	tempData = append(tempData, returnData(2, "Network", "Network connection is offline", v.NetworkOfflineCount))
	tempData = append(tempData, returnData(3, "IR", "Bottom IR sensor is fixed", v.IrBottomCountFixed))
	tempData = append(tempData, returnData(4, "IR", "Bottom IR sensor is in ON state for too long", v.IrBottomCountON))
	tempData = append(tempData, returnData(5, "IR", "Top IR sensor is fixed", v.IrTopCountFixed))
	tempData = append(tempData, returnData(6, "IR", "Top IR sensor is in ON state for too long", v.IrTopCountON))
	tempData = append(tempData, returnData(7, "Video", "Failed to upload aisle image", v.FailedToUploadAisleImageCount))
	tempData = append(tempData, returnData(8, "Video", "Upload failed with status -1: Skipping uploading videos, PANIC is triggered, shopping is done", v.UploadFailedCount))
	tempData = append(tempData, returnData(9, "Broadcast", "Failed to connect to slave device!", v.FailedToConnectToSlaveCount))
	tempData = append(tempData, returnData(10, "Broadcast", "There's been no response from slave", v.NoResponseFromSlaveCount))
	tempData = append(tempData, returnData(11, "UPLOAD_CSV", "catch exception: HTTPSConnectionPool", v.HttpsCounnectionPoolCount))
	tempData = append(tempData, returnData(12, "UPLOAD_CSV", "Failed to upload csv file, catch exception", v.UploadFailedCount))
	tempData = append(tempData, returnData(13, "CAMERAS", "Error! Captured 0 frames", v.ZeroFramesCapturedCount))
	tempData = append(tempData, returnData(14, "CAMERAS", "STOP. Failed to initialize 1 cameras", v.FailedToInitializeOneCamCount))
	tempData = append(tempData, returnData(15, "CAMERAS", "Device not found for `top` camera", v.TopCamErrorCount))
	tempData = append(tempData, returnData(16, "CAMERAS", "Device not found for `bottom` camera", v.BottomCamErrorCount))
	tempData = append(tempData, returnData(17, "Arduino", "Reboot of WeightModule: Starting...", v.WeightModuleStartingCount))
	tempData = append(tempData, returnData(18, "Arduino", "Reboot of WeightModule: Startup is complete", v.WeightModuleCompleteCount))
	tempData = append(tempData, returnData(19, "Arduino", "Application failed to connect to scale's UART device", v.FailedToConnectToUARTCount))
	tempData = append(tempData, returnData(20, "Socket", "Socket connection lost,reconnecting..", v.SocketConnectionLostCount))
	tempData = append(tempData, returnData(21, "Cart", "Xcall ID mismatch", v.XcallIdMismatchCount))

	// writeCountToCsv(tempData, filename)

	return tempData
}
