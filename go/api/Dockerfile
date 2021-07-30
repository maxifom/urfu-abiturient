FROM golang:1.16-alpine as build-env
RUN apk add --no-cache build-base
WORKDIR /goman/
ADD . .
RUN GOBIN=/goman/apps/ GOOS=linux GOARCH=amd64 go install -v -a -tags netgo -mod vendor -ldflags "-w" ./cmd/...

FROM alpine:3
RUN apk add --no-cache ca-certificates tzdata
COPY --from=build-env /goman/apps/ /apps
WORKDIR /apps/
ENV PATH $PATH:/apps
EXPOSE 8080
ENTRYPOINT [ "server" ]
