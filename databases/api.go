package databases

import (
	"fmt"
	"log"
)

func GetVersionInfo() Version {
	var version Version
	get, err := Engine.Where("id = ?", 0).Get(&version)
	if err != nil {
		log.Fatalf("Failed to query data: %v", err)
	}
	if !get {
		log.Println("No data found")
	}
	return version
}

func GetImageFileList(id string) []Image {
	var tag Index
	Engine.Where("id = ?", id).Get(&tag)
	//fmt.Println(tag.Index)
	var images []Image
	err := Engine.Where("id = ?", tag.Index).Find(&images)
	if err != nil {
		fmt.Println("GetImageFileList Err:", err)
	}
	return images
}

func GetStartIndex() Index {
	var index Index
	Engine.Asc("id").Limit(1).Get(&index)
	return index
}

// GetImageTotal 获取索引总数
func GetImageTotal() (count int64) {
	count, _ = Engine.Count(new(Index))
	return count
}
