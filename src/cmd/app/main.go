package main

import (
	"database/sql"
	"fmt"
	"nehnutelnosti-sk/src/internal/email"
	"nehnutelnosti-sk/src/internal/uri"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}

func main() {
	uri1, err := uri.NewUrlBuilder().
		WithPlace("kosice").
		WithSize("3-izbove-byty").
		WithArea(55, 70).
		WithPrice(150000, 180000).
		Build()
	if err != nil {
		panic("wrong parameters for uri")
	}

	db, err := initDB()
	if err != nil {
		panic(err)
	}

	email, err := initEmail()
	if err != nil {
		panic(err)
	}

	app := App{
		uri: []string{
			uri1,
		},
		db:    db,
		email: email,
	}

	err = app.CheckUpdated()
	if err != nil {
		panic(err)
	}
}

func initDB() (*sql.DB, error) {
	_ = os.Mkdir("./data", 0755)

	sqlite, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return nil, err
	}

	return sqlite, nil
}

func initEmail() (*email.Email, error) {
	account, account_ok := os.LookupEnv("EMAIL_FROM")
	if !account_ok {
		return nil, fmt.Errorf("email account not found in env")
	}

	to, to_ok := os.LookupEnv("EMAIL_TO")
	if !to_ok {
		return nil, fmt.Errorf("email to not found in env")
	}

	subject, subject_ok := os.LookupEnv("SUBJECT")
	if !subject_ok {
		return nil, fmt.Errorf("email subject not found in env")
	}

	passwd, passwd_ok := os.LookupEnv("GMAIL_APP_PASSWORD")
	if !passwd_ok {
		return nil, fmt.Errorf("email password not found in env")
	}

	return email.NewEmail(account, to, subject, passwd), nil
}
