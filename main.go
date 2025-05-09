package main

import (
	"interview/db"
	"interview/handler"
	"interview/repository"
	"interview/usecase"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(". env not found")
	}

	db, err := db.InitDB()
	if err != nil{
		log.Fatalf("failed connect to db: %v", err)
	}
	defer db.Close()

	router := gin.Default()

	carRepository := repository.NewCarRepository(db)
	carUsecase := usecase.NewCarUsecase(carRepository)
	carHandler := handler.NewCarHandler(carUsecase)

	api := router.Group("/api")
	{
		api.GET("/cars", carHandler.GetAllCars)
		api.GET("/cars/:id", carHandler.GetCarById)
		api.POST("/cars", carHandler.CreateCar)
		api.DELETE("/cars/:id", carHandler.DeleteCar)
		api.PATCH("/cars/:id", carHandler.UpdateCar)
	}

	port := os.Getenv("PORTSERVER")
	if port == ""{
		port = "8080"
	}

	srv := &http.Server{
		Addr: ":"+port,
		Handler: router,
	}

	srv.ListenAndServe()
}
