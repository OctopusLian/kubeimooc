#第一阶段构建
FROM node:lts as builder
WORKDIR /app/kubeimooc-web
#拷贝源码
COPY . .
#安装依赖
RUN npm install --registry=https://registry.npm.taobao.org

# 开始构建
RUN npm run build:prod

# 第二阶段构建
FROM nginx:alpine
COPY --from=builder /app/kubeimooc-web/dist/ /usr/share/nginx/html/
COPY --from=builder /app/kubeimooc-web/default.conf.template /etc/nginx/templates/default.conf.template
EXPOSE 80
