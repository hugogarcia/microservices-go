FROM alpine:3.19.0

RUN mkdir /app

COPY loggerApp /app

EXPOSE 8282

CMD ["/app/loggerApp"]