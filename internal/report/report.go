package report

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ReportData struct {
	Index    int    `results:"index"`
	Name     string `results:"name"`
	Hour     string `results:"hour"`
	Category string `results:"category"`
}

type File struct {
	Content []ReportData
}

func NewReport(content []ReportData) *File {
	return &File{
		Content: content,
	}
}

func Contains(seen map[int]bool, index int) bool {
	return seen[index]
}

func (f *File) Format(data []ReportData) []byte {
	var builder strings.Builder

	seen := make(map[int]bool)

	for _, element := range data {
		if !Contains(seen, element.Index) {
			seen[element.Index] = true
			builder.WriteString(fmt.Sprintf("%-1d - %s - %s - %s\n", element.Index, element.Name, element.Hour, element.Category))
		}
	}

	return []byte(builder.String())
}

func (f *File) SaveReport() {
	fileName := fmt.Sprintf("relatório-inconsistências-%s.txt", time.Now().Format("02-01-2006"))
	filePath := filepath.Join("Z:\\", "RobôCOP", "Relatórios", fileName)

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("error creating file: %v\n", err)
		return
	}

	defer file.Close()

	_, err = file.Write(f.Format(f.Content))
	if err != nil {
		fmt.Printf("error writing in file: %v\n", err)
		return
	}
}
