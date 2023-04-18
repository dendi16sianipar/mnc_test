package main

import (
	"fmt"
	"log"
	"mnc_test/controllers"
	"mnc_test/repositories"
	"mnc_test/usecase"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	dbHost := "localhost"
	dbPort := "5432"
	dbName := "merchant_bank"
	dbUser := "postgres"
	dbPass := "dendi16"
	conStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", conStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	g := gin.Default()

	customerRepo := repositories.NewCustomerRepo(db)
	customerUsecase := usecase.NewCustomerUsecase(customerRepo)
	controllers.NewCustomerController(&g.RouterGroup, customerUsecase)

	transactionRepo := repositories.NewPaymentRepo(db)
	transactionUsecase := usecase.NewPaymentUsecase(transactionRepo)
	controllers.NewPaymentUsecase(&g.RouterGroup, transactionUsecase)

	err = g.Run(":8080")
	if err != nil {
		log.Fatal("Unable to run application")
	}
}
