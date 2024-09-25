package jitapi

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// DbConfig is used to pass the values to connect to database
type DbConfig struct {
	DBType     string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSslMode  string
}

// Data is a struct used to store the data that is being fetched from the database
type Data struct {
	Id        string    `json:"id"`
	Source    string    `json:"source"`
	Info      Info      `json:"info"`
	StoreId   string    `json:"store_id"`
	CartId    string    `json:"cart_id"`
	CreatedAt time.Time `json:"created_at"`
	ErrorCode string    `json:"error_code"`
}

// Info is sub-struct
type Info struct {
	XcallId       string      `json:"xcall_id"`
	BoardRole     string      `json:"board_role"`
	RaspiIssueId  interface{} `json:"raspi_issue_id"`
	RaspiIssueMsg string      `json:"raspi_issue_msg"`
	// Source        string      `json:"source"`
	// CartId        string      `json:"cart_id"`
	// StoreId       string      `json:"store_id"`
	// ErrorCode     int         `json:"error_code"`
}

// Value: Converts the Info struct to a JSON byte slice for storage in the database.
func (i Info) Value() (driver.Value, error) {
	return json.Marshal(i)
}

// Scan: Converts the JSON byte slice from the database back into the Info struct.
func (i *Info) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}
	return json.Unmarshal(bytes, &i)
}

type Count struct {
	IrBottomCountFixed            int
	IrBottomCountON               int
	IrTopCountFixed               int
	IrTopCountON                  int
	NetworkOnlineCount            int
	NetworkOfflineCount           int
	FailedToUploadAisleImageCount int
	UploadFailedCount             int
	FailedToConnectToSlaveCount   int
	NoResponseFromSlaveCount      int
	XcallIdMismatchCount          int
	ZeroFramesCapturedCount       int
	FailedToInitializeOneCamCount int
	TopCamErrorCount              int
	BottomCamErrorCount           int
	WeightModuleStartingCount     int
	WeightModuleCompleteCount     int
	FailedToConnectToUARTCount    int
	SocketConnectionLostCount     int
	HttpsCounnectionPoolCount     int
	CsvUploadFailedCount          int
}

type Helper struct {
	FilteredData []Data
	MasterData   []Data
	SlaveData    []Data
	Count        int
	MasterCount  int
	SlaveCount   int
	Result       [][]string
}

type CsvData struct {
	StoreID       string
	CartID        string
	XcallID       string
	BoardRole     string
	RaspiIssueMsg string
	Reason        string
	Count         int
}
