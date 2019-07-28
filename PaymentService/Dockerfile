FROM ubuntu:14.04

RUN mkdir /app

ADD PaymentService /
RUN mkdir /app/log
RUN mkdir /app/config
RUN apt-get update
RUN apt-get install -y ca-certificates

CMD ["/PaymentService","app/config/config.toml"]