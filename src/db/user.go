package db

import "scrips/src/model"

/**
创建用户
*/
func CreateUser(user *model.User) {
	DB.Create(user)
}

/**
找回密码
*/
func EditPassword(email string, newPsd string) {
	u := model.User{}
	u.Email = email
	DB.First(&u)
	u.Password = newPsd
	DB.Save(&u)
}

/**
查找用户
*/
func SearchUser(email string) *model.User {

	user := &model.User{Email: email}

	if DB.Find(user).RecordNotFound() {
		return nil
	}

	return user
}

/**
通过ID查找用户
*/
func SearchUserById(id int) *model.User {

	user := &model.User{}
	user.ID = uint(id)

	if DB.Find(user).RecordNotFound() {
		return nil
	}

	return user
}

/**
通过ID查找用户
*/
func SaveUserInfo(user *model.User) *model.User {
	DB.Save(user)
	return user
}
