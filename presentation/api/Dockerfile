FROM golang:1.16-alpine

WORKDIR /go/src/

COPY go.mod .
COPY go.sum .
COPY ["docs", "docs"]
COPY ["application", "application"]
COPY ["domain", "domain"]
COPY ["infrastructure", "infrastructure"]
COPY ["presentation/api", "."]

RUN ls
RUN go clean --modcache
RUN GOOS=linux go build main.go

EXPOSE 5000
ENTRYPOINT ["./main"]