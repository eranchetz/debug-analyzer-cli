package main

import (
	"fmt"
	"log"
	"flag"
)

var (
	debugmode = false
	filename = ""
	) 

func init() {
		flag.BoolVar(&debugmode, "debugmode", false, "run in debug mode")
		flag.StringVar(&filename, "input", "", "debug.json file to parse")
	}

func main() {
	flag.Parse()

	fmt.Println("Starting debug file analysis...")

	an, err := NewAnalysis(filename)
	if err != nil {
		log.Fatal("Could not open debug file ", err)
	}
	an.Analyze()
	an.ToString()
}
