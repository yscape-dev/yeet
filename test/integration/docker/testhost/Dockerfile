FROM debian:bookworm

WORKDIR /work

RUN apt-get update --fix-missing && apt-get -y install openssh-client openssh-server

RUN ln -s /shared/ssh/id_rsa.pub /root/.ssh/authorized_keys

COPY boot.sh .

CMD ["./boot.sh"]
