const express = require("express");
const JWTmiddleware = require("../helpers/jwtVerifyMiddleware");
const DetectionContract = require("../../fabric/contracts/detection");

const router = new express.Router();

router.get("/api/detection/check/:id", JWTmiddleware, async (req, res) => {
    try {
        let data = await DetectionContract.CheckReading(
            { username: req.user.username, organization: req.user.organization },
            [req.params.id, req.body.data.Cipher]
        );
        res.status(200).send({
            detection: data,
        });
    } catch (error) {
        res.status(404).send({ message: "Asset NOT found!" });
    }
});

module.exports = router;
