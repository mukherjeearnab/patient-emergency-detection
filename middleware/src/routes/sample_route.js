const express = require("express");
const md5 = require("md5");
const JWTmiddleware = require("../helpers/jwtVerifyMiddleware");
// const Contract = require("../../fabric/contracts/contract1");

const router = new express.Router();

router.get("/api/main/sample_route/get/:id", JWTmiddleware, async (req, res) => {
    const ID = req.params.id;

    try {
        res.status(200).send(data);
    } catch (error) {
        res.status(404).send({ message: "Asset NOT found!" });
    }
});

router.get("/api/main", async (req, res) => {
    try {
        res.status(200).send("hello");
    } catch (error) {
        res.status(404).send({ message: "Asset NOT found!" });
    }
});

router.post("/api/main/sample_route/post", JWTmiddleware, async (req, res) => {
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
