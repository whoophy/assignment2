package main

// import (
// 	"encoding/json"
// 	"net/http"
// )

// type JSONSuccessResult struct {
// 	Code    int         `json:"code" example:"200"`
// 	Message string      `json:"message" example:"Success"`
// 	Data    interface{} `json:"data"`
// }

// type JSONbadRequest struct {
// 	Code    int         `json:"code" example:"400"`
// 	Message string      `json:"message" example:"Wrong Parameter"`
// 	Data    interface{} `json:"data"`
// }

// func SuccessResponse(w http.ResponseWriter, data interface{}) {
// 	json.NewEncoder(w).Encode(JSONSuccessResult{
// 		Code:    200,
// 		Message: "Success get data",
// 		Data:    data,
// 	})
// 	return
// }
