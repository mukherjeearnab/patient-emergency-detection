#!/bin/bash

function one_line_pem {
    echo "`awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1`"
}

function json_ccp {
    local PP=$(one_line_pem $5)
    local CP=$(one_line_pem $6)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${ORGMSP}/$2/" \
        -e "s/\${P0PORT}/$3/" \
        -e "s/\${CAPORT}/$4/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        ../../connections/ccp-template.json 
}

echo "RUNNING FOR PEER0"

ORG=patient
ORGMSP=Patient
P0PORT=7051
CAPORT=7059
PEERPEM=../crypto-config/peerOrganizations/patient.health.com/tlsca/tlsca.patient.health.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/patient.health.com/ca/ca.patient.health.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P0PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-patient.json

ORG=phc
ORGMSP=PHC
P0PORT=8051
CAPORT=8059
PEERPEM=../crypto-config/peerOrganizations/phc.health.com/tlsca/tlsca.phc.health.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/phc.health.com/ca/ca.phc.health.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P0PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-phc.json
ORG=hospital
ORGMSP=Hospital
P0PORT=9051
CAPORT=9059
PEERPEM=../crypto-config/peerOrganizations/hospital.health.com/tlsca/tlsca.hospital.health.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/hospital.health.com/ca/ca.hospital.health.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P0PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-hospital.json

ORG=healthadmin
ORGMSP=HealthAdmin
P0PORT=10051
CAPORT=10059
PEERPEM=../crypto-config/peerOrganizations/healthadmin.health.com/tlsca/tlsca.healthadmin.health.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/healthadmin.health.com/ca/ca.healthadmin.health.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P0PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-healthadmin.json

echo "RUNNING FOR PEER1"

ORG=patient
ORGMSP=Patient
P1PORT=7053
CAPORT=7059
PEERPEM=../crypto-config/peerOrganizations/patient.health.com/tlsca/tlsca.patient.health.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/patient.health.com/ca/ca.patient.health.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P1PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-patient.json

ORG=phc
ORGMSP=PHC
P1PORT=8053
CAPORT=8059
PEERPEM=../crypto-config/peerOrganizations/phc.health.com/tlsca/tlsca.phc.health.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/phc.health.com/ca/ca.phc.health.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P1PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-phc.json
ORG=hospital
ORGMSP=Hospital
P1PORT=9053
CAPORT=9059
PEERPEM=../crypto-config/peerOrganizations/hospital.health.com/tlsca/tlsca.hospital.health.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/hospital.health.com/ca/ca.hospital.health.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P1PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-hospital.json

ORG=healthadmin
ORGMSP=HealthAdmin
P1PORT=10053
CAPORT=10059
PEERPEM=../crypto-config/peerOrganizations/healthadmin.health.com/tlsca/tlsca.healthadmin.health.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/healthadmin.health.com/ca/ca.healthadmin.health.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P1PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-healthadmin.json
