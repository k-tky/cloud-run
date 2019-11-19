FROM golang:1.13 as base

WORKDIR /app

COPY cloud_run.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o cloud_run

FROM alpine

COPY --from=base /app/cloud_run /cloud_run

CMD [ "/cloud_run" ]
