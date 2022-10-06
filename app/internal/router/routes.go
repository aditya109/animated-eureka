package router

import (
	"net/http"

	h "github.com/aditya109/animated-eureka/internal/handlers"
	"github.com/go-openapi/runtime/middleware"
)

// Route is container template for server routes
type Route struct {
	Name            string
	Method          string
	Pattern         string
	HandlerFunction http.HandlerFunc
	Handler         http.Handler
}

type routes []Route

var routeList = routes{

	// swagger:route GET / home welcome
	// Returns a welcome message
	// responses:
	// 	200: WelcomeResponse
	Route{
		"Welcome",
		"GET",
		"/",
		h.WelcomeHandler,
		nil,
	},

	// swagger:route GET /docs docs docs
	// Returns swagger specification uunder OpenAPIv3 documeted APIs
	Route{
		"swaggerDocumentation",
		"GET",
		"/docs",
		nil,
		middleware.Redoc(middleware.RedocOpts{
			SpecURL: "/swagger.yaml",
		}, nil),
	},

	Route{
		"Swagger JSON",
		"GET",
		"/swagger.yaml",
		nil,
		http.FileServer(http.Dir("./api/swagger")),
	},


	// swagger:route POST /virtualbond virtualBond virtualbond
	// Returns a virtual bond id, with corresponding a-faction-id, b-faction-id, uid, bond-end-time
	// responses:
	// 	200: PostGetVidResponse
	Route{
		"PostVirtualBond",
		"POST",
		"/virtualbond",
		h.PostVirtualBondHandler,
		nil,
	},

	// swagger:route POST /bulkvirtualbonds addBulkVirtualBond bulkvirtualbonds
	// Uploads bulk of virtual bonds in available_virtual_bond table
	// responses:
	// 	200: PostBulkVirtualBondsResponse
	Route{
		"PostBulkVirtualBonds",
		"POST",
		"/bulkvirtualbonds",
		h.PostBulkVirtualBondsHandler,
		nil,
	},
}
