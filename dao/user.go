package dao

import "log"

type User struct {
	ID       int64  // 主键
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	//创建时间，时间戳
	CreateTime int64 `gorm:"column:createtime"`
}

func (u User) TableName() string {
	return "users"
}

func SaveUser(user *User) {
	err := DB.Create(user).Error
	if err != nil {
		log.Println("insert user error", err)
	}
}

func GetById(id int64) User {
	var user User
	err := DB.Where("id=?", id).First(&user).Error
	if err != nil {
		log.Println("get user by id", err)
	}
	return user
}

func GetAll() []User {
	var user []User
	err := DB.Find(&user).Error
	if err != nil {
		log.Println("get user by error ", err)
	}
	return user
}

func UpdateUser(id int64) {
	err := DB.Model(&User{}).Where("id=?", id).Update("username", "lisi").Error
	if err != nil {
		log.Println("update user by id error ", err)
	}
}

func DeleteUser(id int64) {
	err := DB.Where("id=?", id).Delete(&User{}).Error
	if err != nil {
		log.Println("delete user by id error ", err)
	}
}
