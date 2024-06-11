package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lkphuong/room-management/internal/modules/auth"
	"github.com/lkphuong/room-management/internal/modules/receipt"
	"github.com/lkphuong/room-management/internal/modules/room"
	"github.com/lkphuong/room-management/internal/modules/store"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	httpPort := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))

	fmt.Println("Server running on port", httpPort)

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	r.Use(cors.New(config))

	api := r.Group("/api")
	{
		auth.AuthRoutes(api)
		room.RoomRoutes(api)
		store.StoreRoutes(api)
		receipt.ReceiptRoutes(api)
	}

	r.Run(httpPort)

	select {}
}
