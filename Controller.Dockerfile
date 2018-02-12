FROM alpine

LABEL vendor="mtenrero.com"
LABEL maintainer="Marcos Tenrero"

COPY ./releases/atq-linux-amd64 /opt/atq/atq

ENTRYPOINT [ "/bin/ash", "-c", "/opt/atq/atq", "-mode controller"]