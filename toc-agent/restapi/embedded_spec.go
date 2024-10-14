// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Worker API to create backup and restoration",
    "title": "toc-agent",
    "termsOfService": "http://gyk4j.github.io/terms/",
    "contact": {
      "email": "147011991+gyk4j@users.noreply.github.com"
    },
    "license": {
      "name": "Eclipse Public License 2.0 (EPL)",
      "url": "https://www.eclipse.org/legal/epl-2.0/"
    },
    "version": "1.0.0"
  },
  "host": "0.0.0.0:8888",
  "basePath": "/v1",
  "paths": {
    "/backups": {
      "get": {
        "description": "Retrieve information and state of all backups and their snapshots",
        "produces": [
          "application/json"
        ],
        "tags": [
          "backup"
        ],
        "summary": "Query backup info",
        "operationId": "getBackup",
        "responses": {
          "200": {
            "description": "Backup info retrieved",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Backup"
              }
            }
          },
          "403": {
            "description": "Forbidden from querying backup info",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "404": {
            "description": "Backup not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "405": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "500": {
            "description": "TOC controller error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "503": {
            "description": "Service unavailable",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      },
      "post": {
        "description": "Create a new backup snapshot",
        "produces": [
          "application/json"
        ],
        "tags": [
          "backup"
        ],
        "summary": "Create a new backup snapshot",
        "operationId": "newBackup",
        "responses": {
          "200": {
            "description": "Backup started",
            "schema": {
              "$ref": "#/definitions/Backup"
            }
          },
          "403": {
            "description": "Forbidden from creating duplicate/repeat backup",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "404": {
            "description": "Main backup server not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "405": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "500": {
            "description": "Main TOC controller error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "502": {
            "description": "Bad gateway. Main backup server error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "503": {
            "description": "Service unavailable on main TOC controller",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "504": {
            "description": "Gateway timeout. Main backup server did not reply.",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    },
    "/backups/{backupId}": {
      "get": {
        "description": "Query a single backup",
        "produces": [
          "application/json"
        ],
        "tags": [
          "backup"
        ],
        "summary": "Find backup by ID",
        "operationId": "getBackupById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "x-exportParamName": "BackupId",
            "description": "ID of backup to return",
            "name": "backupId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Backup"
            }
          },
          "400": {
            "description": "Invalid ID supplied",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "404": {
            "description": "Backup not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    },
    "/restorations": {
      "get": {
        "description": "Retrieve information and state of all restorations, their backup and snapshots",
        "produces": [
          "application/json"
        ],
        "tags": [
          "restoration"
        ],
        "summary": "Query restoration info",
        "operationId": "getRestoration",
        "responses": {
          "200": {
            "description": "Restoration info retrieved",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Restoration"
              }
            }
          },
          "403": {
            "description": "Forbidden from querying restoration info",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "404": {
            "description": "Restoration not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "405": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "500": {
            "description": "TOC controller error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "503": {
            "description": "Service unavailable",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      },
      "post": {
        "description": "Run backup restoration and return immediately without waiting",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "restoration"
        ],
        "summary": "Start a backup restoration attempt",
        "operationId": "newRestoration",
        "parameters": [
          {
            "x-exportParamName": "Body",
            "description": "Backup and associated snapshots for restoration",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Backup"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Restoration started",
            "schema": {
              "$ref": "#/definitions/Restoration"
            }
          },
          "403": {
            "description": "Forbidden from creating duplicate/repeat restoration",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "404": {
            "description": "Stepup backup server not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "405": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "500": {
            "description": "Stepup TOC controller error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "502": {
            "description": "Bad gateway. Stepup backup server error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "503": {
            "description": "Service unavailable on stepup TOC controller",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "504": {
            "description": "Gateway timeout. Stepup backup server did not reply.",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    },
    "/restorations/{restorationId}": {
      "get": {
        "description": "Query a single restoration",
        "produces": [
          "application/json"
        ],
        "tags": [
          "restoration"
        ],
        "summary": "Find restoration by ID",
        "operationId": "getRestorationById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "x-exportParamName": "RestorationId",
            "description": "ID of restoration to return",
            "name": "restorationId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Restoration"
            }
          },
          "400": {
            "description": "Invalid ID supplied",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "404": {
            "description": "Restoration not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    },
    "/transfers": {
      "get": {
        "description": "Retrieve status of all transfers",
        "produces": [
          "application/json"
        ],
        "tags": [
          "transfer"
        ],
        "summary": "Query transfer status",
        "operationId": "getTransfer",
        "responses": {
          "200": {
            "description": "Transfer info retrieved",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Transfer"
              }
            }
          },
          "403": {
            "description": "Forbidden from querying transfer status",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "404": {
            "description": "Transfer not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "405": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "500": {
            "description": "TOC controller error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "503": {
            "description": "Service unavailable",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      },
      "post": {
        "description": "Run transfer and return immediately without waiting",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "transfer"
        ],
        "summary": "Start a backup transfer attempt",
        "operationId": "newTransfer",
        "parameters": [
          {
            "x-exportParamName": "Body",
            "description": "Backup and associated snapshots for restoration",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Backup"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Transfer started",
            "schema": {
              "$ref": "#/definitions/Transfer"
            }
          },
          "403": {
            "description": "Forbidden from creating duplicate/repeat transfer",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "404": {
            "description": "Stepup backup server not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "405": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "500": {
            "description": "Stepup TOC controller error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "502": {
            "description": "Bad gateway. Stepup backup server error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "503": {
            "description": "Service unavailable on stepup TOC controller",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "504": {
            "description": "Gateway timeout. Stepup backup server did not reply.",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    },
    "/transfers/{transferId}": {
      "get": {
        "description": "Query a single transfer",
        "produces": [
          "application/json"
        ],
        "tags": [
          "transfer"
        ],
        "summary": "Find transfer by ID",
        "operationId": "getTransferById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "x-exportParamName": "TransferId",
            "description": "ID of transfer to return",
            "name": "transferId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Transfer"
            }
          },
          "400": {
            "description": "Invalid ID supplied",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "404": {
            "description": "Transfer not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ApiResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32",
          "x-omitempty": false
        },
        "message": {
          "type": "string",
          "x-omitempty": false
        },
        "type": {
          "description": "Severity level",
          "type": "string",
          "enum": [
            "error",
            "warning",
            "info",
            "debug",
            "trace"
          ],
          "x-omitempty": false
        }
      },
      "example": {
        "code": 0,
        "message": "message",
        "type": "info"
      }
    },
    "Backup": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "x-omitempty": false
        },
        "snapshots": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Snapshot"
          },
          "x-omitempty": false
        },
        "time": {
          "type": "string",
          "format": "date-time",
          "x-omitempty": false
        }
      },
      "xml": {
        "name": "Backup"
      },
      "example": {
        "id": 1,
        "time": "2024-12-31T23:59:59Z"
      }
    },
    "Restoration": {
      "type": "object",
      "properties": {
        "backup": {
          "x-omitempty": false,
          "$ref": "#/definitions/Backup"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "x-omitempty": false
        },
        "status": {
          "description": "Restoration Status",
          "type": "string",
          "enum": [
            "queued",
            "in-progress",
            "completed",
            "failed"
          ],
          "x-omitempty": false
        }
      },
      "xml": {
        "name": "Restoration"
      },
      "example": {
        "id": 1
      }
    },
    "Snapshot": {
      "type": "object",
      "properties": {
        "complete": {
          "type": "boolean",
          "x-omitempty": false
        },
        "file": {
          "type": "string",
          "x-omitempty": false
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "x-omitempty": false
        },
        "status": {
          "description": "Snapshot Status",
          "type": "string",
          "enum": [
            "queued",
            "in-progress",
            "completed",
            "failed"
          ],
          "x-omitempty": false
        }
      },
      "xml": {
        "name": "snapshot"
      },
      "example": {
        "complete": true,
        "file": "/var/opt/toc/archive/20241231T235959Z-server.tar.gz",
        "id": 1,
        "status": "completed"
      }
    },
    "Transfer": {
      "type": "object",
      "properties": {
        "backup": {
          "x-omitempty": false,
          "$ref": "#/definitions/Backup"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "x-omitempty": false
        },
        "status": {
          "description": "Transfer Status",
          "type": "string",
          "enum": [
            "queued",
            "in-progress",
            "completed",
            "failed"
          ],
          "x-omitempty": false
        }
      },
      "xml": {
        "name": "Transfer"
      },
      "example": {
        "id": 1,
        "status": "failed"
      }
    }
  },
  "tags": [
    {
      "description": "For running backup jobs and querying backup job status",
      "name": "backup",
      "externalDocs": {
        "description": "Find out more about TOC",
        "url": "https://github.com/gyk4j/toc"
      }
    },
    {
      "description": "For running restoration jobs and querying restoration job status",
      "name": "restoration",
      "externalDocs": {
        "description": "Find out more about TOC",
        "url": "https://github.com/gyk4j/toc"
      }
    },
    {
      "description": "For running backup transfer jobs and querying transfer status",
      "name": "transfer",
      "externalDocs": {
        "description": "Find out more about TOC",
        "url": "https://github.com/gyk4j/toc"
      }
    }
  ],
  "externalDocs": {
    "description": "Find out more about TOC",
    "url": "https://github.com/gyk4j/toc"
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Worker API to create backup and restoration",
    "title": "toc-agent",
    "termsOfService": "http://gyk4j.github.io/terms/",
    "contact": {
      "email": "147011991+gyk4j@users.noreply.github.com"
    },
    "license": {
      "name": "Eclipse Public License 2.0 (EPL)",
      "url": "https://www.eclipse.org/legal/epl-2.0/"
    },
    "version": "1.0.0"
  },
  "host": "0.0.0.0:8888",
  "basePath": "/v1",
  "paths": {
    "/backups": {
      "get": {
        "description": "Retrieve information and state of all backups and their snapshots",
        "produces": [
          "application/json"
        ],
        "tags": [
          "backup"
        ],
        "summary": "Query backup info",
        "operationId": "getBackup",
        "responses": {
          "200": {
            "description": "Backup info retrieved",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Backup"
              }
            }
          },
          "403": {
            "description": "Forbidden from querying backup info",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "404": {
            "description": "Backup not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "405": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "500": {
            "description": "TOC controller error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "503": {
            "description": "Service unavailable",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      },
      "post": {
        "description": "Create a new backup snapshot",
        "produces": [
          "application/json"
        ],
        "tags": [
          "backup"
        ],
        "summary": "Create a new backup snapshot",
        "operationId": "newBackup",
        "responses": {
          "200": {
            "description": "Backup started",
            "schema": {
              "$ref": "#/definitions/Backup"
            }
          },
          "403": {
            "description": "Forbidden from creating duplicate/repeat backup",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "404": {
            "description": "Main backup server not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "405": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "500": {
            "description": "Main TOC controller error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "502": {
            "description": "Bad gateway. Main backup server error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "503": {
            "description": "Service unavailable on main TOC controller",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "504": {
            "description": "Gateway timeout. Main backup server did not reply.",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    },
    "/backups/{backupId}": {
      "get": {
        "description": "Query a single backup",
        "produces": [
          "application/json"
        ],
        "tags": [
          "backup"
        ],
        "summary": "Find backup by ID",
        "operationId": "getBackupById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "x-exportParamName": "BackupId",
            "description": "ID of backup to return",
            "name": "backupId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Backup"
            }
          },
          "400": {
            "description": "Invalid ID supplied",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "404": {
            "description": "Backup not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    },
    "/restorations": {
      "get": {
        "description": "Retrieve information and state of all restorations, their backup and snapshots",
        "produces": [
          "application/json"
        ],
        "tags": [
          "restoration"
        ],
        "summary": "Query restoration info",
        "operationId": "getRestoration",
        "responses": {
          "200": {
            "description": "Restoration info retrieved",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Restoration"
              }
            }
          },
          "403": {
            "description": "Forbidden from querying restoration info",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "404": {
            "description": "Restoration not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "405": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "500": {
            "description": "TOC controller error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "503": {
            "description": "Service unavailable",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      },
      "post": {
        "description": "Run backup restoration and return immediately without waiting",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "restoration"
        ],
        "summary": "Start a backup restoration attempt",
        "operationId": "newRestoration",
        "parameters": [
          {
            "x-exportParamName": "Body",
            "description": "Backup and associated snapshots for restoration",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Backup"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Restoration started",
            "schema": {
              "$ref": "#/definitions/Restoration"
            }
          },
          "403": {
            "description": "Forbidden from creating duplicate/repeat restoration",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "404": {
            "description": "Stepup backup server not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "405": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "500": {
            "description": "Stepup TOC controller error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "502": {
            "description": "Bad gateway. Stepup backup server error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "503": {
            "description": "Service unavailable on stepup TOC controller",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "504": {
            "description": "Gateway timeout. Stepup backup server did not reply.",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    },
    "/restorations/{restorationId}": {
      "get": {
        "description": "Query a single restoration",
        "produces": [
          "application/json"
        ],
        "tags": [
          "restoration"
        ],
        "summary": "Find restoration by ID",
        "operationId": "getRestorationById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "x-exportParamName": "RestorationId",
            "description": "ID of restoration to return",
            "name": "restorationId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Restoration"
            }
          },
          "400": {
            "description": "Invalid ID supplied",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "404": {
            "description": "Restoration not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    },
    "/transfers": {
      "get": {
        "description": "Retrieve status of all transfers",
        "produces": [
          "application/json"
        ],
        "tags": [
          "transfer"
        ],
        "summary": "Query transfer status",
        "operationId": "getTransfer",
        "responses": {
          "200": {
            "description": "Transfer info retrieved",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Transfer"
              }
            }
          },
          "403": {
            "description": "Forbidden from querying transfer status",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "404": {
            "description": "Transfer not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "405": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "500": {
            "description": "TOC controller error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "503": {
            "description": "Service unavailable",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      },
      "post": {
        "description": "Run transfer and return immediately without waiting",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "transfer"
        ],
        "summary": "Start a backup transfer attempt",
        "operationId": "newTransfer",
        "parameters": [
          {
            "x-exportParamName": "Body",
            "description": "Backup and associated snapshots for restoration",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Backup"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Transfer started",
            "schema": {
              "$ref": "#/definitions/Transfer"
            }
          },
          "403": {
            "description": "Forbidden from creating duplicate/repeat transfer",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "404": {
            "description": "Stepup backup server not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "405": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "500": {
            "description": "Stepup TOC controller error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "502": {
            "description": "Bad gateway. Stepup backup server error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "503": {
            "description": "Service unavailable on stepup TOC controller",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "504": {
            "description": "Gateway timeout. Stepup backup server did not reply.",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    },
    "/transfers/{transferId}": {
      "get": {
        "description": "Query a single transfer",
        "produces": [
          "application/json"
        ],
        "tags": [
          "transfer"
        ],
        "summary": "Find transfer by ID",
        "operationId": "getTransferById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "x-exportParamName": "TransferId",
            "description": "ID of transfer to return",
            "name": "transferId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Transfer"
            }
          },
          "400": {
            "description": "Invalid ID supplied",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "404": {
            "description": "Transfer not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ApiResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32",
          "x-omitempty": false
        },
        "message": {
          "type": "string",
          "x-omitempty": false
        },
        "type": {
          "description": "Severity level",
          "type": "string",
          "enum": [
            "error",
            "warning",
            "info",
            "debug",
            "trace"
          ],
          "x-omitempty": false
        }
      },
      "example": {
        "code": 0,
        "message": "message",
        "type": "info"
      }
    },
    "Backup": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "x-omitempty": false
        },
        "snapshots": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Snapshot"
          },
          "x-omitempty": false
        },
        "time": {
          "type": "string",
          "format": "date-time",
          "x-omitempty": false
        }
      },
      "xml": {
        "name": "Backup"
      },
      "example": {
        "id": 1,
        "time": "2024-12-31T23:59:59Z"
      }
    },
    "Restoration": {
      "type": "object",
      "properties": {
        "backup": {
          "x-omitempty": false,
          "$ref": "#/definitions/Backup"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "x-omitempty": false
        },
        "status": {
          "description": "Restoration Status",
          "type": "string",
          "enum": [
            "queued",
            "in-progress",
            "completed",
            "failed"
          ],
          "x-omitempty": false
        }
      },
      "xml": {
        "name": "Restoration"
      },
      "example": {
        "id": 1
      }
    },
    "Snapshot": {
      "type": "object",
      "properties": {
        "complete": {
          "type": "boolean",
          "x-omitempty": false
        },
        "file": {
          "type": "string",
          "x-omitempty": false
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "x-omitempty": false
        },
        "status": {
          "description": "Snapshot Status",
          "type": "string",
          "enum": [
            "queued",
            "in-progress",
            "completed",
            "failed"
          ],
          "x-omitempty": false
        }
      },
      "xml": {
        "name": "snapshot"
      },
      "example": {
        "complete": true,
        "file": "/var/opt/toc/archive/20241231T235959Z-server.tar.gz",
        "id": 1,
        "status": "completed"
      }
    },
    "Transfer": {
      "type": "object",
      "properties": {
        "backup": {
          "x-omitempty": false,
          "$ref": "#/definitions/Backup"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "x-omitempty": false
        },
        "status": {
          "description": "Transfer Status",
          "type": "string",
          "enum": [
            "queued",
            "in-progress",
            "completed",
            "failed"
          ],
          "x-omitempty": false
        }
      },
      "xml": {
        "name": "Transfer"
      },
      "example": {
        "id": 1,
        "status": "failed"
      }
    }
  },
  "tags": [
    {
      "description": "For running backup jobs and querying backup job status",
      "name": "backup",
      "externalDocs": {
        "description": "Find out more about TOC",
        "url": "https://github.com/gyk4j/toc"
      }
    },
    {
      "description": "For running restoration jobs and querying restoration job status",
      "name": "restoration",
      "externalDocs": {
        "description": "Find out more about TOC",
        "url": "https://github.com/gyk4j/toc"
      }
    },
    {
      "description": "For running backup transfer jobs and querying transfer status",
      "name": "transfer",
      "externalDocs": {
        "description": "Find out more about TOC",
        "url": "https://github.com/gyk4j/toc"
      }
    }
  ],
  "externalDocs": {
    "description": "Find out more about TOC",
    "url": "https://github.com/gyk4j/toc"
  }
}`))
}
