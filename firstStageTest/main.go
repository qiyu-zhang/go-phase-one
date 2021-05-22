package main

import (
	"firstStageTest/controllers"
	"firstStageTest/models"
	"github.com/gin-gonic/gin"
)

func init() {
	models.RegisterDB()
}

func main() {
	r := gin.Default()
	//html加载模板
	//	r.LoadHTMLGlob("templates/*")
	//渲染模板
	//c.HTML(http.StatusOK, "login.html", gin.H{
	//
	//	})
	r.GET("/", controllers.HomeGet)
	r.POST("/", controllers.HomePost)
	r.GET("/login", controllers.LoginGet)
	r.POST("/login", controllers.LoginPost)
	r.GET("/register", controllers.RegisterGet)
	r.POST("/register", controllers.RegisterPost)
	r.GET("/topic", controllers.TopicGet)
	r.POST("/topic", controllers.TopicPost)
	r.GET("/topicInsert", controllers.TopicInsertGet)
	r.POST("/topicInsert", controllers.TopicInsertPost)
	r.GET("/topicSearch", controllers.TopicSearchGet)
	r.POST("/topicSearch", controllers.TopicSearchPost)
	r.Run(":8080")
}

//func main() {
//	//models.RegisterDB()
//	var t models.Topic
//	for i := 0; i < 10; i++ {
//		time.Sleep(time.Second)
//		t = models.Topic{Model: gorm.Model{}, Name: "test:" + strconv.Itoa(i+1), Content: "test", ThumbsUp: 0, Uid: uint(i)}
//		models.InsertTopic(&t)
//	}
//	//models.DeleteTopic(3)
//	//t := models.Topic{Model: gorm.Model{ID: 2}, Name: "test:" + strconv.Itoa(1), Content: "test update", ThumbsUp: 0}
//	//models.UpdateTopic(t)
//	//fmt.Println(models.QueryTopic("1"))
//
//	var c models.Comment
//	for i := 0; i < 10; i++ {
//		time.Sleep(time.Second)
//		c = models.Comment{Model: gorm.Model{}, Uid: uint(i + 1), Content: "test", ThumbsUp: 0, TopicId: uint(i)}
//		models.InsertComment(&c)
//	}
//	//models.DeleteComment(3)
//	//c = models.Comment{Model: gorm.Model{ID: 1}, Uid: uint(3), Content: "test update", ThumbsUp: 0, TopicId: uint(1)}
//	//models.UpdateComment(c)
//	//fmt.Println(models.QueryComment(uint(1)))
//	//models.DeleteAllComment(uint(5))
//	//fmt.Println(models.QueryComment(uint(5)))
//
//	var R models.Reply
//	for i := 0; i < 10; i++ {
//		time.Sleep(time.Second)
//		R = models.Reply{Model: gorm.Model{}, Uid: uint(i + 1), Content: "test", ThumbsUp: 0, CommentId: uint(i)}
//		models.InsertReply(&R)
//	}
//	//models.DeleteReply(3)
//	//R = models.Reply{Model: gorm.Model{ID: 1}, Uid: uint(3), Content: "test update", ThumbsUp: 0, CommentId: uint(1)}
//	//models.UpdateReply(R)
//	//fmt.Println(models.QueryReply(uint(1)))
//	//models.DeleteAllReply(uint(5))
//	//fmt.Println(models.QueryReply(uint(5)))
//
//	var u models.User
//	for i := 0; i < 10; i++ {
//		time.Sleep(time.Second)
//		u = models.User{Model: gorm.Model{}, Name: strconv.Itoa(i + 1), Password: strconv.Itoa(i * 111)}
//		models.InsertUser(&u)
//	}
//	//models.DeleteUser(3)
//	//u = models.User{Model: gorm.Model{ID: 1}, Name: strconv.Itoa(0), Password: strconv.Itoa(123)}
//	//models.UpdateUser(u)
//	//fmt.Println(models.QueryUser("1"))
//
//	var tut models.ThumbsUpTopic
//	for i := 0; i < 10; i++ {
//		time.Sleep(time.Second)
//		tut = models.ThumbsUpTopic{Tid: uint(3), Uid: uint(i + 1), CreateAt: time.Now()}
//		models.InsertThumbsUp(tut)
//	}
//	//tu = models.ThumbsUpTopic{Tid: uint(3), Uid: uint(4)}
//	//models.DeleteThumbsUp(t)
//	//tu = models.ThumbsUpTopic{Tid: uint(3), Uid: uint(3)}
//	//fmt.Println(models.QueryThumbsUp(tu))
//	var tuc models.ThumbsUpComment
//	for i := 0; i < 10; i++ {
//		time.Sleep(time.Second)
//		tuc = models.ThumbsUpComment{Cid: uint(2), Uid: uint(i + 1), CreateAt: time.Now()}
//		models.InsertThumbsUp(tuc)
//	}
//
//	var tur models.ThumbsUpReply
//	for i := 0; i < 10; i++ {
//		time.Sleep(time.Second)
//		tur = models.ThumbsUpReply{Rid: uint(4), Uid: uint(i + 1), CreateAt: time.Now()}
//		models.InsertThumbsUp(tur)
//	}
//}
