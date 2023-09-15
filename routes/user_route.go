package routes

import (
	"github.com/dastardlyjockey/hngx-2/controllers"
	"github.com/go-chi/chi/v5"
)

func UserRoute(route *chi.Mux) {
	route.Post("/", controllers.ApiCfg.CreateUser)
	route.Get("/{user_id}", controllers.ApiCfg.GetUserByID)
	route.Patch("/{user_id}", controllers.ApiCfg.UpdateUserNameById)
	route.Delete("/{user_id}", controllers.ApiCfg.DeleteUserById)
}
