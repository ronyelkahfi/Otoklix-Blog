package main

import (
	"log"
	_blogUsecase "otoklix-blog/business/blogs"
	_blogController "otoklix-blog/controllers/blogs"
	_dbDriver "otoklix-blog/drivers/databases"
	_blogRepo "otoklix-blog/drivers/databases/blogs"
	"time"

	_route "otoklix-blog/apps/routes"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("apps/configs/config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString("database.user"),
		DB_Password: viper.GetString("database.pass"),
		DB_Host:     viper.GetString("database.host"),
		DB_Port:     viper.GetString("database.port"),
		DB_Database: viper.GetString("database.name"),
	}

	db := configDB.ConnectDB()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	blogRepo := _blogRepo.NewBlogRepository(db)
	blogUseCase := _blogUsecase.NewUseCase(blogRepo, timeoutContext)
	blogController := _blogController.NewController(blogUseCase)

	e := echo.New()
	routesInit := _route.ControllerList{
		BlogController: *blogController,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
