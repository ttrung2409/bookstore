# Build
FROM golang:1.18.1-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /dist

# Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /dist /dist

CMD ["/dist"]