# /etc/samba/

The `/etc/samba/` directory contains the configuration files for Samba, the suite of programs that allows a Linux system to share files and printers with Windows systems.

### Key Files in `/etc/samba/`

1. **`smb.conf`**
2. **`secrets.tdb`**
3. **`lmhosts`**

### `smb.conf`

The primary configuration file for Samba. It controls the behavior of the `smbd` and `nmbd` services and defines shared resources such as directories and printers.

#### Key Sections of `smb.conf`

1. **Global Settings**
2. **Share Definitions**

#### Example `smb.conf`

```ini
[global]
   workgroup = WORKGROUP
   server string = Samba Server
   netbios name = hostname
   security = user
   map to guest = bad user
   dns proxy = no

[homes]
   comment = Home Directories
   browseable = no
   writable = yes

[public]
   path = /srv/samba/public
   public = yes
   only guest = yes
   writable = yes
   printable = no
```

#### Key Parameters in `smb.conf`

- **Global Settings**
  - `workgroup`: Defines the Windows workgroup or domain.
  - `server string`: Describes the server.
  - `netbios name`: The NetBIOS name of the server.
  - `security`: The security mode (`user`, `share`, etc.).
  - `map to guest`: How to handle guest users.
  - `dns proxy`: Whether to use DNS for NetBIOS name resolution.

- **Share Definitions**
  - `path`: The directory path to be shared.
  - `browseable`: Whether the share is visible in network browsers.
  - `writable`: Whether the share is writable.
  - `public`: Whether the share is accessible without authentication.
  - `printable`: Whether the share is a printer.

### `secrets.tdb`

This file stores sensitive information, such as user passwords and service credentials. It is maintained by Samba and should not be manually edited.

### `lmhosts`

This file provides a static table for mapping NetBIOS names to IP addresses, similar to `/etc/hosts` for DNS.

#### Example `lmhosts`

```plaintext
192.168.1.10    server1
192.168.1.20    server2
```

### Managing Samba Configuration

- **Testing the Configuration**

  ```bash
  testparm
  ```

  This command checks the syntax of `smb.conf` and ensures there are no errors.

- **Reloading the Configuration**

  ```bash
  sudo systemctl reload smbd
  sudo systemctl reload nmbd
  ```

  This reloads the configuration without restarting the services.

### Summary of `/etc/samba/`

| File            | Purpose                                                     |
|-----------------|-------------------------------------------------------------|
| `smb.conf`      | Main configuration file for Samba services                  |
| `secrets.tdb`   | Database file for storing sensitive information             |
| `lmhosts`       | Static table for NetBIOS name to IP address mapping         |

### Conclusion

The `/etc/samba/` directory is crucial for configuring and managing Samba, enabling file and printer sharing between Linux and Windows systems. Understanding the structure and key files in this directory allows for effective Samba configuration and maintenance.
