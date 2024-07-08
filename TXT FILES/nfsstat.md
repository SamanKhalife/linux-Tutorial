# nfsstat

NFS (Network File System) and its related utilities. `nfsstat` is a command-line utility used to display NFS client and server statistics. Here's an overview of how `nfsstat` works and its typical usage:

### nfsstat Command Overview

#### Purpose
- **nfsstat** is used to report NFS client and server statistics, including various counters and performance metrics related to NFS operations.

#### Usage
- **Syntax**: `nfsstat [options]`
- **Options**:
  - `-c`: Displays client statistics.
  - `-s`: Displays server statistics.
  - `-m`: Displays mount information.
  - `-n`: Displays NFSv4 statistics.
  - `-o`: Displays NFSv4 operation statistics.
  - `-r`: Displays RPC (Remote Procedure Call) statistics.
  - `-t`: Displays TCP statistics for NFS.

#### Typical Output
- **Client Statistics**:
  - Shows information such as number of NFS operations (reads, writes, etc.), caching statistics, and errors encountered by NFS clients.

- **Server Statistics**:
  - Provides details on server-side NFS operations, including read/write operations, cache utilization, and errors encountered by NFS servers.

- **Mount Information**:
  - Lists NFS mounts and associated details like mount options and NFS version used.

- **NFSv4 and RPC Statistics**:
  - Provides NFSv4-specific statistics and RPC performance metrics, including timings and counts for various RPC operations.

- **TCP Statistics**:
  - Displays statistics related to NFS over TCP connections, such as connection states, retransmissions, and errors.

#### Examples
- Display client statistics:
  ```bash
  nfsstat -c
  ```

- Display server statistics:
  ```bash
  nfsstat -s
  ```

- Display NFSv4 operation statistics:
  ```bash
  nfsstat -n
  ```

#### Usage Notes
- **Privileges**: `nfsstat` typically requires superuser (root) privileges to access all statistics, especially server-side information.
- **Interpretation**: Understanding the output requires familiarity with NFS concepts and performance metrics, such as latency, throughput, and error rates.

### Conclusion
`nfsstat` is a valuable tool for administrators and developers managing NFS environments. It provides insights into NFS performance and usage, helping diagnose issues, optimize configurations, and monitor NFS operations effectively. Understanding its various options and interpreting its output is crucial for maintaining reliable NFS services.

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

