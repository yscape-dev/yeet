FROM debian:bookworm

WORKDIR /work

RUN apt-get update --fix-missing && apt-get -y install openssh-client openssh-server

COPY authorized_keys /root/.ssh/authorized_keys
COPY boot.sh .

EXPOSE 22

CMD ["./boot.sh"]
