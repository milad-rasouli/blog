| **Plugin**                  | **Description**                                                                                                                                   |
|-----------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------|
| **abf_plugin.so**            | ACL-based forwarding (ABF) plugin for policy-based routing using ACLs.                                                                            |
| **acl_plugin.so**            | Implements access control lists (ACLs) to filter or permit traffic.                                                                               |
| **adl_plugin.so**            | Adaptive Load Distribution plugin for optimizing traffic across multiple paths.                                                                  |
| **af_packet_plugin.so**      | Interface plugin for Linux AF_PACKET to send and receive packets.                                                                                 |
| **af_xdp_plugin.so**         | Interface plugin using Linux AF_XDP (eXpress Data Path) for high-performance packet processing.                                                   |
| **arping_plugin.so**         | Provides support for generating and handling ARP requests/replies.                                                                                |
| **avf_plugin.so**            | AVF (Advanced Virtual Function) plugin for Intel SR-IOV virtual functions.                                                                        |
| **bpf_trace_filter_plugin.so**| Provides filtering and tracing support using BPF (Berkeley Packet Filter).                                                                        |
| **bufmon_plugin.so**         | Buffer monitor plugin for analyzing and debugging packet buffer usage in VPP.                                                                     |
| **cdp_plugin.so**            | Cisco Discovery Protocol (CDP) plugin to support neighbor discovery in Cisco networks.                                                            |
| **cnat_plugin.so**           | Container Network Address Translation (CNAT) plugin for translating container network addresses.                                                  |
| **crypto_ipsecmb_plugin.so** | Intel IPSec Multi-Buffer plugin for accelerating IPsec encryption/decryption.                                                                     |
| **crypto_native_plugin.so**  | Implements native cryptographic operations.                                                                                                       |
| **crypto_openssl_plugin.so** | Plugin to accelerate cryptographic operations using OpenSSL libraries.                                                                            |
| **crypto_sw_scheduler_plugin.so**| Software-based cryptographic operation scheduling.                                                                                             |
| **ct6_plugin.so**            | Connection Tracking for IPv6.                                                                                                                     |
| **det44_plugin.so**          | Deterministic NAT44 plugin for deterministic address translation in IPv4 networks.                                                                |
| **dev_ena_plugin.so**        | Plugin to support ENA (Elastic Network Adapter) used in AWS for enhanced networking.                                                              |
| **dev_iavf_plugin.so**       | Plugin for Intel Adaptive Virtual Function (IAVF) devices.                                                                                        |
| **dhcp_plugin.so**           | DHCP (Dynamic Host Configuration Protocol) plugin for dynamic IP address assignment.                                                              |
| **dispatch_trace_plugin.so** | Traces the dispatching of packets for debugging and analysis.                                                                                     |
| **dma_intel_plugin.so**      | Direct Memory Access (DMA) plugin for Intel devices.                                                                                              |
| **dns_plugin.so**            | DNS (Domain Name System) plugin for handling DNS queries.                                                                                         |
| **dslite_plugin.so**         | Dual-Stack Lite plugin for IPv4 over IPv6 tunneling.                                                                                              |
| **fateshare_plugin.so**      | Provides fate-sharing capability to link failure or success of multiple connections.                                                              |
| **first_plugin.so**          | A test or experimental plugin.                                                                                                                    |
| **flowprobe_plugin.so**      | Flow monitoring and analysis plugin for collecting flow data.                                                                                     |
| **geneve_plugin.so**         | Implements GENEVE (Generic Network Virtualization Encapsulation) tunneling protocol.                                                              |
| **gre_plugin.so**            | Generic Routing Encapsulation (GRE) plugin for creating GRE tunnels.                                                                              |
| **gtpu_plugin.so**           | GTP-U (GPRS Tunneling Protocol User Plane) plugin for mobile networks.                                                                            |
| **hs_apps_plugin.so**        | Host Stack applications plugin.                                                                                                                   |
| **hsi_plugin.so**            | High-Speed Internet plugin for network performance analysis.                                                                                      |
| **http_plugin.so**           | Basic HTTP (Hypertext Transfer Protocol) plugin.                                                                                                  |
| **http_static_plugin.so**    | Static HTTP server plugin for serving static content.                                                                                             |
| **idpf_plugin.so**           | Intel Dynamic Device Personalization Framework plugin.                                                                                            |
| **igmp_plugin.so**           | Internet Group Management Protocol (IGMP) plugin for multicast group management.                                                                  |
| **ikev2_plugin.so**          | Internet Key Exchange version 2 (IKEv2) plugin for IPsec key management.                                                                          |
| **ila_plugin.so**            | Identifier-Locator Addressing plugin for separating host identity from network location.                                                          |
| **ioam_plugin.so**           | In-band OAM (Operations, Administration, and Maintenance) plugin for performance monitoring.                                                      |
| **ip_session_redirect_plugin.so**| Redirects IP sessions based on predefined rules.                                                                                               |
| **l2tp_plugin.so**           | Layer 2 Tunneling Protocol (L2TP) plugin for tunneling L2 traffic.                                                                                |
| **lacp_plugin.so**           | Link Aggregation Control Protocol (LACP) plugin for bundling multiple interfaces.                                                                 |
| **l3xc_plugin.so**           | Layer 3 cross-connect plugin for forwarding between Layer 3 interfaces.                                                                           |
| **lisp_plugin.so**           | Locator/ID Separation Protocol (LISP) plugin for network virtualization.                                                                          |
| **lldp_plugin.so**           | Link Layer Discovery Protocol (LLDP) plugin for neighbor discovery.                                                                                |
| **linux_cp_plugin.so**       | Linux Control Plane plugin for managing Linux-side interfaces in VPP.                                                                              |
| **map_plugin.so**            | Mapping of Address and Port (MAP) plugin for IPv4-IPv6 coexistence.                                                                                |
| **memif_plugin.so**          | Memory Interface (memif) plugin for high-speed inter-process communication (IPC).                                                                  |
| **mss_clamp_plugin.so**      | Plugin for clamping the maximum segment size (MSS) in TCP connections.                                                                            |
| **nat44_ei_plugin.so**       | NAT44 plugin with endpoint-independent mapping for address translation in IPv4.                                                                    |
| **nat64_plugin.so**          | Network Address Translation 64 (NAT64) plugin for translating IPv6 to IPv4.                                                                       |
| **npt66_plugin.so**          | NAT66 plugin for translating IPv6 addresses.                                                                                                      |
| **nsh_plugin.so**            | Network Service Header (NSH) plugin for service chaining.                                                                                         |
| **osi_plugin.so**            | OSI layer support plugin.                                                                                                                         |
| **ping_plugin.so**           | Plugin for sending ICMP echo (ping) requests.                                                                                                     |
| **pppoe_plugin.so**          | Point-to-Point Protocol over Ethernet (PPPoE) plugin for broadband connections.                                                                   |
| **prom_plugin.so**           | Prometheus metrics plugin for collecting network statistics.                                                                                      |
| **quic_plugin.so**           | QUIC protocol plugin for secure and fast data transfer.                                                                                           |
| **rdma_plugin.so**           | Remote Direct Memory Access (RDMA) plugin for high-speed network communication.                                                                   |
| **srv6ad_plugin.so**         | Segment Routing over IPv6 (SRv6) Advanced plugin for network programming.                                                                         |
| **srv6mobile_plugin.so**     | SRv6 plugin optimized for mobile network deployments.                                                                                             |
| **stn_plugin.so**            | Source NAT (SNAT) plugin for address translation in specific scenarios.                                                                           |
| **tlsmbedtls_plugin.so**     | TLS encryption plugin using the MbedTLS library.                                                                                                  |
| **tlsopenssl_plugin.so**     | TLS encryption plugin using the OpenSSL library.                                                                                                  |
| **tracedump_plugin.so**      | Trace dump plugin for exporting packet traces.                                                                                                    |
| **tracenode_plugin.so**      | Traces individual nodes for performance analysis and debugging.                                                                                   |
| **unittest_plugin.so**       | Plugin for running unit tests in VPP.                                                                                                             |
| **urpf_plugin.so**           | Unicast Reverse Path Forwarding plugin for anti-spoofing.                                                                                         |
| **vhost_plugin.so**          | Plugin for handling vhost-user interfaces (typically used in virtualized environments).                                                            |
| **vrrp_plugin.so**           | Virtual Router Redundancy Protocol (VRRP) plugin for high availability.                                                                           |
| **vxlan_plugin.so**          | VXLAN (Virtual Extensible LAN) plugin for overlay networks.                                                                                       |
| **wireguard_plugin.so**      | Plugin for the WireGuard VPN protocol.                                                                                                            |
