#! /bin/bash

sudo docker exec etcd-1 etcdctl --endpoints=http://etcd-1:2379,http://etcd-2:2379,http://etcd-3:2379 endpoint health 