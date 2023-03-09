FROM golang:1.19

ENV TZ /usr/share/zoneinfo/Asia/Tokyo

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

# ModuleモードをON
ENV GO111MODULE=on

COPY . .
EXPOSE 8080

# Airをインストールし、コンテナ起動時に実行する
RUN go install github.com/cosmtrek/air@latest
# import文を見て、足りないものは追加、すでに使っていないものは削除する
RUN go mod tidy
CMD ["air"]