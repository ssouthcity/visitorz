ARG GO_VERSION=1.21.6

FROM golang:${GO_VERSION} AS build

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/visitorz .

FROM scratch

COPY --from=build /bin/visitorz /bin/

EXPOSE 5000

ENTRYPOINT [ "/bin/visitorz" ]
