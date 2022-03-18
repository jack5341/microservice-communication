import { broker } from "./client.js";
import { logger } from "../configs/configs";

export async function Subscriber(topic) {
    if (!topic) {
        logger.warning("topic is not defined");
        return;
    }

    try {
        const consumer = broker.consumer();

        await consumer.connect();
        await consumer.subscribe({ topic });

        await consumer.run({
            eachMessage: async ({ topic, partition, message }) => {
                logger.info(`consumed event ${message.value.toString()}`);
            },
        });
    } catch (error) {
        logger.error(error);
        return;
    }
}
