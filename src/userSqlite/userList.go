package userSqlite

import "github.com/xormplus/xorm"
import (
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"encoding/json"
)

type User struct {
	Id   string  `xorm:"VARCHAR(20)"`
	Name string `xorm:"VARCHAR(20)"`
	UserSex  int64
}

type UserList struct {
	UserAge string `xorm:"VARCHAR(20)"`
	UserSex  int64
}

var engin1 *xorm.Engine

func initSql()  {
	engin1 , _ = xorm.NewEngine("sqlite3","./test.db")

}

func InsertData(data User) bool {
	initSql()
	engin1.CreateTables(&User{})
	_,err :=  engin1.Insert(data)
	if err != nil {
		return false
	}
	return true
}

func QueryData() string {
	initSql()
	engin1.CreateTables(&User{},&UserList{})

	InsertData(User{Id:"2",Name:"kitty",UserSex:1})
	engin1.Insert(UserList{
		UserAge:"20",
		UserSex:2,
	})

	sql := "select a.*,b.* from user a inner join user_list b on a.user_sex = b.user_sex"
	result,_ := engin1.QueryString(sql)


	kk := map[string] interface{}{}
	if len(result) ==0 {
		kk["content"] = []string{}
	}else {
		kk["content"] = result
	}
	kk["code"] = 200

	byte,_ := json.Marshal(kk)

	reslutSlice := []User {}

	for _,v := range result {

		resultMap := make(map[string](string))
		for key,value := range v{
			resultMap[key] = value
			users := User{}
			users.Id  = key
			users.Name = value

			reslutSlice = append(reslutSlice, users)
		}





	}

	fmt.Print(string(byte))
	return string(byte)


}