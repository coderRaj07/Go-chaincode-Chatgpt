package main

import (
    "fmt"
    "github.com/hyperledger/fabric-chaincode-go/shim"
    pb "github.com/hyperledger/fabric-protos-go/peer"
)

type AssetTransferChaincode struct {
}

func (t *AssetTransferChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
    return shim.Success(nil)
}

func (t *AssetTransferChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    function, args := stub.GetFunctionAndParameters()

    if function == "transfer" {
        // Transfer an asset from one owner to another
        return transfer(stub, args)
    }

    return shim.Error("Invalid function name")
}

func transfer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    if len(args) != 3 {
        return shim.Error("Incorrect number of arguments. Expecting 3")
    }

    asset := args[0]
    from := args[1]
    to := args[2]

    // Get the current asset owner from the state database
    assetOwnerBytes, err := stub.GetState(asset)
    if err != nil {
        return shim.Error(fmt.Sprintf("Failed to get state for asset %s", asset))
    }
    if assetOwnerBytes == nil {
        return shim.Error(fmt.Sprintf("Asset %s does not exist", asset))
    }
    assetOwner := string(assetOwnerBytes)

    // Verify that the sender is the current owner of the asset
    if assetOwner != from {
        return shim.Error(fmt.Sprintf("Sender %s is not the current owner of asset %s", from, asset))
    }

    // Transfer the asset to the new owner
    err = stub.PutState(asset, []byte(to))
    if err != nil {
        return shim.Error(fmt.Sprintf("Failed to update state for asset %s", asset))
    }

    return shim.Success(nil)
}
