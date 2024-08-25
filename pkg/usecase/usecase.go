package usecase

import (
	"errors"
	"fmt"
	"io"
	"time"
	"vrwizards/pkg/db"
	"vrwizards/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func ParseExcel(f io.Reader) ([]models.Record, error) {
	file, err := excelize.OpenReader(f)
	if err != nil {
		return nil, errors.New("failed to open Excel file")
	}
	sheets := file.GetSheetList()
	fmt.Println("Sheets are",sheets)

	var records []models.Record
	rows, err := file.GetRows("uk-500")
	if err != nil {
		return nil, errors.New("failed to read rows from Excel file")
	}

	for i, row := range rows {
		if i == 0 {
			continue // Skip header row
		}

		if len(row) < 10 {
			return nil, errors.New("unexpected number of columns in Excel file")
		}

		record := models.Record{
			FirstName: row[0],
			LastName:  row[1],
			Company:   row[2],
			Address:   row[3],
			City:      row[4],
			County:    row[5],
			Postal:    row[6],
			Phone:     row[7],
			Email:     row[8],
			Web:       row[9],
		}

		records = append(records, record)
	}

	return records, nil
}

func CacheRecords(records []models.Record) error {
	err := db.Redis.Set(&gin.Context{}, "records", records, 5*time.Minute).Err()
	return err
}
