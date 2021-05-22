package models

/*数据库相关表项结构体及注册函数*/
import (
	//"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

//话题
type Topic struct {
	gorm.Model
	Name     string `gorm:"primary_key;"`
	Uid      uint   //发布人
	ThumbsUp int64  //点赞数
	Content  string
}

//评论(对话题)
type Comment struct {
	gorm.Model
	Uid      uint   //评论用户
	Content  string //内容
	ThumbsUp int64
	TopicId  uint //评论的话题的id
}

//回复(对评论)
type Reply struct {
	gorm.Model
	Uid       uint   //回复用户
	Content   string //内容
	ThumbsUp  int64
	CommentId uint //回复的评论的id
}

//话题点赞
type ThumbsUpTopic struct {
	Tid      uint `gorm:"primary_key;default:0" sql:"type:INT(10) UNSIGNED NOT NULL"`
	Uid      uint `gorm:"primary_key;default:0" sql:"type:INT(10) UNSIGNED NOT NULL"`
	CreateAt time.Time
}

//评论点赞
type ThumbsUpComment struct {
	Cid      uint `gorm:"primary_key;" sql:"type:INT(10) UNSIGNED NOT NULL"`
	Uid      uint `gorm:"primary_key;" sql:"type:INT(10) UNSIGNED NOT NULL"`
	CreateAt time.Time
}

//回复点赞
type ThumbsUpReply struct {
	Rid      uint `gorm:"primary_key;" sql:"type:INT(10) UNSIGNED NOT NULL"`
	Uid      uint `gorm:"primary_key;" sql:"type:INT(10) UNSIGNED NOT NULL"`
	CreateAt time.Time
}

//用户
type User struct {
	gorm.Model
	Name     string `gorm:"primary_key"`
	Password string
}

//注册相关表项
func RegisterDB() {
	db, err := gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Topic{})
	db.AutoMigrate(&ThumbsUpTopic{})
	db.AutoMigrate(&ThumbsUpComment{})
	db.AutoMigrate(&ThumbsUpReply{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Reply{})
	db.AutoMigrate(&User{})
	defer db.Close()
}
