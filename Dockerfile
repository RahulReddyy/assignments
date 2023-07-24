FROM golang:1.20 as build

WORKDIR /app

COPY . .

RUN go build -o ginapi

# FROM scratch

# COPY --from=build /app/ginapi .
# EXPOSE 8080

ENTRYPOINT ["./ginapi"]
