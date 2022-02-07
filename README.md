# maven

Drone持续集成Maven插件，功能

- 测试
- 打包
- 发布

## 使用

非常简单，只需要在`.drone.yml`里增加配置

```yaml
- name: 发布到Maven仓库
  image: dronestock/maven
  setttings:
    username: xxx
    password: xxx
    token: xxx
```
