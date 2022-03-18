import winston from "winston";
import CONSTANTS from "../constants/constants.js";

export const logger = winston.createLogger({
    format: winston.format.json(),
    defaultMeta: { service: CONSTANTS.SERVICE_NAME },
    transports: [
        new winston.transports.Console(),
        new winston.transports.File({ filename: "error.log", level: "error" }),
        new winston.transports.File({ filename: "combined.log" }),
    ],
});
