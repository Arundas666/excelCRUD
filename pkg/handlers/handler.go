package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
	"vrwizards/pkg/db"
	"vrwizards/pkg/models"
	"vrwizards/pkg/repository"
	"vrwizards/pkg/usecase"

	"github.com/gin-gonic/gin"
)
func ImportData(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
        return
    }

    f, err := file.Open()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not open file"})
        return
    }
    defer f.Close()

    records, err := usecase.ParseExcel(f)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    const maxConcurrent = 10
    sem := make(chan struct{}, maxConcurrent)
    var wg sync.WaitGroup
    var mu sync.Mutex
    var errors []string

    tx := db.DB.Begin()
    if tx.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
        return
    }

    for _, record := range records {
        wg.Add(1)
        sem <- struct{}{}
        go func(record models.Record) {
            defer wg.Done()
            defer func() { <-sem }()

            if err := tx.Create(&record).Error; err != nil {
                mu.Lock()
                errors = append(errors, fmt.Sprintf("Failed to insert record: %v", err))
                mu.Unlock()
            }
        }(record)
    }

    wg.Wait()

    if len(errors) > 0 {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{"errors": errors})
    } else {
        tx.Commit()
        c.JSON(http.StatusOK, gin.H{"message": "Data imported successfully"})
    }
}


func GetData(c *gin.Context) {
	records, err := db.Redis.Get(c, "records").Result()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"data": records})
		return
	}

	recordsDB, err := repository.GetRecords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve records"})
		return
	}

	db.Redis.Set(c, "records", recordsDB, 5*time.Minute)
	c.JSON(http.StatusOK, gin.H{"data": recordsDB})

}

func UpdateData(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var record models.Record
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := repository.UpdateRecord(id, record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update record"})
		return
	}

	db.Redis.Del(c, "records")

	c.JSON(http.StatusOK, gin.H{"message": "Record updated successfully"})
}

func DeleteData(c *gin.Context) {

}
