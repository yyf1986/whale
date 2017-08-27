#### 为了快速搭建环境，简化操作
# whale v1

- agent
  - [x] 容器的创建、删除、停止、启动（采用host方式）
  - [ ] 容器的创建支持资源限制（cpu，内存等）
  - [x] 启动时，向server注册信息（ip，agent服务端口，资源信息，docker的状态）
  - [x] 随机生成一个端口号
  - [ ] 实时更新宿主机资源信息
  - [ ] image的相关操作
  - [ ] 资源管理，在server请求时候给出应答，并锁定已分配资源（端口）
  - [x] docker服务不可用时，提醒server
  - [x] 容器的信息展示（id、image、name、status、port）
  - [x] 容器删除时，端口需要释放
- server
  - [x] 提供注册接口给agent注册
  - [ ] 自动根据一定的规则，选择合适的agent创建容器
  - [x] 检测注册的agent及其docker，是不是都是可用的，剔除不可用的
  - [x] 提供agent信息展示
  - [x] 提供给web，agent的ip列表，必须是docker服务可用的
- web
  - [x] 选择宿主机，申请端口号，然后将端口号信息填写到环境变量中
  - [x] 选择宿主机自动从server获取，并且docker服务是可用的
  - [ ] 选择image自动从server获取
  
