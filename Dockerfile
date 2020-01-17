FROM golang:latest

# 设置代理
ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/RESTfulGo
COPY . $GOPATH/src/github.com/RESTfulGo
RUN pwd
RUN make
#RUN go build -v .

EXPOSE 7777
ENTRYPOINT ["./RESTfulGo"]