package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct {}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "helloWorld" {
		return t.helloWorld(stub, args)
	}

	return shim.Error(fmt.Sprintf("Unknown function %s", function))
}

func (t *SimpleChaincode) helloWorld(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	return shim.Success([]byte(fmt.Sprintf("Hello world, %s!", args[0])))
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
