package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

var DbConnection *sql.DB

type Co2DataOrigin struct {
	Date string
	Time string
	Data int
}

type Co2Data struct {
	Date time.Time
	Time time.Time
	Data int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./co2ex.sqlite")
	defer DbConnection.Close()

	cmd := "SELECT * FROM DATE_CO2"
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()
	var cco2DataOrigin []Co2DataOrigin
	for rows.Next() {
		var co2DataOrigin Co2DataOrigin
		err := rows.Scan(&co2DataOrigin.Date, &co2DataOrigin.Time, &co2DataOrigin.Data)
		if err != nil {
			log.Println(err)
		}
		cco2DataOrigin = append(cco2DataOrigin, co2DataOrigin)
	}

	var cc2Data []Co2Data

	const format1 = "2020/01/01"
	const format2 = "17:34:42"
	for _, c := range cco2DataOrigin {
		var co2Data Co2Data
		var err error
		co2Data.Date, err = time.Parse(format1, c.Date)
		co2Data.Time, err = time.Parse(format2, c.Time)
		co2Data.Data = c.Data
		if err != nil {
			log.Println(err)
		}
	}
	for _, c := range cc2Data {
		fmt.Println(c.Date, c.Time, c.Data)
	}

}
