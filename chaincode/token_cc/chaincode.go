package main

import (
	"crypto/x509"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Chaincode is the definition of the chaincode structure.
type Chaincode struct {
}

// Definition of the Asset structure
type token struct {
	Type    string `json:"Type"`
	ID      string `json:"ID"`
	Token   string `json:"Token"`
	Creator string `json:"Creator"`
}

// Asset Prefixes
const tokenKey = "token-"

// Init function.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke function.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	fmt.Println("Invoke()", fcn, params)

	if fcn == "createToken" {
		return cc.createToken(stub, params)
	} else if fcn == "updateToken" {
		return cc.updateToken(stub, params)
	} else if fcn == "readToken" {
		return cc.readToken(stub, params)
	} else {
		fmt.Println("Invoke() did not find func: " + fcn)
		return shim.Error("Received unknown function invocation!")
	}
}

// Function to Create a New Token Profile.
func (cc *Chaincode) createToken(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, creatorID, _ := getTxCreatorInfo(stub)
	if !authenticatePHC(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Set Number of Params
	paramCount := 2

	// Check if sufficient Params passed
	if len(params) != paramCount {
		return shim.Error(fmt.Sprintf("Incorrect number of params. Expecting %d!", paramCount))
	}

	// Check if Params are non-empty
	for a := 0; a < paramCount; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Params must be a non-empty string")
		}
	}

	// Copy the Values from params[]
	ID := params[0]
	Token := params[1]

	// Generate Asset Key
	assetKey := tokenKey + ID

	// Check if Token exists with Key => assetKey
	tokenAsBytes, err := stub.GetState(assetKey)
	if err != nil {
		return shim.Error("Failed to check if Token exists!")
	} else if tokenAsBytes != nil {
		return shim.Error("Token Already Exists!")
	}

	// Generate Token from params provided
	token := &token{"token",
		ID, Token, creatorID}

	// Convert to JSON bytes
	tokenJSONasBytes, err := json.Marshal(token)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Token with Key => assetKey
	err = stub.PutState(assetKey, tokenJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(tokenJSONasBytes)
}

// Function to Read a Token Profile.
func (cc *Chaincode) readToken(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, _, _ := getTxCreatorInfo(stub)
	if !authenticatePatient(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Set Number of Params
	paramCount := 1

	// Check if sufficient Params passed
	if len(params) != paramCount {
		return shim.Error(fmt.Sprintf("Incorrect number of params. Expecting %d!", paramCount))
	}

	// Check if Params are non-empty
	for a := 0; a < paramCount; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Params must be a non-empty string")
		}
	}

	// Generate Asset Key
	assetKey := tokenKey + params[0]

	// Get State of Asset with Key => assetKey
	tokenAsBytes, err := stub.GetState(assetKey)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + params[0] + "\"}"
		return shim.Error(jsonResp)
	} else if tokenAsBytes == nil {
		jsonResp := "{\"Error\":\"Token does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Returned on successful execution of the function
	return shim.Success(tokenAsBytes)
}

// Function to Update Token Profile.
func (cc *Chaincode) updateToken(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, creatorID, _ := getTxCreatorInfo(stub)
	if !authenticatePHC(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Set Number of Params
	paramCount := 2

	// Check if sufficient Params passed
	if len(params) != paramCount {
		return shim.Error(fmt.Sprintf("Incorrect number of params. Expecting %d!", paramCount))
	}

	// Check if Params are non-empty
	for a := 0; a < paramCount; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Params must be a non-empty string")
		}
	}

	// Copy the Values from params[]
	ID := params[0]
	Token := params[1]

	// Generate Asset Key
	assetKey := tokenKey + ID

	// Check if Token exists with Key => assetKey
	tokenAsBytes, err := stub.GetState(assetKey)
	if err != nil {
		return shim.Error("Failed to get Token Details!")
	} else if tokenAsBytes == nil {
		return shim.Error("Error: Token Does NOT Exist!")
	}

	// Create Update struct var
	tokenToUpdate := token{}
	err = json.Unmarshal(tokenAsBytes, &tokenToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Update Token
	tokenToUpdate.Token = Token
	tokenToUpdate.Creator = creatorID

	// Convert to JSON bytes
	tokenJSONasBytes, err := json.Marshal(tokenToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Token with Key => assetKey
	err = stub.PutState(assetKey, tokenJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(tokenJSONasBytes)
}

// ---------------------------------------------
// Helper Functions
// ---------------------------------------------

// Authentication
// ++++++++++++++

// Get Tx Creator Info
func getTxCreatorInfo(stub shim.ChaincodeStubInterface) (string, string, string, error) {
	var mspid string
	var err error
	var cert *x509.Certificate
	mspid, err = cid.GetMSPID(stub)

	if err != nil {
		fmt.Printf("Error getting MSP identity: %sn", err.Error())
		return "", "", "", err
	}

	cert, err = cid.GetX509Certificate(stub)
	if err != nil {
		fmt.Printf("Error getting client certificate: %sn", err.Error())
		return "", "", "", err
	}

	return mspid, cert.Issuer.CommonName, cert.Subject.CommonName, nil
}

// Authenticate => Token
func authenticatePatient(mspID string, certCN string) bool {
	return (mspID == "PatientMSP") && (certCN == "ca.patient.health.com")
}

// Authenticate => PHC
func authenticatePHC(mspID string, certCN string) bool {
	return (mspID == "PHCMSP") && (certCN == "ca.phc.health.com")
}
