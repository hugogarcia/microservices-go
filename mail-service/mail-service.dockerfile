FROM alpine:3.19.0

WORKDIR /app

COPY mailApp .
COPY /template /app/template

EXPOSE 8383

CMD ["/app/mailApp"]