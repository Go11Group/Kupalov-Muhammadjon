package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5432/test?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	minTime := 12345.00
	maxTime := 0.00000
	avgTime := 0.00000
	for i := 0; i < 1000; i++{
		rows, err := db.Query(`explain (analyse)
		select * 
		from Users 
		where id = 'e874176f-0fc1-4e4f-9751-bf21c78034f0' and 
		first_name  = 'Luis' and last_name='Boyle' and
		followers = 222780 and is_employed = true;`)
		if err != nil {
			panic(err)
		}
		var desc string
		for rows.Next(){
			err = rows.Scan(&desc)
			if err != nil {
				panic(err)
			}
		}
		time := desc[len(desc)-8:len(desc)-3]
		ft, _ := strconv.ParseFloat(time, 64)
		minTime = min(minTime, ft)
		maxTime = max(maxTime, ft)
		avgTime += ft
	}
	fmt.Println(minTime)
	fmt.Println(maxTime)
	fmt.Println(avgTime/1000.00)

}

func max(a, b float64) float64{
	if a > b{
		return a
	}
	return b
}
func min(a, b float64) float64{
	if a > b{
		return b
	}
	return a
}