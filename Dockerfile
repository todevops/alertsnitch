FROM debian:stretch

RUN apt update && apt install -y tzdata
COPY alertsnitch /alertsnitch

ENV TZ Asia/Shanghai

ENTRYPOINT [ "/alertsnitch" ]