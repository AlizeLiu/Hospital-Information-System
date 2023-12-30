package main

import (
	"backend/Patient"
	"github.com/gin-gonic/gin"
)

func collectRoute(r *gin.Engine) *gin.Engine {
	r.GET("/patient/addPatient", Patient.PatientRegister)
	r.GET("/admin/login", Patient.Login)

	return r
}
