package dialect

import (
	"fmt"
	"reflect"
	"time"
)

type mysql struct{}

var _ Dialect = (*mysql)(nil)

func init() {
	RegisterDialect("mysql", &mysql{})
}
func (m *mysql) DataTypeOf(typ reflect.Value) string {
	switch typ.Kind() {
	case reflect.Bool:
		return "BOOLEAN"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "INT"
	case reflect.Int64, reflect.Uint64:
		return "BIGINT"
	case reflect.Float32, reflect.Float64:
		return "DOUBLE"
	case reflect.String:
		return "VARCHAR(255)"
	case reflect.Array, reflect.Slice:
		return "BLOB"
	case reflect.Struct:
		if _, ok := typ.Interface().(time.Time); ok {
			return "DATETIME"
		}
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind()))
}

func (m *mysql) TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	sql := "SELECT table_name FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = ?"
	return sql, args
}
