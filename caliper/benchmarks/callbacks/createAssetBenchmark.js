"use strict";

module.exports.info = "Create PatientProfiles";

const contractID = "patientprofile_cc";
const version = "1.0";

let bc, ctx, clientArgs, clientIdx;

module.exports.init = async function (blockchain, context, args) {
    bc = blockchain;
    ctx = context;
    clientArgs = args;
    clientIdx = context.clientIdx.toString();
};

module.exports.run = async function () {
    for (let i = 0; i < clientArgs.assets; i++) {
        try {
            const assetID = `PP_${clientIdx}_${i}_${Date.now()}`;
            // console.log(`Client ${clientIdx}: Creating PatientProfile ${assetID}`);
            const myArgs = {
                chaincodeFunction: "createPatientProfile",
                invokerIdentity: "Admin@patient.health.com",
                chaincodeArguments: [assetID, "patient1"],
            };
            return await bc.bcObj.invokeSmartContract(ctx, contractID, version, myArgs);
        } catch (error) {
            console.log(`Client ${clientIdx}: Smart Contract threw with error: ${error}`);
        }
    }
};

module.exports.end = async function () {};
