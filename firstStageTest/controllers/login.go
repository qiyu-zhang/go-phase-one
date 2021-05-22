package controllers

/*
登录页面
用于处理登入操作
*/
import (
	"firstStageTest/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//登出函数
func LogOut(c *gin.Context) {
	c.SetCookie("name", "", -1, "/", "localhost", false, true)
	c.SetCookie("pwd", "", -1, "/", "localhost", false, true)
	c.SetCookie("uid", "", 0, "/", "localhost", false, true)
	fmt.Println("logout")
	c.Redirect(http.StatusMovedPermanently, "/")
}

//处理与get相关的操作，
//参数:exit == true; 执行退出函数
//参数:exit != true;返回登录页面
func LoginGet(c *gin.Context) {
	if c.Query("exit") == "true" {
		LogOut(c)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"page": "登录页面",
		})
	}
}

//处理与post相关的操作，
//获取信息后，与数据库数据比对，
//用户名查询无结果，返回注册提示；；
//正确就设置cookie;错误，返回密码错误信息；
//输入:User{}(JSON格式)
//输出:"查无此人，请检查用户名" 或 "密码错误" 或 设置cookie
func LoginPost(c *gin.Context) {
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatal(err)
	}
	ck := models.QueryUser(user.Name)
	if ck == nil {
		c.JSON(http.StatusOK, "查无此人，请检查用户名")
		return
	}
	if ck.Password == user.Password {
		c.SetCookie("name", user.Name, 1000, "/", "localhost", false, true)
		c.SetCookie("pwd", user.Password, 1000, "/", "localhost", false, true)
		c.SetCookie("uid", strconv.Itoa(int(ck.ID)), 1000, "/", "localhost", false, true)
		c.Redirect(http.StatusMovedPermanently, "/")

	} else {
		c.JSON(http.StatusOK, "密码错误")
	}

}

//用于判断是否有账户已经登录
//输出：true/false(账户存在情况),若为true，更新cookie
func CheckAccount(c *gin.Context) bool {
	cookieName, err := c.Request.Cookie("name")
	if err != nil {
		return false
	}
	ck := models.QueryUser(cookieName.Value)
	if ck == nil {
		return false
	}
	cookiePwd, err := c.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	cookieId, err := c.Request.Cookie("uid")
	if err != nil {
		return false
	}
	if ck.Password == cookiePwd.Value {
		c.SetCookie(cookieName.Name, cookieName.Value, 1000, cookieName.Path, cookieName.Domain, cookieName.Secure, cookieName.HttpOnly)
		c.SetCookie(cookiePwd.Name, cookiePwd.Value, 1000, cookiePwd.Path, cookiePwd.Domain, cookiePwd.Secure, cookiePwd.HttpOnly)
		c.SetCookie(cookieId.Name, cookieId.Value, 1000, cookieId.Path, cookieId.Domain, cookieId.Secure, cookieId.HttpOnly)
		return true
	} else {
		return false
	}
}
