package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "urlshortener",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS shortenedurl (shortcode VARCHAR(10) PRIMARY KEY, url LONGTEXT)")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	url := "https://example.com"
	shortenedUrl := ShortenUrl(url)
	fmt.Println(*shortenedUrl)
	_, err = db.Exec("INSERT INTO shortenedurl (shortcode, url) VALUES (?, ?)", shortenedUrl, url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	rows, err := db.Query("SELECT url FROM shortenedurl WHERE shortcode = ?", shortenedUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var urls []string
	for rows.Next() {
		var url string
		rows.Scan(&url)
		urls = append(urls, url)
	}
	fmt.Println("URLs:", urls)
}

func ShortenUrl(url string) *string {
	id := GetPsuedoRandomIntegerId()
	s, err := encodeBase36(id)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &s
}

func encodeBase36(num int64) (string, error) {
	return strconv.FormatInt(num, 36), nil
}

func decodeBase36(s string) (int64, error) {
	return strconv.ParseInt(s, 36, 64)
}

type TicketServer struct {
	ID      int64
	Start   int64
	End     int64
	Current int64
}

var ticketServer []TicketServer = []TicketServer{
	{
		ID:      1,
		Start:   1000000000,
		End:     3000000000,
		Current: 1000000000,
	},
	{
		ID:      2,
		Start:   3000000000,
		End:     6000000000,
		Current: 3000000000,
	},
	{
		ID:      3,
		Start:   6000000000,
		End:     9000000000,
		Current: 6000000000,
	},
	{
		ID:      4,
		Start:   9000000000,
		End:     12000000000,
		Current: 9000000000,
	},
}

func GetPsuedoRandomIntegerId() int64 {
	index := generateRandomNumber(len(ticketServer))
	current := ticketServer[index].Current
	ticketServer[index].Current += 1
	return current
}

func generateRandomNumber(count int) int {
	randomNum := rand.Intn(count)
	return randomNum
}
