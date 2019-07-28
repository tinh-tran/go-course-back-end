FROM ubuntu:14.04

RUN mkdir /app

ADD employee /
RUN mkdir /app/log
RUN mkdir /app/config

CMD ["/employee","app/config/config.toml"]