FROM golang:1.23-alpine3.20 AS build

WORKDIR /app

RUN apk add git


ENV APP_NAME="sahtian-backend"
ARG COMMIT_HASH="latest"
ARG COMMIT_TIME="latest"
ARG VERSION="dev"
ARG GOPRIVATE
ARG GOPROXY
ARG GIT_TOKEN
ARG GIT_USER
ARG GONOSUMDB

RUN mkdir /out
COPY . /app/

RUN go build  \
    -ldflags "-Xgithub.com/BigDwarf/sahtian/version.Version=${VERSION} -Xgithub.com/BigDwarf/sahtian/version.CommitHash=${COMMIT_HASH} -Xgithub.com/BigDwarf/sahtian/version.CommitTime=${COMMIT_TIME}"  \
    -o /out/sahtian-backend  \
   github.com/BigDwarf/sahtian/cmd

FROM alpine:3.20

WORKDIR /app

COPY --from=build /out/sahtian-backend /app/
COPY --from=build /app/config/default.yml /app/config/

EXPOSE 8080
EXPOSE 8081
ENTRYPOINT ["/app/sahtian-backend", "serve"]
