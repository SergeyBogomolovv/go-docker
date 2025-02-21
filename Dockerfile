FROM golang:1.24 AS build

RUN useradd -u 1001 golang

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /main .

FROM scratch

WORKDIR /

COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /main /main

USER golang

EXPOSE 8080

CMD ["/main"]