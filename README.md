# cos

Drone持续集成腾讯云对象存储插件，功能

- 清理存储空间
- 上传文件
- 配置静态网站

## 使用

非常简单，只需要在`.drone.yml`里增加配置

```yaml
steps:
  - name: 上传到腾讯云
  image: ccr.ccs.tencentyun.com/dronestock/cos
  settings:
    secret:
      id: xxx
      key: xxx
```

更多使用教程，请参考[使用文档](https://www.dronestock.tech/plugin/stock/cos)

## 交流

![微信群](https://www.dronestock.tech/communication/wxwork.jpg)

## 捐助

![支持宝](https://github.com/storezhang/donate/raw/master/alipay-small.jpg)
![微信](https://github.com/storezhang/donate/raw/master/weipay-small.jpg)

## 感谢`Jetbrains`

本项目通过`Jetbrains开源许可IDE`编写源代码，特此感谢

[![Jetbrains图标](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg)](https://www.jetbrains.com/?from=dronestock/cos)
