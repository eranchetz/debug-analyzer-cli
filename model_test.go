package main

import (
	"strings"

	"testing"
)

const jsonStream = `[{
"level": "DEBUG",
"logger": "logger0",
"thread": "main0",
"timestamp": "2017-12-05T20:43:33.511+0000", "hostname": "host0",
"user": "root0",
"action": "Edit File",
"filename": "myfile.txt",
"filepath": "/opt/special_system0/data", "result": "success",
"appId": "special_system0"
},
{
"level": "DEBUG",
"logger": "logger0",
"thread": "main0",
"timestamp": "2017-12-05T20:45:33.511+0000", "hostname": "host1",
"user": "root1",
"action": "Create File",
"filename": "myfile1.txt",
"filepath": "/opt/special_system0/data", "result": "failure",
"appId": "special_system0"
}
]`

func TestAnalysisTotalRecords(t *testing.T) {
	jsonReader := strings.NewReader(jsonStream)
	a := &Analysis{
		fileReader:      jsonReader,
		NumberOfActions: make(map[string]int),
	}
	a.Analyze()

	if a.TotalLogRecords != 2 {
		t.Error("expected number of total items is wrong")
	}
}
