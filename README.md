# RESTAPI Order
## Aplikasi CRUD sederhana

Query yang saya gunakan untuk Membuat table order
```sh
CREATE TABLE orders(
order_id SERIAL PRIMARY KEY,
costumer_name varchar(50) NOT NULL,
ordered_at TIMESTAMP WITH TIME ZONE NOT NULL
);
```
Query yang saya gunakan untuk Membuat table Items
```sh
CREATE TABLE items(
item_id SERIAL PRIMARY KEY,
item_code varchar(30) NOT NULL,
description TEXT NOT NULL,
quantity INT NOT NULL,
order_id INT,
CONSTRAINT fk_orders FOREIGN KEY(order_id) REFERENCES orders(order_id)
);
```


## Program
Semua fungsi CRUD berjalan secara semestinya, test saya lampirkan pada postman
## BUG
- Pada saat parsing data dari json ke struct, salah satu field pada struct tidak terbaca sehingga saya memberikan default yang ditunjukan pada
-> order.CostumerName = "Hallo" //pada fungsi UpdateOrder di main.go
-> order.CostumerName = "template" //pada fungsi CreateOrder di main.go
- Swagger baru bisa yang getOrder, namun seluruh test sudah bisa berjalan dengan baik di postman

