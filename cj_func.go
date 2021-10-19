package cj_func

import "time"

func Change_date(add_day int) string {
	dt := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	//fmt.Println(dt.GoString())
	dt2 := dt.AddDate(0, 0, add_day)
	dt_date := dt2.Format("2006-01-02")
	//fmt.Println("MM-DD-YYYY : ", dt.Format("2006-01-02"))
	//fmt.Println("MM-DD-YYYY : ", dt_date)
	return dt_date
}
