ARG go_version=1.19
ARG alpine_version=3.16

# Server binary builder
FROM golang:${go_version}-alpine${alpine_version} as server_builder
ARG go_swagger_version=0.27.0

RUN apk add --no-cache git
RUN apk add --update --no-cache curl build-base
RUN curl -sSLo /usr/local/bin/swagger https://github.com/go-swagger/go-swagger/releases/download/v${go_swagger_version}/swagger_linux_amd64
RUN chmod +x /usr/local/bin/swagger

RUN apk update && apk add openssh

ENV REPO_DIR ${GOPATH}/src/quocbang/go-swagger-api
ENV SERVER_DIR ${REPO_DIR}/server
ENV APP_NAME go-swagger-api

COPY . ${REPO_DIR}/
COPY ./swagger.yml ${REPO_DIR}/
COPY ./server ${SERVER_DIR}/

WORKDIR ${SERVER_DIR}

# RUN go generate .
RUN ls -l
# RUN go mod download
# RUN go vet ./...
RUN go build -race -ldflags "-extldflags '-static'" -o /opt/go/server ./server/swagger/cmd/go-swagger-api-server

RUN ls -l /opt/go/server

COPY /opt/go/server /root/server

WORKDIR /root/

CMD ["/bin/sh"]