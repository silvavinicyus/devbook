package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		Uri:                    "/posts",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePost,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/posts",
		Method:                 http.MethodGet,
		Function:               controllers.FindPosts,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/posts/{id}",
		Method:                 http.MethodGet,
		Function:               controllers.FindPost,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/posts/{id}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePost,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/posts/{id}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletePost,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/posts/users/{id}",
		Method:                 http.MethodGet,
		Function:               controllers.FindUserPosts,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/posts/{id}/like",
		Method:                 http.MethodPost,
		Function:               controllers.LikePost,
		RequiresAuthentication: true,
	},
	{
		Uri:                    "/posts/{id}/unlike",
		Method:                 http.MethodPost,
		Function:               controllers.UnlikePost,
		RequiresAuthentication: true,
	},
}
