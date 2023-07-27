
Sure, here is the help output for the dig command:

dig [options] [domain]

Perform DNS lookups.

Options:

-h, --help           Show this help message.
-t, --type=TYPE   Set the DNS record type.
-c, --class=CLASS   Set the DNS record class.
-p, --port=PORT   Set the DNS port number.
-q, --query=QUERY   Set the DNS query.
-x, --xfr   Perform a zone transfer.
-v, --verbose   Increase verbosity.
-w, --width=WIDTH   Set the output width.

Examples:
```
    dig @8.8.8.8 google.com
    dig -t A google.com
    dig -x google.com
```

The dig command is a powerful tool that can be used to perform DNS lookups. It can be used to see which DNS servers are authoritative for a domain, and to get the IP addresses of hosts.


```
dig [options] [domain]

Perform DNS lookups.

Options:

-h, --help           Show this help message.
-t, --type=TYPE   Set the DNS record type.
-c, --class=CLASS   Set the DNS record class.
-p, --port=PORT   Set the DNS port number.
-q, --query=QUERY   Set the DNS query.
-x, --xfr   Perform a zone transfer.
-v, --verbose   Increase verbosity.
-w, --width=WIDTH   Set the output width.

Examples:

    dig @8.8.8.8 google.com
    dig -t A google.com
    dig -x google.com
```

## breakdown

```
-h, --help: This option shows this help message.
-t, --type=TYPE: This option sets the DNS record type. The possible values for TYPE are:
A: A host address.
AAAA: A host address for IPv6.
CNAME: A canonical name.
MX: Mail exchanger.
NS: Name server.
PTR: Pointer.
SOA: Start of authority.
-c, --class=CLASS: This option sets the DNS record class. The possible values for CLASS are:
IN: Internet.
CH: Chaos.
HS: Hesiod.
-p, --port=PORT: This option sets the DNS port number. The default port number is 53.
-q, --query=QUERY: This option sets the DNS query. The query can be a domain name or a specific DNS record type.
-x, --xfr: This option performs a zone transfer. A zone transfer is a request for all of the DNS records for a zone.
-v, --verbose: This option increases verbosity. The more verbose the output, the more detailed it will be.
-w, --width=WIDTH: This option sets the output width. The width is the maximum number of characters per line.
```
