# Kafka Note Docker...

建立topic
docker-compose exec broker \
kafka-topics --bootstrap-server broker:9092 \
             --create \
             --topic test

docker-compose exec broker \
 --list --zookeeper localhost:2181

docker-compose exec broker \
 --describe --zookeeper  localhost:2181 --topic topicName

查看一個Topic 有多少個Partition
kakfa-topic.sh --list topic topicName --zookeeper zookeeper.servers.list


對topic發送消息
docker-compose exec --interactive --tty broker \
kafka-console-producer --bootstrap-server broker:9092 \
                       --topic topicName

aaaaaaddasdasd

閱讀topic內的消息
docker-compose exec --interactive --tty broker \
kafka-console-consumer --bootstrap-server broker:9092 \
                       --topic topicName\
                       --from-beginning
                       
                       
                       
                       
## Kafka 系統的角色
- Broker ：一台kafka服務器就是一個broker。一個集群由多個broker組成。一個broker可以容納多個topic
- topic： 可以理解為一個MQ消息隊列的名字
- Partition：為了實現擴展性，一個非常大的topic可以分佈到多個 broker（即服務器）上，一個topic可以分為多個partition，每個partition是一個有序的隊列。 partition中的每條消息 都會被分配一個有序的id（offset）。 kafka只保證按一個partition中的順序將消息發給consumer，不保證一個topic的整體 （多個partition間）的順序。也就是說，一個topic在集群中可以有多個partition，那麼分區的策略是什麼？ (消息發送到哪個分區上，有兩種基本的策略，一是採用Key Hash算法，一是採用Round Robin算法)
- Offset：kafka的存儲文件都是按照offset.kafka來命名，用offset做名字的好處是方便查找。例如你想找位於2049的位置，只要找到2048.kafka的文件即可。當然the first offset就是00000000000.kafka
- Producer ：消息生產者，就是向kafka broker發消息的客戶端。
- Consumer ：消息消費者，向kafka broker取消息的客戶端
- Consumer Group （CG）：消息系統有兩類，一是廣播，二是訂閱發布。廣播是把消息發送給所有的消費者；發布訂閱是把消息只發送給訂閱者。 Kafka通過Consumer Group組合實現了這兩種機制： 實現一個topic消息廣播（發給所有的consumer）和單播（發給任意一個consumer）。一個 topic可以有多個CG。 topic的消息會復制（不是真的複制，是概念上的）到所有的CG，但每個CG只會把消息發給該CG中的一個 consumer（這是實現一個Topic多Consumer的關鍵點：為一個Topic定義一個CG，CG下定義多個Consumer）。如果需要實現廣播，只要每個consumer有一個獨立的CG就可以了。要實現單播只要所有的consumer在同一個CG。用CG還 可以將consumer進行自由的分組而不需要多次發送消息到不同的topic。典型的應用場景是，多個Consumer來讀取一個Topic(理想情況下是一個Consumer讀取Topic的一個Partition）,那麼可以讓這些Consumer屬於同一個Consumer Group即可實現消息的多Consumer並行處理，原理是Kafka將一個消息發佈出去後，ConsumerGroup中的Consumers可以通過Round Robin的方式進行消費(Consumers之間的負載均衡使用Zookeeper來實現)


## Zookeeper 在Kakfa中扮演的角色
- kafka使用zookeeper來實現動態的集群擴展，不需要更改客戶端（producer和consumer）的配置。 
- broker會在zookeeper註冊並保持相關的元數據（topic，partition信息等）更新。
- 而客戶端會在zookeeper上註冊相關的watcher。一旦zookeeper發生變化，客戶端能及時感知並作出相應調整。這樣就保證了添加或去除broker時，各broker間仍能自動實現負載均衡。這裡的客戶端指的是Kafka的消息生產端(Producer)和消息消費端(Consumer)
- Broker端使用zookeeper來註冊broker信息,以及監測partition leader存活性.
- Consumer端使用zookeeper用來註冊consumer信息,其中包括consumer消費的partition列表等,同時也用來發現broker列表,並和partition leader建立socket連接,並獲取消息.
- Zookeer和Producer沒有建立關係，只和Brokers、Consumers建立關係以實現負載均衡，即同一個Consumer Group中的Consumers可以實現負載均衡


                       
## kafka 注意事項

### Partition

- Topic有多個Partition，消息分配到某個Partition的依據是Key Hash或者Round Robin。



- 正常來說，每個partition 能處理的吞吐為幾MB/s（仍需要基於根據本地環境測試後獲取準確指標），增加更多的partitions意味著：

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

#### Partition 只能增加不能減少

### Replication Factor

此參數決定的是records複製的數目，建議至少 設置為2，一般是3，最高設置為4。更高的replication factor（假設數目為N）意味著：

系統更穩定（允許N-1個broker宕機）
更多的副本（如果acks=all，則會造成較高的延時）
系統磁盤的使用率會更高（一般若是RF為3，則相對於RF為2時，會佔據更多50% 的磁盤空間）
 

調整準則：

以3為起始（當然至少需要有3個brokers，同時也不建議一個Kafka 集群中節點數少於3個節點）
如果replication 性能成為了瓶頸或是一個issue，則建議使用一個性能更好的broker，而不是降低RF的數目

#### Replication Factor 在 prod環境中起始值，請勿少於3

慎用，得考慮資料的增長速度，避免儲存資料量太大，造成空間不足

相對的會延長kafka的寫入時間


## How to Use
<!-- 建立Topic及其Partition數量 -->
	config.CreateTopic("test1", 10)

<!-- 一次寫入一大批僅有value的資料 -->
1.
	s := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		s[i] = "value " + gotool.IntToStr(i)
	}
	config.WriteMessages("test3", s...)

2.
    config.WriteMessages("test3", "da", "da", "der", "ma", "te", "sen")

<!-- 一次寫入一大批的帶著key&Value資料 -->
	 m := map[string]string{}
	 for i := 0; i < 10000; i++ {
	 	m[gotool.IntToStr(i)+"@player"] = "value " + gotool.IntToStr(i)
	 }
	 config.WriteMessagesKeyValue("test1", m)