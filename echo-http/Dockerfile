FROM golang:1.18

ENV TZ /usr/share/zoneinfo/Asia/Tokyo

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

# ModuleモードをON
ENV GO111MODULE=on

COPY . .
EXPOSE 8080

# Airをインストールし、コンテナ起動時に実行する
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]