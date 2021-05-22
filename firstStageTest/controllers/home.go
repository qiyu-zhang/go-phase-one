package controllers

/*
主页面
该页面只有10个点赞数最高的话题
*/
import (
	"firstStageTest/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//处理与get相关的操作，
//主页面：检测是否登录，返回对应信息
//status:true(登录），false(未登录）
//page:主页面html
//message: []topic数据（点赞最高的10个数据，如果数据库数据充足）
func HomeGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  CheckAccount(c),
		"page":    "主页面html",
		"message": models.QueryTop10Topic(),
	})
}

//处理与post相关的操作，
//点赞时调用相关数据库操作；
//输入:op == thumbsUp(指明是点赞操作),tid == 点击的文章的id
//输出:未登录：提示信息
//    已登录：添加或删除相应用户，文章的点赞情况，并对话题的点赞数进行加减
func HomePost(c *gin.Context) {
	if c.Query("op") == "thumbsUp" {
		if CheckAccount(c) {
			name, err := c.Cookie("name")
			if err != nil {
				log.Fatal(err)
			}
			user := models.QueryUser(name)
			tid, err := strconv.Atoi(c.Query("tid"))
			if err != nil {
				log.Fatal(err)
			}
			tu := models.ThumbsUpTopic{Tid: uint(tid), Uid: user.ID}
			if models.QueryThumbsUp(tu) != nil {
				models.DeleteThumbsUp(tu)
			} else {
				models.InsertThumbsUp(tu)
			}
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "请登录",
			})
		}

	}

}
