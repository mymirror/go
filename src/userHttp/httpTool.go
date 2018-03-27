package userHttp

import (
	"net/http"
	"fmt"
	"userSqlite"
)

type person struct {
	Name string  `json:"name"`
}


func WebService()  {
	go http.HandleFunc("/test",loginTask)
	go http.HandleFunc("/ss",loginTask1)

	err := http.ListenAndServe("192.168.2.16:8080",nil)

	if err != nil {
		//失败
		fmt.Print("error")
	}
}


func loginTask(w http.ResponseWriter, req *http.Request)  {
		req.ParseForm()//解析客户端上传的数据
		userName , found1 := req.Form["userName1"]
		password , found2 := req.Form["passWord1"]

	if !found1 || !found2 {
		fmt.Fprint(w,"userName 或者 passord 为空")
		return ;
	}

	user := userName[0]
	psd := password[0]

	data := userSqlite.User{}
	data.Name = user
	data.Id = psd

	 resultStr := userSqlite.QueryData()

	 fmt.Println("1")
	fmt.Println(resultStr)
	fmt.Fprint(w,resultStr)


}

func loginTask1(w http.ResponseWriter, req *http.Request)  {
	req.ParseForm()//解析客户端上传的数据
	userName , found1 := req.Form["userName"]
	password , found2 := req.Form["passWord"]

	if !found1 || !found2 {
		fmt.Fprint(w,"userName 或者 passord 为空")
		return ;
	}

	user := userName[0]
	psd := password[0]

	data := userSqlite.User{}
	data.Name = user
	data.Id = psd

	resultStr := userSqlite.QueryData()

	fmt.Println("1")
	fmt.Println(resultStr)
	fmt.Fprint(w,"123244")


}