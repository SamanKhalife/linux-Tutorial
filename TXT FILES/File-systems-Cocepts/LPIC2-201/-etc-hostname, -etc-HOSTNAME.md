# /etc/hostname and /etc/HOSTNAME

The files `/etc/hostname` and `/etc/HOSTNAME` are used to configure the system's hostname in Unix-like operating systems. The location and usage of these files can vary between different Linux distributions. Here's an overview of these files and how they are used:

#### /etc/hostname

This file is commonly used in Debian-based distributions (such as Debian, Ubuntu, and their derivatives) to set the system's hostname.

##### Contents of /etc/hostname

The `/etc/hostname` file typically contains a single line with the hostname of the system. For example:

```sh
myhostname
```

In this example, `myhostname` is the name of the system.

##### Setting the Hostname

1. **Edit the /etc/hostname File**:
   ```sh
   sudo nano /etc/hostname
   ```
   Enter the desired hostname and save the file.

2. **Apply the New Hostname**:
   To apply the new hostname without rebooting, use the `hostname` command:
   ```sh
   sudo hostname -F /etc/hostname
   ```

3. **Verify the Hostname**:
   ```sh
   hostname
   ```

##### Additional Configuration

It is also good practice to update the `/etc/hosts` file to map the new hostname to `127.0.1.1`. For example:

```sh
127.0.0.1   localhost
127.0.1.1   myhostname
```

#### /etc/HOSTNAME

This file is typically used in Red Hat-based distributions (such as RHEL, CentOS, Fedora, and their derivatives) to set the system's hostname.

##### Contents of /etc/HOSTNAME

The `/etc/HOSTNAME` file also contains a single line with the hostname of the system. For example:

```sh
myhostname
```

##### Setting the Hostname

1. **Edit the /etc/HOSTNAME File**:
   ```sh
   sudo nano /etc/HOSTNAME
   ```
   Enter the desired hostname and save the file.

2. **Apply the New Hostname**:
   To apply the new hostname without rebooting, use the `hostnamectl` command:
   ```sh
   sudo hostnamectl set-hostname myhostname
   ```

3. **Verify the Hostname**:
   ```sh
   hostname
   ```

##### Additional Configuration

It is also good practice to update the `/etc/hosts` file to map the new hostname to `127.0.0.1` or `127.0.1.1`. For example:

```sh
127.0.0.1   localhost
127.0.1.1   myhostname
```

### Key Differences and Usage

1. **File Locations**:
   - Debian-based systems: `/etc/hostname`.
   - Red Hat-based systems: `/etc/HOSTNAME`.

2. **Hostname Commands**:
   - Debian-based systems: `hostname -F /etc/hostname`.
   - Red Hat-based systems: `hostnamectl set-hostname <hostname>`.

3. **Configuration Syntax**:
   - Both files use a single line to specify the hostname.

### Conclusion

The `/etc/hostname` and `/etc/HOSTNAME` files are essential for setting the system's hostname in Unix-like operating systems. Understanding their locations and usage in different Linux distributions helps in managing and configuring hostnames effectively. Remember to update the `/etc/hosts` file to ensure proper hostname resolution.
