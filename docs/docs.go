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
        "/upload": {
            "post": {
                "description": "Uploads a file to the specified folder within the server's storage directory.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "Upload file",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Folder where the file will be uploaded",
                        "name": "folder",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "File to upload",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Upload file success",
                        "schema": {
                            "$ref": "#/definitions/api.uploadFileResponse"
                        }
                    }
                }
            }
        },
        "/{foldername}/{filename}": {
            "get": {
                "description": "Retrieve file data from specified folder",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "file"
                ],
                "summary": "Get file",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Folder Name",
                        "name": "foldername",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "File Name",
                        "name": "filename",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "File retrieved successfully",
                        "schema": {
                            "type": "file"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.uploadFileResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Segment3d App API Documentation",
	Description:      "This is a documentation for Segment3d App API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
