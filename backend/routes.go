package main

import (
	"backend/Doctor"
	"backend/Patient"
	"backend/Public"
	"github.com/gin-gonic/gin"
)

func collectRoute(r *gin.Engine) *gin.Engine {
	r.GET("/patient/addPatient", Patient.PatientRegister)
	r.GET("/admin/login", Patient.Login)
	r.GET("/info", Public.GetLoginInf)
	r.GET("/login/getMenu", Public.GetMenu)
	r.GET("/doctor/findPatientById", Patient.FindPatientById)
	r.GET("/admin/addDoctor", Doctor.DoctorInit)
	r.GET("/admin/findAllDoctors", Doctor.FindAllDoctor)
	r.GET("/admin/deleteDoctor", Doctor.DeleteDoctor)
	r.GET("/admin/findDoctor", Doctor.GetFindDoctor)
	r.POST("/admin/modifyDoctor", Doctor.DoctorInfo)
	r.GET("/admin/findAllPatients", Patient.FindAllPatient)
	r.GET("/admin/deletePatient", Patient.DeletePatient)
	r.GET("/patient/findDoctorBySection", Doctor.FindDoctorBySection)
	r.GET("/arrange/findByTime", Doctor.FindByTime)
	r.GET("/order/findOrderTime", Doctor.FindOrderTime)
	return r
}
