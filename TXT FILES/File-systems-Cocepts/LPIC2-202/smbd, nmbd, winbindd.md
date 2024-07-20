# SMB Daemons: smbd, nmbd, winbindd

The SMB protocol daemons, smbd, nmbd, and winbindd, are core components of the Samba suite, which provides file and print services for various Microsoft Windows clients. Here’s an overview of each daemon, their purposes, and configuration options:

#### smbd

`smbd` is the server daemon that provides file and print services to SMB/CIFS clients. It handles the file-sharing and printing capabilities of Samba.

- **Purpose**: 
  - Provides file sharing.
  - Provides print services.
  - Manages authentication and authorization.

- **Configuration**:
  - Configured in the `smb.conf` file, typically located in `/etc/samba/smb.conf`.

- **Key Options**:
  - **[global]** section: Configures global settings.
  - **[share]** sections: Configures individual shares.

- **Common Usage**:
  - Start: `sudo systemctl start smbd`
  - Stop: `sudo systemctl stop smbd`
  - Restart: `sudo systemctl restart smbd`

- **Example Configuration**:
  ```conf
  [global]
      workgroup = WORKGROUP
      server string = Samba Server
      security = user

  [share]
      path = /srv/samba/share
      writable = yes
      guest ok = no
  ```

#### nmbd

`nmbd` is the server daemon that provides NetBIOS name service to SMB/CIFS clients. It handles the network browsing and name resolution.

- **Purpose**:
  - Provides NetBIOS name service.
  - Facilitates network browsing.
  - Supports WINS (Windows Internet Name Service).

- **Configuration**:
  - Configured in the `smb.conf` file, often in conjunction with `smbd`.

- **Key Options**:
  - **[global]** section: Configures global settings related to network browsing and name resolution.

- **Common Usage**:
  - Start: `sudo systemctl start nmbd`
  - Stop: `sudo systemctl stop nmbd`
  - Restart: `sudo systemctl restart nmbd`

- **Example Configuration**:
  ```conf
  [global]
      workgroup = WORKGROUP
      server string = Samba Server
      security = user
      wins support = yes
  ```

#### winbindd

`winbindd` is the server daemon that provides a unified login for Windows and Unix accounts. It allows Unix systems to use Windows domain accounts for authentication.

- **Purpose**:
  - Integrates Unix systems with Windows NT-based systems.
  - Provides access to Windows domain users and groups.
  - Facilitates single sign-on (SSO) in mixed environments.

- **Configuration**:
  - Configured in the `smb.conf` file, especially in environments where Samba integrates with Windows domains.

- **Key Options**:
  - **[global]** section: Configures settings for domain membership and winbind integration.

- **Common Usage**:
  - Start: `sudo systemctl start winbind`
  - Stop: `sudo systemctl stop winbind`
  - Restart: `sudo systemctl restart winbind`

- **Example Configuration**:
  ```conf
  [global]
      workgroup = WORKGROUP
      security = ads
      realm = EXAMPLE.COM
      password server = adserver.example.com
      idmap config * : backend = tdb
      idmap config * : range = 1000-2000
      idmap config EXAMPLE : backend = rid
      idmap config EXAMPLE : range = 2001-100000
  ```

### Installation and Management

#### Installation

To install Samba, which includes `smbd`, `nmbd`, and `winbindd`, use your package manager. For example, on Debian-based systems:

```bash
sudo apt update
sudo apt install samba winbind
```

#### Management

Start, stop, and enable the services using `systemctl`:

```bash
# Start services
sudo systemctl start smbd
sudo systemctl start nmbd
sudo systemctl start winbind

# Stop services
sudo systemctl stop smbd
sudo systemctl stop nmbd
sudo systemctl stop winbind

# Enable services to start on boot
sudo systemctl enable smbd
sudo systemctl enable nmbd
sudo systemctl enable winbind

# Check status
sudo systemctl status smbd
sudo systemctl status nmbd
sudo systemctl status winbind
```

### Use Cases

- **File and Print Sharing**: `smbd` is essential for sharing files and printers with Windows clients.
- **Network Browsing and Name Resolution**: `nmbd` allows Windows clients to see the Samba server in their network neighborhood.
- **Domain Integration**: `winbindd` enables Unix systems to authenticate users against a Windows domain, facilitating SSO.

### Conclusion

Understanding and configuring `smbd`, `nmbd`, and `winbindd` is crucial for administrators managing a mixed network environment. Proper configuration in the `smb.conf` file ensures seamless integration and reliable service for both Unix and Windows clients.

