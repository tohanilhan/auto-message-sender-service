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
            "name": "Tohan İlhan",
            "url": "https://github.com/tohanilhan/auto-message-sender-service",
            "email": "atahantohan@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/change-auto-sending/{option}": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This is the endpoint for changing the behaviour of the message sending service.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Config"
                ],
                "summary": "Can be used to start/stop auto messages sending feature.",
                "parameters": [
                    {
                        "enum": [
                            "ON",
                            "OFF"
                        ],
                        "type": "string",
                        "description": "Option to on/off auto sending messages",
                        "name": "option",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns the current status of the auto sending feature.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/messages": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Can be used to get all sent messages from Redis.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Messages"
                ],
                "summary": "Get all sent messages from Redis.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Message"
                            }
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "Check if the server is alive",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ping"
                ],
                "summary": "Check if the server is alive",
                "responses": {
                    "200": {
                        "description": "Returns PONG if the server is alive",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Message": {
            "type": "object",
            "properties": {
                "messageId": {
                    "type": "string"
                },
                "sent_time": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "x-ins-api-auth-key",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Auto Message Sender Service API",
	Description:      "This is an API for automatic message sender service. Routes are protected with an API key so please use an API key to access private routes. You can find the API key in the .env file of the project. You can also find the API key in the source code of the project on GitHub. The API key should be passed in the header of the request with the key name x-ins-api-auth-key.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}