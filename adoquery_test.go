package adoquery

import (
	"fmt"
	"testing"
)

func TestADOQuery(t *testing.T) {
	conn := Connection{
		Driver: SQLServer,
		Dsn:    "sqlserver://sa:Sl81262299@zs2019.gdfzjy.com:9304?database=zs2020_backup&encrypt=disable",
	}
	ado := New(conn)
	ado.Sql = "select top 5 bpdzld, inckmc, fsjs, fssl from F_BH_bpdzit where F_BH_bpdzit.bpdzld > 0 order by bpdzld desc"
	ado.Open()
	fmt.Println(ado.JSON())
	fmt.Println(ado.RowsAffected)
	ado.Close()
	ado.Sql = "select top 10 bpdzld, inckmc, fsjs, fssl from F_BH_bpdzit where F_BH_bpdzit.bpdzld > 0 order by bpdzld desc"
	ado.Open()
	fmt.Println(ado.JSON())
	fmt.Println(ado.RowsAffected)
	ado.Close()
}
