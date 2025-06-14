// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "登录获取 jwt\ncode == 1102 , 需刷新 jwt;\ncode == 1200 , 需重新登录后跳转;\ncode == 1101 , 再次请求; (基本不需要)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "认证"
                ],
                "summary": "登录接口",
                "parameters": [
                    {
                        "description": "登录信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.authInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/auth/refresh_token": {
            "post": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "刷新 jwt",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "认证"
                ],
                "summary": "刷新token接口",
                "parameters": [
                    {
                        "description": "刷新tokenn信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.refreshInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/user/": {
            "get": {
                "description": "用户模拟接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户接口",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/user/2": {
            "get": {
                "description": "模拟panic",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户接口2",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.authInfo": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "auth.refreshInfo": {
            "type": "object",
            "required": [
                "refresh_jwt"
            ],
            "properties": {
                "refresh_jwt": {
                    "type": "string",
                    "minLength": 1
                }
            }
        }
    },
    "securityDefinitions": {
        "OAuth2Password": {
            "description": "OAuth protects our entity endpoints",
            "type": "oauth2",
            "flow": "password",
            "tokenUrl": "/api/v1.0/auth/login"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
