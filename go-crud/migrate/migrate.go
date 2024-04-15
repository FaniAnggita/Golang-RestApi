package main

import (
	"github.com/FaniAnggita/go-crud/initializers"
	"github.com/FaniAnggita/go-crud/models"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectTODB()
	
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
