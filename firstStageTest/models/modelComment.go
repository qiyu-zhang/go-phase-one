package models

/*评论相关操作*/
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"strconv"
)

//查询对应话题下的所有评论，输入：话题id
func QueryComment(id uint) []Comment {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	c := make([]Comment, 0)
	if err := db.Where("Topic_Id = ?", id).Find(&c).Error; err != nil {
		fmt.Println("not found " + strconv.Itoa(int(id)))
		fmt.Println(err)
	}
	return c
}

//查询相应评论,输入：评论id
func QueryCommentWithId(id uint) *Comment {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	c := new(Comment)
	if err := db.Where("id = ?", id).Find(&c).Error; err != nil {
		fmt.Println("not found " + strconv.Itoa(int(id)))
		fmt.Println(err)
	}
	return c
}

//插入评论
func InsertComment(c *Comment) error {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	if err := db.Create(c).Error; err != nil {
		log.Fatal("插入失败 ", err)
		return err
	}
	return nil
}

//删除一个评论,输入:评论id
func DeleteComment(id uint) error {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	if err := db.Unscoped().Where("id = ?", id).Delete(Comment{}).Error; err != nil {
		log.Fatal(err)
		return err
	}
	if err := db.Unscoped().Where("cid = ?", id).Delete(ThumbsUpComment{}).Error; err != nil {
		log.Fatal(err)
		return err
	}
	if err := DeleteAllReply(id); err != nil {
		return err
	}
	return nil
}

//删除一个话题下的所有评论，输入：话题id
func DeleteAllComment(id uint) error {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	c := QueryComment(id)
	for _, tc := range c {
		DeleteComment(tc.ID)
	}
	//if err := db.Unscoped().Where("Topic_Id = ?", id).Delete(Comment{}).Error; err != nil {
	//	log.Fatal(err)
	//	return err
	//}

	return nil
}

//更新评论
func UpdateComment(c Comment) error {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	if err := db.Save(&c).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
