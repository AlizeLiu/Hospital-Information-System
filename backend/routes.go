package main

import (
	"backend/Admin"
	"backend/Doctor"
	"backend/Patient"
	"backend/Public"
	"github.com/gin-gonic/gin"
)

func collectRoute(r *gin.Engine) *gin.Engine {
	r.GET("/order/orderPeople", Public.OrderPeople)
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
	r.GET("/patient/addOrder", Patient.AddOrder)
	r.GET("/admin/addRole", Admin.AddRole)
	r.GET("/doctor/findOrderByNull", Doctor.OrderbyNull)
	r.GET("/drug/findAllDrugs", Admin.FindAllDrug)
	r.GET("/check/findAllChecks", Admin.FindAllCheck)
	r.GET("/drug/addDrug", Admin.AddDrugs)
	r.GET("/check/findCheck", Admin.AddChecks)
	r.GET("/drug/reduceDrugNumber", Admin.ReduceDrugs)
	return r
}
