package databases

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

/*
通过数据库文件生成模型
xorm reverse sqlite3 test.db "" ./models
*/

type Image struct {
	Id       string `xorm:"index VARCHAR(255)"`
	Filename string `xorm:"VARCHAR(255)"`
	Filepath string `xorm:"VARCHAR(255)"`
}

type Version struct {
	Id      int    `xorm:"not null pk INTEGER"`
	Version string `xorm:"VARCHAR(10)"`
	Verify  string `xorm:"VARCHAR(255)"`
}

type Index struct {
	Id    int64  `xorm:"'id' pk autoincr"`
	Index string `xorm:"VARCHAR(255)"`
}

func SyncTable() {
	err := Engine.Sync2(new(Image))
	if err != nil {
		fmt.Println("Image 同步错误:", err)
	}
	err2 := Engine.Sync2(new(Version))
	if err2 != nil {
		fmt.Println("Version 同步错误:", err)
	}
	err3 := Engine.Sync2(new(Index))
	if err3 != nil {
		fmt.Println("Index 同步错误:", err)
	}
}

func GetTimeFtm() (timed string) {
	// 获取当前时间
	now := time.Now()

	// 格式化当前时间为YYYYMMDDHMS
	formattedTime := now.Format("20060102150405")

	return formattedTime
}
func hashSHA1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
func CreateVersion() {
	_, err := Engine.Where("1=1").Delete(new(Version))
	if err != nil {
		log.Fatalf("Failed to delete records: %v", err)
	}
	versionNo := GetTimeFtm()
	v := &Version{
		Id:      0,
		Version: versionNo,
		Verify:  hashSHA1(versionNo),
	}
	Engine.Insert(v)

}
