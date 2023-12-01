package routes

import (
	"api/src/controllers"
	"net/http"
)

var rotaLogin = Route{
	Uri:                    "/login",
	Method:                 http.MethodPost,
	Function:               controllers.Login,
	RequiresAuthentication: false,
}
