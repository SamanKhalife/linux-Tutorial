# iptable

## 1. Filter Table
Filter is default table for iptables. So, if you don’t define you own table, you’ll be using filter table. Iptables’s filter table has the following built-in chains.

- INPUT chain – Incoming to firewall. For packets coming to the local server.
- OUTPUT chain – Outgoing from firewall. For packets generated locally and going out of the local server.
- FORWARD chain – Packet for another NIC on the local server. For packets routed through the local server.



## 2. NAT table
Iptable’s NAT table has the following built-in chains.

- PREROUTING chain – Alters packets before routing. i.e Packet translation happens immediately after the packet comes to the system (and before routing). This helps to translate the destination ip address of the packets to something that matches the routing on the local server. This is used for DNAT (destination NAT).

- POSTROUTING chain – Alters packets after routing. i.e Packet translation happens when the packets are leaving the system. This helps to translate the source ip address of the packets to something that might match the routing on the desintation server. This is used for SNAT (source NAT).

- OUTPUT chain – NAT for locally generated packets on the firewall.



## 3. Mangle table
Iptables’s Mangle table is for specialized packet alteration. This alters QOS bits in the TCP header. Mangle table has the following built-in chains.

- PREROUTING chain:Rules in this chain apply to packets as they just arrive on the network interface.
- OUTPUT chain:Rules in this chain apply to packets just before they’re given to a local process.
- FORWARD chain:The rules here apply to packets just after they’ve been produced by a process.
- INPUT chain:The rules here apply to any packets that are routed through the current host.
- POSTROUTING chain:The rules in this chain apply to packets as they just leave the network interface.

## 4. Raw table (mostly unused)
Iptable’s Raw table is for configuration excemptions. Raw table has the following built-in chains.

- PREROUTING chain
- OUTPUT chain

## 5. The security table (mostly unused)
The security table is consulted after the filter table to implement mandatory access control network rules. i.e. — it is used by SELinux to set SELinux context markers on packets.


# Chains

## Composition
A chain is a string of rules. When a packet is received, it triggers a netfilter hook, which corresponds to a default chain. iptables finds the appropriate table, then runs it through the chain and traverses a list of rules until it finds a match.
A rule is a statement that tells the system what to do with a packet. Rules can block one type of packet, or forward another type of packet.
The outcome of a rule or where a packet is sent, is called a target.

The default policy of a chain is also a target. The default policy of a chain is what happens to a packet that matches none of the rules. By default, all chains have a default policy of allowing packets.

## Targets

A target is a decision of what to do with a packet. Typically, this is to accept it, drop it, or reject it (which sends an error back to the sender).

- ACCEPT: This causes iptables to accept the packet.
- DROP: iptables drops the packet. To anyone trying to connect to your system, it would appear like the system didn’t even exist.
- REJECT: iptables “rejects” the packet. It sends a “connection reset” packet in case of TCP, or a “destination host unreachable” packet in case of UDP or ICMP.


# main syntax for iptables
```
iptables -t {type of table} -options {chain points} {condition or matching component} {action}
```
for viewing more options visit here 

[Linux.die.net:iptables](https://linux.die.net/man/8/iptables)
