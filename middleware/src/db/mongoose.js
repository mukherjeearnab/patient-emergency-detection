const mongoose = require("mongoose");

const connectionString = `mongodb://${process.env.MONGODB_URI}`;

mongoose.connect(connectionString, {
    useNewUrlParser: true,
    useCreateIndex: true,
    useFindAndModify: false,
    useUnifiedTopology: true,
});
