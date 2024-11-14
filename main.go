package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 加载HTML模板
	r.LoadHTMLGlob("templates/*")

	// 使用db
	// 创建 LevelDBWrapper 实例
	dbWrapper, err := NewLevelDBWrapper("data/notes.db")
	if err != nil {
		log.Fatal(err)
	}
	defer dbWrapper.Close()

	r.GET("/", func(c *gin.Context) {

		// 生成一个随机地址并重定向
		randomAddress := generateRandomAddress()

		// 插入数据
		if err := dbWrapper.Put(randomAddress, TextData{
			Text: "Hello, World!",
		}); err != nil {
			log.Fatal(err)
		}

		c.Redirect(http.StatusFound, "/"+randomAddress)
	})

	r.GET("/:address", func(c *gin.Context) {
		address := c.Param("address")
		// 查询数据
		if textData, err := dbWrapper.Get(address); err == nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"RandomAddress": address,
				"Text":          textData.Text,
			})
		} else {
			log.Println(err)
			// 如果没有找到对应的文本数据,则显示404页面
			c.HTML(http.StatusNotFound, "404.html", nil)
		}
	})

	r.POST("/save/:address", func(c *gin.Context) {
		address := c.Param("address")
		// 获取提交的文本内容
		text := c.PostForm("text")

		// 保存文本内容
		// 插入数据
		if err := dbWrapper.Put(address, &TextData{
			Text: text,
		}); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Saved text:", text, "at address:", address)
		c.Redirect(http.StatusFound, "/"+address)
	})

	r.Run()
}

func generateRandomAddress() string {
	// 设置种子,使每次生成的地址不同
	rand.Seed(time.Now().UnixNano())
	// 生成8位随机字母数字字符串
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomAddress := make([]byte, 8)
	for i := range randomAddress {
		randomAddress[i] = chars[rand.Intn(len(chars))]
	}
	return string(randomAddress)
}
