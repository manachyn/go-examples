You can run the server like this:

apt-get update
apt-get install -y redis-server
echo "bind *" >> /etc/redis/redis.conf
systemctl restart redis-server
sysctl -w fs.file-max=11000000
sysctl -w fs.nr_open=11000000
ulimit -n 11000000
sysctl -w net.ipv4.tcp_mem="100000000 100000000 100000000"
sysctl -w net.core.somaxconn=10000
sysctl -w net.ipv4.tcp_max_syn_backlog=10000
10m-server

You can run the client like this:

sysctl -w fs.file-max=11000000
sysctl -w fs.nr_open=11000000
ulimit -n 11000000
sysctl -w net.ipv4.ip_local_port_range="1025 65535"
sysctl -w net.ipv4.tcp_mem="100000000 100000000 100000000"
10m-client <ip address> <number of connections>