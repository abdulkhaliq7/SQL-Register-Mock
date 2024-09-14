package sqlserver

import (
	"fmt"

	"github.com/abdulkhaliq7/db/sql"
)

func init() {

	sql.Register("sqlserver", &Sqlserver{})
}

type Sqlserver struct {
}

func (s Sqlserver) Open(name string) {

	fmt.Printf("%v Database", name)
}
