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

	// ルートハンドラの定義
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// サーバー起動
	router.Run(":8080")
}
