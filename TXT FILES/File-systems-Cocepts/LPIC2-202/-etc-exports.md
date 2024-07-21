# /etc/exports

The `/etc/exports` file is a configuration file used by the Network File System (NFS) server on Unix-like operating systems to specify which directories are to be shared (exported) over the network and the access permissions for these directories.

### Structure of `/etc/exports`

Each line in the `/etc/exports` file specifies a directory to be shared and the options for sharing it. The general syntax is:

```plaintext
directory  client(options) client(options) ...
```

- **directory**: The path to the directory on the NFS server that is to be shared.
- **client**: The hostname or IP address of the client that is allowed to access the shared directory.
- **options**: A set of options specifying how the directory is shared and what permissions the client has.

### Example `/etc/exports` File

```plaintext
# Share /srv/nfs to a specific client with read-write access
/srv/nfs 192.168.1.100(rw,sync,no_subtree_check)

# Share /home directory to all clients in the 192.168.1.0/24 network with read-only access
/home 192.168.1.0/24(ro,sync,no_subtree_check)

# Share /var/nfs with two specific clients with read-write access and root squashing
/var/nfs 192.168.1.101(rw,sync,no_subtree_check,root_squash) 192.168.1.102(rw,sync,no_subtree_check,root_squash)
```

### Options Explained

- **rw**: Allow read and write access to the shared directory.
- **ro**: Allow read-only access to the shared directory.
- **sync**: Ensure that changes to the shared directory are immediately written to disk.
- **no_subtree_check**: Disable subtree checking, which improves performance.
- **root_squash**: Map requests from root on the client to an anonymous UID/GID on the server, enhancing security.
- **no_root_squash**: Allow root on the client to have root privileges on the server.
- **anonuid**: Specify the UID of the anonymous account.
- **anongid**: Specify the GID of the anonymous account.
- **all_squash**: Map all user and group IDs to the anonymous user.

### Managing NFS Exports

After editing `/etc/exports`, use the following commands to apply changes and manage the NFS server:

#### Apply Export Changes

```bash
sudo exportfs -ra
```

#### Show Current NFS Exports

```bash
sudo exportfs -v
```

#### Start NFS Server

```bash
sudo systemctl start nfs-server
```

#### Enable NFS Server at Boot

```bash
sudo systemctl enable nfs-server
```

#### Restart NFS Server

```bash
sudo systemctl restart nfs-server
```

### Setting Up an NFS Server and Client

#### NFS Server Setup

1. **Install NFS Server Package**

   On Debian/Ubuntu:
   ```bash
   sudo apt update
   sudo apt install nfs-kernel-server
   ```

   On CentOS/RHEL:
   ```bash
   sudo yum install nfs-utils
   ```

2. **Edit `/etc/exports`**

   ```bash
   sudo nano /etc/exports
   ```

   Add the directories you want to share.

3. **Apply Export Changes**

   ```bash
   sudo exportfs -ra
   ```

4. **Start and Enable NFS Server**

   ```bash
   sudo systemctl start nfs-server
   sudo systemctl enable nfs-server
   ```

#### NFS Client Setup

1. **Install NFS Client Package**

   On Debian/Ubuntu:
   ```bash
   sudo apt update
   sudo apt install nfs-common
   ```

   On CentOS/RHEL:
   ```bash
   sudo yum install nfs-utils
   ```

2. **Create Mount Point**

   ```bash
   sudo mkdir -p /mnt/nfs
   ```

3. **Mount NFS Share**

   ```bash
   sudo mount -t nfs 192.168.1.100:/srv/nfs /mnt/nfs
   ```

4. **Verify Mount**

   ```bash
   df -h | grep nfs
   ```

5. **Automatic Mount at Boot**

   Add the following line to `/etc/fstab`:

   ```plaintext
   192.168.1.100:/srv/nfs /mnt/nfs nfs defaults 0 0
   ```

### Security Considerations

- **Restrict Access**: Only allow specific clients or networks to access NFS shares.
- **Use Firewalls**: Ensure that only trusted IP addresses can connect to the NFS server.
- **Secure Options**: Use `root_squash` and other secure options to limit potential damage from compromised clients.
- **Network Security**: Consider using a VPN or secure network for NFS traffic.

### Conclusion

The `/etc/exports` file is a vital part of configuring an NFS server, allowing administrators to specify which directories are shared and the permissions granted to different clients. Proper management and security considerations are essential to ensure reliable and secure NFS operations.
