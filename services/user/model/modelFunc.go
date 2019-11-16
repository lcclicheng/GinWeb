package model

//获取用户信息
func GetUserInfo(userName string) (User, error) {
	//连接数据库
	var user User
	err := GlobalDB.Where("name=?", userName).Find(&user).Error
	return user, err
}

//更新用户名
func UpdateUserName(oldName, newName string) error {
	//更新
	return GlobalDB.Model(new(User)).
		Where("name=?", oldName).
			Update("name", newName).Error
}
