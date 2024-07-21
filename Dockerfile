FROM node:latest as web

WORKDIR /web
COPY ./web .

RUN npm config set registry https://registry.npmmirror.com

RUN npm install
RUN npm run build

FROM golang:1.22.5-alpine3.20 as server

ENV GOPROXY https://goproxy.cn/

WORKDIR /server
COPY ./server .

RUN go mod download
RUN go version
RUN go build

FROM nginx:alpine

COPY --from=web /web/dist /usr/share/nginx/html
COPY --from=server /server/server /

RUN sed -i '/index.htm;$/a\        try_files $uri $uri/ /index.html;' /etc/nginx/conf.d/default.conf
RUN sed -i '/\/404.html;$/i\    location /api {\n        proxy_pass   http://127.0.0.1:8080;\n        rewrite "^/api/(.*)$" /$1 break;\n    }\n' /etc/nginx/conf.d/default.conf

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x entrypoint.sh

EXPOSE 80
CMD ["/entrypoint.sh"]