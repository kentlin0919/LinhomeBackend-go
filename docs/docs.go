// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Skyler"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/Login": {
            "post": {
                "description": "使用者登入功能",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "登入",
                "parameters": [
                    {
                        "description": "使用者登入資訊",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/pojo.Account"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登入成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "請求失敗",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/Login1": {
            "post": {
                "description": "使用者登入功能",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "登入",
                "parameters": [
                    {
                        "type": "string",
                        "description": "使用者名稱",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密碼",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登入成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "請求失敗",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "pojo.Account": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "192.168.0.158:9000",
	BasePath:         "/v1",
	Schemes:          []string{"http"},
	Title:            "Linhome",
	Description:      "swagger test example",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
