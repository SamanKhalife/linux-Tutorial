# nslookup

The nslookup command is a network administration command-line tool for querying the Domain Name System (DNS) to obtain the mapping between domain name and IP address, or other DNS records.

The syntax for the nslookup command is:

`nslookup [options] hostname [server]`

The hostname is the domain name or IP address that you want to query. The server is the DNS server that you want to use to query the DNS. If you do not specify a server, the default DNS server will be used.

# help 

```
nslookup [options] hostname [server]

Options:

-?, --help     Print this help message.
-a, --all       List all records.
-c, --cache     Use the local cache.
-d, --debug      Print debugging information.
-f, --family    Set address family to AF_INET or AF_INET6.
-i, --info      List additional information about records.
-k, --keep-query Continue to query until the answer has expired.
-n, --no-cache  Do not use the local cache.
-p, --port      Set the port number to use for queries.
-r, --recurse   Recursively query for all records.
-s, --server    Specify the name server to use.
-t, --type      Specify the type of record to query.
-v, --verbose    Print verbose output.
-w, --timeout   Set the timeout value in seconds.

For more information, see the nslookup man page.
```

