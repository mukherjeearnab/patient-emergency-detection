const express = require("express");
const JWTmiddleware = require("../helpers/jwtVerifyMiddleware");
// const DetectionContract = require("../../fabric/contracts/detection");

const router = new express.Router();

router.get("/api/location/set/:id", JWTmiddleware, async (req, res) => {
    try {
        // let data = await DetectionContract.CheckReading(
        //     { username: req.user.username, organization: req.user.organization },
        //     [req.params.id, req.query.cipher]
        // );
        res.status(200).send({
            location: req.body.location,
            message: "Got Location, Init Ememgency.",
        });
    } catch (error) {
        res.status(500).send({ message: "Server Error!" });
    }
});

module.exports = router;
