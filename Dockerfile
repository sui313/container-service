FROM ubuntu:18.04

COPY ./containerService /
EXPOSE 8888

ENTRYPOINT ["/containerService"]
