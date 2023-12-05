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
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/users/{id}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/users/{id}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/users/{id}/follow",
		Method:                 http.MethodPost,
		Function:               controllers.FollowUser,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/users/{id}/unfollow",
		Method:                 http.MethodPost,
		Function:               controllers.UnfollowUser,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/users/{id}/followers",
		Method:                 http.MethodGet,
		Function:               controllers.FindFollowers,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/users/{id}/following",
		Method:                 http.MethodGet,
		Function:               controllers.FindFollowing,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/users/{id}/password",
		Method:                 http.MethodPost,
		Function:               controllers.UpdatePassword,
		RequiresAuthentication: true,
	},
}
