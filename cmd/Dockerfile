FROM golang:1.15.2-alpine3.12 AS base-builder
ENV CGO_ENABLED=0

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY cmd cmd/
COPY pkg pkg/
COPY internal internal/
RUN go test -short ./...
ENV GOOS=linux
ENV GOARCH=amd64

# Base Server
FROM scratch as base-server
COPY --from=base-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base-builder /usr/local/go/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO=/zoneinfo.zip

# product
FROM base-builder as product-builder
RUN go build -o /go/bin/product-api cmd/product/main.go

FROM base-server as product-api
COPY --from=product-builder /go/bin/product-api /product-api
ENTRYPOINT ["/product-api"]

# store
FROM base-builder as store-builder
RUN go build -o /go/bin/store-api cmd/store/main.go

FROM base-server as store-api
COPY --from=store-builder /go/bin/store-api /store-api
ENTRYPOINT ["/store-api"]

# user
FROM base-builder as user-builder
RUN go build -o /go/bin/user-api cmd/user/main.go

FROM base-server as user-api
COPY --from=user-builder /go/bin/user-api /user-api
ENTRYPOINT ["/user-api"]

# admin
FROM base-builder as admin-builder
RUN go build -o /go/bin/admin-api cmd/admin/main.go

FROM base-server as admin-api
COPY --from=admin-builder /go/bin/admin-api /admin-api
ENTRYPOINT ["/admin-api"]