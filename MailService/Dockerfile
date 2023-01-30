FROM golang:1.19 as modules

ADD ./go.mod ./go.sum /cmd/
RUN cd /cmd && go mod download

FROM golang:1.19-alpine as builder

COPY --from=modules /go/pkg /go/pkg

RUN mkdir -p /cmd
COPY . /cmd
WORKDIR /cmd

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -ldflags '-s -w -extldflags "-static"' \
    -o /bin/cmd ./cmd/mail/main.go

FROM alpine:latest as image

COPY --from=builder /bin/cmd .
EXPOSE 8082

ENTRYPOINT [ "./cmd" ]