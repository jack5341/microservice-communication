import { Kafka } from "kafkajs";

import CONSTANTS from "../constants/constants.js";

export const broker = new Kafka({
    broker: [...CONSTANTS.BROKER_URL.split(",")],
    sasl: {
        mechanism: "plain",
        username: CONSTANTS.BROKER_USER,
        password: CONSTANTS.BROKER_PASS,
    },
});
