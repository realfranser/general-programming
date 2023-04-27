package org.example.events;

import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.extern.slf4j.Slf4j;
import org.example.models.UserModel;

import java.util.Map;

@Slf4j
public class UserEventHandler extends KafkaEventHandlerAbstract<UserModel> {

    public UserEventHandler(final ObjectMapper objectMapper) {
        super(objectMapper);
    }

    @Override
    protected void processInternalKafkaEvent(final Map<String, String> properties, final String timestamp) {
        final var userModel = mapProperties(properties, timestamp);
        System.out.printf("UserEvent handled properly with value: %s", userModel.toString());
    }

    @Override
    protected UserModel mapProperties(final Map<String, String> properties, final String timestamp) {
        return UserModel.builder()
                .name(properties.get("name"))
                .age(Integer.parseInt(properties.get("age")))
                .timestamp(timestamp)
                .build();
    }
}
