#!/bin/bash
apt-get update
apt-get install -y gcc git
cd /tmp
wget https://go.dev/dl/go1.18.2.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.18.2.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
source /opt/yottadb/current/ydb_env_set
cd /home/gocode
go get lang.yottadb.com/go/yottadb
source $(pkg-config --variable=prefix yottadb)/ydb_env_set
go get -t lang.yottadb.com/go/yottadb
go build
go test lang.yottadb.com/go/yottadb
/opt/yottadb/current/ydb <<< 'H'
tail -f /dev/null
