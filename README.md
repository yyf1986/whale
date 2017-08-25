# dockerapi
- agent 实现docker的管理，信息收集
- api 对外提供服务

```
graph TD
  A[Web]--> B{api}
  B --> |one| C[agent]
  B --> |two| D[agent]
  B --> |three| E[agent]
```
