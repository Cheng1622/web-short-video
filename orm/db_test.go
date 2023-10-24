package orm

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	InitSqlite()
	fmt.Println("success:", DB)

}
