# drone

`Drone`持续集成插件

## 使用方法

使用非常简单，在`Drone`的配置文件`.drone.yaml`中作出如下配置即可

```yaml
steps:
  - name: 编译
    image: ccr.ccs.tencentyun.com/pangum/pangu
    # image: dronestock/pangu
    pull: always
    settings:
      output: server
```
