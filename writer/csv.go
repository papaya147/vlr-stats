package writer

import (
	"bytes"
	"encoding/csv"
	"os"
)

func AppendToCSV(data Serialisable, file *os.File) error {
	var csvBuffer bytes.Buffer
	writer := csv.NewWriter(&csvBuffer)
	defer writer.Flush()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	if fileInfo.Size() == 0 {
		if err := writer.Write(data.GetHeaders()); err != nil {
			return err
		}
	}

	records := data.GetRecords()
	for _, record := range records {
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	_, err = file.Write(csvBuffer.Bytes())
	return err
}
