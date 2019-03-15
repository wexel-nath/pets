package api

import (
	"net/http"

	"github.com/wexel-nath/pets/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

func GetRouter() *httprouter.Router {
	router := httprouter.New()

	for _, path := range getPaths() {
		router.Handle(path.method, path.pattern, middlewareWrapper(path))
	}

	return router
}

type path struct {
	method  string
	pattern string
	handler httprouter.Handle
}

func getPaths() []path {
	return []path{
		{
			method:  http.MethodGet,
			pattern: "/healthz",
			handler: healthzHandler,
		},

		// Pet Routes
		{
			method:  http.MethodPost,
			pattern: "/pet",
			handler: addPetHandler,
		},
		{
			method:  http.MethodPut,
			pattern: "/pet",
			handler: updatePetHandler,
		},
		{
			method:  http.MethodGet,
			pattern: "/pet/:path",
			handler: petGetHandler,
		},
		{
			method:  http.MethodPost,
			pattern: "/pet/:petId",
			handler: updatePetWithFormDataHandler,
		},
		{
			method:  http.MethodDelete,
			pattern: "/pet/:petId",
			handler: deletePetHandler,
		},
		{
			method:  http.MethodPost,
			pattern: "/pet/:petId/uploadImage",
			handler: uploadPetImageHandler,
		},

		// Order Routes
		{
			method:  http.MethodGet,
			pattern: "/store/inventory",
			handler: getInventoryHandler,
		},
		{
			method:  http.MethodPost,
			pattern: "/store/order",
			handler: placeOrderHandler,
		},
		{
			method:  http.MethodGet,
			pattern: "/store/order/:orderId",
			handler: findOrderByIDHandler,
		},
		{
			method:  http.MethodDelete,
			pattern: "/store/order/:orderId",
			handler: deleteOrderHandler,
		},

		// User Routes
		{
			method:  http.MethodPost,
			pattern: "/user",
			handler: createUserHandler,
		},
		{
			method:  http.MethodPost,
			pattern: "/user/createWithArray",
			handler: createUserWithArrayHandler,
		},
		{
			method:  http.MethodPost,
			pattern: "/user/createWithList",
			handler: createUserWithListHandler,
		},
		{
			method:  http.MethodGet,
			pattern: "/user/:path",
			handler: userGetHandler,
		},
		{
			method:  http.MethodPut,
			pattern: "/user/:username",
			handler: updateUserHandler,
		},
		{
			method:  http.MethodDelete,
			pattern: "/user/:username",
			handler: deleteUserHandler,
		},
	}
}

func healthzHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_, err := w.Write([]byte(`{"status":"ok"}`))
	if err != nil {
		logger.Error(err)
	}
}
