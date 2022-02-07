# cos

Drone持续集成腾讯云对象存储插件，功能

- 清理存储空间
- 上传文件
- 配置静态网站

## 使用

非常简单，只需要在`.drone.yml`里增加配置

```yaml
- name: 上传到腾讯云
  image: dronestock/cos
  setttings:
    secret_id: xxx
    secret_key: xxx
    base_url: xxx
```
