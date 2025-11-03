#!/bin/bash

#进入命名空间执行命名,这里是进入到了一个docker创建的容器的mnt命名空间查看python版本
#root@myHost:/proc/1# nsenter -t 78955 -m /bin/bash
#root@myHost:/# python --version
#Python 3.12.12
#root@myHost:/# exit
#exit
#root@myHost:/proc/1# python --version
#Command 'python' not found, did you mean:
#  command 'python3' from deb python3
#  command 'python' from deb python-is-python3
#root@myHost:/proc/1# lsns -t mnt | grep 78955
#4026532230 mnt       1 78955 root            /usr/local/bin/python3.12 /usr/local/bin/uvicorn main:app --host 0.0.0.0 --port 8000
