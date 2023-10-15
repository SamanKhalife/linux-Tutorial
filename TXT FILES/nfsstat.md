# 

The nfsstat command can be used to troubleshoot NFS performance problems. If you are seeing a high number of bad calls or retransmissions, it may indicate a problem with the network or the server.

here is the output of the nfsstat command without any options:

```
nfsstat

Server: 192.168.1.100
Version: 4.0.0

Client NFS   RPC   Calls          Badcalls           Retrans
--------  -------- ------  --------  --------  --------
Total        49996    0          0          0
Lookups      24935    0          0          0
Reads        19222    0          0          0
Writes       6839     0          0          0
Creats       500       0          0          0
Removes       12       0          0          0
```

This output shows the following information:

The server IP address
The NFS version
The number of NFS calls made by the client
The number of bad calls (calls that were rejected by the server)
The number of retransmissions (RPC calls that had to be sent again because they were not acknowledged by the server)




# help 

```
Usage: nfsstat [OPTION]...

  -m, --mounts          Show statistics on mounted NFS filesystems
  -c, --client          Show NFS client statistics
  -s, --server          Show NFS server statistics
  -2                    Show NFS version 2 statistics
  -3                    Show NFS version 3 statistics
  -4                    Show NFS version 4 statistics
  -o [facility]         Show statistics on particular facilities.
     nfs                NFS protocol information
     rpc                General RPC information
     net                Network layer statistics
     fh                 Usage information on the server's file handle cache
     io                 Usage information on the server's io statistics
     ra                 Usage information on the server's read ahead cache
     rc                 Usage information on the server's request reply cache
     all                Select all of the above
  -v, --verbose, --all  Same as '-o all'
  -r, --rpc             Show RPC statistics
  -n, --nfs             Show NFS statistics
  -Z[#], --sleep[=#]    Collects stats until interrupted.
                            Cumulative stats are then printed
                            If # is provided, stats will be output every
                            # seconds.
  -S, --since file      Shows difference between current stats and those in 'file'
  -l, --list            Prints stats in list format
  --version             Show program version
  --help                What you just did

```

