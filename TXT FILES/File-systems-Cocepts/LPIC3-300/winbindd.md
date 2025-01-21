# winbindd

**winbindd** is a component of Samba that allows a Linux or Unix system to communicate with Windows domain controllers (Active Directory or NT domains). It enables Unix-based systems to participate in a Windows-based network by performing tasks such as:

- **User and Group Authentication**: Allows Linux systems to authenticate users from Windows domains.
- **SID to UID/GID Mapping**: Maps Windows security identifiers (SIDs) to Unix user IDs (UIDs) and group IDs (GIDs).
- **Single Sign-On (SSO)**: Integrates Unix systems into a Windows Active Directory (AD) domain for single sign-on capabilities.
- **Name Resolution**: Resolves Windows domain user and group names to Unix counterparts.

### **Core Functions of winbindd**:
1. **Domain Integration**: Integrates a Linux machine into a Windows domain (Active Directory or NT-style domains).
2. **User and Group Management**: Manages domain user and group accounts, allowing domain users to log in to the Linux system using their AD credentials.
3. **Authentication**: Authenticates users against a Windows Domain Controller (DC), supporting both NTLM and Kerberos authentication protocols.
4. **User and Group Enumeration**: Lists users and groups from a Windows domain on the Linux system.
5. **SSO in Mixed Environments**: Provides seamless login capabilities across Windows and Linux systems when integrated with Kerberos.

### **Installation of winbindd**:
To install `winbindd` on your Linux system, use the following package managers:

```bash
# On Ubuntu/Debian:
sudo apt install winbind samba libpam-winbind libnss-winbind

# On CentOS/RedHat:
sudo yum install samba-winbind samba-winbind-clients
```

### **Configuration of winbindd**:

To configure `winbindd`, you'll need to edit several key files:

1. **/etc/samba/smb.conf**: Configures the integration of the Linux system with the Windows domain.
2. **/etc/nsswitch.conf**: Ensures that domain users and groups are available for login.
3. **/etc/krb5.conf** (if using Kerberos): Configures Kerberos for authentication.

**Example smb.conf for winbindd Integration:**

```bash
[global]
   workgroup = EXAMPLE
   security = ads
   realm = EXAMPLE.COM
   dedicated keytab file = /etc/krb5.keytab
   kerberos method = secrets and keytab
   winbind use default domain = true
   winbind enum users = yes
   winbind enum groups = yes
   template shell = /bin/bash
   template homedir = /home/%U
   idmap config * : backend = tdb
   idmap config * : range = 10000-20000
```

### **Step-by-Step Setup for winbindd with Active Directory**:

1. **Join Linux System to AD Domain**:

```bash
sudo realm join --user=Administrator ad.example.com
```

2. **Enable winbindd and related services**:

```bash
sudo systemctl enable winbind
sudo systemctl start winbind
```

3. **Modify `/etc/nsswitch.conf`**:
Ensure the following entries are present to enable winbind for user and group lookups:

```bash
passwd:    compat winbind
group:     compat winbind
```

4. **Testing the Domain Join**:

Check if the system has successfully joined the domain by running:

```bash
wbinfo -u   # List domain users
wbinfo -g   # List domain groups
```

You can also verify that users can be resolved correctly:

```bash
getent passwd 'EXAMPLE\username'
```

### **Mapping Windows Users and Groups to Unix**:
`winbindd` maps Windows users and groups to Unix IDs, which are required for file ownership and permission settings on Linux systems.

- **UID/GID Mapping**: `winbindd` maps Windows Security Identifiers (SIDs) to Unix UIDs/GIDs using a mapping backend.
- **ID Mapping Backends**:
  - **tdb**: Local TDB database for mapping.
  - **ad**: Retrieves Unix attributes (UID/GID) from AD.
  - **rid**: Derives UIDs/GIDs based on the user's RID.

**Example smb.conf configuration for mapping:**

```bash
idmap config * : backend = tdb
idmap config * : range = 10000-20000
idmap config EXAMPLE : backend = rid
idmap config EXAMPLE : range = 10000-30000
```

### **Enabling Kerberos for SSO**:
To fully integrate into a Windows Active Directory domain with SSO, you need to configure Kerberos:

1. **Edit /etc/krb5.conf** to include your domain’s Kerberos details:

```bash
[libdefaults]
   default_realm = EXAMPLE.COM
   dns_lookup_realm = false
   dns_lookup_kdc = true

[realms]
   EXAMPLE.COM = {
       kdc = ad.example.com
       admin_server = ad.example.com
   }

[domain_realm]
   .example.com = EXAMPLE.COM
   example.com = EXAMPLE.COM
```

2. **Obtain a Kerberos ticket** using the `kinit` command:

```bash
kinit username@EXAMPLE.COM
```

3. **Verify ticket status**:

```bash
klist
```

### **winbindd Logs and Troubleshooting**:
The logs for `winbindd` can be found in the `/var/log/samba/` directory. Common log files include:

- **log.winbindd**: Contains logs for the winbind daemon.
- **log.wb-***: Per-domain log files for winbindd interactions with specific domains.

If users cannot log in or there are problems with domain integration, check these logs for errors.

### **Common Issues and Fixes**:
1. **User Cannot Authenticate**:
   - Ensure that the system has successfully joined the domain.
   - Check Kerberos ticket status (`klist`).
   - Verify `/etc/nsswitch.conf` is correctly configured.

2. **UID/GID Mismatch**:
   - Adjust the ID mapping configuration in `smb.conf`.
   - Use the `wbinfo` command to test UID/GID resolution for domain users and groups.

3. **winbindd Fails to Start**:
   - Check for configuration errors in `smb.conf` using `testparm`.
   - Look for logs in `/var/log/samba/log.winbindd`.


