package controller

import "github.com/setiadijoe/fullstack/app/middlewares"

func (svr *Server) initializeRoutes() {

	// Home Route
	svr.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(svr.Home)).Methods("GET")

	// Login Route
	svr.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(svr.Login)).Methods("POST")

	//Users routes
	svr.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(svr.CreateUser)).Methods("POST")
	svr.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(svr.GetUsers)).Methods("GET")
	svr.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(svr.GetUser)).Methods("GET")
	svr.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(svr.UpdateUser))).Methods("PUT")
	svr.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(svr.DeleteUser)).Methods("DELETE")

	//Posts routes
	svr.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(svr.CreatePost)).Methods("POST")
	svr.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(svr.GetPosts)).Methods("GET")
	svr.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(svr.GetPost)).Methods("GET")
	svr.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(svr.UpdatePost))).Methods("PUT")
	svr.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(svr.DeletePost)).Methods("DELETE")
}
