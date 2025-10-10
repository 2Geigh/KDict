package main

import (
	"fmt"
	"hash"
	"log"
	"os"
	"time"

	"net/http"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "github.com/gin-gonic/gin"
)

// A repository (pattern) is a class that encapsulates the logic needed to access data sources
// i.e. it abstracts database operations
// this allows for greater code readability and easier maintenance by reducing duplicate code between controllers
type Repository struct {
	DB *gorm.DB // a pointer to GORM's database connection type
	// A database connection is a large object, so just pointing to it instead saves memory
}

type Meetup struct {
	id                int       `json:"id"`
	time              time.Time `json:"time"`
	place             string    `json:"place"`
	friends_attending []Friend  `json:"friends_attending"`
}

type Friend struct {
	id                          int       `json:"id"`
	name                        string    `json:"name"`
	birthday                    time.Time `json:"birthday"`
	days_since_last_interaction int       `json:"days_since_last_interaction"`
	days_since_last_meetup      int       `json:"days_since_last_meetup"`
	phone_number                string    `json:"phone_number"` // To be encrypted
	meetup_plans                []Meetup  `json:"meetup_plans"`
}

type User struct {
	id                          int               `json:"id"`
	username                    string            `json:"username"`
	password                    hash.Hash64       `json:"password"` // To be encrypted
	days_since_last_interaction int               `json:"days_since_last_interaction`
	days_since_last_meetup      int               `json:"days_since_last_meetup"`
	meetup_plans                []Meetup          `json:"meetup_plans"`
	recievesNotifications       (map[string]bool) `json:"recievesNotifications"`
}

func root(writer http.ResponseWriter, request *http.Request) { // We pass the pointer of the HTTP request to avoid copying over a potentially large request that could slow down the server

	// HTTP requests and respones are transmitted as raw bytes, therefore all requests/respones need to be converted to bytes
	// type1(variable) returns a copy of variable converted to type type1
	writer.Write([]byte("'/'에 오신 것을 환영합니다."))

}

func init() {
	// Get database env variables
	godotenv.Load()
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	// db_URL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db_user, db_password, db_host, db_port, db_name)

	// Connect to database
	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai`, db_host, db_user, db_password, db_name, db_port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal(`Error loading .env file: %s`, err)
	// }

	// pool, err := pgxpool

	fmt.Println(db_host, db_port, db_user, db_password, db_name)

	http.HandleFunc("/", root)

	http.ListenAndServe(":3000", nil)

}
