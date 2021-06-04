FROM alpine

MAINTAINER storezhang "storezhang@gmail.com"
LABEL architecture="AMD64/x86_64" version="latest" build="2021-06-04"
LABEL Description="腾讯云对象存储插件，支持文件上传以及静态网页功能。"



# 复制文件
COPY cos /bin



ENTRYPOINT /bin/cos
