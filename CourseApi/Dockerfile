FROM ubuntu:14.04

RUN mkdir /app

ADD CourseApi /
RUN mkdir /app/log
RUN mkdir /app/config

CMD ["/CourseApi","app/config/config.toml"]