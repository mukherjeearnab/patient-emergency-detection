#!/bin/bash
ls
cd ..
export IMAGE_TAG=1.4
echo "Generating cryto material for peers..."

mkdir ./channel-artifacts

cryptogen generate --config=./crypto-config.yaml

echo "Generating channel artifacts and genesis block..."
configtxgen -profile HEALTHOrdererGenesis -outputBlock ./channel-artifacts/genesis.block
configtxgen -profile HEALTHChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID mainchannel

CURRENT_DIR=$PWD
cd ./base
cp docker-compose-base-template.yaml docker-compose-base.yaml
OPTS="-i"
cd $CURRENT_DIR
cd ./crypto-config/peerOrganizations/patient.health.com/ca/
PRIV_KEY=$(ls *_sk)
cd $CURRENT_DIR
cd ./base
sed $OPTS "s/CA1_PRIVATE_KEY/${PRIV_KEY}/g" docker-compose-base.yaml

cd $CURRENT_DIR
cd ./crypto-config/peerOrganizations/phc.health.com/ca/
PRIV_KEY=$(ls *_sk)
cd $CURRENT_DIR
cd ./base
sed $OPTS "s/CA2_PRIVATE_KEY/${PRIV_KEY}/g" docker-compose-base.yaml

cd $CURRENT_DIR
cd ./crypto-config/peerOrganizations/hospital.health.com/ca/
PRIV_KEY=$(ls *_sk)
cd $CURRENT_DIR
cd ./base
sed $OPTS "s/CA3_PRIVATE_KEY/${PRIV_KEY}/g" docker-compose-base.yaml

cd $CURRENT_DIR
cd ./crypto-config/peerOrganizations/healthadmin.health.com/ca/
PRIV_KEY=$(ls *_sk)
cd $CURRENT_DIR
cd ./base
sed $OPTS "s/CA4_PRIVATE_KEY/${PRIV_KEY}/g" docker-compose-base.yaml

cd $CURRENT_DIR
cd ./scripts
bash ./auxiliary/ccp-generate.sh

# Install Go dependencies
bash ./auxiliary/go-mod.sh
