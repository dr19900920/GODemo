package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"controllers"
	"logger"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {

		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"用户登录",
		"POST",
		"/login",
		controllers.Login,
	},
	Route{
		"学员工单统计数据",
		"GET",
		"/statistic/counts",
		controllers.Counts,
	},
	Route{
		"学员相关数据字典",
		"GET",
		"/dict/student",
		controllers.StudentDict,
	},
}


