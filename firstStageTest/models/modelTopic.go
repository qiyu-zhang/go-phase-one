package models

/*话题相关操作*/
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

//模糊查询相应话题,输入：名称一部分
func QueryTopic(name string) []Topic {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	t := make([]Topic, 0)
	//db.Where("name LIKE ?", "%jin%").Find(&users)
	//fmt.Println(db.HasTable(&Topic{}))
	if err := db.Where("Name LIKE ?", "%"+name+"%").Find(&t).Error; err != nil {
		fmt.Println("not found " + name)
		fmt.Println(err)
	}
	//fmt.Println("last")
	return t
}

//查询相应话题,输入：话题id
func QueryTopicWithId(id uint) *Topic {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	t := new(Topic)
	//fmt.Println(db.HasTable(&Topic{}))
	if err := db.Where("id = ?", id).Find(&t).Error; err != nil {
		fmt.Println("not found Topic")
		fmt.Println(err)
	}
	//fmt.Println("last")
	return t
}

//查询点赞数较多的（最多10个）话题
func QueryTop10Topic() []Topic {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	t := make([]Topic, 0)
	//db.Where("name LIKE ?", "%jin%").Find(&users)
	//db.Order("age desc, name").Find(&users)
	db.Order("thumbs_up desc").Find(&t).Limit(10)
	return t
}

//插入话题
func InsertTopic(t *Topic) error {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	if err := db.Create(t).Error; err != nil {
		log.Fatal("插入失败 ", err)
		return err
	}
	return nil
}

//删除特定话题及对应的评论/回复,点赞情况,输入:话题id
func DeleteTopic(id uint) error {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	if err := db.Unscoped().Where("id = ?", id).Delete(Topic{}).Error; err != nil {
		log.Fatal(err)
		return err
	}
	if err := db.Unscoped().Where("tid = ?", id).Delete(ThumbsUpTopic{}).Error; err != nil {
		log.Fatal(err)
		return err
	}
	DeleteAllComment(id)
	return nil
}

//删除所有话题
func DeleteAllTopic() error {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	if err := db.Unscoped().Delete(Topic{}).Error; err != nil {
		log.Fatal(err)
		return err
	}
	if err := db.Unscoped().Delete(Comment{}).Error; err != nil {
		log.Fatal(err)
		return err
	}
	if err := db.Unscoped().Delete(Reply{}).Error; err != nil {
		log.Fatal(err)
		return err
	}
	if err := db.Unscoped().Delete(ThumbsUpComment{}).Error; err != nil {
		log.Fatal(err)
		return err
	}
	if err := db.Unscoped().Delete(ThumbsUpReply{}).Error; err != nil {
		log.Fatal(err)
		return err
	}
	if err := db.Unscoped().Delete(ThumbsUpTopic{}).Error; err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

//更新话题
func UpdateTopic(t Topic) error {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	if err := db.Save(&t).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
