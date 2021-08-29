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
type patient struct {
	Type     string `json:"Type"`
	ID       string `json:"ID"`
	Name     string `json:"Name"`
	Verified bool   `json:"Verified"`
}

// Asset Prefixes
const patientKey = "patient-"

// Init function.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke function.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	fmt.Println("Invoke()", fcn, params)

	if fcn == "createPatientProfile" {
		return cc.createPatientProfile(stub, params)
	} else if fcn == "updatePatientProfile" {
		return cc.updatePatientProfile(stub, params)
	} else if fcn == "activatePatientProfile" {
		return cc.activatePatientProfile(stub, params)
	} else if fcn == "readPatientProfile" {
		return cc.readPatientProfile(stub, params)
	} else {
		fmt.Println("Invoke() did not find func: " + fcn)
		return shim.Error("Received unknown function invocation!")
	}
}

// Function to Create a New Patient Profile.
func (cc *Chaincode) createPatientProfile(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, _, _ := getTxCreatorInfo(stub)
	if !authenticatePatient(creatorOrg, creatorCertIssuer) {
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
	Name := params[1]

	// Generate Asset Key
	assetKey := patientKey + ID

	// Check if Patient exists with Key => assetKey
	patientAsBytes, err := stub.GetState(assetKey)
	if err != nil {
		return shim.Error("Failed to check if Patient exists!")
	} else if patientAsBytes != nil {
		return shim.Error("Patient Already Exists!")
	}

	// Generate Patient from params provided
	patient := &patient{"patientProfile",
		ID, Name, false}

	// Convert to JSON bytes
	patientJSONasBytes, err := json.Marshal(patient)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Patient with Key => assetKey
	err = stub.PutState(assetKey, patientJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(patientJSONasBytes)
}

// Function to Read a Patient Profile.
func (cc *Chaincode) readPatientProfile(stub shim.ChaincodeStubInterface, params []string) sc.Response {
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
	assetKey := patientKey + params[0]

	// Get State of Asset with Key => assetKey
	patientAsBytes, err := stub.GetState(assetKey)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + params[0] + "\"}"
		return shim.Error(jsonResp)
	} else if patientAsBytes == nil {
		jsonResp := "{\"Error\":\"Patient does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Returned on successful execution of the function
	return shim.Success(patientAsBytes)
}

// Function to Update Patient Profile.
func (cc *Chaincode) updatePatientProfile(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, _, _ := getTxCreatorInfo(stub)
	if !authenticatePatient(creatorOrg, creatorCertIssuer) {
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
	Name := params[1]

	// Generate Asset Key
	assetKey := patientKey + ID

	// Check if Patient exists with Key => assetKey
	patientAsBytes, err := stub.GetState(assetKey)
	if err != nil {
		return shim.Error("Failed to get Patient Details!")
	} else if patientAsBytes == nil {
		return shim.Error("Error: Patient Does NOT Exist!")
	}

	// Create Update struct var
	patientToUpdate := patient{}
	err = json.Unmarshal(patientAsBytes, &patientToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Update Patient
	patientToUpdate.Name = Name

	// Convert to JSON bytes
	patientJSONasBytes, err := json.Marshal(patientToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Patient with Key => assetKey
	err = stub.PutState(assetKey, patientJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(patientJSONasBytes)
}

// Function to Activate Patient Profile.
func (cc *Chaincode) activatePatientProfile(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, _, _ := getTxCreatorInfo(stub)
	if !authenticatePHC(creatorOrg, creatorCertIssuer) {
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
	ID := params[0]

	// Generate Asset Key
	assetKey := patientKey + ID

	// Check if Patient exists with Key => assetKey
	patientAsBytes, err := stub.GetState(assetKey)
	if err != nil {
		return shim.Error("Failed to get Patient Details!")
	} else if patientAsBytes == nil {
		return shim.Error("Error: Patient Does NOT Exist!")
	}

	// Create Update struct var
	patientToUpdate := patient{}
	err = json.Unmarshal(patientAsBytes, &patientToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Update Patient
	patientToUpdate.Verified = true

	// Convert to JSON bytes
	patientJSONasBytes, err := json.Marshal(patientToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Patient with Key => assetKey
	err = stub.PutState(assetKey, patientJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(patientJSONasBytes)
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

// Authenticate => Patient
func authenticatePatient(mspID string, certCN string) bool {
	return (mspID == "PatientMSP") && (certCN == "ca.patient.health.com")
}

// Authenticate => PHC
func authenticatePHC(mspID string, certCN string) bool {
	return (mspID == "PHCMSP") && (certCN == "ca.phc.health.com")
}
