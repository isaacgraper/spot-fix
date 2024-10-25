package report

import (
	"encoding/json"
	"fmt"
	"os"
)

type File struct {
	FileName string
	Content  []byte
}

type ReportData struct {
	Index    int    `results:"index"`
	Name     string `results:"name"`
	Hour     string `results:"hour"`
	Category string `results:"category"`
}

func FormatReport(content []byte) string {
	var data []ReportData

	if err := json.Unmarshal(content, &data); err != nil {
		return ""
	}

	var formattedText string
	for _, entry := range data {
		formattedText += fmt.Sprintf("%-30s %-30s %-20s", entry.Name, entry.Hour, entry.Category)
	}

	return formattedText
}

func NewFile(fileName string, content []byte) *File {
	return &File{
		FileName: fileName,
		Content:  content,
	}
}

func (f *File) SaveFile() error {
	file, err := os.OpenFile(f.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file for appending: %v", err)
	}
	defer file.Close()

	_, err = file.Write(f.Content)
	if err != nil {
		return fmt.Errorf("error inserting new content: %v", err)
	}
	return nil
}
