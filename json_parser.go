package json_parser

import (
    "encoding/json"
    "os"
)

func ConvertToJSON(data []map[string]string) ([]byte, error) {
    return json.MarshalIndent(data, "", "    ")
}

func WriteJSONFile(jsonData []byte, outFile string) error {
    file, err := os.Create(outFile)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.Write(jsonData)
    if err != nil {
        return err
    }

    return nil
}
