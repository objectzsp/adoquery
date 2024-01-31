package adoquery

import (
	"database/sql"
	"encoding/json"
	"reflect"
	"strconv"
)

type ADOQuery struct {
	Database    Connection
	Sql         string
	Description string

	Json         []byte
	Error        error
	RowsAffected int64
}

func New(conn Connection) *ADOQuery {
	return &ADOQuery{
		Database:    conn,
		Sql:         "",
		Description: "A New ADOQuery",
	}
}

// 打开数据库连接并使用内置SQL字段查询并记录结果集
func (ado *ADOQuery) Open() {
	ado.RowsAffected = 0
	ado.Error = nil
	err := ado.Database.Connect()
	if err != nil {
		ado.Error = err
		return
	}
	rows, err := ado.Database.db.Raw(ado.Sql).Rows()
	if err != nil {
		ado.Error = err
		return
	}
	ado.scanRows(rows)
}

// 关闭结果集并关闭数据库连接
func (ado *ADOQuery) Close() {
	ado.Error = ado.Database.Disconnect()
}

func (ado *ADOQuery) scanRows(rows *sql.Rows) {
	var (
		res         = make([]map[string]interface{}, 0)
		colTypes, _ = rows.ColumnTypes()
		value       = make([]interface{}, len(colTypes))
		parma       = make([]interface{}, len(colTypes))
	)
	for i, colType := range colTypes {
		value[i] = reflect.New(colType.ScanType())
		parma[i] = reflect.ValueOf(&value[i]).Interface()
	}
	for rows.Next() {
		ado.RowsAffected++
		rows.Scan(parma...)
		record := make(map[string]interface{})
		for i, colType := range colTypes {
			if value[i] == nil {
				record[colType.Name()] = ""
			} else {
				switch value[i].(type) {
				case []byte:
					f, _ := strconv.ParseFloat(string(value[i].([]byte)), 64)
					record[colType.Name()] = f
				default:
					record[colType.Name()] = value[i]
				}
			}
		}
		res = append(res, record)
	}
	ado.Json, _ = json.Marshal(res)
}

func (ado *ADOQuery) JSON() string {
	return string(ado.Json)
}
