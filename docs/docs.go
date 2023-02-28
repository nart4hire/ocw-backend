// Code generated by swaggo/swag. DO NOT EDIT
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
        "/": {
            "get": {
                "description": "Give server index page response",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "common"
                ],
                "summary": "Index page",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.BaseResponse"
                        }
                    }
                }
            }
        },
        "/admin/user": {
            "get": {
                "description": "Get all users from database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Get All User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.BaseResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a user to database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Add User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.BaseResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a user from database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Delete User By Id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.BaseResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update a user from database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Update User By Id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.BaseResponse"
                        }
                    }
                }
            }
        },
        "/admin/user/{id}": {
            "get": {
                "description": "Get a user from database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Get User By Email",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.BaseResponse"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Login and generate new pair of token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login payload",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/login.LoginRequestPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/web.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/login.LoginResponsePayload"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Input",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/web.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "403": {
                        "description": "Login Credential Error",
                        "schema": {
                            "$ref": "#/definitions/web.BaseResponse"
                        }
                    },
                    "415": {
                        "description": "Not a json request",
                        "schema": {
                            "$ref": "#/definitions/web.BaseResponse"
                        }
                    },
                    "422": {
                        "description": "Invalid JSON input",
                        "schema": {
                            "$ref": "#/definitions/web.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "Unknown Internal Error",
                        "schema": {
                            "$ref": "#/definitions/web.BaseResponse"
                        }
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "description": "Generate new access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Refresh Token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Refresh token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/web.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/refresh.RefreshResponsePayload"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.BaseResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/web.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.BaseResponse"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Do Email Verification",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Email Verification",
                "parameters": [
                    {
                        "description": "Register Payload",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/verification.VerificationRequestPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.BaseResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.BaseResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "login.LoginRequestPayload": {
            "description": "Information that should be available when do a login process",
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "description": "User Email",
                    "type": "string",
                    "example": "someone@example.com"
                },
                "password": {
                    "description": "User Password",
                    "type": "string",
                    "example": "secret"
                }
            }
        },
        "login.LoginResponsePayload": {
            "description": "Login response when process success",
            "type": "object",
            "properties": {
                "access_token": {
                    "description": "Token that used to access the resources",
                    "type": "string"
                },
                "refresh_token": {
                    "description": "Token that used to generate new access token",
                    "type": "string"
                }
            }
        },
        "refresh.RefreshResponsePayload": {
            "description": "Refresh endpoint response when process success",
            "type": "object",
            "properties": {
                "access_token": {
                    "description": "Token that used to access the resources",
                    "type": "string"
                }
            }
        },
        "register.RegisterRequestPayload": {
            "description": "Information that should be available when do a registration process",
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "password_validation"
            ],
            "properties": {
                "email": {
                    "description": "User Email",
                    "type": "string",
                    "example": "someone@example.com"
                },
                "name": {
                    "description": "User name",
                    "type": "string",
                    "example": "someone"
                },
                "password": {
                    "description": "User Password",
                    "type": "string",
                    "example": "secret"
                },
                "password_validation": {
                    "description": "User Password Validation, must be same as user",
                    "type": "string",
                    "example": "secret"
                }
            }
        },
        "verification.VerificationRequestPayload": {
            "description": "Information that should be passed when request verify",
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "description": "User Email",
                    "type": "string",
                    "example": "someone@example.com"
                }
            }
        },
        "web.BaseResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string",
                    "enum": [
                        "success",
                        "failed"
                    ]
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Open Courseware Application",
	Description:      "This is Open Couseware backend",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
