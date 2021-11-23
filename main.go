package main

import (
	"flag"
	"log"

	"github.com/benjamin-guibert/budget-api/controllers"
	"github.com/benjamin-guibert/budget-api/infrastructures"
	"github.com/benjamin-guibert/budget-api/repositories/database"
	"github.com/benjamin-guibert/budget-api/services"
	"github.com/joho/godotenv"
)

var DB database.Database
var Router infrastructures.Router

var MonthBudgetRepository database.MonthBudgetRepository

var MonthBudgetService services.MonthBudgetService

var MonthBudgetController controllers.MonthBudgetController

func main() {
	seed := flag.Bool("s", false, "")
	flag.BoolVar(seed, "seed", *seed, "Seed the database")

	flag.Parse()

	initDatabase()
	initServices()
	initControllers()
	initRouter()

	startDatabase()
	defer DB.Stop()

	migrateDatabase()
	if *seed {
		seedDatabase()
	} else {
		startRouter()
	}
	log.Println("Application stopped")
}

func init() {
	godotenv.Load()
}

func initDatabase() {
	DB = database.NewDatabase()
	MonthBudgetRepository = database.NewMonthBudgetRepository(DB)
}

func startDatabase() {
	err := DB.Start(nil)
	if err != nil {
		panic(err)
	}
}

func migrateDatabase() {
	err := database.Migrate(DB)
	if err != nil {
		panic(err)
	}
}

func seedDatabase() {
	err := database.Seed(MonthBudgetRepository)
	if err != nil {
		panic(err)
	}
}

func initServices() {
	MonthBudgetService = services.NewMonthBudgetService(MonthBudgetRepository)
}

func initControllers() {
	MonthBudgetController = controllers.NewMonthBudgetController(MonthBudgetService)
}

func initRouter() {
	Router = infrastructures.NewRouter()
	infrastructures.AddRoutes(Router, MonthBudgetController)
}

func startRouter() {
	err := Router.Start()
	if err != nil {
		panic(err)
	}
}
