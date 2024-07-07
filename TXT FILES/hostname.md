# hostname

The `hostname` command in Linux is used to display or set the system's hostname, which is a unique identifier assigned to a machine in a network. The hostname is used to identify the machine within a network and is essential for network communication and administration tasks.

### Basic Usage

```sh
hostname [options] [new_hostname]
```

- **options**: Optional flags to modify the command's behavior.
- **new_hostname**: The new hostname you want to set for the machine.

### Common Options and Examples

1. **Display the Current Hostname**:
   ```sh
   hostname
   ```
   This command displays the current hostname of the system.

2. **Set a New Hostname**:
   ```sh
   sudo hostname new_hostname
   ```
   This command sets the system's hostname to `new_hostname`. Note that you need superuser privileges to change the hostname.

3. **Display the Short Hostname**:
   ```sh
   hostname -s
   ```
   This command displays the short hostname (up to the first dot).

4. **Display the FQDN (Fully Qualified Domain Name)**:
   ```sh
   hostname -f
   ```
   This command displays the system's fully qualified domain name (FQDN).

5. **Display the Alias Names**:
   ```sh
   hostname -a
   ```
   This command displays the alias names of the host.

6. **Display the IP Address**:
   ```sh
   hostname -i
   ```
   This command displays the IP address associated with the hostname.

7. **Display the DNS Domain Name**:
   ```sh
   hostname -d
   ```
   This command displays the DNS domain name of the system.

8. **Display All Hostnames**:
   ```sh
   hostname -A
   ```
   This command displays all configured hostnames.

### Example Usage

1. **Check the Current Hostname**:
   ```sh
   $ hostname
   myhostname
   ```

2. **Change the Hostname Temporarily**:
   ```sh
   $ sudo hostname newhostname
   ```
   This change will last until the next reboot.

3. **Set the Hostname Permanently**:
   To set the hostname permanently, you need to edit the `/etc/hostname` file and the `/etc/hosts` file.

   ```sh
   $ sudo nano /etc/hostname
   ```
   Change the content to the new hostname and save the file.

   ```sh
   $ sudo nano /etc/hosts
   ```
   Update the line that starts with `127.0.1.1` to match the new hostname:
   ```
   127.0.1.1   newhostname
   ```

   After making these changes, apply them by running:
   ```sh
   $ sudo systemctl restart systemd-logind.service
   ```

### Practical Applications

- **Network Configuration**: Setting a hostname is essential for identifying machines in a network, facilitating easier communication and management.
- **Server Management**: Hostnames help in managing multiple servers and services, making it easier to identify and configure them.
- **Troubleshooting**: Knowing the hostname can assist in diagnosing network issues and identifying the source of problems.

### Conclusion

The `hostname` command is a straightforward yet essential tool for managing and configuring the system's network identity. For detailed information, you can refer to the man page:

```sh
man hostname
```
# help 

```

```

