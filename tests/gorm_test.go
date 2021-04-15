package tests

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func init() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func TestMakce(t *testing.T) {

}
