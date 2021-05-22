package models

/*点赞相关操作*/
import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
	"time"
)

//查询点赞情况;
//输入:(ThumbsUpTopic/ThumbsUpComment/ThumbsUpReply)对应结构体;
//输出：存在(返回结构体)，不存在(返回nil)
func QueryThumbsUp(tu interface{}) interface{} {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	switch value := tu.(type) {
	case ThumbsUpTopic:
		if !db.Find(&value).RecordNotFound() {
			return value
		}
	case ThumbsUpReply:
		if !db.Find(&value).RecordNotFound() {
			return value
		}
	case ThumbsUpComment:
		if !db.Find(&value).RecordNotFound() {
			return value
		}
	}
	return nil
}

//插入点赞情况,对相应表项的ThumbsUp进行加一
//输入:(ThumbsUpTopic/ThumbsUpComment/ThumbsUpReply)对应结构体;
//输出：error
func InsertThumbsUp(tu interface{}) error {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	switch value := tu.(type) {
	case ThumbsUpTopic:
		t := Topic{}
		if db.Where("id = ?", value.Tid).First(&t).RecordNotFound() {
			fmt.Println("not found topic id = " + strconv.Itoa(int(value.Tid)))
			return err
		}
		t.ThumbsUp++
		UpdateTopic(t)
		value.CreateAt = time.Now()
		if err := db.Create(&value).Error; err != nil {
			log.Fatal("插入失败 ", err)
			return err
		}
	case ThumbsUpReply:
		r := Reply{}
		if db.Where("id = ?", value.Rid).First(&r).RecordNotFound() {
			log.Fatal("回复不在数据库")
			return err

		}
		r.ThumbsUp++
		UpdateReply(r)
		value.CreateAt = time.Now()
		if err := db.Create(&value).Error; err != nil {
			log.Fatal("插入失败 ", err)
			return err
		}
	case ThumbsUpComment:
		c := Comment{}
		if db.Where("id = ?", value.Cid).First(&c).RecordNotFound() {
			return err
		}
		c.ThumbsUp++
		UpdateComment(c)
		value.CreateAt = time.Now()
		if err := db.Create(&value).Error; err != nil {
			log.Fatal("插入失败 ", err)
			return err
		}
	}
	return nil
}

//删除点赞情况;对相应表项的ThumbsUp进行减一
//输入:(ThumbsUpTopic/ThumbsUpComment/ThumbsUpReply)对应结构体;
//输出:error
func DeleteThumbsUp(tu interface{}) error {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	switch value := tu.(type) {
	case ThumbsUpTopic:
		if err := db.Delete(&value).Error; err != nil {
			log.Fatal(err)
			return err
		}
		t := Topic{}
		if db.Where("id = ?", value.Tid).First(&t).RecordNotFound() {
			fmt.Println("not found topic id = " + strconv.Itoa(int(value.Tid)))
			return err
		}
		t.ThumbsUp--
		UpdateTopic(t)
	case ThumbsUpReply:
		if err := db.Delete(&value).Error; err != nil {
			log.Fatal(err)
			return err
		}
		r := Reply{}
		if db.Where("id = ?", value.Rid).First(&r).RecordNotFound() {
			log.Fatal("回复不在数据库")
			return err

		}
		r.ThumbsUp--
		UpdateReply(r)
	case ThumbsUpComment:
		if err := db.Delete(&value).Error; err != nil {
			log.Fatal(err)
			return err
		}
		c := Comment{}
		if db.Where("id = ?", value.Cid).First(&c).RecordNotFound() {
			return err
		}
		c.ThumbsUp--
		UpdateComment(c)
	}
	return nil
}
