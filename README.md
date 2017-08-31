#### 为了快速搭建环境，简化操作
# install(需要go环境)
- agent
```
git clone https://github.com/yyf1986/whale-agent.git
cd whale-agent
GO15VENDOREXPERIMENT=1 go build
```
- server
```
git clone https://github.com/yyf1986/whale.git
cd whale
GO15VENDOREXPERIMENT=1 go build
build完后，将生成的whale文件和swagger目录，一起拷贝到需要运行的机器上
```
# run 
- [whale-agent](https://github.com/yyf1986/yyf1986.github.io/blob/master/attachment/whale-agent)
- [whale](https://github.com/yyf1986/yyf1986.github.io/blob/master/attachment/whale.tar.gz)
```
./whale-agent -ip 10.11.20.111 -p 12345 -whaleserver 10.11.20.112:12346
./whale -p 12346 (default port is 12346)
```
# whale v1

- agent
  - [x] 容器的创建、删除、停止、启动（采用host方式）
  - [x] 启动时，向server注册信息（ip，agent服务端口，资源信息，docker的状态）
  - [x] 随机生成一个端口号
  - [x] 提供端口号删除，查询
  - [x] 实时更新宿主机相关状态信息
  - [x] 获取宿主机可用image
  - [x] 获取容器信息（id、image、name、status、port）

- server
  - [x] 提供注册接口给agent注册
  - [x] 检测注册的agent是不是都是可用的，剔除不可用的
  - [x] 对agent的能力进行封装，统一提供服务
  
- web
  - 自动从server获取可用宿主机
  - 根据选择的宿主机，展示可用镜像
  - 是否需要端口，是：点击生成端口；否：不做操作
  - 将端口信息等参数写入env
  
# whale v2
- agent
  - [ ] 容器的创建支持资源限制（cpu，内存等）
  - [ ] image的相关操作
  - [ ] 资源管理，在server请求时候给出应答，并锁定已分配资源（端口）
  - [ ] 容器删除时，端口需要释放
  - [ ] docker服务不可用时，提醒server
  - [ ] 只接受server的请求
- server
  - [ ] 自动根据一定的规则，选择合适的agent创建容器
