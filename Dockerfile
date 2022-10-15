FROM golang:1.18 as build
WORKDIR /go/src/cicd-apiserver
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build /go/src/cicd-apiserver/cicd-apiserver /
ENTRYPOINT ["/cicd-apiserver"]