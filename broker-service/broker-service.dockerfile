FROM alpine:3.19.0

RUN mkdir /app

COPY brokerApp /app

EXPOSE 8080

CMD ["/app/brokerApp"]