package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:           "/users",
		Method:        http.MethodPost,
		Function:      controllers.CreateUser,
		Authorization: false,
	},
	{
		URI:           "/users",
		Method:        http.MethodGet,
		Function:      controllers.GetUsers,
		Authorization: false,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodGet,
		Function:      controllers.GetUser,
		Authorization: false,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodPut,
		Function:      controllers.UpdateUser,
		Authorization: false,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodDelete,
		Function:      controllers.DeleteUser,
		Authorization: false,
	},
}
