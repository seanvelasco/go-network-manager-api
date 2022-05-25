FROM go:18-alpine

WORKDIR network-manager-api/

ENV DBUS_SYSTEM_BUS_ADDRESS=unix:path=/host/run/dbus/system_bus_socket

COPY go.mod .
COPY go.sum .
COPY *.go .
COPY networkmanager .

RUN go mod download

RUN go build -o networkmanager

EXPOSE 8888

CMD [ "/networkmanager" ]

