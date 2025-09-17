FROM golang:alpine AS builder

WORKDIR /build

COPY . /build/

RUN apk add git && \
    go get -d

ARG TAG "none"
ARG BUILDDATE 

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-X main.BuildVersion=$BUILDDATE -X main.GitVersion=$TAG -extldflags \"-static\"" -o main .


FROM alpine
LABEL maintainer="Andreas Peters <support@aventer.biz>"
LABEL org.opencontainers.image.title="go-cron" 
LABEL org.opencontainers.image.description="Easy an simple cron service as container."
LABEL org.opencontainers.image.vendor="AVENTER UG (haftungsbeschränkt)"
LABEL org.opencontainers.image.source="https://github.com/AVENTER-UG/"

ENV DOCKER_RUNNING=true

RUN apk add --no-cache ca-certificates
RUN adduser -S -D -H -h /app appuser
USER appuser

COPY --from=builder /build/main /app/

EXPOSE 10000

WORKDIR "/app"

CMD ["./main"]
