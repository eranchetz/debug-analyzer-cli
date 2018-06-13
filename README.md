# Debug Log Analysis CLI 

## Introduction

> This toll will parse a debug (see template below) and returns relevant statistics such as : 
 - Total Number of Debug Messages
 - Number of Successful events
 - Histogram of different events  

## Code Samples


```bash
debug-analyzer-cli -input debug.log 
```

Output example: 
```bash
Total Records: 1
Successful %: 100
Actions hisotogram {
  "Edit File": 1
}
```

## Installation

> You can fetch the binary from the [release page](release page)

> From source:

```
go build 
```
