################################################################################
#  BUILD IMAGE                                                                 #
################################################################################
FROM golang:1.17-alpine AS BUILD_IMAGE

COPY [".", "/opt/source/"]

WORKDIR /opt/source

RUN apk add --no-cache build-base

RUN GOOS=linux GOARCH=amd64 go build -a -ldflags="-w -s" -o /opt/bin/backend cmd/backend/main.go

################################################################################
#  RUNTIME IMAGE                                                               #
################################################################################
FROM alpine:3.12

COPY --from=BUILD_IMAGE /opt/bin/backend /opt/bin/backend

ENTRYPOINT ["/opt/bin/backend"]