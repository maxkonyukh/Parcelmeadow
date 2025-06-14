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
  "swagger": "2.0",
  "info": {
    "description": "This API contains services related to fetching information about local delivery in Parcelmeadow.",
    "title": "parcelmeadow API",
    "version": "1.0"
  },
  "host": "localhost:9090",
  "basePath": "/",
  "paths": {
    "/v1/parcels": {
      "get": {
        "description": "Fetches today's parcels.",
        "produces": [
          "application/json"
        ],
        "summary": "parcels",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/GetTodayParcelsV1Response"
            }
          },
          "400": {
            "description": "Bad request"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      },
      "post": {
        "description": "Creates a parcel.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "parcels",
        "parameters": [
          {
            "$ref": "#/parameters/CreateParcelV1Request"
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "400": {
            "description": "Bad request"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/v1/parcels/routes": {
      "get": {
        "description": "Fetches the routes for today's deliveries.",
        "produces": [
          "application/json"
        ],
        "summary": "routes",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/GetTodayRoutesV1Response"
            }
          },
          "400": {
            "description": "Bad request"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    }
  },
  "definitions": {
    "GetTodayParcelsV1Response": {
      "description": "Response containing today's parcels",
      "type": "object",
      "required": [
        "parcels"
      ],
      "properties": {
        "parcels": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ParcelV1"
          }
        }
      }
    },
    "GetTodayRoutesV1Response": {
      "description": "Response containing today's routes",
      "type": "object",
      "required": [
        "routes"
      ],
      "properties": {
        "routes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/RouteV1"
          }
        }
      }
    },
    "ParcelV1": {
      "description": "Represents a parcel in the parcelmeadow system",
      "type": "object",
      "properties": {
        "address": {
          "description": "Delivery address for the parcel",
          "type": "string",
          "example": "123 Main St"
        },
        "id": {
          "description": "Unique identifier for the parcel",
          "type": "string",
          "example": "P001"
        },
        "postCode": {
          "description": "Postal code for the delivery address",
          "type": "string",
          "example": "0465"
        },
        "status": {
          "description": "Current status of the parcel",
          "type": "string",
          "example": "in-transit"
        },
        "weight": {
          "description": "Weight of the parcel in kilograms",
          "type": "number",
          "format": "float",
          "example": 2.3
        }
      }
    },
    "RouteV1": {
      "description": "Represents a delivery route in the parcelmeadow system",
      "type": "object",
      "properties": {
        "id": {
          "description": "Unique identifier for the route",
          "type": "string",
          "example": "R001"
        },
        "stops": {
          "description": "List of stops for this route",
          "type": "array",
          "items": {
            "$ref": "#/definitions/StopV1"
          }
        }
      }
    },
    "StopV1": {
      "description": "Represents a stop in a delivery route",
      "type": "object",
      "properties": {
        "address": {
          "description": "Address of the stop",
          "type": "string",
          "example": "456 Elm St"
        },
        "id": {
          "description": "Unique identifier for the stop",
          "type": "string",
          "example": "S001"
        },
        "parcels": {
          "description": "List of parcels to be delivered at this stop",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ParcelV1"
          }
        }
      }
    }
  },
  "parameters": {
    "CreateParcelV1Request": {
      "description": "Request to create a new parcel",
      "name": "CreateParcelV1Request",
      "in": "body",
      "schema": {
        "$ref": "#/definitions/ParcelV1"
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "description": "This API contains services related to fetching information about local delivery in Parcelmeadow.",
    "title": "parcelmeadow API",
    "version": "1.0"
  },
  "host": "localhost:9090",
  "basePath": "/",
  "paths": {
    "/v1/parcels": {
      "get": {
        "description": "Fetches today's parcels.",
        "produces": [
          "application/json"
        ],
        "summary": "parcels",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/GetTodayParcelsV1Response"
            }
          },
          "400": {
            "description": "Bad request"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      },
      "post": {
        "description": "Creates a parcel.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "parcels",
        "parameters": [
          {
            "description": "Request to create a new parcel",
            "name": "CreateParcelV1Request",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/ParcelV1"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "400": {
            "description": "Bad request"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/v1/parcels/routes": {
      "get": {
        "description": "Fetches the routes for today's deliveries.",
        "produces": [
          "application/json"
        ],
        "summary": "routes",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/GetTodayRoutesV1Response"
            }
          },
          "400": {
            "description": "Bad request"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    }
  },
  "definitions": {
    "GetTodayParcelsV1Response": {
      "description": "Response containing today's parcels",
      "type": "object",
      "required": [
        "parcels"
      ],
      "properties": {
        "parcels": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ParcelV1"
          }
        }
      }
    },
    "GetTodayRoutesV1Response": {
      "description": "Response containing today's routes",
      "type": "object",
      "required": [
        "routes"
      ],
      "properties": {
        "routes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/RouteV1"
          }
        }
      }
    },
    "ParcelV1": {
      "description": "Represents a parcel in the parcelmeadow system",
      "type": "object",
      "properties": {
        "address": {
          "description": "Delivery address for the parcel",
          "type": "string",
          "example": "123 Main St"
        },
        "id": {
          "description": "Unique identifier for the parcel",
          "type": "string",
          "example": "P001"
        },
        "postCode": {
          "description": "Postal code for the delivery address",
          "type": "string",
          "example": "0465"
        },
        "status": {
          "description": "Current status of the parcel",
          "type": "string",
          "example": "in-transit"
        },
        "weight": {
          "description": "Weight of the parcel in kilograms",
          "type": "number",
          "format": "float",
          "example": 2.3
        }
      }
    },
    "RouteV1": {
      "description": "Represents a delivery route in the parcelmeadow system",
      "type": "object",
      "properties": {
        "id": {
          "description": "Unique identifier for the route",
          "type": "string",
          "example": "R001"
        },
        "stops": {
          "description": "List of stops for this route",
          "type": "array",
          "items": {
            "$ref": "#/definitions/StopV1"
          }
        }
      }
    },
    "StopV1": {
      "description": "Represents a stop in a delivery route",
      "type": "object",
      "properties": {
        "address": {
          "description": "Address of the stop",
          "type": "string",
          "example": "456 Elm St"
        },
        "id": {
          "description": "Unique identifier for the stop",
          "type": "string",
          "example": "S001"
        },
        "parcels": {
          "description": "List of parcels to be delivered at this stop",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ParcelV1"
          }
        }
      }
    }
  },
  "parameters": {
    "CreateParcelV1Request": {
      "description": "Request to create a new parcel",
      "name": "CreateParcelV1Request",
      "in": "body",
      "schema": {
        "$ref": "#/definitions/ParcelV1"
      }
    }
  }
}`))
}
