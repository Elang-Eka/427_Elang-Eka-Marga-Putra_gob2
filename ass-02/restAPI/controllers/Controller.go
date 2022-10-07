package controllers

// import package
import (
	"net/http"
	"restapi/structs"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// struct data
type Data struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func (idb *InDB) CreateOrders(c *gin.Context) {

	// init variabel
	var (
		order  structs.Order
		item   structs.Item
		result gin.H
	)

	// form post on postman
	customer_name := c.PostForm("customer_name")
	item_code := c.PostForm("item_code")
	description := c.PostForm("description")
	stringQty := c.PostForm("quantity")

	// convert string to int
	quantity, _ := strconv.Atoi(stringQty)

	// save data on variable
	order.Ordered_at = time.Now()
	order.Customer_Name = customer_name
	item.Item_Code = item_code
	item.Description = description
	item.Quantity = quantity

	// array of struct
	var data []Data
	dataStruct := Data{
		ItemCode:    item.Item_Code,
		Description: item.Description,
		Quantity:    quantity,
	}
	// add object to array
	data = append(data, dataStruct)

	// save to database order
	err := idb.DB.Create(&order).Error
	if err != nil {
		result = gin.H{
			"result": "Create Data Not Successfully",
		}
		c.JSON(http.StatusNoContent, result)
	} else {
		item.Order_Id = order.ID

		// save to database item
		err = idb.DB.Create(&item).Error
		if err != nil {
			result = gin.H{
				"result": "Create Data Not Successfully",
			}
			c.JSON(http.StatusNoContent, result)
		} else {

			// response success
			c.JSON(http.StatusOK, gin.H{
				"orderId":       order.ID,
				"orderedAt":     order.Ordered_at,
				"customer_name": order.Customer_Name,
				"items":         data,
			})
		}
	}
}

// get all data order
func (idb *InDB) GetOrders(c *gin.Context) {

	// init variabel
	var (
		orders []structs.Order
		result gin.H
	)
	// find all orders
	idb.DB.Find(&orders)

	// if orders zero/0 response message
	if len(orders) <= 0 {
		result = gin.H{
			"result":  nil,
			"count":   0,
			"message": "data not found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {

		// response get found order
		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
		c.JSON(http.StatusOK, result)
	}
}

// update order by orderId
func (idb *InDB) UpdateOrder(c *gin.Context) {

	// req param :orderId
	id := c.Param("orderId")

	// struct data2
	type Data2 struct {
		LineItemId  int    `json:"lineItemId"`
		ItemCode    string `json:"itemCode"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
	}

	// variable init
	var (
		order    structs.Order
		item     structs.Item
		newOrder structs.Order
		result   gin.H
	)

	// check id order
	err := idb.DB.Where("id = ?", id).First(&order).Error
	if err != nil {

		// response data order not found
		result = gin.H{
			"result": "data order not found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {

		// check order_id on item
		err = idb.DB.Where("order_id = ?", id).First(&item).Error
		if err != nil {

			// response data item not found
			result = gin.H{
				"result": "data item not found",
			}
			c.JSON(http.StatusNotFound, result)

		} else {

			// save data to new data variable
			orderid, _ := strconv.Atoi(id)
			customer_name := c.PostForm("customer_name")
			newOrder.ID = orderid
			newOrder.Customer_Name = customer_name
			newOrder.Ordered_at = order.Ordered_at

			// array of struct
			var data []Data2
			dataStruct := Data2{
				LineItemId:  item.ID,
				ItemCode:    item.Item_Code,
				Description: item.Description,
				Quantity:    item.Quantity,
			}

			// add object to array
			data = append(data, dataStruct)

			// update data
			err = idb.DB.Model(&order).Updates(newOrder).Error
			if err != nil {

				// response update failed
				result = gin.H{
					"result": "Update Failed",
				}
				c.JSON(http.StatusNoContent, result)
			} else {

				// response update success
				result = gin.H{
					"result":        "successfully updated data",
					"orderId":       order.ID,
					"orderedAt":     order.Ordered_at,
					"customer_name": order.Customer_Name,
					"data":          data,
				}
				c.JSON(http.StatusOK, result)
			}
		}
	}

}

// delete data with {id}
func (idb *InDB) DeleteOrder(c *gin.Context) {

	// init variable
	var (
		order  structs.Order
		item   structs.Item
		result gin.H
	)

	// param orderId
	orderId := c.Param("orderId")

	// querry check id on order
	err := idb.DB.Where("id = ?", orderId).First(&order).Error
	if err != nil {

		// response data order not found
		result = gin.H{
			"result": "data order not found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {

		// querry check order_id on item
		err = idb.DB.Where("order_id = ?", orderId).First(&item).Error
		if err != nil {

			// response data item not found
			result = gin.H{
				"result": "data item not found",
			}
			c.JSON(http.StatusNotFound, result)
		} else {

			// Delete data on db, table item
			err = idb.DB.Delete(&item).Error
			if err != nil {

				// response delete data item not success
				result = gin.H{
					"result": "Delete item Failed",
				}
				c.JSON(http.StatusBadRequest, result)
			} else {

				// Delete data on db, table order
				err = idb.DB.Delete(&order).Error
				if err != nil {

					// response delete data order not success
					result = gin.H{
						"result": "Delete order Failed",
					}
					c.JSON(http.StatusBadRequest, result)
				} else {

					// request delete data success
					result = gin.H{
						"result": "data deleted successfully",
						"data":   order,
					}
					c.JSON(http.StatusOK, result)
				}
			}
		}
	}
}
