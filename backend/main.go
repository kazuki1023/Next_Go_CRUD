package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/kazuki1023/Next_Go_CRUD/ent"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	// Ginフレームワークのデフォルトの設定を使用してルータを作成
	router := gin.Default()

	// PostgreSQLへの接続
	client, err := ent.Open("postgres", "host=db port=5432 user=postgres dbname=db password=password sslmode=disable")

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// CORS設定
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	// ルートハンドラの定義
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// eventの情報を取得
	router.GET("/events", func(c *gin.Context) {
		events, err := client.Event.Query().All(context.Background())
		if err != nil {
			log.Fatalf("failed get events: %v", err)
		}
		c.JSON(200, events)
	})

	// サーバー起動
	router.Run(":8080")
}
