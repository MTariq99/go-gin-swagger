// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/hello": {
            "get": {
                "description": "Return Hello response",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hello"
                ],
                "summary": "Show a hello message",
                "responses": {
                    "200": {
                        "description": "hello world",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/ping": {
            "get": {
                "description": "Get ping pong response",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "Show a ping pong message",
                "responses": {
                    "200": {
                        "description": "pong",
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
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/",
	Schemes:          []string{},
	Title:            "",
	Description:      "this is an example of using swagger with gin framework in golang",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,

}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
