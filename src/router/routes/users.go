package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		Uri:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		Uri:                    "/users",
		Method:                 http.MethodGet,
		Function:               controllers.FindUsers,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/users/{id}",
		Method:                 http.MethodGet,
		Function:               controllers.FindUser,
		RequiresAuthentication: false,
	},
	{
		Uri:                    "/users/{id}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		RequiresAuthentication: false,
	},
	{
		Uri:                    "/users/{id}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		RequiresAuthentication: false,
	},
}
