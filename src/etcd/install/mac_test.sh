# write,read to etcd
ETCDCTL_API=3 /tmp/etcd-download-test/etcdctl
ETCDCTL_API=3 /tmp/etcd-download-test/etcdctl --endpoints=localhost:2379 put foo bar
ETCDCTL_API=3 /tmp/etcd-download-test/etcdctl --endpoints=localhost:2379 get foo
