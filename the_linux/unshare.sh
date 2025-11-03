#!/bin/bash

# 创建新的命名空间,例如创建一个新的网络命名空间运行bash
#root@myHost:/proc/1# unshare -n /bin/bash
#root@myHost:/proc/1# ip a
#1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
#    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
