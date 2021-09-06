"use strict";

module.exports.info = "Create and Read PatientProfile";

const contractID = "patientprofile_cc";
const version = "1.0";

let bc, ctx, clientArgs, clientIdx;

module.exports.init = async function (blockchain, context, args) {
    bc = blockchain;
    ctx = context;
    clientArgs = args;
    clientIdx = context.clientIdx.toString();
    for (let i = 0; i < clientArgs.assets; i++) {
        try {
            const assetID = `PP_${clientIdx}_${i}_${clientArgs.seed}`;
            console.log(`Client ${clientIdx}: Creating PatientProfile ${assetID}`);
            const myArgs = {
                chaincodeFunction: "createPatientProfile",
                invokerIdentity: "Admin@patient.health.com",
                chaincodeArguments: [assetID, "patient"],
            };
            await bc.bcObj.invokeSmartContract(ctx, contractID, version, myArgs);
        } catch (error) {
            console.log(`Client ${clientIdx}: Smart Contract threw with error: ${error}`);
        }
    }
};

module.exports.run = function () {
    const randomId = Math.floor(Math.random() * clientArgs.assets);
    const myArgs = {
        chaincodeFunction: "readPatientProfile",
        invokerIdentity: "Admin@patient.health.com",
        chaincodeArguments: [`PP_${clientIdx}_${randomId}_${clientArgs.seed}`],
    };
    return bc.bcObj.querySmartContract(ctx, contractID, version, myArgs);
};

module.exports.end = async function () {};
