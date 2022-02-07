FROM storezhang/alpine


LABEL author="storezhang<华寅>"
LABEL email="storezhang@gmail.com"
LABEL qq="160290688"
LABEL wechat="storezhang"
LABEL description="Drone持续集成Maven插件，支持测试、打包、发布等常规功能"


# 复制文件
COPY cos /bin


RUN set -ex \
    \
    \
    \
    # 增加执行权限
    && chmod +x /bin/cos \
    \
    \
    \
    && rm -rf /var/cache/apk/*



# 执行命令
ENTRYPOINT /bin/cos
