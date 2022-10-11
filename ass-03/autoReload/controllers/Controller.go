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

// get data
func (idb *InDB) Getdata(c *gin.Context) {

	// init variabel
	var (
		reload models.AutoReload
		result gin.H
	)
	// req param :orderId
	id := c.Param("id")

	err := idb.DB.Where("id = ?", id).First(&reload).Error
	if err != nil {
		// response data item not found
		result = gin.H{
			"result": "Data not found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {
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
		c.JSON(http.StatusOK, gin.H{
			"messageWater": msg1,
			"messageWind":  msg2,
			"status":       dataStruct,
		})
	}
}

func (idb *InDB) UpdateData(c *gin.Context) {

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
