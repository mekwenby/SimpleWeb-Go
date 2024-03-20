package databases

import "fmt"

/*
通过数据库文件生成模型
xorm reverse sqlite3 test.db "" ./models
*/

type User struct {
	Id       int64 `xorm:"pk autoincr"`
	Username string
	Password string
	Token    string `xorm:"default 'def'"`
	UserType string
}

func CreateUser(name string, passwd string, usertype string) error {
	user := User{
		Username: name,
		Password: passwd,
		UserType: usertype,
	}

	_, err := Engine.Insert(&user)
	return err
}

func SyncTable() {
	err := Engine.Sync2(new(User))
	if err != nil {
		fmt.Println("User 同步错误:", err)
	}
}
