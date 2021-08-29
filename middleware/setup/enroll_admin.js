const adminInfo = [
    { name: "patient", ca: "ca.patient.health.com", msp: "PatientMSP" },
    { name: "phc", ca: "ca.phc.health.com", msp: "PHCMSP" },
    { name: "hospital", ca: "ca.hospital.health.com", msp: "HospitalMSP" },
    { name: "healthadmin", ca: "ca.healthadmin.health.com", msp: "HealthAdminMSP" },
];

const FabricAPI = require("../fabric/api");

const main = async () => {
    try {
        adminInfo.forEach(async (organization) => {
            console.log("Enrolling", organization);
            await FabricAPI.Account.EnrollAdmin(organization);
        });
    } catch (error) {
        console.error("Enrollment Failed!", error);
    }
};

main();
