package main

import (
	"co2-sensor-web/app/controllers"
	"database/sql"
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
	Time time.Time
	Data int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./co2ex.sqlite")
	defer DbConnection.Close()

	cmd := "SELECT * FROM DATE_CO2"
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()
	var co2DataOriginStructArr []Co2DataOrigin
	for rows.Next() {
		var co2DataOrigin Co2DataOrigin
		err := rows.Scan(&co2DataOrigin.Date, &co2DataOrigin.Time, &co2DataOrigin.Data)
		if err != nil {
			log.Println(err)
		}
		co2DataOriginStructArr = append(co2DataOriginStructArr, co2DataOrigin)
	}

	var co2DataStructArr []Co2Data
	for _, c := range co2DataOriginStructArr {
		var co2DataStruct Co2Data
		str1 := c.Date + " " + c.Time
		time1, err := time.Parse("2006/01/02 15:04:05", str1)
		if err != nil {
			log.Println(err)
		}
		co2DataStruct.Time = time1
		co2DataStruct.Data = c.Data
		co2DataStructArr = append(co2DataStructArr, co2DataStruct)
	}

	//for _, c := range co2DataStructArr {
	//	fmt.Println(c.Time, c.Data)
	//}
	controllers.StartWebServer()
}
