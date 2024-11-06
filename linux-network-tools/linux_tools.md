| Tool        | Command Example                                | Description                                                                                     |
|-------------|------------------------------------------------|-------------------------------------------------------------------------------------------------|
| **tcpdump** | `tcpdump -i eth0 -c 10`                        | Capture 10 packets on `eth0` interface.                                                         |
|             | `tcpdump -i eth0 host 192.168.1.1`            | Capture packets to/from a specific IP address.                                                  |
| **nmap**    | `nmap -sP 192.168.1.0/24`                      | Ping scan for live hosts in a subnet.                                                           |
|             | `nmap -sS -p 1-1000 192.168.1.1`              | Perform a stealth scan on ports 1-1000 on a host.                                               |
| **tracepath** | `tracepath google.com`                      | Trace the route packets take to reach `google.com`.                                             |
| **net-tools** | `ifconfig -a`                                | Display all interfaces and their configurations.                                                |
|             | `route -n`                                    | Display the kernel routing table.                                                               |
| **traceroute** | `traceroute -n google.com`                  | Display the route packets take without resolving hostnames.                                     |
| **curl**    | `curl -I http://example.com`                   | Fetch HTTP headers from a URL.                                                                  |
|             | `curl -O http://example.com/file.zip`         | Download a file from a URL.                                                                     |
| **wget**    | `wget -r -np http://example.com`               | Recursively download files from a website (no parent links).                                    |
| **nmtui**   | `nmtui`                                        | Open the NetworkManager text user interface to manage connections.                              |
| **iproute2** | `ip addr show`                                | Display IP addresses and interfaces.                                                            |
|             | `ip route add default via 192.168.1.1`        | Add a default gateway.                                                                          |
| **bridge-utils** | `brctl addbr br0`                         | Create a new bridge `br0`.                                                                      |
|             | `brctl addif br0 eth0`                        | Add interface `eth0` to bridge `br0`.                                                           |
| **dnsutils** | `dig example.com`                             | Perform a DNS lookup for a domain.                                                              |
|             | `nslookup example.com`                        | Another way to perform DNS lookup.                                                              |
| **iftop**   | `sudo iftop -i eth0`                           | Monitor bandwidth usage on `eth0` interface.                                                    |
| **vnstat**  | `vnstat -i eth0`                               | Display bandwidth usage statistics for `eth0`.                                                  |
|             | `vnstat -d`                                   | Show daily traffic statistics.                                                                  |
| **iptraf-ng** | `sudo iptraf-ng`                             | Launch `iptraf-ng` for interactive real-time traffic monitoring.                                |
| **htop**    | `htop`                                         | Open `htop` for interactive system monitoring.                                                  |
| **frr**     | `vtysh`                                        | Enter FRR shell for managing routing protocols like OSPF and BGP.                               |
|             | `show ip route`                               | (In `vtysh`) Display IP routing table.                                                          |
| **bmon**    | `bmon`                                         | Launch `bmon` for graphical real-time bandwidth monitoring.                                     |
| **ethtool** | `ethtool eth0`                                 | Display settings of the network interface `eth0`.                                               |
|             | `ethtool -s eth0 speed 100 duplex full`       | Set speed to 100Mbps and full duplex for `eth0`.                                                |
| **ipvsadm** | `ipvsadm -Ln`                                  | List configured IPVS services.                                                                  |
| **ncat**    | `ncat -l 1234`                                 | Listen for incoming connections on port `1234`.                                                 |
|             | `ncat -z -v 192.168.1.1 1-1000`               | Scan ports 1-1000 on a remote host.                                                             |
| **fail2ban** | `fail2ban-client status`                      | Check status of Fail2Ban jails.                                                                 |
| **snort**   | `snort -i eth0 -c /etc/snort/snort.conf`       | Start Snort with a configuration file.                                                          |
| **wireshark** | `wireshark &`                                | Open Wireshark GUI.                                                                             |
| **tshark**  | `tshark -i eth0 -c 10`                         | Capture 10 packets on `eth0` in terminal.                                                       |
| **ngrep**   | `ngrep 'GET' tcp and port 80`                  | Capture HTTP GET requests on port 80.                                                           |
| **iperf3**  | `iperf3 -s`                                    | Start in server mode.                                                                           |
|             | `iperf3 -c 192.168.1.1`                       | Run a bandwidth test to server at 192.168.1.1.                                                  |
| **mtr**     | `mtr google.com`                               | Run a continuous traceroute to `google.com`.                                                    |
| **ncdu**    | `ncdu /var/log`                                | Analyze disk usage in `/var/log` directory.                                                     |
