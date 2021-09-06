const TPEContract = require("../fabric/contracts/tpe");

const main = async () => {
    try {
        await TPEContract.SetConfig({ username: "admin", organization: "phc" }, [5, 28.1]);
    } catch (error) {
        console.error("TPE Config set!", error);
    }
};

main();
