package routes

import (
	"net/http"
	"vp_alp/controllers"

	"github.com/labstack/echo/v4"
)

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	name := c.Param("name")
	return c.String(http.StatusOK, "Hello, "+name)
}

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Gengs!")
	})

	e.GET("/user", func(c echo.Context) error {
		return c.String(http.StatusOK, "ini user page")
	})

	e.GET("user/:name", getUser)

	e.GET("/customer", controllers.FetchAllCustomer)

	e.GET("/customer-usernm/:username", controllers.GetCustomerByUsername)

	e.POST("/customer", controllers.StoreCustomer)

	e.PATCH("/customer", controllers.UpdateCustomer)

	// e.DELETE("/customer/:c_id", controllers.DeleteCustomer)

	e.PUT("/customer/:c_id", controllers.DeleteCustomer)

	e.GET("/technician-all", controllers.FetchAllTechnician)

	e.GET("/technician/:name", controllers.GetTechnicianByName)

	e.GET("/technician-location/:k_id", controllers.GetTechnicianByLocation)

	e.GET("/technician-order/:t_id/:status", controllers.GetTechnicianOrder)

	e.GET("/customer-order/:c_id/:status", controllers.GetCustomerOrder)

	e.POST("/technician", controllers.StoreTechnician)

	e.PATCH("/technician", controllers.UpdateTechnician)

	e.PATCH("/technician-rate/:t_id/:rate", controllers.UpdateTechnicianRate)

	e.PUT("/technician/:t_id", controllers.DeleteTechnician)

	// e.DELETE("/technician/:t_id", controllers.DeleteTechnician)

	e.GET("/order", controllers.FetchAllOrder)

	e.POST("/order", controllers.StoreOrder)

	e.PATCH("/order", controllers.UpdateOrder)

	e.POST("/login-technician", controllers.CheckLoginTechnician)

	e.POST("/login-customer", controllers.CheckLoginCustomer)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)

	e.GET("/order-by-id/:o_id", controllers.GetOrderByID)

	e.GET("/technician-by-id/:t_id", controllers.GetTechnicianByID)

	e.GET("/customer-by-id/:c_id", controllers.GetCustomerByID)

	e.GET("/kecamatan", controllers.FetchAllKecamatan)

	e.GET("/wilayah", controllers.FetchAllWiayah)

	return e
}
