package gormExtra

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"database/sql/driver"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var SchemaNamerGenericSafe = schema.NamingStrategy{
	NameReplacer: strings.NewReplacer("[", "_", "]", "_", ".", "_", "/", "_"),
}

type JSON[T any] struct {
	V T
}

func AsJSON[T any](value T) *JSON[T] {
	return &JSON[T]{
		V: value,
	}
}

func (j *JSON[T]) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan JSON value '%v' of type %v", src, reflect.TypeOf(src))
	}

	return json.Unmarshal(bytes, &j.V)
}

func (j *JSON[T]) Value() (driver.Value, error) {
	return json.Marshal(j.V)
}

func (JSON[T]) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case "mysql", "sqlite":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}
