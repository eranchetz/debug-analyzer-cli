package main

type DebugLine struct {
	Action    string `json:"action"`
	AppID     string `json:"appId"`
	Filename  string `json:"filename"`
	Filepath  string `json:"filepath"`
	Hostname  string `json:"hostname"`
	Level     string `json:"level"`
	Logger    string `json:"logger"`
	Result    string `json:"result"`
	Thread    string `json:"thread"`
	Timestamp string `json:"timestamp"`
	User      string `json:"user"`
}
