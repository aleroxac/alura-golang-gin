FROM golang:1.20.4-alpine3.17 AS base


FROM base as build
WORKDIR /build
COPY . /build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o skill


FROM scratch as main
WORKDIR /app
COPY --from=build /build/skill /app/skill
ENTRYPOINT [ "/app/skill" ]