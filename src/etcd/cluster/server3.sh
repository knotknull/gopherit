ETCDCTL_API=3 /tmp/etcd-download-test/etcd --name infra3 --initial-advertise-peer-urls http://127.0.0.1:32380 \
  --listen-peer-urls http://127.0.0.1:32380 \
  --listen-client-urls http://127.0.0.1:32379,http://127.0.0.1:32379 \
  --advertise-client-urls http://127.0.0.1:32379 \
  --initial-cluster-token etcd-cluster-1 \
  --initial-cluster infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380 \
  --initial-cluster-state new
