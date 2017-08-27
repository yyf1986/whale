#### 为了快速搭建环境，简化操作
# whale v1

- agent
  - [x] 容器的创建、删除、停止、启动（采用host方式）
  - [x] 启动时，向server注册信息（ip，agent服务端口，资源信息，docker的状态）
  - [ ] 实时更新宿主机资源信息
  - [ ] image的相关操作
  - [ ] 资源的管理，在api请求时候给出应答，并锁定以分配资源（端口）
  - [x] docker服务不可用时，提醒server
- server
  - [x] 提供注册接口给agent注册
  - [x] 可以先选择ip来创建容器，端口先自己来填写
  - [ ] 自动根据一定的规则，选择合适的agent创建容器
