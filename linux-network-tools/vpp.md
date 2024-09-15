``` bash
lcp  create GE2 host-if e0

lcp  create GE8 host-if e1



set interface state GE2 up

set interface state GE8 up

set interface mtu packet 1500 GE8
```


## create tap interface and ping it

``` bash 
set interface state GE8 up
set interface mtu packet 9000 GE8
set interface ip address GE8 10.10.2.15/24
lcp create GE8 host-if e0
ip link set e0 up mtu 9000
ip addr add 10.10.2.15/24 dev e0
# set interface ip address GE8 2001:db8:0:1::1/64
ip addr add 2001:db8:0:1::1/64 dev e0

```
