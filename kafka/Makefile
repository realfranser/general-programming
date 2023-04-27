# Constants
SHELL := /bin/bash
KAFKA_SCRIPT := ./kafka-commands.sh
TOPIC_NAME := test-topic

# Commands
create-cluster-id:
	@bash $(KAFKA_SCRIPT) create_cluster_id

format-log-directories:
	@bash $(KAFKA_SCRIPT) format_log_directories

start-kafka-server:
	@bash $(KAFKA_SCRIPT) start_kafka_server

create-topic:
	@bash $(KAFKA_SCRIPT) create_topic $(TOPIC_NAME)

describe-topic:
	@bash $(KAFKA_SCRIPT) describe_topic $(TOPIC_NAME)

write-events:
	@bash $(KAFKA_SCRIPT) write_events $(TOPIC_NAME)

create-consumer:
	@bash $(KAFKA_SCRIPT) create_consumer $(TOPIC_NAME)

generate-java-archetype:
	@bash $(KAFKA_SCRIPT) generate_java_archetype

fresh-start-kafka-server: create-cluster-id format-log-directories start-kafka-server
