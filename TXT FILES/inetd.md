# inetd

The inetd command is a superserver that listens for incoming connections on a variety of ports and then forks a child process to handle each connection. The child process then invokes the appropriate service to handle the connection.

# help 

```
usage: inetd [-dEil] [-q len] [-R rate] [configuration_file]

inetd [options]

Internet superserver.

Options:

-d, --debug     Print debugging information.
-f, --foreground Run in foreground.
-l, --listen     Listen for connections on all interfaces.
-n, --nowait     Do not fork a child process for each connection.
-r, --reuseaddr  Reuse the listening socket.
-v, --verbose    Print verbose information.
-h, --help       Show this help message.

For more information, see the inetd man page
```

breakdown 

```
-d: This option tells inetd to print debugging information.
-f: This option tells inetd to run in foreground.
-l: This option tells inetd to listen for connections on all interfaces.
-n: This option tells inetd to do not fork a child process for each connection.
-r: This option tells inetd to reuse the listening socket.
-v: This option tells inetd to print verbose information.
-h: This option shows this help message.
```
