# LPIC-3 Mixed Environments 300 Objectives.md
## Topic 301: Samba Basics
### 301.1 Samba Concepts and Architecture * [weight: 2]
#### Description: Candidates should understand the essential concepts of Samba, including the various Samba server processes and networking protocols used by Samba when acting in various roles. Samba version 4.8 or higher is covered.

Key Knowledge Areas:

Understand the roles of the various Samba daemons and components
Understand key issues regarding heterogeneous networks
Understand the networking services used with SMB/CIFS and Active Directory, including their ports
Understand the major features of SMB protocol versions 1.0, 2.0, 2.1 and 3.0
Understand of Samba 3 and Samba 4 differences
Awareness of Samba VFS modules
Awareness of Samba Clustering and CTDB
Partial list of the used files, terms and utilities:

* [smbd]()

* [nmbd](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/nmbd.md)

* [samba](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba.md)

* [winbindd](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/winbindd.md)


### 301.2 Samba Configuration * [weight: 4]

#### Description: Candidates should be able to configure the Samba daemons.

Key Knowledge Areas:

Manage Samba server file-based configuration
Manage of Samba server registry-based configuration
Manage of Samba configuration parameters and variables
Understand Samba server roles and security modes
Configure Samba to use TLS
Check the validity of a Samba configuration
Troubleshoot and debug configuration problems with Samba
Understand Windows tools used to configure a Samba Server
The following is a partial list of the used files, terms and utilities:

* [smb.conf](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/smb.conf.md)

* [security](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/security.md)

* [server role](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/server-role.md)

* [server string](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/server-string.md)

* [server services](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/server-services.md)

* [tls enabled](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/tls-enabled.md)

* [tls keyfile](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/tls-keyfile.md)

* [tls certfile](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/tls-certfile.md)

* [tls dh params file](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/tls-dh-params-file.md)

* [tls cafile](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/tls-cafile.md)

* [config backend](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/config-backend.md)

* [registry shares](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/registry-shares.md)

* [include](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/include.md)

* [vfs objects](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/vfs-objects.md)

* [samba-regedit](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-regedit.md)

* [HKLM\Software\Samba](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/HKLM-Software-Samba.md)

* [REG_SZ, REG_MULTI_SZ](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/REG_SZ%2C%20REG_MULTI_SZ.md)

* [testparm](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/testparm.md)

* [net registry including relevant subcommands](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/net-registry-including-relevant-subcommands.md)

* [Microsoft RSAT Tools](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/Microsoft-RSAT-Tools.md)

* [Microsoft MMC](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/Microsoft-MMC.md)

* [Microsoft ADSI Edit](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/Microsoft-ADSI-Edit.md)

* [Microsoft LDP](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/Microsoft-LDP.md)

* [Microsoft Regedit](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/Microsoft-Regedit.md)

### 301.3 Regular Samba Maintenance * [weight: 2]

#### Description: Candidates should know the various tools and utilities that are part of a Samba installation.

Key Knowledge Areas:

Start and stop Samba services on domain controllers and file servers
Monitor and interact with running Samba daemons
Backup and restore TDB files
Backup and restore an Active Directory domain controller
Understand backup and recovery strategies for Active Directory domain controllers
Understand the impact of virtualization on Active Directory domain controllers
The following is a partial list of the used files, terms and utilities:

* [systemctl](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/systemctl.md)

* [smbcontrol (including relevant message types)](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/smbcontrol(including-relevant-message-types).md)

* [smbstatus](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/smbstatus.md)

* [tdbbackup](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/tdbbackup.md)

* [tdbrestore](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/tdbrestore.md)

* [samba-tool domain backup (including subcommands)](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-tool-domain-backup(including-subcommands).md)

* [Virtual Machine Generation Identifier](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/Virtual-Machine-Generation-Identifier.md)

* [Virtual Machine Snapshots](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/Virtual-Machine-Snapshots.md)


### 301.4 Troubleshooting Samba * [weight: 3]

#### Description: Candidates should be able to analyze and troubleshoot Samba issues. This includes accessing and modifying the LDAP content of a Samba server hosting an Active directory as well as working with trivial database files. Furthermore, candidates should be able to create a renamed clone of an existing Active Directory for debugging.

Key Knowledge Areas:

