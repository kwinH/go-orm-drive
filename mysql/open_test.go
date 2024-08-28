package mysql

import (
	"fmt"
	"github.com/kwinh/go-orm"
	"testing"
)

type User struct {
	orm.Model
	UserName string
	Password string
	Nickname string
	Status   int8
	Avatar   string
}

var db *orm.DB

func init() {
	var err error
	db, err = orm.Open(Open("root:root@tcp(127.0.0.1:3306)/gin?parseTime=true"))

	if err != nil {
		fmt.Printf("%#v", err.Error())
		panic("连接数据库失败")
	}
}

func TestDB_Migrate(t *testing.T) {
	value := User{}

	db.Migrate.Auto(value, true, true)
}
