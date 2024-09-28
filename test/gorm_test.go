package test

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestGorm(t *testing.T) {
	myDb := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "111111", "localhost:3306", "zap-netdisk")
	_, err := gorm.Open(mysql.Open(myDb), &gorm.Config{})
	if err != nil {
		t.Fatal(err.Error())
	}
	// conn.First()
}

func TestBytes(t *testing.T) {
	f()
	fmt.Println("end")
}

func f() {
	defer func() {
		fmt.Println("defer2 start")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("defer2 end")
	}()
	defer func() {
		fmt.Println("defer1 start") // 3 Execute2
		if err := recover(); err != nil {
			panic(err) // 4 第二次panic()
		}
		fmt.Println("defer1 end")
	}()
	for {
		fmt.Println("func begin") // 1 Execute1
		a := []string{"a", "b"}
		fmt.Println(a[3]) // 2 第一次panic(Out of range)
		fmt.Println("func end")
		time.Sleep(1 * time.Second)
	}
}