Configure Samba logging, including setting log levels for specific debug classes and client-specific logging
Query and modify the Samba password database
Understand the contents of important TDB files
List and edit TDB file content
Identify TDB file corruption
Access and modify objects in a Samba LDAP directory
Enable and use the LDAP recycle bin
Confirm the integrity of a domain controller’s database
Create a renamed clone of a domain controller
Awareness of Samba eventlog shipping
Use rpcclient to query information on a Samba server

The following is a partial list of the used files, terms and utilities:

* [smb.conf:](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/smb.conf.md)

* [log level](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/log-level.md)

* [debuglevel](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/debuglevel.md)

* [/var/log/samba/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/varlogsamba.md)

* [smbpasswd](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/smbpasswd.md)

* [pdbedit](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/pdbedit.md)

* [registry.tdb](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/registry.tdb.md)

* [secrets.tdb](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/secrets.tdb.md)

* [tdbdump](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/tdbdump.md)

* [tdbtool](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/tdbtool.md)

* [ldbsearch](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/ldbsearch.md)

* [ldbmodify](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/ldbmodify.md)

* [ldbedit](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/ldbedit.md)

* [ldbadd](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/ldbadd.md)

* [ldbdel](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/ldbdel.md)

* [LDIF](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/LDIF.md)

* [samba-tool dbcheck](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-tool-dbcheck.md)

* [samba-tool domain backup (including relevant subcommands)](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-tool-domain-backup(including-relevant-subcommands).md)

* [rpcclient](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/rpcclient.md)

## Topic 302: Samba and Active Directory Domains
### 302.1 Samba as Active Directory Domain Controller * [weight: 5]

#### Description: Candidates should be able to configure Samba as an Active Directory domain controller. This includes managing an Active Directory domain.

Key Knowledge Areas:

Understand the concepts of Active Directory
Understand the principles of the network services used by Active Directory (i.e. DNS, Kerberos, NTP and LDAP and CIFS and MS-RPC)
Set up a new Active Directory domain using Samba
Add a Samba domain controller to an existing Active Directory domain
Demote and remove online and offline domain controllers
Verify AD replication
Understand and query the global catalog and the partial attribute set
Understand and configure domain functional levels
Understand and configure Active Directory forest and domain trusts
Understand and configure Active Directory sites, including subnet assignments
Understand and manage FSMO roles, including their impact in case of an outage
Configure authentication audit logging
Configure SYSVOL replication using rsync or robocopy
Integrate Samba with ntpd
Awareness of Windows NT4 domains
The following is a partial list of the used files, terms and utilities:

* [smb.conf:](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/smb.conf.md)

* [server role](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/server-role.md)

* [log level](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/log-level.md)

* [samba-tool domain (including relevant subcommands)](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-tool-domain(including-relevant-subcommands).md)

* [samba-tool fsmo (including relevant subcommands)](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-tool-fsmo(including-relevant-subcommands).md)

* [samba-tool drs (including relevant subcommands)](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-tool-drs(including-relevant-subcommands).md)

* [samba-tool sites (including relevant subcommands)](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/rsync.conf.md)

* [rsync](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/rsync.md)

* [rsync.conf](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/rsync.conf.md)

* [/var/lib/samba/sysvol](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/varlibsambasysvol.md)

* [robocopy](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/robocopy.md)

* [ntpd.conf](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/ntpd.conf.md)

* [ntpsigndsocket](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/ntpsigndsocket.md)

### 302.2 Active Directory Name Resolution * [weight: 2]

#### Description: Candidates should be familiar with the internal DNS server of Samba.

Key Knowledge Areas:

Understand and manage DNS for Samba as an AD domain controller
Manage DNS records in Samba DNS
DNS forwarding
Standardized names in an Active Directory
Multicast DNS
Awareness of BIND9 DLZ DNS back end
Awareness of NetBIOS name resolution and WINS
The following is a partial list of the used files, terms and utilities:

* [smb.conf:](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/smb.conf.md)

* [dns forwarder](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/dns-forwarder.md)

* [allow dns updates](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/allow-dns-updates.md)

* [multicst dns register](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/multicst-dns-register.md)

* [samba-tool dns (with subcommands)](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-tool-dns-(with-subcommands).md)

* [samba_dnsupdate](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba_dnsupdate.md)

* [dig](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/dig.md)

* [host](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/host.md)

* [/etc/resolv.conf](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-300/etcresolv.conf.md)


### 302.3 Active Directory User Management * [weight: 4]

