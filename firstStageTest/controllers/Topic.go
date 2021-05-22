package controllers

/*
话题页面
返回话题数据及相应评论数据
可post请求对评论的回复
可记录点赞信息
点赞时调用相关数据库操作；
评论时，将收到的信息存储或更新并返回;
回复：将收到的消息存入数据库或更新并返回。
删除自己的话题及相应的评论及回复
删除自己的评论及相应的回复
删除自己的的回复
*/
import (
	"firstStageTest/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//返回话题页面
//根据传回的话题ID，在数据库中检索该话题，及对应评论并返回
//可返回指定评论的回复
//参数：status:true/false(登录情况); message: 评论的回复 或 话题信息(Topic)及评论信息
//输入: op : replySearch ; id : 指定评论的id 或 本页面话题id
//输出: op == replySearch 则 返回指定评论的回复 否则 返回页面初始信息：
func TopicGet(c *gin.Context) {
	//返回指定评论的回复
	if c.Query("op") == "replySearch" {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  CheckAccount(c),
			"message": models.QueryReply(uint(id)),
		})
	} else {
		//返回话题页面 : 根据传回的话题ID，在数据库中检索该话题，及对应评论并返回
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": CheckAccount(c),
			"page":   "话题页面",
			"message": gin.H{
				"topic":   models.QueryTopicWithId(uint(id)),
				"comment": models.QueryComment(uint(id)),
			},
		})
	}
}

//点赞时调用相关数据库操作；
//评论时，将收到的信息存储或更新并返回;
//回复：将收到的消息存入数据库或更新并返回。
//删除自己的话题及相应的评论及回复
//删除自己的评论及相应的回复
//删除自己的的回复
//输入:op : post类型(thumbsUp/comment/reply/commentUpdate/replyUpdate/deleteTopic/deleteComment/deleteReply),
//    id: topic/comment/reply的id, JSON:Comment{} 或 Reply{}
//    subTextType : topic/comment/reply
//输出:点赞:刷新页面， 评论/回复:返回对应结构体
func TopicPost(c *gin.Context) {
	if CheckAccount(c) {
		//确认是否是点赞,op ,id,subTextType
		if c.Query("op") == "thumbsUp" {
			name, err := c.Cookie("name")
			if err != nil {
				log.Fatal(err)
			}
			user := models.QueryUser(name)
			id, err := strconv.Atoi(c.Query("id")) //获取点赞对象topic/comment/reply的id
			if err != nil {
				log.Fatal(err)
			}
			var tu interface{}
			//确认数据类型
			subType := c.Query("subTextType")
			switch subType {
			case "topic":
				tu = models.ThumbsUpTopic{Tid: uint(id), Uid: user.ID}
			case "comment":
				tu = models.ThumbsUpComment{Cid: uint(id), Uid: user.ID}
			case "reply":
				tu = models.ThumbsUpReply{Rid: uint(id), Uid: user.ID}
			}

			if models.QueryThumbsUp(tu) != nil {
				models.DeleteThumbsUp(tu)
			} else {
				models.InsertThumbsUp(tu)
			}
			switch subType {
			case "topic":
				c.JSON(http.StatusOK, gin.H{
					"status":  CheckAccount(c),
					"message": models.QueryTopicWithId(uint(id)),
				})
			case "comment":
				c.JSON(http.StatusOK, gin.H{
					"status":  CheckAccount(c),
					"message": models.QueryCommentWithId(uint(id)),
				})
			case "reply":
				c.JSON(http.StatusOK, gin.H{
					"status":  CheckAccount(c),
					"message": models.QueryReplyWithId(uint(id)),
				})
			}
		}
		//评论时，将收到的信息存储并返回;
		if c.Query("op") == "comment" {
			comment := models.Comment{}
			c.BindJSON(&comment)
			if err := models.InsertComment(&comment); err != nil {
				log.Fatal("评论时")
				log.Fatal(err)
			}
			c.JSON(http.StatusOK, gin.H{
				"status":  true,
				"message": comment,
			})
		}
		//回复：将收到的消息存入数据库并返回。
		if c.Query("op") == "reply" {
			reply := models.Reply{}
			c.BindJSON(&reply)
			if err := models.InsertReply(&reply); err != nil {
				log.Fatal("回复时")
				log.Fatal(err)
			}
			c.JSON(http.StatusOK, gin.H{
				"status":  true,
				"message": reply,
			})
		}
		//评论：将收到的消息更新入数据库并返回。
		if c.Query("op") == "commentUpdate" {
			comment := models.Comment{}
			c.BindJSON(&comment)
			if err := models.UpdateComment(comment); err != nil {
				log.Fatal("评论时")
				log.Fatal(err)
			}
			c.JSON(http.StatusOK, gin.H{
				"status":  true,
				"message": comment,
			})
		}
		//回复：将收到的消息更新入数据库并返回。
		if c.Query("op") == "replyUpdate" {
			reply := models.Reply{}
			c.BindJSON(&reply)
			if err := models.UpdateReply(reply); err != nil {
				log.Fatal("回复时")
				log.Fatal(err)
			}
			c.JSON(http.StatusOK, gin.H{
				"status":  true,
				"message": reply,
			})
		}
		//删除自己的话题及相应的评论及回复
		if c.Query("op") == "deleteTopic" {
			Uid, err := c.Cookie("uid")
			if err != nil {
				log.Fatal(err)
			}
			var uid int
			if uid, err = strconv.Atoi(Uid); err != nil {
				log.Fatal(err)
			}
			var tid int
			if tid, err = strconv.Atoi(c.Query("id")); err != nil {
				log.Fatal(err)
			}
			t := models.QueryTopicWithId(uint(tid))
			if t.Uid == uint(uid) {
				models.DeleteTopic(uint(tid))
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status":  true,
					"message": "非本人的话题",
				})
			}

		}
		//删除自己的评论及相应的回复
		if c.Query("op") == "deleteComment" {
			Uid, err := c.Cookie("uid")
			if err != nil {
				log.Fatal(err)
			}
			var uid int
			if uid, err = strconv.Atoi(Uid); err != nil {
				log.Fatal(err)
			}
			var cid int
			if cid, err = strconv.Atoi(c.Query("id")); err != nil {
				log.Fatal(err)
			}
			comment := models.QueryCommentWithId(uint(cid))
			if comment.Uid == uint(uid) {
				models.DeleteComment(uint(cid))
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status":  true,
					"message": "非本人的评论",
				})
			}

		}
		//删除自己的的回复
		if c.Query("op") == "deleteReply" {
			Uid, err := c.Cookie("uid")
			if err != nil {
				log.Fatal(err)
			}
			var uid int
			if uid, err = strconv.Atoi(Uid); err != nil {
				log.Fatal(err)
			}
			var rid int
			if rid, err = strconv.Atoi(c.Query("id")); err != nil {
				log.Fatal(err)
			}
			r := models.QueryReplyWithId(uint(rid))
			if r.Uid == uint(uid) {
				models.DeleteReply(uint(rid))
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status":  true,
					"message": "非本人的回复",
				})
			}

		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "请登录",
		})
	}

}
