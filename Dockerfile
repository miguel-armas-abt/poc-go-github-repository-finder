ARG BINARY=application

FROM golang:1.23-alpine AS builder

ARG BINARY

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o "${BINARY}" ./main.go

FROM scratch
COPY --from=builder /app/"${BINARY}" /"${BINARY}"
EXPOSE 8080
ENTRYPOINT ["/${BINARY}"]