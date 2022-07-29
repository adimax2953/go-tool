# go-tool#

Packet go-tool implements a way to use go more easily

## Kafka Note Docker...

建立topic
docker-compose exec broker \
kafka-topics --bootstrap-server broker:9092 \
             --create \
             --topic test

docker-compose exec broker \
 --list --zookeeper localhost:2181


docker-compose exec broker \
 --describe --zookeeper  localhost:2181 --topic test




對topic發送消息
docker-compose exec --interactive --tty broker \
kafka-console-producer --bootstrap-server broker:9092 \
                       --topic test

aaaaaaddasdasd

閱讀topic內的消息
docker-compose exec --interactive --tty broker \
kafka-console-consumer --bootstrap-server broker:9092 \
                       --topic test\
                       --from-beginning



