package v1 

const (
swagger = `{
  "swagger": "2.0",
  "info": {
    "title": "api.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/record": {
      "get": {
        "operationId": "GetRecord",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/v1Entries"
            }
          }
        },
        "parameters": [
          {
            "name": "name",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "address",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "DNSService"
        ]
      },
      "delete": {
        "operationId": "DeleteRecord",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/v1Entry"
            }
          }
        },
        "tags": [
          "DNSService"
        ]
      },
      "post": {
        "operationId": "PostRecord",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/v1Entry"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v1Entry"
            }
          }
        ],
        "tags": [
          "DNSService"
        ]
      },
      "put": {
        "operationId": "PutRecord",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/v1Entry"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v1Entry"
            }
          }
        ],
        "tags": [
          "DNSService"
        ]
      }
    }
  },
  "definitions": {
    "v1Entries": {
      "type": "object",
      "properties": {
        "entries": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1Entry"
          }
        }
      }
    },
    "v1Entry": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "address": {
          "type": "string"
        }
      }
    }
  }
}
`
)
