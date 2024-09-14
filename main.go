package main

import (
	_ "github.com/abdulkhaliq7/db/sqlserver"
	"github.com/abdulkhaliq7/db/sql"
)

func main() {

	sql.Open("sqlserver")

}
