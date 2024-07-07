# traceroute

The `traceroute` command is a network diagnostic tool used to track the path that a packet takes from the source machine to the destination host. It also provides the round-trip time (RTT) for each hop along the route. This is particularly useful for identifying points of failure or latency within a network.

### How `traceroute` Works

`traceroute` works by sending packets with incrementing Time-To-Live (TTL) values. Each router along the path to the destination host decrements the TTL value by 1 before forwarding the packet. When the TTL value reaches zero, the router returns an ICMP "Time Exceeded" message back to the source. By starting with a TTL of 1 and incrementing by 1 for each subsequent set of packets, `traceroute` can determine the routers along the path to the destination host.

### Basic Syntax

The basic syntax of the `traceroute` command is:

```sh
traceroute [options] destination
```

- **destination**: The IP address or hostname of the target host.

### Common Options

1. **Basic Traceroute Command**:
   ```sh
   traceroute google.com
   ```
   This command traces the path to `google.com`.

2. **Specify Number of Queries per Hop**:
   ```sh
   traceroute -q 3 google.com
   ```
   This sends 3 probe packets per hop (default is 3).

3. **Specify the Maximum Number of Hops**:
   ```sh
   traceroute -m 20 google.com
   ```
   This sets the maximum number of hops to 20 (default is 30).

4. **Specify the Initial TTL Value**:
   ```sh
   traceroute -f 5 google.com
   ```
   This sets the initial TTL value to 5.

5. **Specify the Packet Size**:
   ```sh
   traceroute -s 64 google.com
   ```
   This sets the size of probe packets to 64 bytes.

6. **Use ICMP ECHO Instead of UDP**:
   ```sh
   traceroute -I google.com
   ```
   This uses ICMP ECHO instead of the default UDP packets.

7. **Use TCP SYN Instead of UDP**:
   ```sh
   traceroute -T google.com
   ```
   This uses TCP SYN packets instead of UDP packets.

### Example Usage

1. **Basic Traceroute**:
   ```sh
   traceroute 8.8.8.8
   ```
   This traces the route to Google's public DNS server.

2. **Traceroute with ICMP ECHO**:
   ```sh
   traceroute -I google.com
   ```
   This uses ICMP ECHO requests instead of UDP packets.

3. **Traceroute with a Specific Number of Queries per Hop**:
   ```sh
   traceroute -q 5 google.com
   ```
   This sends 5 queries per hop.

4. **Traceroute with TCP SYN Packets**:
   ```sh
   traceroute -T google.com
   ```
   This uses TCP SYN packets for tracing the route.

### Analyzing Output

The typical output of a `traceroute` command looks like this:

```
traceroute to google.com (172.217.11.142), 30 hops max, 60 byte packets
 1  router.local (192.168.1.1)  1.123 ms  1.072 ms  1.058 ms
 2  10.0.0.1 (10.0.0.1)  2.145 ms  2.123 ms  2.101 ms
 3  * * *
 4  192.0.2.1 (192.0.2.1)  10.123 ms  10.101 ms  10.083 ms
 5  198.51.100.1 (198.51.100.1)  20.123 ms  20.101 ms  20.083 ms
 6  203.0.113.1 (203.0.113.1)  30.123 ms  30.101 ms  30.083 ms
 7  * * *
 8  172.217.11.142 (172.217.11.142)  40.123 ms  40.101 ms  40.083 ms
```

- **Hop Number**: The first column indicates the hop number.
- **Hostname and IP Address**: The second column shows the hostname and IP address of the router at each hop.
- **Round-Trip Times**: The subsequent columns show the round-trip time for each of the three queries sent to each hop.

### Conclusion

The `traceroute` command is a powerful tool for diagnosing network issues, particularly for identifying where delays or failures occur along the route to a destination. It is particularly useful for network administrators and engineers in pinpointing problematic routers or network segments. For more detailed information, consult the `traceroute` man page:

```sh
man traceroute
```

# help 

