# syslogd





# help 

```
Usage: syslogd [OPTION...] 
Log system messages.

  -4, --ipv4                 restrict to IPv4 transport (default)
  -6, --ipv6                 restrict to IPv6 transport
  -a SOCKET                  add unix socket to listen to (up to 19)
  -b, --bind=ADDR            bind listener to this address/name
  -B, --bind-port=PORT       bind listener to this port
  -d, --debug                print debug information (implies --no-detach)
  -D, --rcdir=DIR            override configuration directory (default:
                             /etc/syslog.d)
  -f, --rcfile=FILE          override configuration file (default:
                             /etc/syslog.conf)
  -h, --hop                  forward messages from remote hosts
      --ipany                allow transport with IPv4 and IPv6
  -l HOSTLIST                log hosts in HOSTLIST by their hostname
  -m, --mark=INTVL           specify timestamp interval in minutes (0 for no
                             timestamping)
      --no-forward           do not forward any messages (overrides --hop)
      --no-klog              do not listen to kernel log device /proc/kmsg
      --no-unixaf            do not listen on unix domain sockets (overrides -a
                             and -p)
  -n, --no-detach            do not enter daemon mode
  -p, --socket=FILE          override default unix domain socket /dev/log
  -P, --pidfile=FILE         override pidfile (default: /run/syslog.pid)
  -r, --inet                 receive remote messages via internet domain socket
                            
  -s DOMAINLIST              list of domains which should be stripped from the
                             FQDN of hosts before logging their name
  -S, --sync                 force a file sync on every line
  -T, --local-time           set local time on received messages
  -?, --help                 give this help list
      --usage                give a short usage message
  -V, --version              print program version

```



## breakdown

```

```
