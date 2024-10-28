package report

import (
	"fmt"
	"os"
	"strings"
)

type ReportData struct {
	Index    int    `results:"index"`
	Name     string `results:"name"`
	Hour     string `results:"hour"`
	Category string `results:"category"`
}

type File struct {
	FileName string
	Content  []ReportData
}

func NewReport(fileName string, content []ReportData) *File {
	return &File{
		FileName: fileName,
		Content:  content,
	}
}

func Contains(data []ReportData, record ReportData) bool {
	for _, r := range data {
		if r.Index == record.Index {
			return true
		}
	}
	return false
}

func (f *File) Format(data []ReportData) []byte {
	var builder strings.Builder

	for _, element := range data {
		if Contains(data, element) {
			builder.WriteString(fmt.Sprintf("Relatório de Inconsistências:\n\n%d - %-40s %-10s %-30s\n", element.Index, element.Name, element.Hour, element.Category))
		} else {
			continue
		}
	}

	return []byte(builder.String())
}

func (f *File) SaveReport() error {
	file, err := os.OpenFile(f.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file for appending: %v", err)
	}
	defer file.Close()

	_, err = file.Write(f.Format(f.Content))
	if err != nil {
		return fmt.Errorf("error inserting new content: %v", err)
	}
	return nil
}
