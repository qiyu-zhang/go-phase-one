package controllers

/*
插入话题
更新相应的话题
*/
import (
	"firstStageTest/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//输出:插入话题页面的html和登录情况
func TopicInsertGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  CheckAccount(c),
		"message": "话题插入页面",
	})
}

//将收到的数据存入数据库
//更新相应的话题
//输入:Topic{}（JSON格式）
//输出:Topic结构体（JSON格式） 重定向到topic页面
func TopicInsertPost(c *gin.Context) {
	//将收到的数据存入数据库
	if c.Query("op") == "insert" {
		if CheckAccount(c) {
			t := &models.Topic{}
			err := c.BindJSON(t)
			if err != nil {
				log.Fatal(err)
			}
			if err := models.InsertTopic(t); err != nil {
				log.Fatal(err)
			}
			c.Redirect(http.StatusMovedPermanently, "/topic?id="+strconv.Itoa(int(t.ID)))
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "请登录",
			})
		}
	}
	//更新相应的话题
	if c.Query("op") == "update" {
		if CheckAccount(c) {
			t := models.Topic{}
			err := c.BindJSON(&t)
			if err != nil {
				log.Fatal(err)
			}
			if err := models.UpdateTopic(t); err != nil {
				log.Fatal(err)
			}
			c.Redirect(http.StatusMovedPermanently, "/topic?id="+strconv.Itoa(int(t.ID)))
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "请登录",
			})
		}
	}
}
