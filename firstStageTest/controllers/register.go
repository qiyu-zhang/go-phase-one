package controllers

/*注册：
接收前端传回的信息，对信息进行数据库查询，
有就返回用户名存在等信息;
无，则存入数据库，并设置cookie；
*/
import (
	"firstStageTest/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//处理与get相关的操作，
//获取注册页面
//参数：page : 注册页面
func RegisterGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"page": "注册页面",
	})
}

//处理与post相关的操作；
//接收前端传回的信息，对信息进行数据库查询;
//有就返回用户名存在等信息；
//无，则存入数据库，并设置cookie；
//输入:user结构体(JSON)格式
//输出: "用户已存在" 或 "插入失败" 或 设置cookie
func RegisterPost(c *gin.Context) {
	user := models.User{}
	c.BindJSON(&user)
	if models.QueryUser(user.Name) != nil {
		c.JSON(http.StatusOK, "用户已存在")
		return
	}
	if err := models.InsertUser(&user); err != nil {
		c.JSON(http.StatusOK, "插入失败")
	} else {
		c.SetCookie("name", user.Name, 1000, "/", "localhost", false, true)
		c.SetCookie("pwd", user.Password, 1000, "/", "localhost", false, true)
		c.SetCookie("uid", strconv.Itoa(int(user.ID)), 1000, "/", "localhost", false, true)
		return
	}
}
