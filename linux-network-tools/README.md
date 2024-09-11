# Linux Network Commands Guide

## Introduction

This guide provides a comprehensive overview of several important Linux network commands, explaining their functionality and practical use. The focus is on tools like `ip`, `tcpdump`, `dpdk-devbind`, `nc`, `ss`, `nmtui`, `hping3`, and others. Each section contains detailed command examples in a table format for easy reference.

---

## 1. `ip` Command

The `ip` command is used for managing network interfaces, routing tables, and other network-related configurations.

### Common `ip` Commands:

| Command                                      | Description                                                                                   |
|----------------------------------------------|-----------------------------------------------------------------------------------------------|
| `ip addr`                                    | Show IP addresses of all network interfaces                                                    |
| `ip link set dev <interface> up/down`        | Bring a network interface up or down                                                           |
| `ip route`                                   | Display the current routing table                                                              |
| `ip -br a`                                   | Show brief info about all interfaces with IP addresses                                         |
| `ip addr add <IP>/<CIDR> dev <iface>`        | Assign a specific IP address to a network interface                                            |
| `ip route add <network> via <gateway>`       | Add a new route to the routing table                                                           |
| `ip link`                                    | Display the state of all network interfaces                                                    |
| `ip neigh show`                              | Show ARP table                                                                                 |
| `ip link set <dev> mtu <value>`              | Set the MTU for a network interface                                                            |
| `ip ro`                                      | Display routing table                                                                          |
| `ip link show <interface>`                   | Show detailed info about a specific network interface                                          |

---

## 2. `tcpdump` Command

`tcpdump` is a powerful packet capture tool used to analyze network traffic.

### Common `tcpdump` Commands:

| Command                                    | Description                                                                                   |
|--------------------------------------------|-----------------------------------------------------------------------------------------------|
| `tcpdump`                                  | Capture packets on all interfaces                                                             |
| `tcpdump -i <interface>`                   | Capture traffic on a specific network interface                                                |
| `tcpdump -c <number>`                      | Capture a limited number of packets                                                            |
| `tcpdump -n`                               | Display IP addresses instead of resolving hostnames                                            |
| `tcpdump -v`                               | Enable verbose mode to display more packet details                                             |
| `tcpdump -w <file.pcap>`                   | Save captured packets to a file for later analysis                                             |
| `tcpdump -r <file.pcap>`                   | Read and analyze packets from a `.pcap` file                                                   |
| `tcpdump -X`                               | Display the packet’s content in both hex and ASCII formats                                     |
| `tcpdump port <port_number>`               | Capture traffic on a specific port                                                             |
| `tcpdump -A`                               | Display packet contents in ASCII format                                                        |
| `tcpdump 'tcp[13] = 0x12'`                 | Capture SYN-ACK packets only                                                                  |
| `tcpdump icmp`                             | Capture only ICMP packets                                                                     |

---

## 3. `dpdk-devbind` Command

The `dpdk-devbind` utility is used for binding/unbinding network devices to/from DPDK-compatible drivers.

### Common `dpdk-devbind` Commands:

| Command                                      | Description                                                                                   |
|----------------------------------------------|-----------------------------------------------------------------------------------------------|
| `dpdk-devbind.py --status`                   | Display the current status of network devices                                                  |
| `dpdk-devbind.py -b <driver> <device>`       | Bind a network device to a specific driver                                                     |
| `dpdk-devbind.py -u <device>`                | Unbind a network device from its current driver                                                |
| `dpdk-devbind.py -s`                         | Display a short status of network devices                                                      |
| `dpdk-devbind.py --help`                     | Show help and usage information                                                                |

---

## 4. `nc` (Netcat) Command

`nc` or `Netcat` is a versatile networking tool that can create TCP/UDP connections and perform port scanning.

### Common `nc` Commands:

| Command                                    | Description                                                                                   |
|--------------------------------------------|-----------------------------------------------------------------------------------------------|
| `nc <hostname> <port>`                     | Open a TCP connection to a remote host                                                         |
| `nc -l -p <port>`                          | Listen for incoming connections on a specific port                                             |
| `nc -zv <hostname> <port_range>`           | Scan for open ports on a host                                                                  |
| `echo "<data>" | nc <hostname> <port>`     | Send data over a TCP connection                                                                |
| `nc -u <hostname> <port>`                  | Send UDP data                                                                                  |
| `nc -l <port> -e /bin/bash`                | Create a reverse shell on a specific port                                                      |
| `nc -v`                                    | Enable verbose mode for more detailed output                                                   |
| `nc -w <seconds> <host> <port>`            | Set a timeout for connections                                                                  |
| `nc -k -l <port>`                          | Keep listening after a connection is closed                                                    |

---

## 5. `ss` Command

The `ss` command is used for displaying socket statistics and is considered a faster and more detailed replacement for `netstat`.

### Common `ss` Commands:

| Command                                    | Description                                                                                   |
|--------------------------------------------|-----------------------------------------------------------------------------------------------|
| `ss -a`                                    | Display all sockets (both listening and non-listening)                                         |
| `ss -t`                                    | Display only TCP sockets                                                                      |
| `ss -u`                                    | Display only UDP sockets                                                                      |
| `ss -l`                                    | Show only listening sockets                                                                   |
| `ss -p`                                    | Display process info for each socket                                                          |
| `ss -s`                                    | Display socket summary                                                                        |
| `ss -tnp`                                  | Display TCP connections with process information                                               |
| `ss -4`                                    | Show only IPv4 connections                                                                    |
| `ss -6`                                    | Show only IPv6 connections                                                                    |
| `ss state established`                     | Show only established TCP connections                                                         |

---

## 6. `nmtui` Command

`nmtui` is a text user interface for configuring network settings through NetworkManager.

### Common `nmtui` Commands:

| Command                                      | Description                                                                                   |
|----------------------------------------------|-----------------------------------------------------------------------------------------------|
| `nmtui`                                      | Launch the interactive network configuration tool                                              |
| `nmtui-connect <SSID>`                       | Connect to a specific wireless network by SSID                                                |
| `nmtui-edit <connection_name>`               | Edit a specific network connection                                                            |
| `nmtui-hostname`                             | Set the system hostname                                                                       |

---

## 7. `hping3` Command

`hping3` is a command-line packet generator and analyzer used for network testing and security auditing.

### Common `hping3` Commands:

| Command                                      | Description                                                                                   |
|----------------------------------------------|-----------------------------------------------------------------------------------------------|
| `hping3 <host>`                              | Send TCP packets to a specific host                                                           |
| `hping3 -S <host> -p <port>`                 | Send SYN packets to a specific port                                                           |
| `hping3 --udp <host>`                        | Send UDP packets                                                                              |
| `hping3 -1 <host>`                           | Send ICMP echo requests (ping)                                                                |
| `hping3 --flood <host>`                      | Send packets as fast as possible (flooding)                                                   |
| `hping3 -R <host>`                           | Send TCP RST packets                                                                          |
| `hping3 -c <count> <host>`                   | Limit the number of packets sent                                                              |

---

This guide serves as a reference for essential Linux networking commands, covering key tools and usage examples to help network administrators and engineers troubleshoot, manage, and monitor networks efficiently.
