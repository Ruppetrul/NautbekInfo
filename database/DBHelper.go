package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type DbSingleton struct {
	Db *sql.DB
}

var instance *DbSingleton
var once sync.Once

func GetDBInstance() (*DbSingleton, error) {
	var err error

	params := map[string]string{
		"user":     os.Getenv("DB_USER"),
		"password": os.Getenv("DB_PASSWORD"),
		"dbname":   os.Getenv("DB_NAME"),
		"sslmode":  os.Getenv("DB_SSL_MODE"),
		"host":     os.Getenv("DB_HOST"),
	}

	var connStr string
	for key, value := range params {
		connStr += fmt.Sprintf("%s=%s ", key, value)
	}
	connStr = strings.TrimSpace(connStr)

	once.Do(func() {
		instance = &DbSingleton{}
		instance.Db, err = sql.Open("postgres", connStr)
	})

	if err != nil {
		log.Fatalf("DB connect error: %v", err)
	}

	return instance, nil
}

func IncrementVisitCount(remoteAddress string, app string) error {
	connection, err := GetDBInstance()

	currentDate := time.Now().Format("2006-01-02")

	_, err = connection.Db.Exec(`
		INSERT INTO user_visits (visit_date, visit_count, visit_ip, app)
		VALUES ($1, 1, $2, $3)
		ON CONFLICT (visit_date, visit_ip)
		DO UPDATE SET visit_count = user_visits.visit_count + 1;
	`, currentDate, remoteAddress, app)

	return err
}

func SaveUserFeedback(remoteAddress string, app string, text string) error {
	connection, err := GetDBInstance()

	currentDate := time.Now().Format("2006-01-02")

	_, err = connection.Db.Exec(`
		INSERT INTO user_feedback (visit_date, visit_ip, app, text)
		VALUES ($1, $2, $3, $4);
	`, currentDate, remoteAddress, app, text)

	return err
}
