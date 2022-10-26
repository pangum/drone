# drone

`Drone`持续集成插件，功能有

- 程序压缩
- 默认优化编译标志
- 编译时增加调试信息
    - 名称
    - 编译时间
    - 分支
    - 版本
    - 运行时
- 编译模式
    - 正式版
    - 调试版

## 使用方法

使用非常简单，在`Drone`的配置文件`.drone.yaml`中作出如下配置即可

```yaml
steps:
  - name: 编译
    image: dronestock/pangu
    pull: always
    settings:
      output: server
```

## 捐助

![支持宝](https://github.com/storezhang/donate/raw/master/alipay-small.jpg)
![微信](https://github.com/storezhang/donate/raw/master/weipay-small.jpg)

## 感谢Jetbrains

本项目通过`Jetbrains开源许可IDE`编写源代码，特此感谢
[![Jetbrains图标](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.png)](https://www.jetbrains.com/?from=pangum/drone)
