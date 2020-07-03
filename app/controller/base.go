package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/setiadijoe/fullstack/app/internal/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql database driver
)

// Server ...
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Initialize ...
func (svr *Server) Initialize(dbDriver, dbUser, dbPassword, dbPort, dbHost, dbName string) {
	var err error
	if dbDriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
		svr.DB, err = gorm.Open(dbDriver, DBURL)
		if nil != err {
			fmt.Printf("cannot_connect_to_%s_database", dbDriver)
			log.Fatal("error:", err)
		} else {
			fmt.Printf("connected_to_%s_database/n", dbDriver)
		}
	} else {
		fmt.Printf("unknown_driver")
		log.Fatal("unknown_driver")
	}

	svr.DB.Debug().AutoMigrate(&models.User{}, &models.Post{})
	svr.Router = mux.NewRouter()

	svr.initializeRoutes()
}

// Run ...
func (svr *Server) Run(addr string) {
	fmt.Println("Listening to port", addr)
	log.Fatal(http.ListenAndServe(addr, svr.Router))
}
