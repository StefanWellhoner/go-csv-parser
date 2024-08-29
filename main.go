package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var headers []string
var delimiter *string

type Data struct {
    Rows    []map[string]string `json:"rows"`
}

func ParseHeaders(line string) {
    headers = strings.Split(line, *delimiter)
}

func BuildMap(line string) map[string]string {
    fields := strings.Split(line, *delimiter)
    m := make(map[string]string)
    for i, field := range fields {
        m[headers[i]] = field
    }
    return m
}

func WriteFile(data []byte, path string) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.Write(data)
    if err != nil {
        return err
    }

    return nil
}

/**
* Flags
* -delim: Delimiter used in the CSV file
* -o: Output file path (optional)
* -h: Show help
*/
func main() {
    delimiter = flag.String("delim", ",", "Delimiter used in the CSV file")
    outFile := flag.String("o", "", "Output file path (optional)")
    help := flag.Bool("h", false, "Show help")

    flag.Parse()

    // Show help
    if *help {
		flag.Usage()
		return
	}
    
    // Check if the CSV file is provided
    if len(flag.Args()) < 1 {
        fmt.Println("Please provide the CSV file name")
        flag.Usage()
        os.Exit(0)
    }
     
    // Open the file
    filePath := flag.Arg(0)

    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error opening file:", err)
        os.Exit(1)
    }
    // Close the file when the function returns
    defer file.Close()
    
    fmt.Printf("Parsing file: %s\nUsing delimiter %s\n", filePath, *delimiter)

    scanner := bufio.NewScanner(file)
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
        os.Exit(1)
    }

    startTime := time.Now()

    linesParsed := 0
    data := Data{}
    for scanner.Scan() {
        if linesParsed == 0 {
            ParseHeaders(scanner.Text())
            linesParsed++
            continue
        }
        data.Rows = append(data.Rows, BuildMap(scanner.Text()))
        linesParsed++
    }

    json, err := json.MarshalIndent(data.Rows, "", "\t")
    if err != nil {
        fmt.Println("Error marshalling data:", err)
        os.Exit(1)
    }
    
    if *outFile != "" {
        if err := WriteFile(json, *outFile); err != nil {
            fmt.Println("Error writing file:", err)
            os.Exit(1)
        }
    } else {
        fmt.Println(string(json))
    }

    fmt.Println("Time taken:", time.Since(startTime))
    fmt.Println("Parsed lines:", linesParsed)
}
