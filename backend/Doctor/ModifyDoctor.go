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

	// 从请求体中获取医生ID
	dId := ctx.PostForm("dId")

	// 使用 DID 查询医生信息
	var doctor model.Doctor
	if err := DB.First(&doctor, "DID = ?", dId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}
	doctor.DGender = ctx.PostForm("dGender")
	doctor.DName = ctx.PostForm("dName")
	doctor.DPost = ctx.PostForm("dPost")
	doctor.DSection = ctx.PostForm("dSection")
	doctor.DPhone, _ = strconv.Atoi(ctx.PostForm("dPhone"))
	doctor.DEmail = ctx.PostForm("dEmail")
	dCard, _ := strconv.Atoi(ctx.PostForm("dCard"))
	dPrice, _ := strconv.Atoi(ctx.PostForm("dPrice"))
	doctor.DCard = dCard
	doctor.DPrice = dPrice
	doctor.DIntroduction = ctx.PostForm("dIntroduction")
	dState, _ := strconv.Atoi(ctx.PostForm("dState"))
	doctor.DState = dState

	if err := DB.Save(&doctor).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update doctor"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Doctor information updated successfully", "doctor": doctor})
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

	eachPart := int(math.Ceil(float64(totalDNum) / 5.0))
	part1 := eachPart
	part2 := eachPart
	part3 := eachPart
	part4 := eachPart
	part5 := totalDNum - (eachPart * 4)

	ctx.JSON(http.StatusOK, gin.H{
		"eTOn": part1,
		"nTOt": part2,
		"tTOe": part3,
		"fTOf": part4,
		"sTOs": part5,
	})
}
