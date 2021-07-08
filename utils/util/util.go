package util

import (
	"fmt"
	"time"
)

func GetCurrDate() time.Time {

	return time.Now()
}

func GetCurrFormatDate() time.Time {

	trxDate := time.Now().Format("2006-01-02 00:00:00")
	date, err := time.Parse("2006-01-02 00:00:00", trxDate)
	fmt.Println("date ==>", date)
	if err != nil {
		fmt.Println("error ", err)
	}
	return date
}
