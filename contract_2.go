package main

import (
    "fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/hyperledger/fabric/protos/peer"
)

type MyChaincode struct {
}

func (t *MyChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
    return shim.Success(nil)
}

func (t *MyChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
    function, args := stub.GetFunctionAndParameters()

    if function == "set" {
        return t.set(stub, args)
    } else if function == "get" {
        return t.get(stub, args)
    }

    return shim.Error("Invalid function name")
}

func (t *MyChaincode) set(stub shim.ChaincodeStubInterface, args []string) peer.Response {
    if len(args) != 2 {
        return shim.Error("Incorrect number of arguments. Expecting 2")
    }

    err := stub.PutState(args[0], []byte(args[1]))
    if err != nil {
        return shim.Error(err.Error())
    }

    return shim.Success(nil)
}

func (t *MyChaincode) get(stub shim.ChaincodeStubInterface, args []string) peer.Response {
    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting 1")
    }

    value, err := stub.GetState(args[0])
    if err != nil {
        return shim.Error(err.Error())
    }
    if value == nil {
        return shim.Error("Key not found")
    }

    return shim.Success(value)
}

func main() {
    err := shim.Start(new(MyChaincode))
    if err != nil {
        fmt.Printf("Error starting chaincode: %s", err)
    }
}
