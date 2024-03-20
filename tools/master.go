package tools

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Dictionary 实现Python 字典类型
// 使用空接口实现万能数据类型字典
type Dictionary map[string]interface{}

// Logo_Slabt 打印Logo
func Logo_Slabt(appName string) {

	text := `
    __  ___       __  
   /  |/  /___   / /__
  / /|_/ // _ \ / //_/
 / /  / //  __// ,<       
/_/  /_/ \___//_/|_|  %s
`

	fmt.Printf(text, appName)

}

// Print_Class 打印数据类型
func Print_Class(obj interface{}) {
	fmt.Printf("%T \n", obj)
}

// Read_File_String 读取文件为字符串
func Read_File_String(file string) (s string) {
	b, _ := os.ReadFile(file) // 调用os.ReadFile读取为二进制切片
	s = string(b[:])          // 二进制切片转字符串
	return s                  // 返回字符串
}

// Read_File_Byte 读取文件以二进制切片形式返回
func Read_File_Byte(file string) []byte {
	b, _ := os.ReadFile(file)
	return b // 直切返回二进制切片
}

// Write_File_String 字符串写入文件
func Write_File_String(file string, text string) {
	err := os.WriteFile(file, []byte(text), os.ModePerm) // os.ModePerm 文件写入模式 Unix permission bits, 0o777
	if err != nil {
		fmt.Println("Write_File_String错误:", err)
	}
}

// Addition_File_String 追加写入字符串,不带换行
func Addition_File_String(file string, text string) {
	openFile, _ := os.OpenFile(file, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777) // os.O_RDWR|os.O_APPEND 写入模式 追加模式
	defer openFile.Close()                                                    // 函数结束时执行
	openFile.WriteString(text)                                                // 写入字符串
	openFile.Close()
}

// 读取json 文件 返回字典
func Read_Json(file string) (jsons Dictionary) {
	jsons = Dictionary{} // 空字典
	s := Read_File_String(file)
	err := json.Unmarshal([]byte(s), &jsons) // 传入空字典指针,用于接收读取后的数据
	if err != nil {
		fmt.Println("Read_Json错误:", err)
	}

	return jsons
}

// 写入json 文件
func Write_Json(file string, text interface{}) {
	jsons, err := json.Marshal(text)
	if err != nil {
		fmt.Println("Write_Json错误:", err)
	} else {
		Write_File_String(file, string(jsons))
	}

}

// 读取 Csv文件
func Read_Csv(file string) (list [][]string) {

	list = [][]string{}
	csv_file, err := os.Open(file) //读取文件
	if err != nil {
		fmt.Println(err)
	}

	defer csv_file.Close() // 延迟关闭文件

	reader := csv.NewReader(csv_file)
	reader.FieldsPerRecord = -1
	csvdata, err := reader.ReadAll() //读取全部数据
	// fmt.Println(csvdata)
	for _, line := range csvdata { //按行打印读取
		list = append(list, line)
	}
	return list
}

// 写入 Csv文件 传入[][]interface{}切片
func Write_Csv(file string, list [][]interface{}) {
	// 写入文件如需换成追加模式请修改 File, err := os.OpenFile(file, os.O_RDWR|os.O_WRONLY|os.O_CREATE, 0777)
	File, err := os.OpenFile(file, os.O_RDWR|os.O_WRONLY|os.O_CREATE, 0777) // 写入模式
	if err != nil {
		fmt.Println("写入Csv文件打开失败！")
	}
	defer File.Close()
	//创建写入接口
	WriterCsv := csv.NewWriter(File)
	for _, v := range list {
		s := Interface_Section(v) // []interface{}转换为[]string
		err := WriterCsv.Write(s)
		if err != nil {
			fmt.Println("写入Csv文件错误:", err)
		}
	}
	//WriterCsv.WriteAll(list)
	WriterCsv.Flush()

}

// 写入 Csv文件 传入[][]string切片
func Write_Csv_Sing(file string, list [][]string) {
	// 写入文件如需换成追加模式请修改 File, err := os.OpenFile(file, os.O_RDWR|os.O_WRONLY|os.O_CREATE, 0777)
	File, err := os.OpenFile(file, os.O_RDWR|os.O_WRONLY|os.O_CREATE, 0777) // 写入模式
	if err != nil {
		fmt.Println("写入Csv文件打开失败！")
	}
	defer File.Close()
	//创建写入接口
	WriterCsv := csv.NewWriter(File)
	WriterCsv.WriteAll(list)
	if err != nil {
		fmt.Println("写入Csv文件错误:", err)
	}

	//WriterCsv.WriteAll(list)
	WriterCsv.Flush()

}

// 将[]interface{} 转换为 []string
func Interface_Section(list []interface{}) (s []string) {
	s = []string{}
	for _, v := range list {
		str := fmt.Sprintf("%v", v)
		s = append(s, str)
	}
	return s
}

// 获得当前时间 YYYY-MM-DD H:M:S
func Get_Localtime() (str string) {
	now := time.Now()
	str = now.Format("2006-01-02 03:04:05")
	return
}

// 获得当前时间 YYYYMMDDHMS
func Get_NumberTime() (str string) {
	now := time.Now()
	str = now.Format("20060102030405")
	return
}

// 获得当前时间戳
func Get_Unix() int64 {
	return time.Now().Unix()
}

// 时间戳转换为时间并格式化输出 YYYY-MM-DD H:M:S
func Unix_to_Time(unix int64) string {
	now := time.Unix(unix, 0)
	return now.Format("2006-01-02 03:04:05")

}

// 创建文件夹,可以多级路径
func MkDir(path string) error {

	err := os.MkdirAll(path, 777)
	return err
}

// 递归删除文件
func RmRf(path string) error {
	err := os.RemoveAll(path)
	return err
}

// 判断文件是否存在
func File_Exist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else {
		return false
	}
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// 获取字符串的MD5
func Get_Str_Md5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

// 获取字符串的SHA256
func Get_Str_Sha256(str string) string {
	sum256 := sha256.Sum256([]byte(str))
	toString := hex.EncodeToString(sum256[:])
	return string(toString[:])
}

// 获取字符串的SHA512
func Get_Str_Sha512(str string) string {
	sum512 := sha512.Sum512([]byte(str))
	toString := hex.EncodeToString(sum512[:])
	return string(toString[:])
}

// 获取字符串的SHA
func Get_Str_Sha(str string) string {
	sum := sha1.Sum([]byte(str))
	toString := hex.EncodeToString(sum[:])
	return string(toString[:])
}

// 获得文件MD5
func Get_File_MD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// 获得文件大小 字节
func Get_File_Size(path string) int64 {
	file, err := os.Stat(path)

	if err == nil {
		return file.Size()
	} else {
		return 0
	}

}

// 获得一个N位随机数字
func Get_Random_Number(int int) (N string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < int; i++ {
		N = N + fmt.Sprintf("%v", rand.Intn(10))
	}
	return N
}

// 随机获得N位字母,返回大写格式和小写格式
func Get_Random_Letters(int int) (N string, upper string) {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < int; i++ {
		index := rand.Intn(24)
		N = N + string(letters[index])
	}
	upper = strings.ToUpper(N)
	return N, upper
}
