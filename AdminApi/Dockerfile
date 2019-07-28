FROM ubuntu:14.04

RUN mkdir /app

ADD AdminApi /
RUN mkdir /app/log
RUN mkdir /app/config

CMD ["/AdminApi","app/config/config.toml"]