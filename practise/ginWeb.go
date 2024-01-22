package practise

/**
Gin 是一套使用 golang 打造的 web 框架，官方自己的介紹如下

 > Gin is a web framework written in Golang.
 > It features a martini-like API with much better performance, up to 40 times faster.
 > If you need performance and good productivity, you will love Gin.

為甚麼說是 martini-like API 呢？ 原因是為 martini 也是一款使用 golang 寫的 web framework，gin 就是基於他的設計原理打造，但是效能比 martini 還要強大框架。
[備註] 由於 martini 目前已經沒有繼續維護了

＃安裝
 > go get github.com/gin-gonic/gin
*/

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// Gin 練習
func GinSetRoute() {
	server := gin.Default()

	// 載入靜態檔案
	server.LoadHTMLGlob("template/*")

	// 設定路由
	server.GET("/", ginPage)
	server.POST("/", ginPost)
	server.POST("/formdata", ginPostFormData)

	server.GET("/google.gpt", googleGpt)

	// 啟動服務
	server.Run(":8888")
}

type IndexData struct {
	Title   string
	Content string
}

// get 頁面
func ginPage(c *gin.Context) {
	// 宣告 view 會用到的資料
	data := new(IndexData)
	data.Title = "首頁"
	data.Content = "我的第一個首頁"

	// 載入 template/index.html
	c.HTML(http.StatusOK, "index.html", data)
}

type PostData struct {
	Message string `json:"message"`
}

// post 頁面
func ginPost(c *gin.Context) {
	var postData PostData
	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseMessage := "You sent: " + postData.Message
	c.JSON(http.StatusOK, gin.H{"message": responseMessage})
}

func ginPostFormData(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 檢查資料是否為空
	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or Password is empty"})
		return
	}

	a := returnTest()

	c.JSON(http.StatusOK, gin.H{
		"message":  "Data received successfully",
		"username": username,
		"password": password,
		"a":        a,
	})
}

func googleGpt(c *gin.Context) {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyDNEb6CPs_MkVSU9gnnVLN1x7O6OyBJ78s"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text("誰是 NBA 的 GOAT？"))
	if err != nil {
		log.Fatal(err)
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("無法將資料轉換成 JSON:", err)
		return
	}

	// 頁面上輸出 google.gpt
	c.String(http.StatusOK, string(respJSON))
}

func GoogleGptTest() {

	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyDNEb6CPs_MkVSU9gnnVLN1x7O6OyBJ78s"))
	if err != nil {
		fmt.Println("E1: ")
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text("2022 MLB 世界大賽冠軍是誰？"))
	if err != nil {
		fmt.Println("E2: ")
		log.Fatal(err)
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("無法將資料轉換成 JSON:", err)
		return
	}

	fmt.Println(string(respJSON))
}
