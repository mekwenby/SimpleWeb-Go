package databases

import "fmt"

/*
用户Api
编写结构体的方法使用
*/

func CreateUser(username, password string) (*User, error) {
	if username == "" || password == "" {
		return nil, fmt.Errorf("用户名和密码不能为空")
	}

	// 检查用户名是否已存在（未被逻辑删除）
	exists, err := Engine.Where("username = ? AND is_deleted = false", username).Exist(&User{})
	if err != nil {
		return nil, fmt.Errorf("检查用户名时出错: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("用户名已存在")
	}
	PasswdSHA256 := StringToSHA256(password)
	// 构建用户对象
	user := &User{
		Username:  username,
		Password:  PasswdSHA256,
		UserType:  "user", // 默认用户类型
		IsDeleted: false,
		Data:      "{}",
	}

	_, err = Engine.Insert(user)
	if err != nil {
		return nil, fmt.Errorf("创建用户失败: %w", err)
	}

	return user, nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	has, err := Engine.Where("username = ? AND is_deleted = false", username).Get(&user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil // 用户未找到
	}
	return &user, nil
}

// VerifyPassword 根据用户名和密码进行验证
func VerifyPassword(username, password string) (*User, error) {
	// 1. 从数据库查找用户
	user, err := GetUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("数据库查询失败: %v", err)
	}
	// 2. 如果用户不存在
	if user == nil {
		return nil, fmt.Errorf("用户不存在")
	}
	// 3. 比较密码
	if user.Password != StringToSHA256(password) {
		return nil, fmt.Errorf("密码错误")
	}
	// 4. 更新用户的 Token 到数据库
	user.Token, err = GenerateToken(user.Username)
	if err != nil {
		return nil, fmt.Errorf("生成 Token 失败: %v", err)
	}
	_, err = Engine.ID(user.ID).Update(user)
	if err != nil {
		fmt.Println("Update User Error:", err)
		return nil, fmt.Errorf("更新用户失败: %v", err)
	}
	return user, nil
}

// FormTokenGetName 验证Token
func FormTokenGetName(token string) (map[string]interface{}, *User, error) {
	var user User
	_, err := Engine.Where("token=?", token).Get(&user)
	if err != nil {
		return nil, nil, err
	}
	if user.ID == 0 {
		return nil, nil, fmt.Errorf("无效的 Token")
	}
	userName, err := VerifyToken(token)
	return userName, &user, nil
}
