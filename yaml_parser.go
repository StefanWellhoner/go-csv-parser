package yaml_parser

import (
    "gopkg.in/yaml.v2"
    "os"
)

func ConvertToYAML(data []map[string]string) ([]byte, error) {
    return yaml.Marshal(data)
}

func WriteYAMLFile(yamlData []byte, outFile string) error {
    file, err := os.Create(outFile)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.Write(yamlData)
    if err != nil {
        return err
    }

    return nil
}
