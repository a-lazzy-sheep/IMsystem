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
        "/index": {
            "get": {
                "tags": [
                    "首页"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/searchFriends": {
            "post": {
                "tags": [
                    "searchfriends"
                ],
                "summary": "search friends",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of the user",
                        "name": "userId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/DeleteUser": {
            "delete": {
                "tags": [
                    "DeleteUser"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the user",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/GetUserList": {
            "get": {
                "tags": [
                    "getlist"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/Login": {
            "post": {
                "tags": [
                    "Login"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email of the user",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password of the user",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/Register": {
            "post": {
                "tags": [
                    "register"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email of the user",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name of the user",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password of the user",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/UpdateUser": {
            "put": {
                "tags": [
                    "UpdateUser"
                ],
                "summary": "Update a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the user",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name of the user",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Email of the user",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password of the user",
                        "name": "password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Phone of the user",
                        "name": "phone",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
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
