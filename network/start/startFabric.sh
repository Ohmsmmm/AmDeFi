#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error
set -e

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1
starttime=$(date +%s)
CC_SRC_LANGUAGE=${1:-"go"}
CC_SRC_LANGUAGE=`echo "$CC_SRC_LANGUAGE" | tr [:upper:] [:lower:]`
if [ "$CC_SRC_LANGUAGE" = "go" -o "$CC_SRC_LANGUAGE" = "golang"  ]; then
	CC_RUNTIME_LANGUAGE=golang
	CC_SRC_PATH=github.com/AmDeFi/go
elif [ "$CC_SRC_LANGUAGE" = "javascript" ]; then
	CC_RUNTIME_LANGUAGE=node # chaincode runtime language is node.js
	CC_SRC_PATH=/opt/GO-PAt/src/github.com/chaincode/AmDeFi/javascript
elif [ "$CC_SRC_LANGUAGE" = "typescript" ]; then
	CC_RUNTIME_LANGUAGE=node # chaincode runtime language is node.js
	CC_SRC_PATH=/opt/GO-PAt/src/github.com/chaincode/AmDeFi/typescript
	echo Compiling TypeScript code into JavaScript ...
	pushd ../chaincode/AmDeFi/typescript
	npm install
	npm run build
	popd
	echo Finished compiling TypeScript code into JavaScript
else
	echo The chaincode language ${CC_SRC_LANGUAGE} is not supported by this script
	echo Supported chaincode languages are: go, javascript, and typescript
	exit 1
fi


# clean the keystore
rm -rf ./hfc-key-store
# clean the wallet  
rm -rf ./../../API-NODE-JS-SDK/blockchain/config/wallet
rm -rf ./../../API-NODE-JS-SDK/blockchain/config/walletorg1
rm -rf ./../../API-NODE-JS-SDK/blockchain/config/walletorg2
rm -rf ./../../API-NODE-JS-SDK/blockchain/config/walletorg3
cat<<EOF
-----------------------------------------------------------------------------------------------------------------------
----------------------clean ./API-NODE-JS-SDK/blockchain/config/wallet Suscess----------------------------------------
-----------------------------------------------------------------------------------------------------------------------
EOF
rm -rf ./operation/org1/wallet
rm -rf ./operation/org2/wallet
rm -rf ./operation/org3/wallet
cat<<EOF 
---------------------------------------------clean wallet Suscess----------------------------------------
EOF
# docker rm -f $(docker ps -aq)

# cat<<EOF 
# ---------------------------------------------rm docker ps ----------------------------------------
# EOF
#launch network:create channel and join peer to channel
cd ../basic-network
./start.sh

cat<<EOF 
============================================================! ! ! KMITL ! ! !=======================================================
EOF

# # launch network; create channel and join peer to channel
# #cd ../first-network
# #echo y | ./byfn.sh down
# #echo y | ./byfn.sh up -a -n -s couchdb
CONFIG_ROOT=/opt/gopath/src/github.com/hyperledger/fabric/peer
#set varible Org1
ORG1_MSPCONFIGPATH=${CONFIG_ROOT}/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
PEER0ORG1_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
PEER1ORG1_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/ca.crt
 
#set varible Org2
ORG2_MSPCONFIGPATH=${CONFIG_ROOT}/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
PEER0ORG2_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
PEER1ORG2_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/ca.crt

#set varible Org3
ORG3_MSPCONFIGPATH=${CONFIG_ROOT}/crypto/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp
PEER0ORG3_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt
PEER1ORG3_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/peerOrganizations/org3.example.com/peers/peer1.org3.example.com/tls/ca.crt

ORDERER_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
set -x

echo "Installing smart contract on peer0.org1.example.com" 
docker exec \
  -e CORE_PEER_LOCALMSPID=Org1MSP \
  -e CORE_PEER_ADDRESS=peer0.org1.example.com:7051 \
  -e CORE_PEER_MSPCONFIGPATH=${ORG1_MSPCONFIGPATH} \
  -e CORE_PEER_TLS_ROOTCERT_FILE=${PEER0ORG1_TLS_ROOTCERT_FILE} \
  cli \
  peer chaincode install \
    -n AmDeFi \
    -v 1.0 \
    -p "$CC_SRC_PATH" \
    -l "$CC_RUNTIME_LANGUAGE"

