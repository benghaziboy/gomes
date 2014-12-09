package gomes

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gomes/docker"
	"log"
	"os"
)

var (
	Conn *sql.DB

	DbContainerName = os.Getenv("DATABASE_CONTAINER_NAME")
)

func init() {
	dburl := os.Getenv("DATABASE_URL")
	if dburl == "" && os.Getenv("ENVIRONMENT") == "drone" {
		dburl = fmt.Sprintf("user=postgres sslmode=disable host=localhost")
	} else if dburl == "" {
		dbip := os.Getenv("FIG_DATABASE   _ADDR")
		if dbip == "" {
			dbip = docker.FindContainerIp(DbContainerName)
		}
		dburl = fmt.Sprintf("user=postgres sslmode=disable host=%s", dbip)
		fmt.Println(dburl)
	}
	var err error
	Conn, err = sql.Open("postgres", dburl)
	if err != nil {
		log.Fatal(err)
	}
}

func CreatePushTokenTable() error {
	_, err := Conn.Exec(createPushTokenTable)
	if err != nil {
		return err
	}

	return nil
}

func InsertPushToken(pt *PushToken) error {
	_, err := Conn.Exec(insertPushToken, pt.Uid, pt.Arn, pt.ArnType, pt.Token)
	if err != nil {
		return err
	}

	return nil
}

func SelectPushToken(uid string) (*PushToken, error) {
	var pushtoken PushToken

	row, err := Conn.Query(selectPushToken, uid)
	if err != nil {
		return nil, err
	}

	err = row.Scan(&pushtoken)
	if err != nil {
		return nil, err
	}

	return &pushtoken, nil
}
