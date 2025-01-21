# Samba

**Samba** is a free and open-source implementation of the **SMB/CIFS** networking protocol, allowing Linux and Unix systems to interact with Windows clients for file and printer sharing. It enables Linux servers to offer resources like shared folders, printers, and other network services to Windows machines and vice versa, creating a seamless cross-platform networking experience.

### **Key Components of Samba**:
1. **smbd**: The server daemon that provides SMB/CIFS file-sharing and print services.
2. **nmbd**: Handles NetBIOS name service requests (name resolution) and network browsing.
3. **winbindd**: Allows Linux systems to authenticate against Windows NT/Active Directory domains.

### **Core Functions of Samba**:
1. **File and Printer Sharing**: Linux/Unix systems can share files and printers with Windows machines.
2. **Domain Integration**: Can act as a Primary Domain Controller (PDC), or integrate with Active Directory, providing authentication and authorization services.
3. **Cross-Platform Compatibility**: Allows for communication between Linux/Unix and Windows systems using the SMB/CIFS protocol.
4. **Network Browsing**: Samba can make Linux/Unix machines visible to Windows machines in network browsing (via `nmbd`).
5. **Active Directory Integration**: Samba can integrate with an Active Directory (AD) infrastructure, enabling single sign-on and centralized management.

### **Installing Samba**:
On most Linux distributions, Samba can be installed using the package manager:
```bash
sudo apt install samba          # On Ubuntu/Debian
sudo yum install samba samba-client  # On CentOS/RedHat
```

### **Key Configuration File**: `/etc/samba/smb.conf`
The main configuration file for Samba is `/etc/samba/smb.conf`. This file defines the behavior of Samba services, including shared resources, network settings, and authentication.

### **Common smb.conf Structure**:

```bash
[global]
   workgroup = WORKGROUP    # Windows workgroup or domain name
   server string = Samba Server %v
   security = user          # User-level security
   map to guest = Bad User  # Guests are treated as "bad users"

[homes]
   comment = Home Directories
   browseable = no          # Hide from network browsing
   writable = yes           # Allow users to write

[printers]
   comment = All Printers
   path = /var/spool/samba
   browseable = no
   printable = yes          # Mark share as a printer

[shared]
   comment = Shared Folder
   path = /srv/samba/shared
   writable = yes           # Allow writing
   browseable = yes         # Show share in network browsing
   guest ok = yes           # Allow guest access
```

### **Starting and Stopping the Samba Services**:
Samba services (`smbd` and `nmbd`) are typically managed using **systemd**:

```bash
sudo systemctl start smbd
sudo systemctl start nmbd

sudo systemctl enable smbd    # Start at boot
sudo systemctl enable nmbd

sudo systemctl status smbd    # Check the status of smbd
```

### **File Sharing with Samba**:
1. Define the shared directory in `/etc/samba/smb.conf`.
2. Set directory permissions (e.g., `chmod` and `chown`).
3. Restart Samba services: `sudo systemctl restart smbd nmbd`.
4. Set up Samba users: 
   ```bash
   sudo smbpasswd -a username
   ```

### **Testing Samba Configuration**:
After making changes to the Samba configuration file, you can validate it using:
```bash
testparm
```
This command checks the `smb.conf` file for any syntax errors or misconfigurations.

### **Samba Security Options**:
- **security = user**: Requires users to authenticate with valid credentials.
- **security = share**: Anyone can access shared resources without authentication (legacy and less secure).
- **map to guest = Bad User**: Users without valid credentials are mapped to a guest account, allowing limited access.

### **Active Directory Integration**:
Samba can be configured as an Active Directory Domain Controller (AD DC) or to join an existing AD domain. This allows for centralized management of users and permissions across both Linux and Windows systems.

**Joining Samba to an Active Directory Domain**:
```bash
sudo realm join ad.example.com -U Administrator
sudo systemctl enable winbind
sudo systemctl start winbind
```

### **Samba Logs**:
- Samba stores its log files in `/var/log/samba/`, which includes logs for `smbd` and `nmbd`.
- **Common log files**:
  - `log.smbd`: Logs for the `smbd` service.
  - `log.nmbd`: Logs for the `nmbd` service.
  - `log.winbindd`: Logs for the `winbindd` service (if applicable).

### **Samba in a Domain Controller Role**:
Samba can act as a **Primary Domain Controller (PDC)**, providing user authentication and network logon services for Windows clients in a network. It can also participate in domain logons and handle user profiles in a Windows domain.

### **Samba Integration with Windows**:
Once configured, Windows machines can access the shared folders by:
- Navigating to `\\<server-ip>` in the Windows file explorer.
- Accessing the shared folders, where users will be prompted for their Samba credentials (unless guest access is allowed).
