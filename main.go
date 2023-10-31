// main.go

package main

import (
	// Import your configuration package
	"healthcare/controllers"
	"healthcare/models" //Import your models package

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize the Gin router
	router := gin.Default()

	//connect to database
	models.InitDB()

	// Set up application routes
	unauthenticated := router.Group("/api")
	{
		unauthenticated.POST("/login", controllers.Login)
		unauthenticated.POST("/signup", controllers.Signup)
	}

	{
		// Appointment management
		router.POST("/appointment/create", controllers.CreateAppointment)
		router.GET("/appointment/:id", controllers.GetAppointment)
		router.PUT("/appointment/approve/:id", controllers.ApproveAppointment)
		router.PUT("/appointment/reject/:id", controllers.RejectAppointment)
		router.PUT("/appointment/update/:id", controllers.UpdateAppointment)
		router.DELETE("/appointment/remove/:id", controllers.RemoveAppointment)
		router.GET("/appointment/status", controllers.GetAppointmentStatus)

		// Admin operations
		router.POST("/admin/diagnostic-center/add", controllers.AddDiagnosticCenter)
		router.PUT("/admin/diagnostic-center/modify/:id", controllers.ModifyDiagnosticCenter)
		router.GET("/admin/diagnostic-center/view/:id", controllers.ViewDiagnosticCenter)
		router.DELETE("/admin/diagnostic-center/remove/:id", controllers.RemoveDiagnosticCenter)
		router.POST("/admin/diagnostic-test/add", controllers.AddDiagnosticTest)
		router.PUT("/admin/diagnostic-test/modify/:id", controllers.ModifyDiagnosticTest)
		router.GET("/admin/diagnostic-test/view/:id", controllers.ViewDiagnosticTest)
		router.DELETE("/admin/diagnostic-test/remove/:id", controllers.RemoveDiagnosticTest)
	}
	// Start the server
	router.Run()
}
