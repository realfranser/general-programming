package org.example.events;

import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
import lombok.Getter;

/** Enum with all the events handled by this service. */
public enum EventTypeEnum {
    USER_EVENT("user-event"),
    PRODUCT_EVENT("product-event");

    private static final Map<String, EventTypeEnum> BY_TYPE = new HashMap<>();

    static {
        for (EventTypeEnum eventType : EventTypeEnum.values()) {
            BY_TYPE.put(eventType.eventName, eventType);
        }
    }

    @Getter private final String eventName;

    EventTypeEnum(final String eventName) {
        this.eventName = eventName;
    }

    /**
     * Gets the event type from the event name.
     *
     * @param eventName of the event.
     * @return event type.
     */
    public static EventTypeEnum toEventType(final String eventName) {
        return Objects.requireNonNull(
                BY_TYPE.get(eventName),
                String.format("Unable to find Event type for event with name: %s", eventName));
    }

}
