FROM golang:1.16 AS build

WORKDIR /build
COPY . .

RUN cd hello && go build -o bin/hello cmd/main.go

# FROM golang:1.16-alpine
FROM gcr.io/distroless/base-debian10

COPY --from=build /build/hello/bin/hello /

ENTRYPOINT ["/hello"]
