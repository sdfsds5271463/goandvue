package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	port := os.Getenv("PORT")
	if port == "" {
	port = "8080"
	}

	// 從環境變數讀取 Redis 位址，若沒設定則預設連本地 localhost
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379" 
	}

	// Redis 連線
	rdb := redis.NewClient(&redis.Options{
	Addr: redisAddr, // Kubernetes 內部 service 名稱 redis:6379
	})

	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// 取得目前數字
	num, err := rdb.Get(ctx, "counter").Int()
	if err == redis.Nil {
		num = 0
	}

	num++

	// 存回 Redis
	rdb.Set(ctx, "counter", num, 0)

	fmt.Fprintf(w, "Hello from Go API:%d", num)
	})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Server running on port", port)
	http.ListenAndServe(":"+port, nil)
} 