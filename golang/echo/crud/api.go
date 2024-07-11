package main

import (
	"encoding/csv"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Record struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func readCSV() ([]Record, error) {
	file, err := os.Open("data.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var result []Record
	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		result = append(result, Record{ID: id, Name: record[1]})
	}

	return result, nil
}

func writeCSV(records []Record) error {
	file, err := os.Create("data.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		err := writer.Write([]string{strconv.Itoa(record.ID), record.Name})
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	e := echo.New()

	e.POST("/create", func(c echo.Context) error {
		var newRecord Record
		if err := c.Bind(&newRecord); err != nil {
			return err
		}

		records, err := readCSV()
		if err != nil {
			return err
		}

		newRecord.ID = len(records) + 1
		records = append(records, newRecord)
		writeCSV(records)

		return c.JSON(http.StatusCreated, newRecord)
	})

	e.GET("/read/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		records, err := readCSV()
		if err != nil {
			return err
		}

		for _, record := range records {
			if record.ID == id {
				return c.JSON(http.StatusOK, record)
			}
		}

		return c.JSON(http.StatusNotFound, echo.Map{"message": "Record not found"})
	})

	e.PUT("/update/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var updatedRecord Record
		if err := c.Bind(&updatedRecord); err != nil {
			return err
		}

		records, err := readCSV()
		if err != nil {
			return err
		}

		for i, record := range records {
			if record.ID == id {
				records[i].Name = updatedRecord.Name
				writeCSV(records)
				return c.JSON(http.StatusOK, records[i])
			}
		}

		return c.JSON(http.StatusNotFound, echo.Map{"message": "Record not found"})
	})

	e.DELETE("/delete/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		records, err := readCSV()
		if err != nil {
			return err
		}

		for i, record := range records {
			if record.ID == id {
				records = append(records[:i], records[i+1:]...)
				writeCSV(records)
				return c.JSON(http.StatusOK, echo.Map{"message": "Record deleted"})
			}
		}

		return c.JSON(http.StatusNotFound, echo.Map{"message": "Record not found"})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
