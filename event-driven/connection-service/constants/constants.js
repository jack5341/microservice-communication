export default {
    BROKER_URL: process.env.BROKER_UR || "",
    TOPIC_NAME: process.env.TOPIC_NAME || "",
    EVENT_TITLE: process.env.EVENT_TITLE || "",
    BROKER_USER: process.env.BROKER_USER || "",
    BROKER_PASS: process.env.BROKER_PASS || "",
    PORT: process.env.PORT || 3000,
    SERVICE_NAME: process.env.SERVICE_NAME || "connection-service",
};