#### Description: Candidates should be able to manage user and group accounts on a standalone server and in a Samba based Active Directory.

Key Knowledge Areas:

Manage user accounts and user group for standalone servers and Samba AD
Knowledge of user account management tools
Delegate administrative permissions in AD to specific users / user groups
Configure password expiration and change requirements
Manage password policies and password setting objects
Understand principals and their identification SID (DN, GUID)
Understand User Principal Name (UPN) and User Principal Name Suffix (UPN Suffix)
Understand and manage Security and Distribution Groups
Understand and manage LDAP attributes of security principals
Understand and manage RFC2307 attributes in a Samba AD
Map Kerberos service principal names to user accounts
Export Kerberos keytabs for a specific principal
Awareness of LDAP Account Manager
The following is a partial list of the used files, terms and utilities:

* [samba-tool user (including relevant subcommands)]()

* [samba-tool group (including relevant subcommands)]()

* [samba-tool domain passwordsettings]()

* [samba-tool domain exportkeytab]()

* [samba-tool spn (including relevant subcommands)]()

* [smbpasswd]()

* [pdbedit]()

* [kinit]()

* [klist]()


### 302.4 Samba Domain Membership * [weight: 4]

#### Description: Candidates should be able to join a Samba server into an existing Active Directory domain and authorize domain users to use the server. This includes installing and configuring the Winbind service.

Key Knowledge Areas:

Join Samba to an existing AD domain
Configure Winbind service, including ID mapping
Understand and configure Winbind ID mapping, including various mapping backends
Configure PAM and NSS to use Winbind
The following is a partial list of the used files, terms and utilities:

* [smb.conf:]()

* [security]()

* [server role]()

* [realm]()

* [workgroup]()

* [idmap config]()

* [winbind enumerate users]()

* [winbind enumerate groups]()

* [winbind offline logon]()

* [winbind separator]()

* [template shell]()

* [template homedir]()

* [allow trusted domains]()

* [idmap_ad]()

* [idmap_autorid]()

* [idmap_ldap]()

* [idmap_rfc2307]()

* [idmap_rid]()

* [idmap_tdb]()

* [idmap_tdb2]()

* [net ads (including relevant subcommands)]()

* [/etc/nsswitch.conf]()

* [/etc/pam.conf]()

* [/etc/pam.d/]()

* [libnss_winbind]()

* [libpam_winbind]()

* [getent]()

* [wbinfo]()


### 302.5 Samba Local User Management * [weight: 2]

#### Description: Candidates should be able to create and manage local user accounts on a stand alone Samba server.

Key Knowledge Areas:

Setup a local password database
Perform password synchronization
Knowledge of different passdb backends
Convert between Samba passdb backends
The following is a partial list of the used files, terms and utilities:

* [smb.conf:]()

* [passdb backend]()

* [/etc/passwd]()

* [/etc/group]()

* [pam_smbpass.so]()

* [smbpasswd]()

* [pdbedit]()


## Topic 303: Samba Share Configuration
### 303.1 File Share Configuration * [weight: 4]

#### Description: Candidates should be able to create and configure CIFS file shares in Samba.

Key Knowledge Areas:

Create and configure CIFS file shares
Manage Samba share access configuration parameters
Use registry based share configuration
Manage profile and user home shares
Plan file service migration
Limit access to IPC$
Awareness of user shares
Awareness of existing VFS modules and their general functionality, including modules to support audit logs and snapshots / shadow copies
The following is a partial list of the used files, terms and utilities:

* [smb.conf:]()

* [path]()

* [browsable]()

* [writable / write ok / read only]()

* [valid users]()

* [invalid users]()

* [read list]()

* [write list]()

* [guest ok]()

* [hosts allow / allow hosts]()

* [hosts deny / deny hosts]()

* [copy]()

* [hide unreadable]()

* [hide unwritable files]()

* [hide dot files]()

* [hide special files]()

* [veto files]()

* [delete veto files]()

* [(homes)]()

* [(IPC$)]()

* [smbcquotas]()


### 303.2 File Share Security * [weight: 3]

#### Description: Candidates should understand file permissions on CIFS shares and on a Linux file system.

Key Knowledge Areas:

Enforce ownership and permissions of files and directories
Manage ACLs for shares and folders
Understand POSIX, Extended POSIX and Windows ACLs
Understand how Samba stores Windows ACLs in Linux ACLs and extended attributes
Configure ACLs for profile and home folder shares
Configure encryption of CIFS connections
The following is a partial list of the used files, terms and utilities:

