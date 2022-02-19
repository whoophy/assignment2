// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Arif Kurniawan",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/order/{id}": {
            "put": {
                "description": "get details of all orders",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "get details list of all orders",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "GetOrders"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete orders by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "delete selected orders",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/orders": {
            "get": {
                "description": "get details of all orders",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "get details list of all orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "GetOrder"
                        }
                    }
                }
            },
            "post": {
                "description": "get details of all orders",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "get details list of all orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.Order"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "structs.Item": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "A random description"
                },
                "itemCode": {
                    "type": "string",
                    "example": "AAA321"
                },
                "itemId": {
                    "type": "integer",
                    "example": 1
                },
                "orderId": {
                    "type": "integer",
                    "example": 1
                },
                "quantity": {
                    "type": "integer",
                    "example": 10
                }
            }
        },
        "structs.Order": {
            "type": "object",
            "properties": {
                "costumerName": {
                    "type": "string",
                    "example": "Arif Kurniawa"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/structs.Item"
                    }
                },
                "orderAt": {
                    "type": "string",
                    "example": "2022-02-17T21:16:38.147532+07:00"
                },
                "orderId": {
                    "type": "integer",
                    "example": 1
                }
            }
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "1.0",
	Host:             "petstore.swagger.io",
	BasePath:         "/v2",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is Orders API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}