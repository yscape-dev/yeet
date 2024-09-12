#!/bin/bash

# Enable tracing
set -x


while [ ! -f /root/.ssh/id_rsa.pub ]; do echo "Waiting for ssh keys"; sleep 1; done

if [ -f /root/.ssh/known_hosts ]; then
    rm /root/.ssh/known_hosts
fi

while ! ssh-keyscan testhost > /root/.ssh/known_hosts 2> /dev/null; do
    echo "Waiting for testhost"
    sleep 1
done

exec sleep infinity

