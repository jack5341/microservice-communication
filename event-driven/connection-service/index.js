import express from "express";
const app = express();

import { logger } from "./configs/configs.js";
import CONSTANTS from "./constants/constants.js";

app.get("/", (req, res) => {
    res.send("Hello World!");
});

app.listen(CONSTANTS.PORT, () => {
    logger.info("Server is running on port " + CONSTANTS.PORT);
});
