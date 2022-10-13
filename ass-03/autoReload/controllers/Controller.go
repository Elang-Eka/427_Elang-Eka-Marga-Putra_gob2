package controllers

// import package
import (
	"autoReload/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// struct data
type Data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func (idb *InDB) PostData(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
	// init variabel
	var (
		reload models.AutoReload
		result gin.H
	)

	// form post on postman
	water := c.PostForm("water")
	wind := c.PostForm("wind")

	intWater, _ := strconv.Atoi(water)
	intWind, _ := strconv.Atoi(wind)

	reload.Water = intWater
	reload.Wind = intWind
	dataStruct := Data{
		Water: reload.Water,
		Wind:  reload.Wind,
	}

	err := idb.DB.Create(&reload).Error
	if err != nil {
		result = gin.H{
			"result": "Create Data Not Successfully",
		}
		c.JSON(http.StatusNoContent, result)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": dataStruct,
		})
	}
}

// get all data
func (idb *InDB) Getdatas(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
	// init variabel
	var (
		reload2 []models.AutoReload
		result  gin.H
	)
	err := idb.DB.Order("id desc").Find(&reload2).Error
	// err := idb.DB.Find(&reload2).Error
	if err != nil {
		result = gin.H{
			"message": "data not found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {
		// response get found order
		result = gin.H{
			"status": reload2,
		}
		c.JSON(http.StatusOK, result)
	}
}

// get 1 data
func (idb *InDB) Getdata(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
	// init variabel
	var (
		reload  models.AutoReload
		reload2 []models.AutoReload
		result  gin.H
	)
	// find all data
	idb.DB.Find(&reload2).Take(&reload2).Last(&reload2)
	// if reload zero/0 response message
	if len(reload2) <= 0 {
		result = gin.H{
			"message": "data not found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {
		idb.DB.Find(&reload).Take(&reload).Last(&reload)
		dataStruct := Data{
			Water: reload.Water,
			Wind:  reload.Wind,
		}
		var msg1, msg2 string
		if dataStruct.Water <= 5 {
			msg1 = "Aman"
		} else if dataStruct.Water >= 6 && dataStruct.Water <= 8 {
			msg1 = "Siaga"
		} else if dataStruct.Water > 8 {
			msg1 = "Bahaya"
		}
		if dataStruct.Wind <= 6 {
			msg2 = "Aman"
		} else if dataStruct.Wind >= 7 && dataStruct.Wind <= 15 {
			msg2 = "Siaga"
		} else if dataStruct.Wind > 15 {
			msg2 = "Bahaya"
		}

		// response get found order
		result = gin.H{
			"status":       dataStruct,
			"MessageWater": msg1,
			"MessageWind":  msg2,
		}
		c.JSON(http.StatusOK, result)
	}

}

// Update Data
func (idb *InDB) UpdateData(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "PUT")
	// req param :orderId
	id := c.Param("id")

	// variable init
	var (
		reload  models.AutoReload
		newData models.AutoReload
		result  gin.H
	)

	// check id order
	err := idb.DB.Where("id = ?", id).First(&reload).Error
	if err != nil {
		// response data order not found
		result = gin.H{
			"result": "data not found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {

		// save data to new data variable
		DWater := c.PostForm("water")
		DWind := c.PostForm("wind")
		intWater, _ := strconv.Atoi(DWater)
		intWind, _ := strconv.Atoi(DWind)
		newData.Water = intWater
		newData.Wind = intWind

		// array of struct
		dataStruct := Data{
			Water: newData.Water,
			Wind:  newData.Wind,
		}

		// update data
		err = idb.DB.Model(&reload).Updates(newData).Error
		if err != nil {

			// response update failed
			result = gin.H{
				"result": "Update Failed",
			}
			c.JSON(http.StatusNoContent, result)
		} else {

			// response update success
			result = gin.H{
				"CodeStatus": http.StatusOK,
				"result":     "successfully updated data",
				"status":     dataStruct,
			}
			c.JSON(http.StatusOK, result)
		}
	}
}
