package urlshortener

import (
	"database/sql"
	"fmt"
)

func ShortenUrl(url string, db *sql.DB) *string {
	id := GetPsuedoRandomIntegerId()
	s, err := EncodeBase36(id)
	if err != nil {
		return nil
	}
	err = insertDataToDB(db, s, url)
	if err != nil {
		return nil
	}
	return &s
}

func insertDataToDB(db *sql.DB, shortcode string, url string) error {
	_, err := db.Exec("INSERT INTO shortenedurl (shortcode, url) VALUES (?, ?)", shortcode, url)
	if err != nil {
		return err
	}
	return nil
}

func GetOriginalUrl(shortcode string, db *sql.DB) (*string, error) {
	rows, err := db.Query("SELECT url FROM shortenedurl WHERE shortcode = ?", shortcode)
	if err != nil {
		return nil, err
	}
	var urls []string
	for rows.Next() {
		var url string
		rows.Scan(&url)
		urls = append(urls, url)
	}
	if len(urls) == 0 {
		return nil, fmt.Errorf("no URL found for shortcode: %s", shortcode)
	}
	return &urls[0], nil
}
