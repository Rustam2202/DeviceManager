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
        "/device": {
            "post": {
                "description": "Add a new device to database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device"
                ],
                "summary": "Add device",
                "parameters": [
                    {
                        "description": "Add Device Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/device.DeviceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponce"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponce"
                        }
                    }
                }
            }
        },
        "/device/{uuid}": {
            "get": {
                "description": "Get device info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device GET"
                ],
                "summary": "Get a device",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Device UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Device"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponce"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "delete": {
                "description": "Delete a device from database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device"
                ],
                "summary": "Delete a device",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Device Id",
                        "name": "uuid",
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
                            "$ref": "#/definitions/handlers.ErrorResponce"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponce"
                        }
                    }
                }
            }
        },
        "/device_email": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device UPDATE"
                ],
                "summary": "Update a device E-mail",
                "parameters": [
                    {
                        "description": "Update Device E-mail Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/device.UpdateEmailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/device.UpdateEmailRequest"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponce"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponce"
                        }
                    }
                }
            }
        },
        "/device_email/{email}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device GET"
                ],
                "summary": "Get devices by Email filter",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Devices Email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Device"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponce"
                        }
                    }
                }
            }
        },
        "/device_geo": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device GET"
                ],
                "summary": "Get devices by Geoposition",
                "parameters": [
                    {
                        "type": "number",
                        "description": "longitude",
                        "name": "longitude",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "latitude",
                        "name": "latitude",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "distance",
                        "name": "distance",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Device"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponce"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponce"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device UPDATE"
                ],
                "summary": "Update a device geolocation",
                "parameters": [
                    {
                        "description": "Update Device geolocation Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/device.UpdateGeolocationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/device.UpdateGeolocationRequest"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponce"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponce"
                        }
                    }
                }
            }
        },
        "/device_lang": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device UPDATE"
                ],
                "summary": "Update a device language",
                "parameters": [
                    {
                        "description": "Update Device language Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/device.UpdateLanguageRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/device.UpdateLanguageRequest"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponce"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponce"
                        }
                    }
                }
            }
        },
        "/device_lang/{language}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device GET"
                ],
                "summary": "Get devices by Language filter",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Devices Language",
                        "name": "language",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Device"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponce"
                        }
                    }
                }
            }
        },
        "/event": {
            "get": {
                "description": "Get events of device from database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "Get events",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID",
                        "name": "uuid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Begin time range",
                        "name": "timeBegin",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "End time range",
                        "name": "timeEnd",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Event"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponce"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponce"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new event from device to database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "Add event",
                "parameters": [
                    {
                        "description": "Add Event Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/event.AddEventRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponce"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponce"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "device.Coordinates": {
            "type": "object",
            "properties": {
                "latitude": {
                    "type": "number",
                    "example": 37.552375
                },
                "longitude": {
                    "type": "number",
                    "example": 55.646574
                }
            }
        },
        "device.DeviceRequest": {
            "type": "object",
            "properties": {
                "coordinates": {
                    "$ref": "#/definitions/device.Coordinates"
                },
                "email": {
                    "type": "string",
                    "example": "example@email.com"
                },
                "language": {
                    "type": "string",
                    "example": "en-US"
                },
                "platform": {
                    "type": "string",
                    "example": "Mozilla/5.0 (Linux; U; Android 2.3.7; en-us; Nexus One Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"
                },
                "uuid": {
                    "type": "string",
                    "example": "550e8400-e29b-41d4-a716-446655440000"
                }
            }
        },
        "device.UpdateEmailRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "device.UpdateGeolocationRequest": {
            "type": "object",
            "properties": {
                "latitude": {
                    "type": "number",
                    "default": 37.552375
                },
                "longitude": {
                    "type": "number",
                    "default": 55.646575
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "device.UpdateLanguageRequest": {
            "type": "object",
            "properties": {
                "language": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "domain.Device": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "language": {
                    "type": "string"
                },
                "location": {
                    "$ref": "#/definitions/domain.Location"
                },
                "platform": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "domain.Event": {
            "type": "object",
            "properties": {
                "attributes": {
                    "type": "array",
                    "items": {}
                },
                "createdAt": {
                    "type": "string"
                },
                "deviceUUID": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "domain.Location": {
            "type": "object",
            "properties": {
                "coordinates": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "event.AddEventRequest": {
            "type": "object",
            "properties": {
                "attributes": {
                    "type": "array",
                    "items": {}
                },
                "name": {
                    "type": "string",
                    "example": "device event"
                },
                "uuid": {
                    "type": "string",
                    "example": "550e8400-e29b-41d4-a716-446655440000"
                }
            }
        },
        "handlers.ErrorResponce": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error": {},
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Device Manager API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
