FROM alpine:latest

WORKDIR /app

COPY app /app/
COPY .env /app/

RUN chmod +x /app/app

CMD ["./app"]
