# /etc/services

The `/etc/services` file on Unix-like operating systems is a crucial configuration file that maps network service names to port numbers and protocols. This file helps the operating system and applications identify which port numbers should be used for specific network services. It serves as a reference for both TCP and UDP port assignments and is used by various network services and tools.

### Structure of `/etc/services`

The file typically contains lines formatted as follows:

```
service_name   port/protocol   [aliases...]
```

Here's a breakdown of each field:

1. **`service_name`**: The name of the service or application using the port.
2. **`port`**: The port number assigned to the service.
3. **`protocol`**: The transport protocol used by the service (`tcp`, `udp`, etc.).
4. **`aliases`** (optional): Alternative names for the service.

### Example Entries

Here's a snippet from a typical `/etc/services` file:

```
http            80/tcp
https           443/tcp
ftp             21/tcp
smtp            25/tcp
ssh             22/tcp
domain          53/udp
```

### Common Entries

- **`http`**
  - **Port Number:** 80
  - **Protocol:** TCP
  - **Description:** Used for HTTP web traffic.

- **`https`**
  - **Port Number:** 443
  - **Protocol:** TCP
  - **Description:** Used for HTTPS secure web traffic.

- **`ftp`**
  - **Port Number:** 21
  - **Protocol:** TCP
  - **Description:** Used for FTP file transfer.

- **`smtp`**
  - **Port Number:** 25
  - **Protocol:** TCP
  - **Description:** Used for SMTP email sending.

- **`ssh`**
  - **Port Number:** 22
  - **Protocol:** TCP
  - **Description:** Used for SSH secure shell access.

- **`domain`**
  - **Port Number:** 53
  - **Protocol:** UDP
  - **Description:** Used for DNS queries.

### Usage

The `/etc/services` file is used by various system utilities and network programs, such as:

- **`netstat`**: Displays network connections, routing tables, and interface statistics. It relies on `/etc/services` to provide human-readable service names.
- **`ss`**: A utility to investigate sockets, also uses `/etc/services` for service name resolution.
- **`nmap`**: A network scanner that uses `/etc/services` to map scanned ports to service names.

### Managing `/etc/services`

- **Editing**: Directly edit the file using a text editor like `vi` or `nano` to add or change service definitions. Ensure to have superuser permissions.
  ```sh
  sudo nano /etc/services
  ```

- **Adding Custom Services**: You can add your own custom services for internal use. Just ensure that they do not conflict with existing entries.

### Example of Adding a Custom Service

To add a custom service:

1. Open the file in a text editor with superuser privileges:
   ```sh
   sudo nano /etc/services
   ```

2. Add a new line for your service:
   ```
   myservice        12345/tcp
   ```

3. Save and exit the editor.

### Conclusion

The `/etc/services` file is a fundamental part of Unix-like systems that helps in mapping network services to their respective ports and protocols. While it is not frequently changed, understanding its format and function is essential for network configuration and troubleshooting.
