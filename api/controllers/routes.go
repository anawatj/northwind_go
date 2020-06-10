package controllers

import "github.com/northwind_go/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/api/v1/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/api/v1/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/api/v1/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/api/v1/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/api/v1/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/api/v1/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Region routes
	s.Router.HandleFunc("/api/v1/regions",middlewares.SetMiddlewareJSON(s.CreateRegion)).Methods("POST")
	s.Router.HandleFunc("/api/v1/regions",middlewares.SetMiddlewareJSON(s.GetRegions)).Methods("GET")

	//Province routes
	s.Router.HandleFunc("/api/v1/regions/{rid}/provinces",middlewares.SetMiddlewareJSON(s.GetProvinces)).Methods("GET")
	
}