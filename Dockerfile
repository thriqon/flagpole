FROM golang:1.27.1 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /flagpole

FROM scratch
CMD ["/flagpole"]
COPY --from=builder /flagpole /flagpole