* [smb.conf]()

* [create mask / create mode]()

* [directory mask / directory mode]()

* [force create mode]()

* [force directory mode]()

* [force user]()

* [force group / group]()

* [profile acls]()

* [inherit acls]()

* [map acl inherit]()

* [store dos attributes]()

* [vfs objects]()

* [smb encrypt]()

* [chown]()

* [chmod]()

* [getfacl]()

* [setfacl]()

* [getfattr]()

* [smbcacls]()

* [sharesec]()

* [SeDiskOperatorPrivilege]()

* [vfs_acl_xattr]()

* [vfs_acl_tdb]()

* [samba-tool ntacl (including subcommands)]()



### 303.3 DFS Share Configuration * [weight: 1]

#### Description: Candidates should be able to create and manage DFS shares in Samba.

Key Knowledge Areas:

Understand DFS
Configure DFS shares
The following is a partial list of the used files, terms and utilities:

* [smb.conf:]()

* [host msdfs]()

* [msdfs root]()

* [msdfs proxy]()

* [ln]()

### 303.4 Print Share Configuration * [weight: 2]

#### Description: Candidates should be able to create and manage print shares in Samba.

Key Knowledge Areas:

Understand Samba printing, including raw printing
Create and configure print shares
Configure integration between Samba and CUPS
Manage Windows print drivers and configure downloading of print drivers
Upload printer drivers using 'Add Print Driver Wizard' in Windows
Preconfigure driver settings
Configure paper sizes and forms
Supported driver versions
Manage GPO options for trusted print servers
Awareness of spoolssd
The following is a partial list of the used files, terms and utilities:

* [smb.conf:]()

* [printing]()

* [printable / print ok]()

* [printcap name / printcap]()

* [spoolss: architecture = Windows x64]()

* [(printers)]()

* [(print$)]()

* [CUPS]()

* [cupsd.conf]()

* [/var/spool/samba/]()

* [smbspool]()

* [rpcclient (to execute topic-related commands (enumdrivers, enumprinters, setdriver))]()

* [net (included topic-related subcommands)]()

* [SePrintOperatorPrivilege]()


## Topic 304: Samba Client Configuration
### 304.1 Linux Authentication Clients * [weight: 5]

#### Description: Candidates should be familiar with management and authentication of user accounts. This includes configuration and use of NSS, PAM, SSSD and Kerberos for both local and remote directories and authentication mechanisms as well as enforcing a password policy.

Key Knowledge Areas:

Understand and configure NSS and PAM
Enforce password complexity policies and periodic password changes
Create home directories for new users
Lock accounts automatically after failed login attempts
Configure NSS and PAM to retrieve information from LDAP
Configure SSSD authentication against Active Directory, IPA, LDAP and Kerberos domains and the local system’s authentication database
Manage local accounts through SSSD
Obtain and manage Kerberos tickets
The following is a partial list of the used files, terms and utilities:

* [/etc/pam.conf]()

* [/etc/pam.d/]()

* [/etc/nsswitch.conf]()

* [/etc/login.defs]()

* [pam_ldap.so]()

* [ldap.conf]()

* [pam_krb5.so]()

* [pam_cracklib.so]()

* [pam_tally2.so]()

* [pam_faillock.so]()

* [pam_mkhomedir.so]()

* [chage]()

* [faillog]()

* [sssd]()

* [sssd.conf]()

* [sss_override]()

* [sss_cache]()

* [sss_debuglevel]()

* [sss_user* and sss_group*]()

* [/var/lib/sss/db/]()

* [krb5.conf]()

* [kinit]()

* [klist]()

* [kdestroy]()



### 304.2 Linux CIFS Clients * [weight: 3]

#### Description: Candidates should be able to use remote CIFS shares from a Linux client. This includes client-side management of CIFS credentials and managing remote ACLs and quotas.

Key Knowledge Areas:

Access remote CIFS shares from a Linux client
Mount remote CIFS shares on a Linux client
Automatically mount home directories
Store and manage CIFS credentials securely
Understand and manage permissions and file ownership of remote CIFS shares
Understand and manage quotas on CIFS shares
The following is a partial list of the used files, terms and utilities:

* [smb.conf]()

* [smbclient (including relevant subcommands)]()

* [mount]()

* [mount.cifs]()

