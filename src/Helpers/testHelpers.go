package helpers

import (
	"testing"

	fixture "github.com/Gabriel-Ivarsson/GORMdbFixture/src/Fixture"
	"gorm.io/gorm"
)

var (
	dbGorm           *gorm.DB
	additionalTables = []string{"exam_users"}
  tables []interface{}
)

// InitFixture function
func InitFixture(DbGorm_p *gorm.DB, additionalTables_p []string, tables_p ...interface{}) {
  dbGorm = DbGorm_p
  additionalTables = additionalTables_p
  tables = tables_p
}

func SetupTables(entries ...interface{}) error {
	var err error
	for _, entry := range entries {
		err = dbGorm.Create(entry).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// FixtureWrapNonCreate function
func FixtureWrap(t *testing.T, tables ...interface{}) error {
	err := fixture.CleanUp(dbGorm, additionalTables, tables...)
	if err != nil {
		t.Fatal(err)
		return err
	}
	return nil
}
