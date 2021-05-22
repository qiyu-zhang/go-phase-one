package models

/*用户相关操作*/
import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

//查询用户;
//输入：名称;
//输出：密码;无结果，就返回""
func QueryUser(name string) *User {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	t := new(User)
	if err := db.Where("name = ?", name).Find(&t).Error; err != nil {
		fmt.Println("not found " + name)
		return nil
	}
	return t
}

//插入用户
func InsertUser(u *User) error {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	if err := db.Create(u).Error; err != nil {
		log.Fatal("插入失败 ", err)
		return err
	}
	return nil
}

//删除特定用户，输入:用户id
func DeleteUser(id uint) error {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	if err := db.Unscoped().Where("id = ?", id).Delete(User{}).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

//更新用户
func UpdateUser(u User) error {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	if err := db.Save(&u).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
