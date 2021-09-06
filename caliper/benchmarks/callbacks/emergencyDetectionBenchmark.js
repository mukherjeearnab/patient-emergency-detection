"use strict";

module.exports.info = "Create and Read PatientProfile";

const contractID = "emergencydetection_cc";
const version = "1.0";

let bc, ctx, clientArgs, clientIdx;

module.exports.init = async function (blockchain, context, args) {
    bc = blockchain;
    ctx = context;
    clientArgs = args;
    clientIdx = context.clientIdx.toString();
        try {
            const assetID = `TOK_${clientIdx}_${i}_${clientArgs.seed}`;
            console.log(`Client ${clientIdx}: Creating Token ${assetID}`);
            const myArgs = {
                chaincodeFunction: "createToken",
                invokerIdentity: "Admin@phc.health.com",
                chaincodeArguments: [
                    assetID,
                    "AQAAAEdGQQAIAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABsf8ALm3/a/nYvmqxdr/L/1u/Ad4IjTP3KOV1eh19q/Grdnun6I4D8MwA6/tnTUP2ktw5W1XPA/UKlDCz2f1j9w/WJlP2jyPy8hRtKRGPc/pE3rfFhT0L+ZEBXLC/vTP/w9YMW/Dtq/B/CIkUEY0L+fL2eozIfpvwaEFFV9btO/Zhxe24/o6j9Sd8CDpALxPz4nRY+6UcW/jgT8vID1yj8G+R0LegHTv+MjgnMfh8i/2NBQS8So4r9Kj5PwItXMv+L+jNUrH7O/3w1Po/vUhr+DcncNLjZ8v4OFWv1qHrK/IIN5nHGbtD9eSKcZpPimv7jvLY6wRLY/SKnjbaoNfr+wH+WRVzGxvy2ky7VQ96S/h4WIcyonhD8b/v/xjiugv9IQ48S5Ap4/wnxa2IToSz8VSPDPLZivP0xTEm1VbO++CeM1gjIG5L8GrVYF2OvsvxYXuCuAV8U/3DUAEsymur9FQBwbxdjEP/hSvgNpCMs/B7kWzlVp2D8uo+HgSDDKP3J2je0U3OG/dfKy4FrC6L/O7F06SZu/P3DN2jAYy7m/dLmFjJ4xxT+0wLV97rjFPzEbq92TB9c/XwBYC8MKxz9ELRgCJBDpP7jKpR/xCvA/6rhADQTrxb9o6RAUZTrJP96qk71mydC/6gnVhNd3y78kjDk+xTfhv/INytklxMi/"
                ],
            };
            await bc.bcObj.invokeSmartContract(ctx, "token_cc", version, myArgs);
        } catch (error) {
            console.log(`Client ${clientIdx}: Smart Contract threw with error: ${error}`);
        }
    
};

module.exports.run = function () {
    const myArgs = {
        chaincodeFunction: "checkReading",
        invokerIdentity: "Admin@patient.health.com",
        chaincodeArguments: [`TOK_${clientIdx}_${i}_${clientArgs.seed}`, "AQAAAEdGQQAIAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAJlokqt2eeI/3lzc/AcU4z82yDpgSzXrP01+n8F8g/s//SY5beoU2z+6Fpds2TnzP6B9L/JToay/U2NZMypt3D9coatP2/7OP+BO3WUSMpY/3PblAKjDwb/Mk3EPEhfQP47gH9iGA+W/qFp3eUbSyz8I+aDpApfnvxuoc4tKDcs/rSHMUTIp4T9K4d8Pl6XUPzChfq4Y4rg/eCOvLbHk8D/jQwxvh3fhv/FbP4qH/us/pOXqXfYg8L9cbbmU3BfgP1ibOYCqReI/vz58cKGH4D/qMYUnqNvmP9OdngQb0fo/iPZimnu4yT9Y/LQvSu3yP5uMYSRKGtW/AgU1qwcP3j/rV7j8B33fPwpuetQ6XtE/gIs9NPZVlD/gKVh1GxLsP9qZOqOnnuO/gSzPlmfD5z9eBKlyxk3vvzqqQWEF5t0/RWwHCLBazT+QwmHpATKmv4SnKXp279u/wGy+bRm6sL/IcaOotNHuv4jXZh2t5LI/XQx2wPc78L/syHFRzr/NP7K0ujy5BOA/ZOosHex04D8k9LkiRDvkP0kXGBTd7fY/4DJyr3h/0T86OofkBlHwP4BW4TVI9MS/XzzUBADt2D+BxW/ddZDcP2oIAeZgocA/QO2dYR6KkD8I5pqybxbkP5IDVdhenOi/whHIBgXz3j9qWAnbMDntvxbE+/JVONk/"],
    };
    return bc.bcObj.querySmartContract(ctx, contractID, version, myArgs);
};

module.exports.end = async function () {};


