# addresses

here is the output of the addresses command:

```
Interface   Address         Type
---------------------------------------------
lo          127.0.0.1     Loopback
enp0s3    192.168.1.100   Ethernet
enp0s3    fe80::a00:27ff:fe88:911f   IPv6
```

The output shows the following information for each network interface:

- The interface name
- The IP address
- The address type

The addresses command is a useful tool for troubleshooting network problems. It can be used to see which interfaces have IP addresses, and what type of addresses they are.

Here are some examples of how to use the addresses command:

```
# To display all addresses:
addresses

# To display addresses for the enp0s3 interface:
addresses -i enp0s3

# To display only IPv4 addresses:
addresses -a | grep -E '^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$'
```


# help 

```
addresses [options]

Display information about network addresses.

Options:

-a, --all               Display all addresses.
-i, --interface=INTERFACE   Display addresses for the specified interface.
-h, --help               Show this help message.
-V, --version            Print version information.

For more information on network addresses, see the following resources:

* The `ifconfig` man page: https://linux.die.net/man/8/ifconfig
* The `ip` man page: https://linux.die.net/man/8/ip
```

## breakdown

```
-a, --all: This option displays all addresses, including IPv4, IPv6, and MAC addresses.
-i, --interface=INTERFACE: This option displays addresses for the specified interface.
-h, --help: This option shows this help message.
-V, --version: This option prints version information.
```
