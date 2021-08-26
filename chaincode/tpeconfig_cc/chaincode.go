package main

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Chaincode is the definition of the chaincode structure.
type Chaincode struct {
}

// Definition of the Asset structure
type tpeConfig struct {
	Type  string  `json:"Type"`
	N     int     `json:"N"`
	Theta float64 `json:"Theta"`
}

// Asset Prefixes
const tpeConfigKey = "tpeConfig"

// Init function.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke function.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	fmt.Println("Invoke()", fcn, params)

	if fcn == "setTPEConfig" {
		return cc.setTPEConfig(stub, params)
	} else if fcn == "readTPEConfig" {
		return cc.readTPEConfig(stub, params)
	} else {
		fmt.Println("Invoke() did not find func: " + fcn)
		return shim.Error("Received unknown function invocation!")
	}
}

// Function to Set TPEConfig
func (cc *Chaincode) setTPEConfig(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, _, _ := getTxCreatorInfo(stub)
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
	N, err := strconv.Atoi(params[0])
	if err != nil {
		return shim.Error("Error: Invalid N!")
	}
	Theta, err := strconv.ParseFloat(params[1], 64)
	if err != nil {
		return shim.Error("Error: Invalid Theta!")
	}

	// Generate Asset Key
	assetKey := tpeConfigKey

	// Generate TPEConfig from params provided
	tpeConfig := &tpeConfig{"tpeConfig",
		N, Theta}

	// Convert to JSON bytes
	tpeConfigJSONasBytes, err := json.Marshal(tpeConfig)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated TPEConfig with Key => assetKey
	err = stub.PutState(assetKey, tpeConfigJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(tpeConfigJSONasBytes)
}

// Function to Read a TPEConfig
func (cc *Chaincode) readTPEConfig(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Set Number of Params
	paramCount := 0

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
	assetKey := tpeConfigKey

	// Get State of Asset with Key => assetKey
	tpeConfigAsBytes, err := stub.GetState(assetKey)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get TPEConfig\"}"
		return shim.Error(jsonResp)
	} else if tpeConfigAsBytes == nil {
		jsonResp := "{\"Error\":\"TPEConfig does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Returned on successful execution of the function
	return shim.Success(tpeConfigAsBytes)
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

// Authenticate => PHC
func authenticatePHC(mspID string, certCN string) bool {
	return (mspID == "PHCMSP") && (certCN == "ca.phc.health.com")
}
