package main

import (
	"assignment/controllers"
	_ "assignment/docs"
	"assignment/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

type JSONSuccessResult struct {
	Code   int         `json:"code" example:"200"`
	Status string      `json:"status" example:"Success"`
	Data   interface{} `json:"data"`
}

// type JSONbadRequest struct {
// 	Code    int         `json:"code" example:"400"`
// 	Message string      `json:"message" example:"Wrong Parameter"`
// 	Data    interface{} `json:"data"`
// }

func SuccessResponse(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(JSONSuccessResult{
		Code:   200,
		Status: "Success",
		Data:   data,
	})
	return
}

type ResponseMessage struct {
	Message string `json:"message" example:"berhasil menambahkan data"`
}

// getOrders godoc
// @Summary      get details list of all orders
// @Description  get details of all orders
// @Tags         orders
// @Accept       json
// @Produce      json
// @Success      200  {object}  JSONSuccessResult{data=structs.Order,code=int,status=string}
// @Router       /orders [get]
func GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orders, err := controllers.GetOrders()

	if err != nil {
		panic(err)
	}
	// json.NewEncoder(w).Encode(orders)
	SuccessResponse(w, orders)
}

// getOrders godoc
// @Summary      get details list of all orders
// @Description  get details of all orders
// @Tags         orders
// @Accept       json
// @Produce      json
// @Success      200 {object}  JSONSuccessResult{data=ResponseMessage,code=int,status=string}
// @Param        id   path      int  true  "Order ID"
// @Param        order body     structs.Order  true  "json data"
// @Router       /order/{id} [put]
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	order_id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	var order structs.Order
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &order)
	if err != nil {
		panic(err)
	}
	order.CostumerName = "Hallo"
	msg := controllers.UpdateOrder(order, order_id)
	fmt.Println(msg)
	var response ResponseMessage
	response.Message = msg
	SuccessResponse(w, response)
}

// DeleteOrders godoc
// @Summary      delete selected orders
// @Description  delete orders by ID
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Order ID"
// @Success      200 {object}  JSONSuccessResult{data=ResponseMessage,code=int,status=string}
// @Router       /order/{id} [delete]
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	order_id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	msg, err := controllers.DeleteOrder(order_id)
	if err != nil {
		panic(err)
	}
	var response ResponseMessage
	response.Message = msg
	SuccessResponse(w, response)
}

// getOrders godoc
// @Summary      get details list of all orders
// @Description  get details of all orders
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        order body     structs.Order  true  "json data"
// @Success      200 {object}  JSONSuccessResult{data=ResponseMessage,code=int,status=string}
// @Router       /orders [post]
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order structs.Order
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &order)
	if err != nil {
		panic(err)
	}
	order.CostumerName = "template"
	msg := controllers.CreateOrder(order)
	fmt.Println(msg)
	var response ResponseMessage
	response.Message = msg
	SuccessResponse(w, response)
}

// @title Swagger Example API
// @version 1.0
// @description This is Orders API
// @termsOfService http://swagger.io/terms/

// @contact.name Arif Kurniawan
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/orders", CreateOrder).Methods("POST", "OPTIONS")
	r.HandleFunc("/order/{id}", DeleteOrder).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/orders", GetOrder).Methods("GET", "OPTIONS")
	r.HandleFunc("/order/{id}", UpdateOrder).Methods("PUT", "OPTIONS")
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	fmt.Println("program sedang berjalan")
	log.Fatal(http.ListenAndServe(":8080", r))
}
