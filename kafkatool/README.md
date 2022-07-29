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
                       
                       
                       
### kafka 注意事項

#### Partition

正常來說，每個partition 能處理的吞吐為幾MB/s（仍需要基於根據本地環境測試後獲取準確指標），增加更多的partitions意味著：

更高的並行度與吞吐
可以擴展更多的（同一個consumer group中的）consumers
若是集群中有較多的brokers，則可更大程度上利用閒置的brokers
但是會造成Zookeeper的更多選舉
也會在Kafka中打開更多的文件
 

調整準則：

一般來說，若是集群較小（小於6個brokers），則配置2 x broker數的partition數。在這裡主要考慮的是之後的擴展。若是集群擴展了一倍（例如12個），則不用擔心會有partition不足的現象發生
一般來說，若是集群較大（大於12個），則配置1 x broker 數的partition數。因為這裡不需要再考慮集群的擴展情況，與broker數相同的partition數已經足夠應付常規場景。若有必要，則再手動調整
考慮最高峰吞吐需要的並行consumer數，調整partition的數目。若是應用場景需要有20個（同一個consumer group中的）consumer並行消費，則據此設置為20個partition
考慮producer所需的吞吐，調整partition數目（如果producer的吞吐非常高，或是在接下來兩年內都比較高，則增加partition的數目）

# 只能增加不能減少

#### Replication Factor

此參數決定的是records複製的數目，建議至少 設置為2，一般是3，最高設置為4。更高的replication factor（假設數目為N）意味著：

系統更穩定（允許N-1個broker宕機）
更多的副本（如果acks=all，則會造成較高的延時）
系統磁盤的使用率會更高（一般若是RF為3，則相對於RF為2時，會佔據更多50% 的磁盤空間）
 

調整準則：

以3為起始（當然至少需要有3個brokers，同時也不建議一個Kafka 集群中節點數少於3個節點）
如果replication 性能成為了瓶頸或是一個issue，則建議使用一個性能更好的broker，而不是降低RF的數目

# prod環境中起始值，請勿少於3

慎用，得考慮資料的增長速度，避免儲存資料量太大，造成空間不足

相對的會延長kafka的寫入時間




