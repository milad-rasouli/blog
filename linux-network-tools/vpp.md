# Create a tap interface
``` bash
lcp  create GE2 host-if e0
lcp  create GE8 host-if e1
set interface state GE2 up
set interface state GE8 up
set interface mtu packet 1500 GE8
```


# Create tap interface and ping it

1. Modify ```vi /etc/vpp/startup.conf```.
``` bash
plugins {
    plugin linux_cp_plugin.so { enable }
    plugin ping_plugin.so { disable }
}

```
2. Perform the following commands on machine 1:
``` bash 
set interface state GE8 up
set interface mtu packet 9000 GE8
set interface ip address GE8 10.10.2.15/24
lcp create GE8 host-if e0
ip link set e0 up mtu 9000
ip addr add 10.10.2.15/24 dev e0
# set interface ip address GE8 2001:db8:0:1::1/64
#ip addr add 2001:db8:0:1::1/64 dev e0

tcpdump -i e0
```
3. Perform the following command on machine 2:

``` bash
ping 10.10.2.15
```

# Ping with lcp-sync

vpp:
``` bash
lcp create GE8 host-if e0
set interface state GE8 up
```

linux:
``` bash
ip a add 10.10.2.20/24 dev e0
ip l set mtu 1500 dev e0
ip l set state up dev e0

tcpdump -i e0
```

# VLAN

``` bash
create sub GE8 1234
set interface mtu packet 9000 GE8.1234
lcp create GE8.1234 host-if e0.1234
set interface state GE8.1234 up
set interface ip address GE8.1234 10.10.2.16/24
# set interface ip address GE8.1234 2001:db8:0:2::1/64
```

# Change interface mtu

``` bash
ip link set enp1s0 down
ip link add link enp1s0 name foo type vlan id 1234
ip link set foo down
## Both interfaces are down, which makes sense because I set them both down
ip link | grep enp1s0


ip link set enp1s0 up
ip link | grep enp1s0
## Both interfaces are up, which doesn't make sense because I only changed one of them!
ip link set foo  mtu 5000
ip a | grep enp1s
2: enp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 5000 qdisc fq_codel state UP group default qlen 1000
    inet 172.25.110.245/24 brd 172.25.110.255 scope global noprefixroute enp1s0
39: foo@enp1s0: <BROADCAST,MULTICAST> mtu 5000 qdisc noqueue state DOWN group default qlen 1000
# ip link set foo  mtu 5000


```
