#!/bin/bash
CHAINCODE=$1
PEER=$2
ORG=$3
MSP=$4
PORT=$5
VERSION=$6

ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/health.com/orderers/orderer0.health.com/msp/tlscacerts/tlsca.health.com-cert.pem
CORE_PEER_LOCALMSPID=$MSP
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/$ORG.health.com/peers/$PEER.$ORG.health.com/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/$ORG.health.com/users/Admin@$ORG.health.com/msp
CORE_PEER_ADDRESS=$PEER.$ORG.health.com:$PORT
CHANNEL_NAME=mainchannel
CORE_PEER_TLS_ENABLED=true

peer chaincode install -n $CHAINCODE -v $VERSION -p $CHAINCODE >&log.txt

cat log.txt
