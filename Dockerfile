FROM go:18-alpine

WORKDIR network-manager-api/

COPY go.mod .
COPY go.sum .
COPY *.go .
COPY networkmanager .

RUN go mod download

RUN go build -o networkmanager

EXPOSE 8888

CMD [ "/networkmanager" ]

