package main 

import (
	"fmt"
)

func debug(msg ...interface{}) {
	if debugmode {
		fmt.Println("DEBUG:", msg)
	} 
}
