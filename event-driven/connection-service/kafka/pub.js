import { broker } from "./client.js";
import { logger } from "../configs/configs";

export async function Publisher(topic, event) {
    if (!topic) {
        logger.warning("topic is not defined");
        return;
    }

    if (!event.key || !event.value) {
        logger.warning("event is not defined");
        return;
    }

    try {
        const producer = broker.producer();

        await producer.connect();
        await producer.send({
            topic: topic,
            messages: [...event],
        });

        await producer.disconnect();

        logger.info("event commited successfully");
    } catch (error) {
        logger.error(error);
    }
}
