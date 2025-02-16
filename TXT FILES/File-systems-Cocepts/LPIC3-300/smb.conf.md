# smb.conf

The **smb.conf** file is the main configuration file for the Samba suite. It controls how Samba interacts with clients, defines shares, user authentication, and integrates with Windows domains or Active Directory environments. Below is an explanation of the key components and directives of the `smb.conf` file, along with an example configuration.

### **Structure of smb.conf**
The `smb.conf` file is divided into sections, each enclosed in square brackets (`[section name]`). The most important sections are:

- **[global]**: Contains global settings that apply to the entire Samba server.
- **[share name]**: Defines a specific share that clients can access.
- **[homes]**: Special section for home directories.

### **Key Directives in the [global] Section**
The `[global]` section holds settings for the overall server behavior, including network configuration, authentication methods, and performance tuning.

1. **Basic Configuration**:
   - **workgroup**: Defines the Windows workgroup or domain the server belongs to.
   - **server string**: Describes the Samba server in network browsing.
   - **netbios name**: The NetBIOS name for the Samba server.

   ```bash
   [global]
       workgroup = MYWORKGROUP
       server string = Samba Server
       netbios name = SAMBASERVER
   ```

2. **Security Settings**:
   - **security**: Specifies the type of security to use. Common values:
     - `user`: Authenticates users based on a Samba password.
     - `ads`: Active Directory security, for joining Samba to an AD domain.
     - `share`: All users have access to shared resources without authentication.
   
   ```bash
   security = user
   ```

   - **encrypt passwords**: Enables password encryption.
   - **passdb backend**: Defines the backend for storing user accounts, such as `tdbsam` or `ldapsam`.

3. **Domain and Active Directory Integration** (for Windows AD environments):
   - **realm**: Defines the Kerberos realm for Active Directory (used with `security = ads`).
   - **kerberos method**: Specifies the Kerberos configuration (e.g., `secrets and keytab`).
   - **winbind use default domain**: Allows winbind to use the domain without prepending it to usernames.
   - **idmap config**: Configures the ID mapping for users and groups between AD and Linux.

   ```bash
   realm = EXAMPLE.COM
   security = ads
   winbind use default domain = yes
   idmap config * : backend = tdb
   idmap config * : range = 10000-20000
   ```

4. **Networking and Access Control**:
   - **interfaces**: Defines the network interfaces that Samba will listen on.
   - **bind interfaces only**: Ensures that Samba only listens on the specified interfaces.
   - **hosts allow**: Specifies which hosts are allowed to connect.

   ```bash
   interfaces = lo eth0
   bind interfaces only = yes
   hosts allow = 192.168.1. 127.
   ```

5. **Logging**:
   - **log file**: Path to the Samba log file.
   - **log level**: Controls verbosity of log messages (higher numbers mean more detailed logs).
   
   ```bash
   log file = /var/log/samba/log.%m
   log level = 1
   ```

### **Share Definitions**

Each share (whether a directory or a printer) is defined in its own section. Here are some key options for share definitions:

1. **Basic Share Options**:
   - **path**: The directory path for the share.
   - **comment**: A description of the share.
   - **read only**: Specifies if the share is read-only (`yes` or `no`).
   - **browseable**: Controls whether the share is visible when browsing the network.

   ```bash
   [myshare]
       path = /srv/samba/myshare
       comment = My File Share
       read only = no
       browseable = yes
   ```

2. **Permissions and Access Control**:
   - **valid users**: Specifies which users or groups can access the share.
   - **write list**: Specifies users or groups allowed to write to the share.
   - **force user/force group**: Forces access to be executed under a specific user or group.
   - **create mask**: Defines the permissions for newly created files.

   ```bash
   [myshare]
       path = /srv/samba/myshare
       valid users = @users
       write list = @users
       force group = users
       create mask = 0660
       directory mask = 0770
   ```

3. **Special Shares**:
   - **[homes]**: A special share for user home directories. When accessed, Samba automatically serves the home directory of the authenticated user.
   - **[printers]**: A section for managing printers shared via Samba.

   ```bash
   [homes]
       comment = Home Directories
       browseable = no
       writable = yes

   [printers]
       comment = All Printers
       path = /var/spool/samba
       printable = yes
       browseable = no
   ```

### **Example smb.conf File**

Below is a sample `smb.conf` file for a typical file server setup that integrates with an Active Directory domain and defines a few common shares:

```bash
[global]
    workgroup = EXAMPLE
    security = ads
    realm = EXAMPLE.COM
    netbios name = SAMBASERVER
    server string = Samba File Server
    log file = /var/log/samba/log.%m
    log level = 1
    max log size = 1000
    winbind use default domain = yes
    winbind enum users = yes
    winbind enum groups = yes
    idmap config * : backend = tdb
    idmap config * : range = 10000-20000
    template shell = /bin/bash

    interfaces = lo eth0
    bind interfaces only = yes
    hosts allow = 192.168.1. 127.
    encrypt passwords = yes

[homes]
    comment = Home Directories
    browseable = no
    writable = yes

[public]
    path = /srv/samba/public
    public = yes
    writable = no
    browseable = yes
    guest ok = yes

[secure]
    path = /srv/samba/secure
    valid users = @securegroup
    write list = @securegroup
    force group = securegroup
    create mask = 0660
    directory mask = 0770
    browseable = yes
    writable = yes
```

### **Testing smb.conf Configuration**
Once you have configured the `smb.conf` file, always test it for syntax errors using the following command:

```bash
testparm
```

This will help verify if the configuration is correct and identify any mistakes.

### **Restart Samba Services**
After making changes to `smb.conf`, restart the Samba services to apply the new configuration:

```bash
# On systemd-based systems:
sudo systemctl restart smbd nmbd winbind
```
