#!/bin/bash

readonly DEFAULT_HOST=localhost
readonly DEFAULT_PORT=9092

create_cluster_id() {
	KAFKA_CLUSTER_ID=$($KAFKA_HOME/bin/kafka-storage.sh "random-uuid")
	export KAFKA_CLUSTER_ID=$KAFKA_CLUSTER_ID
	echo "The new kafka cluster id is: " $KAFKA_CLUSTER_ID
}

format_log_directories() {
	$KAFKA_HOME/bin/kafka-storage.sh "format" -t $KAFKA_CLUSTER_ID -c $KAFKA_HOME/config/kraft/server.properties
}

start_kafka_server() {
	$KAFKA_HOME/bin/kafka-server-start.sh $KAFKA_HOME/config/kraft/server.properties
}

create_topic() {
    $KAFKA_HOME/bin/kafka-topics.sh --create --topic $1 \
        --bootstrap-server $DEFAULT_HOST:$DEFAULT_PORT
}

describe_topic() {
    $KAFKA_HOME/bin/kafka-topics.sh --describe --topic $1 \
        --bootstrap-server $DEFAULT_HOST:$DEFAULT_PORT
}

write_events() {
    $KAFKA_HOME/bin/kafka-console-producer.sh --topic $1 \
        --bootstrap-server $DEFAULT_HOST:$DEFAULT_PORT
}

create_consumer() {
    $KAFKA_HOME/bin/kafka-console-consumer.sh --topic $1 --from-beginning \
        --bootstrap-server $DEFAULT_HOST:$DEFAULT_PORT

}

case $1 in
    "create_cluster_id")
        create_cluster_id
        ;;
    "format_log_directories")
        format_log_directories
        ;;
    "start_kafka_server")
        start_kafka_server
        ;;
    "create_topic")
        create_topic $2
        ;;
    "describe_topic")
        describe_topic $2
        ;;
    "write_events")
        write_events $2
        ;;
    "create_consumer")
        create_consumer $2
        ;;
    *)
        echo "kafka-commands.sh: Invalid Choice"
        ;;
esac
