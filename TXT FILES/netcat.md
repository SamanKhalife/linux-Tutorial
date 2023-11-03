# netcat

Netcat (often abbreviated to nc) is a computer networking utility for reading from and writing to network connections using TCP or UDP. The command is designed to be a dependable back-end that can be used directly or easily driven by other programs and scripts. At the same time, it is a feature-rich network debugging and investigation tool, since it can produce almost any kind of connection its user could need and has a number of built-in capabilities.

Here are some of the things that netcat can be used for:

* Port scanning: Netcat can be used to scan ports on a remote host to see if they are open.
* File transfer: Netcat can be used to transfer files between two hosts.
* Remote command execution: Netcat can be used to execute commands on a remote host.
* Tunneling: Netcat can be used to create a tunnel between two hosts.
* Network debugging: Netcat can be used to debug network problems.

Netcat is a powerful tool that can be used for a variety of tasks. It is a valuable tool for network administrators, security professionals, and anyone who needs to work with networks.

Here are some examples of how netcat can be used:

* To scan ports on a remote host, you can use the following command:

```
nc -n -v -p 80 192.168.1.1
```

This command will scan the port 80 on the host 192.168.1.1. The `-n` option tells netcat to not resolve hostnames, the `-v` option tells netcat to be verbose, and the `-p` option tells netcat to listen on the port 80.

* To transfer a file from one host to another, you can use the following command:

```
nc -c 'cat file.txt' 192.168.1.1 2000
```

This command will transfer the file `file.txt` from the local host to the host 192.168.1.1 on port 2000. The `-c` option tells netcat to read from a file, and the `cat` command will read the contents of the file and send them to the remote host.

* To execute a command on a remote host, you can use the following command:

```
nc -c 'echo "ls"' 192.168.1.1 2000
```

This command will execute the command `ls` on the remote host 192.168.1.1 on port 2000. The `-c` option tells netcat to read from a command, and the `echo` command will echo the string "ls" to the remote host.

* To create a tunnel between two hosts, you can use the following command:

```
nc -l -p 8080 | nc 192.168.1.1 80
```

This command will create a tunnel between the local host and the host 192.168.1.1. The first `nc` command will listen on port 8080 on the local host, and the second `nc` command will connect to the host 192.168.1.1 on port 80. Any data that is sent to the local host on port 8080 will be forwarded to the host 192.168.1.1 on port 80, and vice versa.

* To debug network problems, you can use netcat to connect to a remote host and send specific commands. For example, you can use netcat to connect to a web server and send the `HEAD` command to see if the web server is responding.

Netcat is a powerful tool that can be used for a variety of tasks. It is a valuable tool for network administrators, security professionals, and anyone who needs to work with networks.


# help 

```
usage: nc [-46CDdFhklNnrStUuvZz] [-I length] [-i interval] [-M ttl]
          [-m minttl] [-O length] [-P proxy_username] [-p source_port]
          [-q seconds] [-s sourceaddr] [-T keyword] [-V rtable] [-W recvlimit]
          [-w timeout] [-X proxy_protocol] [-x proxy_address[:port]]
          [destination] [port]
```
