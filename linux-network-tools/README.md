
# Linux Network Commands Guide

## Introduction

This guide explores a variety of essential Linux network commands, offering in-depth insights and practical examples. Each tool is important for network troubleshooting, monitoring, and general management. The guide covers commands like `ip`, `tcpdump`, `dpdk-devbind`, `nc`, and `ss`.

## 1. `ip` Command

The `ip` command is used to manage network devices, routing tables, and network interfaces. It replaces the older `ifconfig` and `route` commands.

### Common Commands:

| Command                                | Description                                                                                   |
|----------------------------------------|-----------------------------------------------------------------------------------------------|
| `ip addr`                              | Show IP addresses assigned to all network interfaces                                           |
| `ip link set dev <interface> up/down`  | Enable or disable a specific interface                                                        |
| `ip route`                             | Display the current routing table                                                             |
| `ip -br a`                             | Show brief info about all interfaces with IP addresses                                        |
| `ip addr add <IP>/<CIDR> dev <iface>`  | Assign a specific IP address to a network interface                                           |
| `ip route add <network> via <gateway>` | Add a new route to the routing table                                                          |

## 2. `tcpdump` Command

`tcpdump` is a packet analyzer tool that captures network traffic on specific interfaces, useful for network diagnostics and troubleshooting.

### Common Commands:

| Command                                        | Description                                                                                           |
|------------------------------------------------|-------------------------------------------------------------------------------------------------------|
| `tcpdump`                                      | Capture packets on all interfaces (needs root privileges)                                              |
| `tcpdump -i <interface>`                       | Capture traffic on a specific interface                                                                |
| `tcpdump -c <number>`                          | Limit packet capture to a specific number                                                              |
| `tcpdump -n`                                   | Don't resolve hostnames, show IPs instead                                                              |
| `tcpdump -v`                                   | Verbose output for more detailed packet info                                                           |
| `tcpdump -w <file.pcap>`                       | Write captured packets to a .pcap file                                                                 |
| `tcpdump -r <file.pcap>`                       | Read and analyze packets from a .pcap file                                                             |

## 3. `dpdk-devbind` Command

The `dpdk-devbind` command manages network interfaces for DPDK applications, binding or unbinding them to/from drivers.

### Common Commands:

| Command                                        | Description                                                                                           |
|------------------------------------------------|-------------------------------------------------------------------------------------------------------|
| `dpdk-devbind.py --status`                     | Show the current status of all network devices and their drivers                                       |
| `dpdk-devbind.py -b <driver> <device>`         | Bind a device to a specific DPDK-compatible driver                                                     |
| `dpdk-devbind.py -u <device>`                  | Unbind a device from its current driver                                                                |

## 4. `nc` (Netcat) Command

The `nc` command (Netcat) is used for network communication, acting as a TCP/UDP client or server.

### Common Commands:

| Command                                        | Description                                                                                           |
|------------------------------------------------|-------------------------------------------------------------------------------------------------------|
| `nc <hostname> <port>`                         | Establish a TCP connection to a remote host                                                            |
| `nc -l -p <port>`                              | Listen for incoming connections on a specific port (server mode)                                       |
| `nc -zv <hostname> <port_range>`               | Port scanning to check for open ports                                                                  |
| `echo "<data>" | nc <hostname> <port>`         | Send data over a TCP connection                                                                        |
| `nc -u <hostname> <port>`                      | Send UDP data                                                                                          |

## 5. `ss` Command

The `ss` command is used for displaying socket statistics and inspecting network connections. It is a faster and more detailed replacement for `netstat`.

### Common Commands:

| Command                                        | Description                                                                                           |
|------------------------------------------------|-------------------------------------------------------------------------------------------------------|
| `ss -a`                                        | Show all sockets (both listening and non-listening)                                                    |
| `ss -t`                                        | Show only TCP sockets                                                                                  |
| `ss -u`                                        | Show only UDP sockets                                                                                  |
| `ss -l`                                        | Show only listening sockets                                                                            |
| `ss -p`                                        | Show process info for each socket                                                                      |
| `ss -t state established`                      | Show only established TCP connections                                                                  |
| `ss -s`                                        | Summary of socket statistics (number of established connections, etc.)                                 |

---

This guide provides a comprehensive view of the most important networking tools in Linux. Each section includes the essential commands that every network or system administrator should know for effective network management and troubleshooting.