* [/etc/fstab]()

* [pam_mount.so]()

* [pam_mount.conf.xml]()

* [cifscreds]()

* [getcifsacl]()

* [setcifsacl]()

* [smbcquotas]()

* [cifsiostat]()

* [smbget]()

* [smbtar]()

### 304.3 Windows Clients * [weight: 3]

#### Description: Candidates should be able to access CIFS and print shares from Windows hosts and join such hosts into an Active Directory domain. Furthermore, candidates should be able to manage Windows hosts using GPOs and access remote Windows hosts.

Key Knowledge Areas:

Understand how to set up and use Windows hosts
Join a Windows host to an Active Directory domain
Access remote CIFS shares from a Windows client
Configure printing to remote printers from a Windows client
Configure file and print shares on a Windows host
Understand the concept, structure and capabilities of GPOs
Create and modify GPOs and apply GPOs to machines or users
Access a remote Windows desktop
Create and configure logon scripts
Configure roaming profiles for Active Directory users
Configure profile folder redirects
The following is a partial list of the used files, terms and utilities:

* [smb.conf:]()

* [logon path]()

* [logon script]()

* [net (Windows command; including all relevant subcommands)]()

* [samba-tool gpo (including all relevant subcommands)]()

* [gpupdate (Windows command)]()

* [rdesktop]()

## Topic 305: Linux Identity Management and File Sharing
### 305.1 FreeIPA Installation and Maintenance * [weight: 2]

#### Description: Candidates should be able to set up and manage a FreeIPA domain using standard settings and default services. This includes setting up replication and joining clients to the domain.

Key Knowledge Areas:

Understand the features, architecture as well as server-side and client-side components of FreeIPA
Install a FreeIPA server
Set up and manage a FreeIPA domain using standard settings and default services
Understand replication topology and configure FreeIPA replication
Join clients to an existing FreeIPA domain
Awareness of ipa-backup
The following is a partial list of the used files, terms and utilities:

* [ipa-server-install]()

* [ipa-replica-prepare]()

* [ipa-replica-install]()

* [ipa-client-install]()

* [ipactl]()

### 305.2 FreeIPA Entity Management * [weight: 4]

#### Description: Candidates should be able manage users, hosts and services in a FreeIPA domain.

Key Knowledge Areas:

Manage user accounts and groups
Manage hosts, hostgroups and services
Understand the principle of IPA access control permissions, privileges and roles
Understand ID views
Awareness of sudo, autofs, SSH, SELinux and NIS integration as well as host based access control in FreeIPA
Awareness of the FreeIPA CA
The following is a partial list of the used files, terms and utilities:

* [ipa (including relevant user-*, stageuser-* and group-* and idview-* subcommands)]()

* [ipa (including relevant host-*, hostgroup-*, service-* and getkeytab subcommands)]()

* [ipa (including relevant permission-*, privilege-*, and role-* subcommands)]()

* [ipctl]()

* [ipa-advice]()



### 305.3 FreeIPA Active Directory Integration * [weight: 2]

#### Description: Candidates should be able to set up a cross-forest trust between a FreeIPA and an Active Directory domain.

Key Knowledge Areas:

Understand and set up FreeIPA and Active Directory integration using Kerberos cross-realm trusts
Configure ID ranges in FreeIPA
Understand and manage external non-POSIX groups in FreeIPA
Awareness of Microsoft Privilege Attribute Certificates and how they are handled by FreeIPA
Awareness of replication based FreeIPA and Active Directory integration
The following is a partial list of the used files, terms and utilities:

* [ipa-adtrust-install]()
* [ipa (including relevant trust-*, idrange-* and group-* subcommands)]()


### 305.4 Network File System * [weight: 3]

#### Description: Candidates should be able to use NFSv4. This includes understanding ID mapping, NFSv4 ACLs and Kerberos authentication for NFS.

Key Knowledge Areas:

Understand major NFSv4 features
Configure and manage an NFSv4 server and clients
Understand and use the NFSv4 pseudo file system
Understand and use NFSv4 ACLs
Use Kerberos for for NFSv4 authentication
The following is a partial list of the used files, terms and utilities:

* [exportfs]()

* [/etc/exports]()

* [/etc/idmapd.conf]()

* [nfs4_editfacl]()

* [nfs4_getfacl]()

* [nfs4_setfacl]()

* [mount (including common NFS mount options)]()

* [/etc/fstab]()
