package Doctor

import (
	"backend/common"
	"backend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func DoctorInfo(ctx *gin.Context) { //修改医生信息

	DB := common.GetDB()

	// 使用 DID 查询医生信息
	var doctor model.DoctorData
	_ = ctx.ShouldBindJSON(&doctor)
	// 从请求体中获取医生ID

	dId := doctor.DID
	var existingDoctor model.Doctor
	if err := DB.First(&existingDoctor, "d_id = ?", dId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}

	originalDoctor := existingDoctor

	// Update doctor information
	existingDoctor.DID = doctor.DID
	existingDoctor.DGender = doctor.DGender
	existingDoctor.DName = doctor.DName
	existingDoctor.DPost = doctor.DPost
	existingDoctor.DSection = doctor.DSection
	existingDoctor.DIntroduction = doctor.DIntroduction
	existingDoctor.DPhone = doctor.DPhone
	existingDoctor.DEmail = doctor.DEmail

	// Convert and update integer fields
	existingDoctor.DCard, _ = strconv.Atoi(doctor.DCard)
	existingDoctor.DState = doctor.DState
	existingDoctor.DPrice, _ = strconv.Atoi(doctor.DPrice)

	// Check which fields have been updated
	updatedFields := make(map[string]interface{})
	if existingDoctor.DID != originalDoctor.DID {
		updatedFields["DID"] = existingDoctor.DID
	}

	if err := DB.Model(&originalDoctor).Updates(doctor).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, "22222")
		return
	}

	ctx.JSON(200, &gin.H{
		"code":  200,
		"error": "11111"})
	return

}

func FindDoctorBySection(ctx *gin.Context) {
	DB := common.GetDB()

	dSection := ctx.Query("dSection")

	var doctors []model.Doctor
	DB.Model(&model.Doctor{}).Where("d_section LIKE ?", dSection).
		Find(&doctors)

	var transformedDoctors []gin.H
	for _, doctor := range doctors {
		transformedDoctors = append(transformedDoctors, gin.H{
			"dPrice":   doctor.DPrice,
			"dName":    doctor.DName,
			"dPost":    doctor.DPost,
			"dId":      doctor.DID,
			"dAvgStar": int(5), // 注意：需要为 Doctor 结构体添加 DAvgStar 字段
			"dGender":  doctor.DGender,
			"dSection": doctor.DSection,
			"dNum":     doctor.DNum,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"doctors": transformedDoctors,
	})

}

func FindByTime(ctx *gin.Context) {
	DB := common.GetDB()

	dSection := ctx.Query("dSection")
	arTime := ctx.Query("arTime")

	// 查询医生信息
	var doctors []model.Doctor
	DB.Where("d_section = ?", dSection).Find(&doctors)

	// 构建响应数据
	var response []gin.H
	for _, doctor := range doctors {
		// 过滤不工作的医生
		nonWorkDays := strings.Split(doctor.DNonWorkDays, ",")
		isAvailable := true
		for _, day := range nonWorkDays {
			if day == getDayOfWeek(arTime) {
				isAvailable = false
				break
			}
		}
		if isAvailable {
			response = append(response, gin.H{
				"dId":      doctor.DID,
				"dName":    doctor.DName,
				"dGender":  doctor.DGender,
				"dPost":    doctor.DPost,
				"dSection": doctor.DSection,
				"dPrice":   doctor.DPrice,
				"dNum":     doctor.DNum,
			})
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"doctors": response,
	})
}

func getDayOfWeek(date string) string {
	layout := "2006-01-02"           // 定义日期格式
	t, _ := time.Parse(layout, date) // 解析日期
	return t.Weekday().String()
}

func FindOrderTime(ctx *gin.Context) {
	DB := common.GetDB()

	Did := ctx.Query("arId")

	var doctor model.Doctor
	if err := DB.Where("d_id = ?", Did).First(&doctor).Error; err != nil {
		fmt.Println("Failed to fetch doctor:", err)
		return
	}
	totalDNum := doctor.DNum

	eachPart := int(math.Ceil(float64(totalDNum) / 6.0))
	part1 := eachPart
	part2 := eachPart
	part3 := eachPart
	part4 := eachPart
	part5 := eachPart
	part6 := totalDNum - (eachPart * 5)

	ctx.JSON(http.StatusOK, gin.H{
		"eTOn": part1,
		"nTOt": part2,
		"tTOe": part3,
		"fTOf": part4,
		"fTOs": part5,
		"sTOs": part6,
	})
}

func OrderbyNull(ctx *gin.Context) {
	DB := common.GetDB()

	dId := ctx.Query("dId")
	//oStart := ctx.Query("oStart")

	var registrations []model.Registration

	// 查询当天的挂号信息
	if err := DB.Where("d_id = ? ", dId).Find(&registrations).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "查询挂号信息失败"})
		return
	}

	// 查询患者和医生信息
	var response []gin.H
	for _, reg := range registrations {
		var patient model.User
		var doctor model.Doctor

		// 查询患者信息
		if err := DB.Where("account = ?", reg.Account).First(&patient).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "查询患者信息失败"})
			return
		}

		// 查询医生信
		if err := DB.Where("d_id = ?", reg.DID).First(&doctor).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "查询医生信息失败"})
			return
		}

		oStart := reg.CreatedAt.Format("2006-01-02 15:04")
		response = append(response, gin.H{
			"oId":    reg.ID,
			"pName":  patient.Username,
			"dName":  doctor.DName,
			"oStart": oStart,
			"pId":    patient.Account,
		})
	}

	ctx.JSON(http.StatusOK, response)
}
