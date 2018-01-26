package query

import (
	"database/sql"
	"fmt"
	"os"
	"runtime"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func setup() {
	runtime.LockOSThread()

	var err error

	if db, err = dbConnect(); err != nil {
		panic(err.Error())
	}
}

func getEnvValue(name string) string {

	if v := os.Getenv(name); len(v) < 1 {
		panic(fmt.Errorf("Cannot find env:%s", name))
	} else {
		return v
	}
}

func dbConnect() (*sql.DB, error) {

	var dsn = getEnvValue("TEST_QUERY_USER") +
		":" + getEnvValue("TEST_QUERY_PASSWORD") + "@" +
		"tcp(" + getEnvValue("TEST_QUERY_HOST") + ")/" +
		getEnvValue("TEST_QUERY_DATABASE")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	return db, err
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()

	db.Close()
	//
	os.Exit(code)
}
