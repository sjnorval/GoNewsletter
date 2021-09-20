package controllers

func (server *Server) initializeRoutes() {

	//Users routes
	server.Router.POST("/users", server.CreateUser)
	server.Router.GET("/users",server.GetUsers)
	server.Router.GET("/users/{id}", server.GetUser)
	server.Router.PUT("/users/{id}", server.UpdateUser)
	//server.Router.DELETE("/users/{id}", server.DeleteUser)

	//Posts routes
	server.Router.POST("/posts", server.CreatePost)
	server.Router.GET("/posts", server.GetPosts)
	server.Router.GET("/posts/{id}", server.GetPost)

	//Topics routes
	server.Router.POST("/topics", server.CreateTopic)
	server.Router.GET("/topics", server.GetTopics)
	server.Router.GET("/topics/{id}", server.GetTopic)
	server.Router.POST("/topics/{topicid}/User/{userid}", server.RegisterUserToTopic)
}