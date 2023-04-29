package main

import (
    "encoding/json"
    "fmt"
    "strconv"

    "github.com/hyperledger/fabric-chaincode-go/shim"
    "github.com/hyperledger/fabric-protos-go/peer"
)

// Define the chaincode structure
type SupplyChainChaincode struct {
}

// Define the product structure
type Product struct {
    ID       string  `json:"id"`
    Name     string  `json:"name"`
    Quantity int     `json:"quantity"`
    Price    float64 `json:"price"`
}

// Define the shipment structure
type Shipment struct {
    ID        string    `json:"id"`
    Products  []Product `json:"products"`
    ShippedTo string    `json:"shippedTo"`
}

// Initialize the chaincode
func (t *SupplyChainChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
    return shim.Success(nil)
}

// Handle chaincode invocations
func (t *SupplyChainChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
    function, args := stub.GetFunctionAndParameters()

    if function == "createProduct" {
        return t.createProduct(stub, args)
    } else if function == "getProduct" {
        return t.getProduct(stub, args)
    } else if function == "createShipment" {
        return t.createShipment(stub, args)
    } else if function == "getShipment" {
        return t.getShipment(stub, args)
    }

    return shim.Error("Invalid function name")
}

// Create a new product
func (t *SupplyChainChaincode) createProduct(stub shim.ChaincodeStubInterface, args []string) peer.Response {
    if len(args) != 4 {
        return shim.Error("Incorrect number of arguments. Expecting 4")
    }

    id := args[0]
    name := args[1]
    quantity, err := strconv.Atoi(args[2])
    if err != nil {
        return shim.Error("Invalid quantity. Must be an integer")
    }
    price, err := strconv.ParseFloat(args[3], 64)
    if err != nil {
        return shim.Error("Invalid price. Must be a float")
    }

    product := &Product{
        ID:       id,
        Name:     name,
        Quantity: quantity,
        Price:    price,
    }
    productBytes, err := json.Marshal(product)
    if err != nil {
        return shim.Error("Failed to marshal product data")
    }

    err = stub.PutState(id, productBytes)
    if err != nil {
        return shim.Error("Failed to put product data on ledger")
    }

    return shim.Success(nil)
}

// Get a product by ID
func (t *SupplyChainChaincode) getProduct(stub shim.ChaincodeStubInterface, args []string) peer.Response {
    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting 1")
    }

    id := args[0]

    productBytes, err := stub.GetState(id)
    if err != nil {
        return shim.Error(fmt.Sprintf("Failed to get product data for ID %s", id))
    }
    if productBytes == nil {
        return shim.Error(fmt.Sprintf("Product with ID %s does not exist", id))
    }

    var product Product
    err = json.Unmarshal(productBytes, &product)
    if err != nil {
        return shim.Error("Failed to unmarshal product data")
    }

    return shim.Success(productBytes)
}

