package gormExtra

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"
)

const environmentVariablePostgresDSN = "GOTEST_POSTGRESQL_DSN"

var gormConfig = &gorm.Config{
	NamingStrategy: SchemaNamerGenericSafe,
}

type table[T any] struct {
	ID    uint `gorm:"primaryKey"`
	Value *JSON[T]
}

func newTable[T any](v T) table[T] {
	return table[T]{
		Value: AsJSON(v),
	}
}

type jsonAbleStruct struct {
	V1 int    `json:"v1"`
	V2 string `json:"v2"`
}

func TestJSON_SQLite(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), gormConfig)
	if !assert.NoError(t, err) {
		return
	}

	testJSONWithAllTypes(t, db)
}

func TestJSON_Postgresql(t *testing.T) {
	dsn, ok := os.LookupEnv(environmentVariablePostgresDSN)
	if !ok {
		t.Skipf("skipping due to missing environment variable '%s'", environmentVariablePostgresDSN)
	}

	db, err := gorm.Open(postgres.Open(dsn))
	if !assert.NoError(t, err) {
		return
	}

	testJSONWithAllTypes(t, db)
}

func testJSONWithAllTypes(t *testing.T, db *gorm.DB) {
	testJSONWithType(t, db, 42)
	testJSONWithType(t, db, "test")
	testJSONWithType(t, db, []string{"a", "b", "c"})
	testJSONWithType(t, db, []int{1, 2, 3})
	testJSONWithType(t, db, map[string]int{"a": 1, "b": 2, "c": 3})
	testJSONWithType(t, db, jsonAbleStruct{
		V1: 42,
		V2: "json rulz",
	})
}

func testJSONWithType[T any](t *testing.T, db *gorm.DB, v T) {
	t.Run(fmt.Sprintf("json handles tpye '%v' correctly", reflect.TypeOf(v)), func(t *testing.T) {
		defer db.Exec(fmt.Sprintf("DROP TABLE \"%s\"", db.NamingStrategy.TableName(reflect.TypeOf(table[T]{}).Name())))

		err := db.AutoMigrate(table[T]{})
		if !assert.NoError(t, err) {
			return
		}

		created := &table[T]{
			Value: AsJSON(v),
		}

		err = db.Create(created).Error
		if !assert.NoErrorf(t, err, "should be able to create a value") {
			return
		}

		var val table[T]
		err = db.Find(&val, created.ID).Error
		if !assert.NoErrorf(t, err, "should be able to query a value") {
			return
		}

		assert.Equal(t, created, &val)
	})
}
