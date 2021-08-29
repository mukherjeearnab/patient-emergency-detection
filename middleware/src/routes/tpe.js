const express = require("express");
const md5 = require("md5");
const JWTmiddleware = require("../helpers/jwtVerifyMiddleware");
const TPEContract = require("../../fabric/contracts/tpe");

const router = new express.Router();

router.get("/api/tpe/config/get", JWTmiddleware, async (req, res) => {
    try {
        let data = TPEContract.GetConfig(
            {
                name: "tpeconfig_cc",
                channel: "mainchannel",
                function: "getConfig",
            },
            { username: "", organization: "" },
            []
        );
        res.status(200).send(data);
    } catch (error) {
        res.status(404).send({ message: "Asset NOT found!" });
    }
});

router.post("/api/tpe/config/set", JWTmiddleware, async (req, res) => {
    try {
        let Data = req.body.data;
        res.status(200).send({
            data: Data,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! Asset NOT Added!" });
    }
});

module.exports = router;
