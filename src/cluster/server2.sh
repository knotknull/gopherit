consul agent -data-dir=/tmp/consul -node=agent-two \
    -bind=172.20.20.11 -enable-script-checks=true -config-dir=/etc/consul.d
