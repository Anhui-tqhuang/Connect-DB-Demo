package main

import (
	"fmt"
	"os"

	"github.com/golang/glog"

	"database/sql"

	_ "github.com/lib/pq"
)

var (
	hostname string
	port     string
	username string
	password string
	database string
	sslmode  string
	cafile   string
)

func init() {
	hostname = os.Getenv("HOSTNAME")
	port = os.Getenv("PORT")
	username = "admin"
	password = os.Getenv("PASSWORD")
	database = os.Getenv("DATABASE")
	sslmode = os.Getenv("SSLMODE")
	cafile = os.Getenv("CAFILE")
}

const pgConnectionString = "host=%s port=%s dbname=%s user=%s password=%s sslmode=%s sslrootcert=%s connect_timeout=6"

func main() {

	connStr := fmt.Sprintf(pgConnectionString, hostname, port, database, username, password, sslmode, cafile)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		glog.Error(err)
	}

	defer db.Close()

	_, err = db.Query("CREATE TABLE IF NOT EXISTS demo ()")
	if err != nil {
		glog.Error(err)
	}

	rows, err := db.Query("select tablename from pg_tables")
	if err != nil {
		glog.Error(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			tablename string
		)
		if err := rows.Scan(&tablename); err != nil {
			glog.Error(err)
		}
		fmt.Printf("%s\n", tablename)
	}
}
