# 

The `ping` command in Linux is used to test the connectivity between your computer and another computer on the network. It sends a series of packets to the other computer and waits for a response. If the other computer responds, the ping command will display the round-trip time (RTT) for each packet.

The syntax of the `ping` command is as follows:

```
ping [options] host
```

The `host` argument specifies the IP address or hostname of the computer that you want to ping.

The `options` argument controls the behavior of the `ping` command. The most common options are as follows:

* `-c`: Specifies the number of packets to send.
* `-i`: Specifies the interval between packets.
* `-t`: Specifies the time to stop sending packets.

For example, the following command will ping the IP address 192.168.1.100 10 times, with an interval of 1 second between packets:

```
ping -c 10 -i 1 192.168.1.100
```

This command will send 10 packets to the IP address 192.168.1.100, with an interval of 1 second between packets. The output of the command will show the RTT for each packet.

The `ping` command is a useful command for troubleshooting network problems. It can be used to see if a computer is reachable on the network, and to see how long it takes for packets to travel between computers.

Here are some additional things to keep in mind about the `ping` command:

* The `ping` command will only work if the other computer is running a network stack that supports ICMP.
* The `ping` command will not work if the other computer is firewalled or if there is a problem with the network infrastructure.
* The `ping` command can be used to stress test a network by sending a large number of packets.

It is important to be aware of these limitations when using the `ping` command, so that you do not get confused by the output or cause problems on the network.


# help 

```
ping: invalid option -- '-'

Usage
  ping [options] <destination>

Options:
  <destination>      dns name or ip address
  -a                 use audible ping
  -A                 use adaptive ping
  -B                 sticky source address
  -c <count>         stop after <count> replies
  -D                 print timestamps
  -d                 use SO_DEBUG socket option
  -f                 flood ping
  -h                 print help and exit
  -I <interface>     either interface name or address
  -i <interval>      seconds between sending each packet
  -L                 suppress loopback of multicast packets
  -l <preload>       send <preload> number of packages while waiting replies
  -m <mark>          tag the packets going out
  -M <pmtud opt>     define mtu discovery, can be one of <do|dont|want>
  -n                 no dns name resolution
  -O                 report outstanding replies
  -p <pattern>       contents of padding byte
  -q                 quiet output
  -Q <tclass>        use quality of service <tclass> bits
  -s <size>          use <size> as number of data bytes to be sent
  -S <size>          use <size> as SO_SNDBUF socket option value
  -t <ttl>           define time to live
  -U                 print user-to-user latency
  -v                 verbose output
  -V                 print version and exit
  -w <deadline>      reply wait <deadline> in seconds
  -W <timeout>       time to wait for response

IPv4 options:
  -4                 use IPv4
  -b                 allow pinging broadcast
  -R                 record route
  -T <timestamp>     define timestamp, can be one of <tsonly|tsandaddr|tsprespec>

IPv6 options:
  -6                 use IPv6
  -F <flowlabel>     define flow label, default is random
  -N <nodeinfo opt>  use icmp6 node info query, try <help> as argument

For more details see ping(8).
```
