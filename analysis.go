package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

/*
The output of the tool should be statistical information, including the following data:
• Total number of logs records
• Number of logs per each action (actions may be any file operation)
• The percentage of actions whose result is “success”
*/

type Analysis struct {
	fileReader       io.Reader
	TotalLogRecords  int
	NumberOfActions  map[string]int
	SucessfulRecords int
}

func NewAnalysis(fileLocation string) (*Analysis, error) {
	file, err := os.Open(fileLocation)
	if err != nil {
		return nil, err
	}

	return &Analysis{
		fileReader:      file,
		NumberOfActions: make(map[string]int),
	}, nil
}

func (a *Analysis) Analyze() *Analysis {
	dec := json.NewDecoder(a.fileReader)
	_, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	// while the array contains values
	for dec.More() {
		var dl DebugLine
		// json decode an array value (one debugline)
		err := dec.Decode(&dl)
		if err != nil {
			log.Println("failed to read json line", err)
		}
		a.TotalLogRecords++
		a.NumberOfActions[dl.Action]++
		if dl.Result == "success" {
			a.SucessfulRecords++
		}
		
		debug("output %#v", dl)
	}

	return nil
}

func (a *Analysis) ToString() {
	
	fmt.Println("Total Records:", a.TotalLogRecords)
	
	fmt.Println("Successful %:", (a.SucessfulRecords/a.TotalLogRecords)*100)

	prettyActions, err := json.MarshalIndent(a.NumberOfActions, "", "  ")
	if err != nil {
    fmt.Println("error:", err)
	}
	fmt.Println("Actions hisotogram:", string(prettyActions))
}
