# smbclient

`smbclient` is a command-line tool in Samba that allows users to interact with Windows network shares and printers. It provides an FTP-like interface to access and manage files on SMB/CIFS servers, including Windows machines and other Samba servers. This tool is useful for transferring files, accessing shared directories, and testing SMB connections.

#### General Usage

```bash
smbclient [OPTION] [SERVICE] [COMMAND]
```

### Common Options

- `-L host`: List shares available on a host.
- `-U user`: Specify the username for authentication.
- `-W workgroup`: Specify the workgroup or domain.
- `-I ip_address`: Specify the IP address of the host.
- `-c command`: Execute a command immediately after connecting.

### Basic Usage Examples

1. **List Shares on a Server**

   ```bash
   smbclient -L //hostname -U username
   ```

   This command lists all available shares on the specified server.

   Example:

   ```bash
   smbclient -L //server -U myuser
   ```

   Output:

   ```plaintext
   Enter MYUSER's password:
   Sharename       Type      Comment
   ---------       ----      -------
   print$          Disk      Printer Drivers
   public          Disk      Public Share
   IPC$            IPC       IPC Service (server)
   ```

2. **Connecting to a Share**

   ```bash
   smbclient //hostname/sharename -U username
   ```

   This command connects to a specific share on the server.

   Example:

   ```bash
   smbclient //server/public -U myuser
   ```

   Output:

   ```plaintext
   Enter MYUSER's password:
   smb: \>
   ```

3. **Uploading a File**

   Once connected to a share, you can upload files using the `put` command.

   ```bash
   smb: \> put localfile remotefile
   ```

   Example:

   ```bash
   smb: \> put /path/to/localfile.txt remotefile.txt
   ```

   Output:

   ```plaintext
   putting file /path/to/localfile.txt as \remotefile.txt (0.0 kb/s) (average 0.0 kb/s)
   ```

4. **Downloading a File**

   Similarly, you can download files using the `get` command.

   ```bash
   smb: \> get remotefile localfile
   ```

   Example:

   ```bash
   smb: \> get remotefile.txt /path/to/localfile.txt
   ```

   Output:

   ```plaintext
   getting file \remotefile.txt of size 12345 as /path/to/localfile.txt (0.0 kb/s) (average 0.0 kb/s)
   ```

5. **Listing Files in a Directory**

   ```bash
   smb: \> ls
   ```

   Example:

   ```bash
   smb: \> ls
   ```

   Output:

   ```plaintext
     .                                   D        0  Mon Jul 19 12:34:56 2021
     ..                                  D        0  Mon Jul 19 12:34:56 2021
     file.txt                            A    12345  Mon Jul 19 12:34:56 2021

                   12345678 blocks of size 1024. 567890 blocks available
   ```

6. **Creating a Directory**

   ```bash
   smb: \> mkdir dirname
   ```

   Example:

   ```bash
   smb: \> mkdir newfolder
   ```

   Output:

   ```plaintext
   smb: \>
   ```

7. **Deleting a File**

   ```bash
   smb: \> del filename
   ```

   Example:

   ```bash
   smb: \> del oldfile.txt
   ```

   Output:

   ```plaintext
   smb: \>
   ```

8. **Executing Multiple Commands**

   You can execute multiple commands by using the `-c` option.

   ```bash
   smbclient //hostname/sharename -U username -c 'command1; command2'
   ```

   Example:

   ```bash
   smbclient //server/public -U myuser -c 'ls; get remotefile.txt /path/to/localfile.txt'
   ```

   Output:

   ```plaintext
   Enter MYUSER's password:
     .                                   D        0  Mon Jul 19 12:34:56 2021
     ..                                  D        0  Mon Jul 19 12:34:56 2021
     remotefile.txt                      A    12345  Mon Jul 19 12:34:56 2021

                   12345678 blocks of size 1024. 567890 blocks available
   getting file \remotefile.txt of size 12345 as /path/to/localfile.txt (0.0 kb/s) (average 0.0 kb/s)
   ```

### Advanced Usage

#### Authenticating with a Password File

Instead of entering a password interactively, you can use a credentials file.

1. **Create a Credentials File**

   ```plaintext
   username=myuser
   password=mypassword
   domain=mydomain
   ```

2. **Use the Credentials File**

   ```bash
   smbclient //hostname/sharename -A /path/to/credentialsfile
   ```

   Example:

   ```bash
   smbclient //server/public -A /path/to/credentials.txt
   ```

#### Using Kerberos Authentication

If your network uses Kerberos for authentication, you can connect using your Kerberos ticket.

1. **Obtain a Kerberos Ticket**

   ```bash
   kinit myuser@MYDOMAIN
   ```

2. **Connect with Kerberos Authentication**

   ```bash
   smbclient //hostname/sharename -k
   ```

   Example:

   ```bash
   smbclient //server/public -k
   ```

### Summary

| Command                              | Purpose                                                  | Example Command                                                |
|--------------------------------------|----------------------------------------------------------|----------------------------------------------------------------|
| `smbclient -L //hostname -U username` | List available shares on a server                        | `smbclient -L //server -U myuser`                              |
| `smbclient //hostname/sharename -U username` | Connect to a specific share                               | `smbclient //server/public -U myuser`                          |
| `smb: \> put localfile remotefile`   | Upload a file to the connected share                      | `put /path/to/localfile.txt remotefile.txt`                    |
| `smb: \> get remotefile localfile`   | Download a file from the connected share                  | `get remotefile.txt /path/to/localfile.txt`                    |
| `smb: \> ls`                         | List files in the current directory                       | `ls`                                                           |
| `smb: \> mkdir dirname`              | Create a new directory                                    | `mkdir newfolder`                                              |
| `smb: \> del filename`               | Delete a file                                             | `del oldfile.txt`                                              |
| `smbclient //hostname/sharename -A /path/to/credentialsfile` | Use a credentials file for authentication                 | `smbclient //server/public -A /path/to/credentials.txt`        |
| `smbclient //hostname/sharename -k`  | Use Kerberos authentication                               | `smbclient //server/public -k`                                 |

### Conclusion

`smbclient` is a powerful tool for interacting with SMB/CIFS shares from the command line. It provides a wide range of options and commands for accessing, managing, and transferring files on network shares. Whether you need to list available shares, connect to a specific share, or perform file operations, `smbclient` offers the functionality needed for effective network file management in a Samba environment. Understanding its options and commands is essential for system administrators and users working with network shares.
