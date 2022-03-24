import winston from "winston";
import CONSTANTS from "../constants/constants.js";

const format = winston.format.combine(
    winston.format.timestamp({ format: "YYYY-MM-DD HH:mm:ss:ms" }),
    winston.format.colorize({ all: true }),
    winston.format.printf((info) => `${info.timestamp} ${info.level}: ${info.message}`)
);

export const logger = winston.createLogger({
    format: format,
    defaultMeta: { service: CONSTANTS.SERVICE_NAME },
    transports: [new winston.transports.Console()],
});
