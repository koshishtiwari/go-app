package main

import(
	"github.com/koshishtiwari/go-app/controllers"
	"github.com/koshishtiwari/go-app/database"
	"github.com/koshishtiwari/go-app/middleware"
	"github.com/koshishtiwari/go-app/routes"
	"github.com/gin-gonic/gin"

)

func main(){
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}

	app := controllers.NewApp(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	//product data, get from products; user data, get from users

	router := gin.New() //creates a new router
	router.Use(gin.Logger())
	

	routes.UserRoutes(router)
	router.User(middleware.Authentication())
}
