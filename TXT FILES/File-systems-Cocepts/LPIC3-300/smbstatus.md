# smbstatus
The `smbstatus` command is a useful tool in **Samba** for monitoring and displaying the current status of the Samba server, including information about active connections, open files, and the status of various services. It's especially helpful for administrators to troubleshoot issues, monitor file access, and check the current Samba server activity.

### Overview of `smbstatus`

- **Purpose**: Provides information about the active sessions, locked files, and file shares being accessed through the Samba server. It can show details such as who is connected, what files they are accessing, and what shares are being used.
  
- **Use Case**: It is commonly used for troubleshooting or monitoring the activity of Samba shares, users, and open files. It is useful for checking who is connected to the server and what files are being accessed or locked.

### General Syntax:
```bash
smbstatus [options]
```

### Key Options for `smbstatus`

- `-b`, `--browsing`: Show information about browsing.
- `-L`, `--locked`: Show a list of locked files.
- `-S`, `--sessions`: Show session information (i.e., users connected to the Samba server).
- `-O`, `--open-files`: Show open file information.
- `-u <user>`, `--user <user>`: Display status information for a specific user.
- `-p`, `--processes`: Show processes using the Samba server.
- `-h`, `--help`: Displays help about the command and available options.

### Example Usage

#### 1. Checking Active Sessions
To see who is currently connected to the Samba server and what shares they are accessing, run:

```bash
smbstatus -S
```

This will display the session information, showing active users, their machine names, and the shares they are connected to.

Example output:
```
Samba version 4.12.5
PID     Username     Group        Machine
----------------------------------------------------
12345   user1        users        192.168.1.10
67890   user2        users        192.168.1.12
```

#### 2. Viewing Locked Files
To see which files are currently locked by users on the Samba server, run:

```bash
smbstatus -L
```

This will list all files that are locked and provide information about the user who locked them, the file path, and the type of lock (read/write).

Example output:
```
Locked files:
Pid     Uid     DenyMode  Access  R/W  Oplock  ShareMode  Path
---------------------------------------------------------------
12345   1001    DENY_ALL  0x0001  R/W  None    Shared     /mnt/data/file.txt
```

#### 3. Listing Open Files
To display a list of all currently open files, use:

```bash
smbstatus -O
```

This will show detailed information about open files on the Samba server, such as which user has the file open, its path, and its lock state.

Example output:
```
Samba version 4.12.5
PID     Username     Group        Machine   File Path
--------------------------------------------------------
12345   user1        users        192.168.1.10 /mnt/data/file1.txt
67890   user2        users        192.168.1.12 /mnt/data/file2.txt
```

#### 4. Displaying File/Share Access by a Specific User
If you want to check the status of a specific user, run:

```bash
smbstatus -u user1
```

This will show only the sessions, open files, or locked files for the user `user1`.

#### 5. Full Overview of Current Samba Activity
To get a full overview of the Samba server status, including sessions, open files, and locked files, run:

```bash
smbstatus
```

This will combine all information, including active users, file locks, and open files.

Example output:
```
Samba version 4.12.5
PID     Username     Group        Machine    File Path              Lock Mode
--------------------------------------------------------------------------------
12345   user1        users        192.168.1.10 /mnt/data/file1.txt   Read-Write
67890   user2        users        192.168.1.12 /mnt/data/file2.txt   Read-Only
```

### Practical Use Cases

1. **Troubleshooting File Access Issues**: If users are having trouble accessing files or if there are issues with locked files, `smbstatus` can help identify which files are locked and by whom.
   
2. **Monitoring Active Sessions**: If you need to check which users are currently connected to a Samba share, `smbstatus -S` can provide a quick overview of all active sessions.
   
3. **Ensuring Share Availability**: You can use `smbstatus -O` to monitor which files are currently being accessed. This can help ensure that files are being used as expected or identify any issues such as stalled processes.

4. **Security Auditing**: By regularly checking the sessions and open files, administrators can spot unauthorized access or unusual file activities.

### Conclusion

The `smbstatus` command is an essential tool for managing and monitoring Samba servers. Whether you need to check active user sessions, track file access, or diagnose issues with locked files, `smbstatus` provides real-time insights into the state of your Samba server. It's a valuable tool for administrators, especially when troubleshooting or performing routine monitoring of a Samba environment.
