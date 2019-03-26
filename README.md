**Query** is a tiny package for executing select-like queries and converting `sql.Rows` into slice of `map[string]string`. The new map is allocated for each row. Keys are column names as seen in sql.



## Install

Use go modules and just
```
import "github.com/rogozhka/query"
``` 

or 

```bash
go get -u github.com/rogozhka/query
```



## Example

```go
package main

import (
	"github.com/rogozhka/query"

	"database/sql"
	"fmt"
	"log"
	
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql",
		"username:pass@tcp(hostname)/database_name")
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic(fmt.Errorf("db is nil"))
	}
	defer db.Close()

	// fetch all the rows into slice of map[string]string
	rows, err := query.Fetch("SELECT * FROM `example_table`", db)
	if err != nil {
		panic(err)
	}

	for _, row := range rows {
		for colName, val := range row {
			log.Printf("%s:%s", colName, val)
		}
	}

	// fetch only 1 row
	// in real project you should use LIMIT 0,1
	// but there is limited version of method ;)
	rows2, err := query.FetchLimited("SELECT * FROM `example_table`", 1, db)
	if err != nil {
		panic(err)
	}

	if len(rows2) > 0 {
		for colName, val := range rows2[0] {
			log.Printf("%s:%s", colName, val)
		}
	}
}
```

