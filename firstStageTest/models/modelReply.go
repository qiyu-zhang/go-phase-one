package models

/*回复相关操作*/
import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)

//查询对应评论下的所有回复，输入：评论id
func QueryReply(id uint) []Reply {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r := make([]Reply, 0)
	if err := db.Where("Comment_Id = ?", id).Find(&r).Error; err != nil {
		fmt.Println("not found " + strconv.Itoa(int(id)))
		fmt.Println(err)
	}
	//fmt.Println("last")
	return r
}

//查询相应回复,输入：回复id
func QueryReplyWithId(id uint) *Reply {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r := new(Reply)
	if err := db.Where("id = ?", id).Find(&r).Error; err != nil {
		fmt.Println("not found " + strconv.Itoa(int(id)))
		fmt.Println(err)
	}
	return r
}

//插入回复
func InsertReply(r *Reply) error {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	if err := db.Create(r).Error; err != nil {
		log.Fatal("插入失败 ", err)
		return err
	}
	return nil
}

//删除一个回复,输入:回复id
func DeleteReply(id uint) error {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	if err := db.Unscoped().Where("id = ?", id).Delete(Reply{}).Error; err != nil {
		log.Fatal(err)
		return err
	}
	if err := db.Unscoped().Where("rid = ?", id).Delete(ThumbsUpReply{}).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

//删除一个评论下的所有回复，输入：评论id
func DeleteAllReply(id uint) error {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	r := QueryReply(id)
	for _, tr := range r {
		DeleteReply(tr.ID)
	}
	//if err := db.Unscoped().Where("Comment_Id = ?", id).Delete(Reply{}).Error; err != nil {
	//	log.Fatal(err)
	//	return err
	//}

	return nil
}

//更新回复
func UpdateReply(r Reply) error {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	if err := db.Save(&r).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
