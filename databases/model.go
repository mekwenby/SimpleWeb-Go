package databases

import (
	"fmt"
	"time"
)

/*
通过数据库文件生成模型
xorm reverse sqlite3 test.db "" ./models
*/

type User struct {
	ID        int64     `xorm:"pk autoincr"`             // 主键并自增
	Username  string    `xorm:"unique notnull"`          // 用户名
	Password  string    `xorm:"notnull"`                 // 密码
	Token     string    `xorm:"index"`                   // 登录 token，可选字段
	UserType  string    `xorm:"notnull default('user')"` // 用户类型，默认是 'user'
	Data      string    `xorm:"notnull"`                 // 用户数据（JSON 或其他格式）
	IsDeleted bool      `xorm:"notnull default(false)"`
	Created   time.Time `xorm:"created"` // 创建时间，由 XORM 自动填充
	Updated   time.Time `xorm:"updated"` // 更新时间，由 XORM 自动更新
}

func SyncTable() {
	err := Engine.Sync2(new(User))
	if err != nil {
		fmt.Println("User 同步错误:", err)
	}
}
