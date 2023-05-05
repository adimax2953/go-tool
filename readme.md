# go-tool


## Install

```console
go get -u -v github.com/adimax2953/go-tool
```

## File Archetecture 
### argtool
  - 參數檢查
  
### bfttool
  - encrypt
    - 加解密相關 
  - geo_ip
    - ip檢查

### googledrivertool
  - google drive 相關功能

### httptool
  - fasthttp為基底做的api client

### iotool
  - 檔案操作相關 

### jsontool
  - 以 github.com/json-iterator/go 開發的json 加解密工具
  
### kafkatool
  - kafka相關開發工具
  - docker-compose
    - 放docker-compose.yml 

### nsqtool
  - nsq相關開發工具
  - docker-compose
     - 放docker-compose.yml 
  
### rmqtool
  - rocketmq相關開發工具
  - docker-compose
    - 放docker-compose.yml 
  
### randtool
- 基於梅森旋轉鏈(github.com/seehuhn/mt19937)開發的rng工具 


### Dependency

- testify

  ```console
    go get -u -v github.com/stretchr/testify
  ```

- log-tool

  ```console
    go get -u -v github.com/adimax2953/log-tool
  ```

- rocketmq-client-go

  ```console
    go get -u -v github.com/apache/rocketmq-client-go/v2
  ```

- json-iterator

  ```console
    go get -u -v github.com/json-iterator/go
  ```