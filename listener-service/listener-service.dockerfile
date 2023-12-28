FROM alpine:3.19.0

RUN mkdir /app

COPY listenerApp /app

CMD ["/app/listenerApp"]