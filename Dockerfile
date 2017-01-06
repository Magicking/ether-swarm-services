FROM golang:1.6

ENV PROJECT_PATH=/go/src/github.com/Magicking/ether-swarm-services/

COPY . $PROJECT_PATH
WORKDIR $PROJECT_PATH

RUN go install -v ./cmd/...

ENTRYPOINT ["/go/bin/etherinfo-server", "--host", "0.0.0.0", "--port", "8090"]
