package printer

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type CSVPrinter struct {
	filePath string
	file     *os.File
	writer   *csv.Writer
}

func NewCSVPrinter() (*CSVPrinter, error) {
	currentTime := time.Now().Format("20060102")
	dirPath := fmt.Sprintf("../../data/%s", currentTime)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return nil, err
	}
	filePath := fmt.Sprintf("%s/test6.csv", dirPath)
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	writer := csv.NewWriter(file)
	return &CSVPrinter{
		filePath: filePath,
		file:     file,
		writer:   writer,
	}, nil
}

func (p *CSVPrinter) Print(fitness float64, fitnessCount uint) {
	record := []string{
		fmt.Sprintf("%.3f", fitness),
		fmt.Sprintf("%d", fitnessCount),
	}
	if err := p.writer.Write(record); err != nil {
		fmt.Printf("Error writing record to csv: %v\n", err)
	}
	p.writer.Flush()
}

func (p *CSVPrinter) Close() {
	p.writer.Flush()
	p.file.Close()
}
