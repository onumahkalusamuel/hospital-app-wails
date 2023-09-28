package internal

import (
	"hospital-app/config"

	"hospital-app/frontend"
	"hospital-app/internal/api"
	"hospital-app/internal/api/delivery"
	"hospital-app/internal/api/invoice"
	"hospital-app/internal/api/patient"
	"hospital-app/internal/api/payment"
	"hospital-app/internal/api/staff"
	"hospital-app/internal/db"
	"hospital-app/internal/middleware"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func WebServer() {

	// get router
	e := echo.New()
	e.Use(echomiddleware.CORS())
	e.Static("files/", config.FilesFolder)

	// setup the database
	db.Init()
	db.InitialData()
	// set route for web files
	frontend.RegisterHandlers(e)

	unauthApi := e.Group("/api", middleware.ActivationAndInitialSetup())
	unauthApi.POST("/login", api.Login)
	unauthApi.POST("/activate", api.Activate)
	unauthApi.POST("/create-admin", api.CreateAdmin)
	unauthApi.POST("/hospital-setup", api.HospitalSetup)
	unauthApi.POST("/hospital-update", api.HospitalUpdate)
	unauthApi.GET("/hospital-details", api.HospitalDetails)
	unauthApi.GET("/installation-code", api.InstallationCode)

	unauthApi.POST("/get-activation-code", api.GetActivationCode) // move endpoint to server when done

	unauthApi.GET("/get-remote-address", api.GetRemoteAddress)

	authApi := e.Group("/api", middleware.ActivationAndInitialSetup(), echojwt.JWT([]byte(config.JwtSecret)), middleware.AppendJwtClaims())
	// profile actions
	authApi.GET("/profile", staff.ReadProfile)
	authApi.PUT("/profile", staff.UpdateProfile)
	// staff crud
	authApi.GET("/staff", staff.ReadAll)
	authApi.POST("/staff", staff.Create)
	authApi.GET("/staff/:id", staff.ReadOne)
	authApi.PUT("/staff/:id", staff.Update)
	authApi.DELETE("/staff/:id", staff.Delete)

	// patients crud and actions
	authApi.GET("/patients", patient.ReadAll)
	authApi.POST("/patients", patient.Create)
	authApi.GET("/patients/:id", patient.ReadOne)
	authApi.PUT("/patients/:id", patient.Update)
	authApi.DELETE("/patients/:id", patient.Delete)
	// patient history
	authApi.POST("/patients/anc-patient", patient.ReadAllUnpaginated)
	authApi.GET("/patients/:patient_id/patient-history", patient.PatientHistory)
	authApi.POST("/patients/:patient_id/patient-history", patient.PatientHistoryCreate)
	authApi.POST("/patients/:patient_id/patient-history/all", patient.PatientHistoryAll)
	authApi.GET("/patients/:patient_id/patient-history/:history_id", patient.PatientHistoryReadOne)
	authApi.PUT("/patients/:patient_id/patient-history/:history_id", patient.PatientHistoryUpdate)
	authApi.DELETE("/patients/:patient_id/patient-history/:history_id", patient.PatientHistoryDelete)

	authApi.POST("/patients/:patient_id/admit-patient", patient.AdmitPatient)
	authApi.POST("/patients/:patient_id/discharge-patient", patient.DischargePatient)
	authApi.POST("/patients/:patient_id/initiate-appointment", patient.InitiateAppointment)

	// patient invoices
	authApi.GET("/patients/:patient_id/invoices", patient.PatientInvoice)
	// deliveries crud and actions
	authApi.POST("/deliveries/all", delivery.ReadAllUnpaginated)
	authApi.GET("/deliveries", delivery.ReadAll)
	authApi.POST("/deliveries", delivery.Create)
	authApi.GET("/deliveries/:id", delivery.ReadOne)
	authApi.PUT("/deliveries/:id", delivery.Update)
	authApi.DELETE("/deliveries/:id", delivery.Delete)
	// invoices crud and actions
	authApi.GET("/invoices", invoice.ReadAll)
	authApi.POST("/invoices", invoice.Create)
	authApi.GET("/invoices/:id", invoice.ReadOne)
	authApi.PUT("/invoices/:id", invoice.Update)
	authApi.DELETE("/invoices/:id", invoice.Delete)
	// payments crud and actions
	authApi.POST("/invoices/:invoice_id/payments", payment.Create)
	authApi.DELETE("/invoices/:invoice_id/payments/:id", payment.Delete)

	e.Logger.Fatal(e.StartServer(config.SERVER_HANDLE))
	// e.Logger.Fatal(e.Start(net.JoinHostPort("0.0.0.0", config.ServerPort)))
}
