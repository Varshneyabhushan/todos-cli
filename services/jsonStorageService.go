package services 

import (
	"encoding/json"
	"os"
)

type JSONStorage struct {
	filePath string
}

func NewJSONStorage(filePath string) *JSONStorage {
	return &JSONStorage{
		filePath: filePath,
	}
}

func (js *JSONStorage) LoadJSON(target interface{}) error {
	if _, err := os.Stat(js.filePath); os.IsNotExist(err) {
		return nil
	}

	fileData, err := os.ReadFile(js.filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileData, target)
	if err != nil {
		return err
	}

	return nil
}

func (js *JSONStorage) SaveJSON(data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(js.filePath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
