package controllers

/*
话题搜索页面
根据输入字段返回对应话题
可对话题点赞
*/
import (
	"firstStageTest/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//返回话题搜索页面
//根据接收字段模糊查询相应话题，并返回对应数据
//输入:op : (search )； 话题名称(string(json))
//输出:登录情况,话题搜索页面 若op == search 则 返回搜索到的话题信息([]Topic）
func TopicSearchGet(c *gin.Context) {
	//返回话题搜索页面
	c.JSON(http.StatusOK, gin.H{
		"status": CheckAccount(c),
		"page":   "话题搜索页面",
	})
	//根据接收字段模糊查询相应话题，并返回对应数据
	if c.Query("op") == "search" {
		s := c.Query("searchText")
		c.JSON(http.StatusOK, gin.H{
			"search":  s,
			"status":  CheckAccount(c),
			"message": models.QueryTopic(s),
		})
	}

}

//点赞时调用相关数据库操作；
//输入:op : (thumbsUp); id : 获取topic的id
//输出:未登录：提示信息
//    已登录：添加或删除相应用户，文章的点赞情况，并对话题的点赞数进行加减
func TopicSearchPost(c *gin.Context) {
	//确认是否是点赞
	if c.Query("op") == "thumbsUp" {
		if CheckAccount(c) {
			name, err := c.Cookie("name")
			if err != nil {
				log.Fatal(err)
			}
			user := models.QueryUser(name)
			id, err := strconv.Atoi(c.Query("id")) //获取topic的id
			if err != nil {
				log.Fatal(err)
			}
			tu := models.ThumbsUpTopic{Tid: uint(id), Uid: user.ID}
			if models.QueryThumbsUp(tu) != nil {
				models.DeleteThumbsUp(tu)
			} else {
				models.InsertThumbsUp(tu)
			}
			c.JSON(http.StatusOK, gin.H{
				"status":  CheckAccount(c),
				"message": models.QueryTopicWithId(uint(id)),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "请登录",
			})
		}

	}
}
