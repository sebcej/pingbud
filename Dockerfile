FROM node:14-alpine as frontend_builder

COPY ./frontend /frontend
WORKDIR /frontend

RUN npm install && npm run build

FROM golang:1.17-alpine as backend_builder

COPY . /app
COPY --from=frontend_builder /frontend /app/frontend

WORKDIR /app

RUN go get && go build -ldflags "-X main.build=`date -u +b%Y%m%d.%H%M%S`"

FROM alpine:latest

WORKDIR /go

COPY --from=backend_builder /app/pingbud /go/pingbud

ENV PINGBUD_DB_PATH=/data/
ENV PINGBUD_SETTINGS_PATH=/data/

EXPOSE 8080

CMD [ "./pingbud" ]