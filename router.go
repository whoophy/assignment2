package main

// // getOrders godoc
// // @Summary      get details list of all orders
// // @Description  get details of all orders
// // @Tags         orders
// // @Accept       json
// // @Produce      json
// // @Success      200  {object}  structs.Order
// // @Router       /orders [get]
// func GetOrder(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application//json")
// 	currentTime := time.Now()
// 	resp := structs.Order{
// 		OrderID:      1,
// 		CostumerName: "Arif Kurniawan",
// 		OrderAt:      currentTime,
// 		Items:        nil,
// 	}
// 	res, err := json.Marshal(&resp)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }
