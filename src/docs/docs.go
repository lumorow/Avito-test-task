// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/segment": {
            "post": {
                "description": "Создание нового сегмента",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Segment"
                ],
                "parameters": [
                    {
                        "description": "Данные сегмента",
                        "name": "segment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SegmentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SegmentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/segment/{slug}": {
            "delete": {
                "description": "Удаление существующего сегмента",
                "tags": [
                    "Segment"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Slug сегмента",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/{uid}": {
            "delete": {
                "description": "Удаление пользователя",
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID пользователя",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/{uid}/segments": {
            "get": {
                "description": "Получение списка сегментов, в которые входит пользователь",
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID пользователя",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserSegmentsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Добавление сегментов пользователю",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID пользователя",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Данные сегментов",
                        "name": "segments",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Segments"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserSegmentsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаление сегментов у пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID пользователя",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Данные сегментов",
                        "name": "segments",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Segments"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/{uid}/{time}": {
            "get": {
                "description": "Получение отчета по пользователю за период",
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID пользователя",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Период пользователя",
                        "name": "time",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.SegmentRequest": {
            "type": "object",
            "properties": {
                "segment_name": {
                    "type": "string"
                }
            }
        },
        "models.SegmentResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "segment_name": {
                    "type": "string"
                }
            }
        },
        "models.Segments": {
            "type": "object",
            "properties": {
                "segments_name": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "models.SuccessResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.UserSegmentsResponse": {
            "type": "object",
            "properties": {
                "segments": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "user_id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Avito-test-task",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
