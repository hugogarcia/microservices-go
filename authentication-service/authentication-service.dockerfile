FROM alpine:3.19.0

RUN mkdir /app

COPY authApp /app

EXPOSE 8181

CMD ["/app/authApp"]