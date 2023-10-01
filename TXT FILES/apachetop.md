# apachetop

ApacheTop is a command-line tool that can be used to monitor the performance of an Apache web server. It displays a graphical representation of the server's traffic and activity.

Here are some examples of how to use the apachetop command:

```
# To monitor the Apache web server on port 80:
apachetop -p 80

# To monitor the Apache web server on port 80 and update the output every 5 seconds:
apachetop -p 80 -i 5

# To show help for the apachetop command:
apachetop -h
```

# help 

```
  -q          keep query strings [no]
  -l          lowercase all URLs [no]
  -s num      keep num path segments of URL [all]
  -p          preserve protocol at front of referrers [no]
  -r          resolve hostnames/IPs into each other [no]

Stats options:
  Supply up to one of the following two. default: [-T 30]
  -H hits     remember stats for this many hits
  -T secs     remember stats for this many seconds

  -d secs     refresh delay in seconds [5]

  -v          show version
  -h          this help

```

## breakdown

```
-p, --port=PORT: This option specifies the port number of the Apache web server. The default port number is 80.
-l, --logfile=LOGFILE: This option specifies the log file that the apachetop command should read. The default log file is /var/log/apache2/access.log.
-i, --interval=INTERVAL: This option specifies the interval at which the apachetop command should update the output. The default interval is 1 second.
-d, --debug: This option enables debug output.
-h, --help: This option shows this help message.
-V, --version: This option prints version information.
```

