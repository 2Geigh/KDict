package main

import (
	"fmt"
	"hash"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Meetup struct {
	id                int
	time              time.Time
	place             string
	friends_attending []Friend
}

type Friend struct {
	id                          int
	name                        string
	birthday                    time.Time
	days_since_last_interaction int
	days_since_last_meetup      int
	phone_number                string // To be encrypted
	meetup_plans                []Meetup
}

type User struct {
	id                                         int
	username                                   string
	password                                   hash.Hash64 // To be encrypted
	days_since_you_last_interacted_with_anyone int
	days_since_you_last_hung_out_with_anyone   int
	meetup_plans                               []Meetup
	recievesNotifications                      (map[string]bool)
}

func root(writer http.ResponseWriter, request *http.Request) { // We pass the pointer of the HTTP request to avoid copying over a potentially large request that could slow down the server

	// HTTP requests and respones are transmitted as raw bytes, therefore all requests/respones need to be converted to bytes
	// type1(variable) returns a copy of variable converted to type type1
	writer.Write([]byte("'/'에 오신 것을 환영합니다."))

}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(`Error loading .env file: %s`, err)
	}

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_URL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db_user, db_password, db_host, db_port, db_name)

	// pool, err := pgxpool

	fmt.Println(db_host, db_port, db_user, db_password, db_name)

	http.HandleFunc("/", root)

	http.ListenAndServe(":3000", nil)

}
