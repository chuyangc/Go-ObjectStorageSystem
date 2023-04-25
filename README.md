基于Go语言实现的分布式对象存储系统
--
Installation:
```shell
# 安装rabbitmq-server
sudo apt-get install rabbitmq-server

# 下载rabbitmqadmin管理工具
sudo rabbitmq-plugins enable rabbitmq_management
wget localhost:15672/cli/rabbitmqadmin
python3 rabbitmqadmin declare exchange name=apiServers type=fanout
python3 rabbitmqadmin declare exchange name=dataServers type=fanout
sudo rabbitmqctl add_user test test
sudo rabbitmqctl set_permissions -p / test ".*" ".*" ".*"
```
Usage:
```shell
curl -v localhost:12345/objects/xx

curl -v localhost:12345/objects/xx -XPUT -d "This is my oss"
```