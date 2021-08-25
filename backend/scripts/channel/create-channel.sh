#!/bin/bash
echo "Creating channel..."
ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/health.com/orderers/orderer0.health.com/msp/tlscacerts/tlsca.health.com-cert.pem
CORE_PEER_LOCALMSPID=PatientMSP
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/patient.health.com/peers/peer0.patient.health.com/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/patient.health.com/users/Admin@patient.health.com/msp
CORE_PEER_ADDRESS=peer0.patient.health.com:7051
CHANNEL_NAME=mainchannel
CORE_PEER_TLS_ENABLED=true
ORDERER_SYSCHAN_ID=syschain

sleep 20
peer channel create -o orderer0.health.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA >&log.txt

cat log.txt

#peer channel create -o orderer0.health.com:7050 /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/health.com/orderers/orderer0.health.com/msp/tlscacerts/tlsca.health.com-cert.pem -c mainchannel -f ./channel-artifacts/channel.tx
sleep 10
echo
echo "Channel created, joining Patient..."
peer channel join -b mainchannel.block
