package tools

type Machine struct {
	Id             int    `json"-"`
	Name           string `json"name"`
	CpuCount       int    `json"cpuCount"`
	TotalDiskSpace int    `json"“totalDiskSpace”"`
}

type errorObject struct {
	Message string `json:"message"`
}
