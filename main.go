package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	json_parser "github.com/StefanWellhoner/csv-parser/parsers/json"
	yaml_parser "github.com/StefanWellhoner/csv-parser/parsers/yaml"
)

var headers []string
var delimiter *string

type Data struct {
	Rows []map[string]string `json:"rows"`
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

/**
* Flags
* -delim: Delimiter used in the CSV file
* -o: Output file path (optional)
* -format: Output format (json or yaml)
* -h: Show help
 */
func main() {
	delimiter = flag.String("delim", ",", "Delimiter used in the CSV file")
	outFile := flag.String("o", "", "Output file path (optional)")
	format := flag.String("format", "json", "Output format (json or yaml)")
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

	switch *format {
	case "yaml":
		yamlData, err := yaml_parser.ConvertToYAML(data.Rows)
		if err != nil {
			fmt.Println("Error converting to Yaml:", err)
			os.Exit(1)
		}

		if *outFile != "" {
			if err := yaml_parser.WriteYAMLFile(yamlData, *outFile); err != nil {
				fmt.Println("Error writing file:", err)
				os.Exit(1)
			}
		} else {
			fmt.Println(string(yamlData))
		}
	case "json":
		jsonData, err := json_parser.ConvertToJSON(data.Rows)
		if err != nil {
			fmt.Println("Error converting to JSON:", err)
			os.Exit(1)
		}

		if *outFile != "" {
			if err := json_parser.WriteJSONFile(jsonData, *outFile); err != nil {
				fmt.Println("Error writing file:", err)
				os.Exit(1)
			}
		} else {
			fmt.Println(string(jsonData))
		}
	default:
		fmt.Println("Invalid format specified. Use 'json' or 'yaml'.")
		os.Exit(1)
	}

	fmt.Println("Time taken:", time.Since(startTime))
	fmt.Println("Parsed lines:", linesParsed)
	if *outFile != "" {
		fmt.Println("Output file:", *outFile)
	}
}
