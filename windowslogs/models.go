package windowslogs

type Logs struct {
	SecurityLogs    []Log
	ApplicationLogs []Log
}

type Log struct {
	ID                int    `json:"Id"`
	Version           int    `json:"Version"`
	Qualifiers        any    `json:"Qualifiers"`
	Level             int    `json:"Level"`
	Task              int    `json:"Task"`
	Opcode            int    `json:"Opcode"`
	Keywords          int64  `json:"Keywords"`
	RecordID          int    `json:"RecordId"`
	ProviderName      string `json:"ProviderName"`
	ProviderID        string `json:"ProviderId"`
	LogName           string `json:"LogName"`
	ProcessID         int    `json:"ProcessId"`
	ThreadID          int    `json:"ThreadId"`
	MachineName       string `json:"MachineName"`
	UserID            any    `json:"UserId"`
	TimeCreated       string `json:"TimeCreated"`
	ActivityID        string `json:"ActivityId"`
	RelatedActivityID any    `json:"RelatedActivityId"`
	ContainerLog      string `json:"ContainerLog"`
	MatchedQueryIds   []any  `json:"MatchedQueryIds"`
	Bookmark          struct {
	} `json:"Bookmark"`
	LevelDisplayName     string   `json:"LevelDisplayName"`
	OpcodeDisplayName    string   `json:"OpcodeDisplayName"`
	TaskDisplayName      string   `json:"TaskDisplayName"`
	KeywordsDisplayNames []string `json:"KeywordsDisplayNames"`
	Properties           []string `json:"Properties"`
	Message              string   `json:"Message"`
}
