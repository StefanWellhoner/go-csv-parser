# CSV Parser in GO
This is a simple CSV parser written in GO. It reads a CSV file and prints the content to the console or to a text file in JSON or Yaml format.

## How to run
1. Clone the repository
2. Run the following command:
```bash
go run main.go
```
3. The content of the CSV file will be printed to the console.

## Flag for the cli application
- `-o` or `--output` : Output the content of the CSV file to a file in JSON or Yaml format. Example:
```bash
go run main.go -o output.json
```
- `-h` or `--help` : Display the help message. Example:
```bash
go run main.go -h
```
- `-delim`: Specify the delimiter of the CSV file. Example:
```bash
go run main.go -delim ; -o output.json
```
- `-format`: Specify the output format of the CSV file. Use `json` or `yaml`. Example:
```bash
go run main.go -format yaml -o output.yaml
```
