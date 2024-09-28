package test

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGorm(t *testing.T) {
	myDb := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "111111", "localhost:3306", "zap-netdisk")
	_, err := gorm.Open(mysql.Open(myDb), &gorm.Config{})
	if err != nil {
		t.Fatal(err.Error())
	}
	// conn.First()
}
