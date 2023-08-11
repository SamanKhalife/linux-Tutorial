The arp command is used to display the Address Resolution Protocol (ARP) table. The ARP table is a list of IP addresses and their corresponding MAC addresses. MAC addresses are unique identifiers assigned to network interfaces. IP addresses are used to identify hosts on a network, while MAC addresses are used to identify specific devices on a network.

For example, to display all entries in the ARP table, you would use the following command:

`arp -a`

To display numeric addresses instead of names, you would use the following command:

`arp -n`

To set a static ARP entry for the IP address 192.168.1.100 to the MAC address 00:11:22:33:44:55, you would use the following command:

`arp -s 192.168.1.100 00:11:22:33:44:55`

To display the ARP table for the interface eth0, you would use the following command:

`arp -i eth0`



# help 

```
arp [options] [interface] [IP address]

Display ARP table, or set an entry.

Options:

-a, --all            Display all entries in the ARP table.
-n, --numeric        Display numeric addresses instead of names.
-s, --static         Set a static ARP entry.
-i, --interface      Interface to use.

```



## breakdown

```
-a, --all: This option displays all entries in the ARP table.
-n, --numeric: This option displays numeric addresses instead of names.
-s, --static: This option sets a static ARP entry.
-i, --interface: This option specifies the interface to use.
```
