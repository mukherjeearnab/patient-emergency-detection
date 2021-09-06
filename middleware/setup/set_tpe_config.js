require("dotenv").config();

const TPEContract = require("../fabric/contracts/tpe");

const main = async () => {
    try {
        await TPEContract.SetConfig({ username: "admin", organization: "phc" }, [
            process.env.TPE_N,
            process.env.TPE_THETA,
        ]);
    } catch (error) {
        console.error("TPE Config set!", error);
    }
};

main();
