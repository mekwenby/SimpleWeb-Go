package databases

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Assembly struct {
	ID       string   // id 属性，类型为字符串
	FileList []string // fileList 属性，类型为字符串切片
}

func GetFileList() []string {

	path := "static/archives"
	files, err := GetAllFiles(path)
	if err != nil {
		var errFiles []string
		return errFiles
	} else {
		return files
	}

}

func GetAllFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func BuildFileAssembly() {
	var err error
	_, err = Engine.Where("1=1").Delete(new(Image))
	_, err = Engine.Where("1=1").Delete(new(Index))

	fileList := GetFileList()
	for i := 0; i < len(fileList); i++ {
		fileName := filepath.Base(fileList[i])
		fileId := fileName[:15]
		file := &Image{
			Id:       fileId,
			Filename: fileName,
			Filepath: fileList[i],
		}
		_, err = Engine.Insert(file)

	}

	var filesId []string
	err = Engine.Table(new(Image)).Select("DISTINCT id").Find(&filesId)
	if err != nil {
		log.Fatalf("Failed to query data: %v", err)
	}
	for i := 0; i < len(filesId); i++ {
		index := &Index{
			Id:    int64(i + 1),
			Index: filesId[i],
		}
		_, err = Engine.Insert(index)

	}

}

func CopyFile(srcPath, dstPath string) error {
	// 读取源文件内容
	srcContent, err := ioutil.ReadFile(srcPath)
	if err != nil {
		log.Fatalf("读取文件失败 %v", srcPath)
		return err
	}

	// 写入目标文件
	err = ioutil.WriteFile(dstPath, srcContent, 0644)
	if err != nil {
		log.Fatalf("复制文件失败 %v", dstPath)
		return err
	}

	return nil
}

func AutoCopy(files []Image) {
	dstDir := "Select"
	for i := range files {
		file := files[i]
		srcContent := file.Filepath
		fileName := file.Filename
		dstPath := filepath.Join(dstDir, fileName)
		err := CopyFile(srcContent, dstPath)
		if err != nil {
			log.Fatalf(fileName, "复制失败")
		}
	}
}