echo "Installing smart contract on peer0.org2.example.com"
docker exec \
  -e CORE_PEER_LOCALMSPID=Org2MSP \
  -e CORE_PEER_ADDRESS=peer0.org2.example.com:7051 \
  -e CORE_PEER_MSPCONFIGPATH=${ORG2_MSPCONFIGPATH} \
  -e CORE_PEER_TLS_ROOTCERT_FILE=${PEER0ORG2_TLS_ROOTCERT_FILE} \
  cli \
  peer chaincode install \
    -n AmDeFi \
    -v 1.0 \
    -p "$CC_SRC_PATH" \
    -l "$CC_RUNTIME_LANGUAGE"

echo "Installing smart contract on peer0.org3.example.com"
docker exec \
  -e CORE_PEER_LOCALMSPID=Org3MSP \
  -e CORE_PEER_ADDRESS=peer0.org3.example.com:7051 \
  -e CORE_PEER_MSPCONFIGPATH=${ORG3_MSPCONFIGPATH} \
  -e CORE_PEER_TLS_ROOTCERT_FILE=${PEER0ORG3_TLS_ROOTCERT_FILE} \
  cli \
  peer chaincode install \
    -n AmDeFi \
    -v 1.0 \
    -p "$CC_SRC_PATH" \
    -l "$CC_RUNTIME_LANGUAGE"
echo "Installing smart contract on peer1.org1.example.com"
docker exec \
  -e CORE_PEER_LOCALMSPID=Org1MSP \
  -e CORE_PEER_ADDRESS=peer1.org1.example.com:7051 \
  -e CORE_PEER_MSPCONFIGPATH=${ORG1_MSPCONFIGPATH} \
  -e CORE_PEER_TLS_ROOTCERT_FILE=${PEER1ORG1_TLS_ROOTCERT_FILE} \
  cli \
  peer chaincode install \
    -n AmDeFi \
    -v 1.0 \
    -p "$CC_SRC_PATH" \
    -l "$CC_RUNTIME_LANGUAGE"
echo "Installing smart contract on peer1.org2.example.com"
docker exec \
  -e CORE_PEER_LOCALMSPID=Org2MSP \
  -e CORE_PEER_ADDRESS=peer1.org2.example.com:7051 \
  -e CORE_PEER_MSPCONFIGPATH=${ORG2_MSPCONFIGPATH} \
  -e CORE_PEER_TLS_ROOTCERT_FILE=${PEER1ORG2_TLS_ROOTCERT_FILE} \
  cli \
  peer chaincode install \
    -n AmDeFi \
    -v 1.0 \
    -p "$CC_SRC_PATH" \
    -l "$CC_RUNTIME_LANGUAGE"
echo "Installing smart contract on peer1.org3.example.com"
docker exec \
  -e CORE_PEER_LOCALMSPID=Org3MSP \
  -e CORE_PEER_ADDRESS=peer1.org3.example.com:7051 \
  -e CORE_PEER_MSPCONFIGPATH=${ORG3_MSPCONFIGPATH} \
  -e CORE_PEER_TLS_ROOTCERT_FILE=${PEER1ORG3_TLS_ROOTCERT_FILE} \
  cli \
  peer chaincode install \
    -n AmDeFi \
    -v 1.0 \
    -p "$CC_SRC_PATH" \
    -l "$CC_RUNTIME_LANGUAGE"


cat<<EOF 
============================================================! ! ! SLEEP NOW! ! !=======================================================
EOF
export FABRIC_START_TIMEOUT=5
#echo ${FABRIC_START_TIMEOUT}
sleep ${FABRIC_START_TIMEOUT}

echo "Instantiating smart contract on AmDeFi"
docker exec \
 -e CORE_PEER_LOCALMSPID=Org1MSP \
 -e CORE_PEER_ADDRESS=peer0.org1.example.com:7051 \
 -e CORE_PEER_MSPCONFIGPATH=${ORG1_MSPCONFIGPATH} \
 -e CORE_PEER_TLS_ROOTCERT_FILE=${PEER0ORG1_TLS_ROOTCERT_FILE} \
 cli \
 peer chaincode instantiate \
   -o orderer.example.com:7050 \
   -C origin \
   -n AmDeFi \
   -l "$CC_RUNTIME_LANGUAGE" \
   -v 1.0 \
      -c '{"function":"init","Args":[]}' \
      -P "AND ('Org1MSP.member','Org2MSP.member','Org3MSP.member')"

   
cat<<EOF 
==========================================================! ! ! BLOCKCHAIN OPEN ! ! !=======================================================
============================================================! ! ! SUCCESSFULLY ! ! !=======================================================
EOF

