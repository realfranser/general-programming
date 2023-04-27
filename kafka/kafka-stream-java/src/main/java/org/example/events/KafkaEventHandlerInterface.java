package org.example.events;

/** Interface for all kinds of Kafka events. */
public interface KafkaEventHandlerInterface {
    void processKafkaEvent(final byte[] data);
}
