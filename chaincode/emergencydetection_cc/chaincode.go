package main

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	sc "github.com/hyperledger/fabric/protos/peer"
	tpe "github.com/mukherjeearnab/gotpe"
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

// Definition of the Asset structure
type tpeConfig struct {
	Type  string  `json:"Type"`
	N     int     `json:"N"`
	Theta float64 `json:"Theta"`
}

// Asset Prefixes
const tokenKey = "token-"
const tpeConfigKey = "tpeConfig"

// Init function.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke function.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	fmt.Println("Invoke()", fcn, params)

	if fcn == "checkReading" {
		return cc.checkReading(stub, params)
	} else {
		fmt.Println("Invoke() did not find func: " + fcn)
		return shim.Error("Received unknown function invocation!")
	}
}

// Function to Create a New Token Profile.
func (cc *Chaincode) checkReading(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, creatorID, _ := getTxCreatorInfo(stub)
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

	// Copy the Values from params[]
	ID := creatorID
	Cipher := params[0]

	// Generate Asset Key
	assetKey := tokenKey + ID

	// Check if Token exists with Key => assetKey
	tokenAsBytes, err := stub.GetState(assetKey)
	if err != nil {
		return shim.Error("Failed to get Token Details!")
	} else if tokenAsBytes == nil {
		return shim.Error("Error: Token Does NOT Exist!")
	}

	// Retrieve Token struct var
	Token := token{}
	err = json.Unmarshal(tokenAsBytes, &Token) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Check if TPE Configuration exists
	tpeConfigAsBytes, err := stub.GetState(tpeConfigKey)
	if err != nil {
		return shim.Error("Failed to get TPE Configuration Details!")
	} else if tpeConfigAsBytes == nil {
		return shim.Error("Error: TPE Configuration Does NOT Exist!")
	}

	// Retrieve TPE Configuration
	TPEConfig := tpeConfig{}
	err = json.Unmarshal(tpeConfigAsBytes, &TPEConfig) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Init TPE instance
	var TPE tpe.TPE

	// Setup TPE instance
	TPE.Setup(TPEConfig.N, TPEConfig.Theta)

	// Decrypt Cipher and obtain result
	detectionStatus := TPE.Decrypt(Cipher, Token.Token)

	// Convert to JSON bytes
	detectionAsBytes, err := json.Marshal(detectionStatus)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(detectionAsBytes)
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
