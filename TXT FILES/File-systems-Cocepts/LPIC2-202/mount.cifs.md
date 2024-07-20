# mount.cifs

`mount.cifs` is a command-line utility that allows you to mount a CIFS (Common Internet File System) or SMB (Server Message Block) network share onto your local filesystem. This is useful for accessing files on remote servers as if they were part of your local filesystem. `mount.cifs` is part of the `cifs-utils` package on many Linux distributions.

### General Usage

```bash
mount -t cifs [OPTIONS] //SERVER/SHARE /MOUNT_POINT
```

### Common Options

- `-o user=USERNAME`: Specify the username for authentication.
- `-o password=PASSWORD`: Specify the password for authentication.
- `-o domain=DOMAIN`: Specify the domain for authentication.
- `-o vers=VERSION`: Specify the SMB protocol version (e.g., `1.0`, `2.0`, `2.1`, `3.0`).
- `-o rw`: Mount the share with read and write permissions.
- `-o ro`: Mount the share with read-only permissions.
- `-o uid=UID`: Set the user ID for the owner of the mounted files.
- `-o gid=GID`: Set the group ID for the owner of the mounted files.
- `-o credentials=FILE`: Use a credentials file for authentication.
- `-o file_mode=MODE`: Set the file mode (permissions) for the mounted files.
- `-o dir_mode=MODE`: Set the directory mode (permissions) for the mounted directories.

### Basic Usage Examples

1. **Mounting a Share with Username and Password**

   ```bash
   sudo mount -t cifs //server/share /mnt/mountpoint -o user=myuser,password=mypassword
   ```

   This command mounts the network share `//server/share` at the local directory `/mnt/mountpoint` using the specified username and password.

2. **Mounting a Share with a Credentials File**

   **Credentials File (`/path/to/credentials`):**

   ```plaintext
   username=myuser
   password=mypassword
   domain=mydomain
   ```

   **Mount Command:**

   ```bash
   sudo mount -t cifs //server/share /mnt/mountpoint -o credentials=/path/to/credentials
   ```

   This command mounts the share using the credentials specified in the file.

3. **Mounting with Read-Only Permissions**

   ```bash
   sudo mount -t cifs //server/share /mnt/mountpoint -o user=myuser,password=mypassword,ro
   ```

   This command mounts the share with read-only permissions.

4. **Specifying File and Directory Modes**

   ```bash
   sudo mount -t cifs //server/share /mnt/mountpoint -o user=myuser,password=mypassword,file_mode=0644,dir_mode=0755
   ```

   This command sets the file permissions to `0644` and the directory permissions to `0755`.

5. **Mounting a Share with a Specific SMB Version**

   ```bash
   sudo mount -t cifs //server/share /mnt/mountpoint -o user=myuser,password=mypassword,vers=3.0
   ```

   This command mounts the share using SMB protocol version `3.0`.

### Unmounting the CIFS Share

To unmount a mounted CIFS share, use the `umount` command:

```bash
sudo umount /mnt/mountpoint
```

### Advanced Usage

#### Mounting with Kerberos Authentication

If your network uses Kerberos for authentication, you can mount the share using Kerberos tickets.

1. **Obtain a Kerberos Ticket**

   ```bash
   kinit myuser@MYDOMAIN
   ```

2. **Mount the Share**

   ```bash
   sudo mount -t cifs //server/share /mnt/mountpoint -o sec=krb5
   ```

#### Auto-Mounting CIFS Shares

You can configure your system to automatically mount CIFS shares at boot by adding entries to the `/etc/fstab` file.

**Example `/etc/fstab` Entry:**

```plaintext
//server/share /mnt/mountpoint cifs credentials=/path/to/credentials,uid=1000,gid=1000,vers=3.0 0 0
```

### Summary

| Command                                                                                          | Purpose                                                  | Example Command                                                                                       |
|--------------------------------------------------------------------------------------------------|----------------------------------------------------------|-------------------------------------------------------------------------------------------------------|
| `sudo mount -t cifs //server/share /mnt/mountpoint -o user=myuser,password=mypassword`            | Mount a CIFS share with a username and password          | `sudo mount -t cifs //server/share /mnt/mountpoint -o user=myuser,password=mypassword`                 |
| `sudo mount -t cifs //server/share /mnt/mountpoint -o credentials=/path/to/credentials`           | Mount a CIFS share using a credentials file              | `sudo mount -t cifs //server/share /mnt/mountpoint -o credentials=/path/to/credentials`                |
| `sudo mount -t cifs //server/share /mnt/mountpoint -o user=myuser,password=mypassword,ro`         | Mount a CIFS share with read-only permissions            | `sudo mount -t cifs //server/share /mnt/mountpoint -o user=myuser,password=mypassword,ro`              |
| `sudo mount -t cifs //server/share /mnt/mountpoint -o user=myuser,password=mypassword,file_mode=0644,dir_mode=0755` | Mount a CIFS share with specific file and directory modes | `sudo mount -t cifs //server/share /mnt/mountpoint -o user=myuser,password=mypassword,file_mode=0644,dir_mode=0755` |
| `sudo mount -t cifs //server/share /mnt/mountpoint -o user=myuser,password=mypassword,vers=3.0`   | Mount a CIFS share using a specific SMB version          | `sudo mount -t cifs //server/share /mnt/mountpoint -o user=myuser,password=mypassword,vers=3.0`        |
| `sudo umount /mnt/mountpoint`                                                                    | Unmount a CIFS share                                     | `sudo umount /mnt/mountpoint`                                                                          |

### Conclusion

`mount.cifs` is a versatile tool for mounting SMB/CIFS network shares on Linux systems. It provides various options for authentication, permissions, and protocol versions, making it suitable for different network environments and use cases. Understanding how to use `mount.cifs` effectively can significantly enhance your ability to access and manage network resources in a Linux environment.
