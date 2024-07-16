# Rogue access points, routers, and DHCP servers
Rogue access points, routers, and DHCP servers pose significant security risks to networks by introducing unauthorized devices that can intercept, manipulate, or disrupt network traffic. Here’s an overview of what these rogue devices are, the threats they pose, and how to mitigate these risks:

### Rogue Access Points

**Definition**: Rogue access points are unauthorized wireless access points that are connected to a network without the knowledge or approval of network administrators. They can be intentionally installed by malicious actors or inadvertently introduced by employees or visitors.

**Threats**:

1. **Man-in-the-Middle Attacks**: Rogue access points can intercept and monitor wireless communications between devices and legitimate access points, potentially capturing sensitive information.

2. **Network Intrusion**: They can provide unauthorized access to network resources, allowing attackers to penetrate network defenses and compromise sensitive data.

3. **Denial of Service (DoS)**: Rogue access points can disrupt network operations by creating conflicts or interference with legitimate access points.

**Mitigation**:

- **Wireless Site Surveys**: Regularly conduct wireless site surveys to detect unauthorized access points and ensure coverage areas are secure.

- **Wireless Intrusion Detection Systems (WIDS)**: Implement WIDS solutions to detect and mitigate rogue access points automatically.

- **802.1X Authentication**: Deploy 802.1X authentication to restrict access to authorized devices only and prevent unauthorized devices from connecting to the network.

- **Network Segmentation**: Segment wireless networks using VLANs to isolate critical assets and limit the impact of potential compromises.

### Rogue Routers

**Definition**: Rogue routers are unauthorized routers connected to a network infrastructure, usually to provide additional network connectivity without proper authorization.

**Threats**:

1. **Network Partitioning**: Rogue routers can create separate network segments that bypass network security controls, allowing unauthorized access to sensitive data.

2. **Traffic Interception**: They can intercept and analyze network traffic passing through them, facilitating data exfiltration or unauthorized monitoring.

3. **Configuration Manipulation**: Rogue routers can alter network settings or redirect traffic to malicious destinations, compromising network integrity.

**Mitigation**:

- **Port Security**: Implement port security measures, such as MAC address filtering, to prevent unauthorized devices from connecting to network switches.

- **Network Monitoring**: Use network monitoring tools to detect unauthorized devices or unusual traffic patterns that may indicate the presence of rogue routers.

- **Network Access Control (NAC)**: Deploy NAC solutions to enforce security policies and restrict unauthorized devices from accessing the network.

- **Regular Audits**: Conduct regular audits and network scans to identify and remove unauthorized devices, including rogue routers, from the network.

### Rogue DHCP Servers

**Definition**: Rogue DHCP servers are unauthorized servers that distribute IP addresses and network configuration parameters to devices on a network, causing conflicts and potential security vulnerabilities.

**Threats**:

1. **IP Address Exhaustion**: Rogue DHCP servers can deplete the available pool of IP addresses, causing network connectivity issues for legitimate devices.

2. **Traffic Redirection**: They can redirect network traffic to malicious destinations or intercept sensitive information transmitted over the network.

3. **Man-in-the-Middle Attacks**: Rogue DHCP servers can facilitate man-in-the-middle attacks by intercepting and modifying DHCP responses sent to client devices.

**Mitigation**:

- **DHCP Snooping**: Enable DHCP snooping on network switches to monitor and filter DHCP messages, blocking unauthorized DHCP server responses.

- **Static DHCP Binding**: Use static DHCP bindings to assign IP addresses to specific MAC addresses, preventing unauthorized devices from obtaining IP addresses dynamically.

- **Segmentation and Filtering**: Segment networks and use access control lists (ACLs) to restrict DHCP traffic to authorized servers only.

- **Network Monitoring**: Continuously monitor network traffic and DHCP server logs for signs of unauthorized DHCP server activity.

### Conclusion

Rogue access points, routers, and DHCP servers represent serious security risks by introducing unauthorized devices into network environments. Implementing proactive detection and mitigation strategies, such as regular network audits, deploying intrusion detection systems, enforcing access controls, and segmenting network traffic, is crucial to safeguarding against these threats. By maintaining vigilance and implementing robust security measures, organizations can mitigate the risks associated with rogue devices and maintain the integrity and security of their networks.
