package main

import "fmt"
import "github.com/hyperledger/fabric/core/chaincode/shim" //The shim package
import "github.com/hyperledger/fabric/protos/peer" //peer.response is in the peer package

//Tokenchaincode represent our chaincode object
type TokenChaincode struct {

}


//Init Implemets the Init method
func (token *TokenChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response{
	fmt.Println("Init Executed");
	return shim.Success(nil);
}

// Invoke method
func (token *TokenChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
	fmt.Println("Invoke succeed");
	return shim.Success(nil)
}

func main(){
	fmt.Println("Started Chaincode")

	err:= shim.Start(new(TokenChaincode))
	if err != nil{
		fmt.Printf("Error starting chaincode: %s", err)
	}
}