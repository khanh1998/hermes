# hermes
An instant messaging app that delivers your messages as fast as Hermes

docker exec -it ksqldb ksql http://ksqldb:8088
DROP STREAM MESSAGE;

CREATE STREAM MESSAGE (senderId INT, clanId INT, channelId INT, message VARCHAR, time INT)
  WITH (KAFKA_TOPIC='message', PARTITIONS=1, FORMAT='JSON');

INSERT INTO MESSAGE (senderId, clanId, channelId, message, time) VALUES (1, 2, 3, 'moving awayt', 1);


SHOW TOPICS;
PRINT message FROM BEGINNING LIMIT 2;

docker exec elasticsearch curl -s -XDELETE "http://localhost:9200/message"

CREATE SINK CONNECTOR SINK_ELASTIC_TEST_02 WITH (
  'connector.class'         = 'io.confluent.connect.elasticsearch.ElasticsearchSinkConnector',
  'connection.url'          = 'http://elasticsearch:9200',
  'key.converter'           = 'org.apache.kafka.connect.storage.StringConverter',
  'value.converter'         = 'org.apache.kafka.connect.json.JsonConverter',
  'value.converter.schemas.enable' = 'false',
  'type.name'               = '_doc',
  'topics'                  = 'message',
  'key.ignore'              = 'true',
  'schema.ignore'           = 'true'
);

describe connector SINK_ELASTIC_TEST_02;

drop connector SINK_ELASTIC_TEST_02;

curl -s http://localhost:9200/message/_search \
    -H 'content-type: application/json' \
    -d '{ "size": 42  }' | jq -c '.hits.hits[]'

curl -s http://localhost:9200/message/_mapping | jq '.'
