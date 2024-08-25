package repository

import (
	"vrwizards/pkg/db"
	"vrwizards/pkg/models"
)

func InsertRecord(record models.Record) error {
    result := db.DB.Create(&record)
     return result.Error
}

func GetRecords() ([]models.Record, error) {
   

    var records []models.Record
    result := db.DB.Find(&records)
    if result.Error != nil {
        return nil, result.Error
    }

    return records, nil
}

func UpdateRecord(id int, record models.Record) error {
    result := db.DB.Model(&models.Record{}).Where("id = ?", id).Updates(record)
    return result.Error
}