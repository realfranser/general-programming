package org.example.events;

import java.io.IOException;
import java.util.Map;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.AllArgsConstructor;

@AllArgsConstructor
public abstract class KafkaEventHandlerAbstract<T> implements KafkaEventHandlerInterface {
    private final static String KAFKA_EVENT_ERROR_MESSAGE =
            "The following error occurred when processing a kafka event: %s";
    private final static String PROPERTIES_ATTRIBUTE = "properties";
    private final static String TIMESTAMP_ATTRIBUTE = "sentAt";

    protected final ObjectMapper objectMapper;

    protected Map<String, String> readProperties(final byte[] data) throws IOException {
        final JsonNode properties = objectMapper.readTree(data).get(PROPERTIES_ATTRIBUTE);
        return objectMapper.readValue(properties.traverse(), Map.class);
    }

    protected String readTimestamp(final byte[] data) throws IOException {
        return objectMapper.readTree(data).get(TIMESTAMP_ATTRIBUTE).asText();
    }

    @Override
    public void processKafkaEvent(final byte[] data) {
        try {
            processInternalKafkaEvent(readProperties(data), readTimestamp(data));
        } catch (IOException e) {
            System.out.printf(KAFKA_EVENT_ERROR_MESSAGE, e.getMessage());
        }
    }

    protected abstract void processInternalKafkaEvent(final Map<String, String> readProperties,
                                                      final String readTimestamp);

    protected abstract T mapProperties(final Map<String, String> properties, final String timestamp);
}
