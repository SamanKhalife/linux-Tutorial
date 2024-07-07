# arp

The `arp` command in Linux is used to manipulate the system's ARP (Address Resolution Protocol) cache. ARP is used to map IP addresses to MAC addresses, which are necessary for network communication within a local network segment. Each time a host needs to communicate with another host on the same network, it uses ARP to find out the MAC address associated with the IP address it wants to communicate with.


1. **Display the ARP Cache**:
   ```sh
   arp -a
   ```
   This command lists all the entries in the ARP cache.

2. **Add a Static ARP Entry**:
   ```sh
   sudo arp -s 192.168.1.2 00:11:22:33:44:55
   ```
   This command adds a static ARP entry for the IP address `192.168.1.2` with the MAC address `00:11:22:33:44:55`.

3. **Delete an ARP Entry**:
   ```sh
   sudo arp -d 192.168.1.2
   ```
   This command deletes the ARP entry for the IP address `192.168.1.2`.

### Example of `arp` Command Output

When you run `arp -a`, you might see output similar to this:

```sh
? (192.168.1.1) at 00:11:22:33:44:55 [ether] on eth0
? (192.168.1.2) at 00:11:22:33:44:56 [ether] on eth0
```

### Using `ip` Command for ARP

The `ip` command from the `iproute2` package also allows you to manage ARP entries, and it is the modern replacement for the `arp` command. Here are equivalent commands using `ip`:

1. **Display the ARP Cache**:
   ```sh
   ip neigh
   ```

2. **Add a Static ARP Entry**:
   ```sh
   sudo ip neigh add 192.168.1.2 lladdr 00:11:22:33:44:55 dev eth0
   ```

3. **Delete an ARP Entry**:
   ```sh
   sudo ip neigh del 192.168.1.2 dev eth0
   ```

### Conclusion

While the `arp` command is useful for viewing and managing ARP cache entries, it is being replaced by the `ip` command for most modern networking tasks. The `ip` command provides a more comprehensive and consistent interface for managing networking configuration, including ARP entries. It is recommended to use `ip` for advanced and up-to-date network management.

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
