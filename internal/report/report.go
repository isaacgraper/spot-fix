package report

import (
	"bytes"
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
	correctedContent := bytes.ReplaceAll(content, []byte("]["), []byte(","))
	var data []ReportData

	if err := json.Unmarshal(correctedContent, &data); err != nil {
		return fmt.Sprintf("error: %v", err)
	}

	var formattedText string
	uniqueRecords := make([]ReportData, 0)

	for _, entry := range data {
		if !ContainsReport(uniqueRecords, entry) {
			uniqueRecords = append(uniqueRecords, entry)
		}
	}

	for _, entry := range uniqueRecords {
		formattedText += fmt.Sprintf("%d - %-40s %-10s %-10s", entry.Index, entry.Name, entry.Hour, entry.Category)
	}

	return formattedText
}

func ContainsReport(data []ReportData, record ReportData) bool {
	for _, r := range data {
		if r.Index == record.Index {
			return true
		}
	}
	return false
}

func NewReport(fileName string, content []byte) *File {
	return &File{
		FileName: fileName,
		Content:  content,
	}
}

func (f *File) SaveReport() error {
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