```
Usage:
  traceroute [ -46dFITnreAUDV ] [ -f first_ttl ] [ -g gate,... ] [ -i device ] [ -m max_ttl ] [ -N squeries ] [ -p port ] [ -t tos ] [ -l flow_label ] [ -w MAX,HERE,NEAR ] [ -q nqueries ] [ -s src_addr ] [ -z sendwait ] [ --fwmark=num ] host [ packetlen ]
Options:
  -4                          Use IPv4
  -6                          Use IPv6
  -d  --debug                 Enable socket level debugging
  -F  --dont-fragment         Do not fragment packets
  -f first_ttl  --first=first_ttl
                              Start from the first_ttl hop (instead from 1)
  -g gate,...  --gateway=gate,...
                              Route packets through the specified gateway
                              (maximum 8 for IPv4 and 127 for IPv6)
  -I  --icmp                  Use ICMP ECHO for tracerouting
  -T  --tcp                   Use TCP SYN for tracerouting (default port is 80)
  -i device  --interface=device
                              Specify a network interface to operate with
  -m max_ttl  --max-hops=max_ttl
                              Set the max number of hops (max TTL to be
                              reached). Default is 30
  -N squeries  --sim-queries=squeries
                              Set the number of probes to be tried
                              simultaneously (default is 16)
  -n                          Do not resolve IP addresses to their domain names
  -p port  --port=port        Set the destination port to use. It is either
                              initial udp port value for "default" method
                              (incremented by each probe, default is 33434), or
                              initial seq for "icmp" (incremented as well,
                              default from 1), or some constant destination
                              port for other methods (with default of 80 for
                              "tcp", 53 for "udp", etc.)
  -t tos  --tos=tos           Set the TOS (IPv4 type of service) or TC (IPv6
                              traffic class) value for outgoing packets
  -l flow_label  --flowlabel=flow_label
                              Use specified flow_label for IPv6 packets
  -w MAX,HERE,NEAR  --wait=MAX,HERE,NEAR
                              Wait for a probe no more than HERE (default 3)
                              times longer than a response from the same hop,
                              or no more than NEAR (default 10) times than some
                              next hop, or MAX (default 5.0) seconds (float
                              point values allowed too)
  -q nqueries  --queries=nqueries
                              Set the number of probes per each hop. Default is
                              3
  -r                          Bypass the normal routing and send directly to a
                              host on an attached network
  -s src_addr  --source=src_addr
                              Use source src_addr for outgoing packets
  -z sendwait  --sendwait=sendwait
                              Minimal time interval between probes (default 0).
                              If the value is more than 10, then it specifies a
                              number in milliseconds, else it is a number of
                              seconds (float point values allowed too)
  -e  --extensions            Show ICMP extensions (if present), including MPLS
  -A  --as-path-lookups       Perform AS path lookups in routing registries and
                              print results directly after the corresponding
                              addresses
  -M name  --module=name      Use specified module (either builtin or external)
                              for traceroute operations. Most methods have
                              their shortcuts (`-I' means `-M icmp' etc.)
  -O OPTS,...  --options=OPTS,...
                              Use module-specific option OPTS for the
                              traceroute module. Several OPTS allowed,
                              separated by comma. If OPTS is "help", print info
                              about available options
  --sport=num                 Use source port num for outgoing packets. Implies
                              `-N 1'
  --fwmark=num                Set firewall mark for outgoing packets
  -U  --udp                   Use UDP to particular port for tracerouting
                              (instead of increasing the port per each probe),
                              default port is 53
  -UL                         Use UDPLITE for tracerouting (default dest port
                              is 53)
  -D  --dccp                  Use DCCP Request for tracerouting (default port
                              is 33434)
  -P prot  --protocol=prot    Use raw packet of protocol prot for tracerouting
  --mtu                       Discover MTU along the path being traced. Implies
                              `-F -N 1'
  --back                      Guess the number of hops in the backward path and
                              print if it differs
  -V  --version               Print version info and exit
  --help                      Read this help and exit

Arguments:
+     host          The host to traceroute to
      packetlen     The full packet length (default is the length of an IP
                    header plus 40). Can be ignored or increased to a minimal
                    allowed value
```



## breakdown

```
-h, --help: This option shows this help message.
-m, --maxhops=NUMBER: This option sets the maximum number of hops. The default is 30.
-n, --noresolve: This option does not resolve hostnames. This means that the traceroute command will only show the IP addresses of the routers or gateways.
-q, --quiet: This option does not print as much output. This can be useful if you are tracing a large number of hops.
-t, --ttl=NUMBER: This option sets the Time To Live (TTL) value. The TTL is a number that is decremented by each router or gateway that the packet passes through. When the TTL reaches 0, the packet is discarded.
-v, --verbose: This option prints more detailed output. This includes the IP address of each router or gateway, as well as the RTT for each hop.

```
