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

* [nmbd](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/nmbd.md)

* [samba](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba.md)

* [winbindd](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/winbindd.md)


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

* [smb.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smb.conf.md)

* [security](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/security.md)

* [server role](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/server-role.md)

* [server string](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/server-string.md)

* [server services](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/server-services.md)

* [tls enabled](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/tls-enabled.md)

* [tls keyfile](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/tls-keyfile.md)

* [tls certfile](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/tls-certfile.md)

* [tls dh params file](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/tls-dh-params-file.md)

* [tls cafile](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/tls-cafile.md)

* [config backend](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/config-backend.md)

* [registry shares](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/registry-shares.md)

* [include](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/include.md)

* [vfs objects](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/vfs-objects.md)

* [samba-regedit](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-regedit.md)

* [HKLM\Software\Samba](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/HKLM-Software-Samba.md)

* [REG_SZ, REG_MULTI_SZ](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/REG_SZ%2C%20REG_MULTI_SZ.md)

* [testparm](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/testparm.md)

* [net registry including relevant subcommands](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/net-registry-including-relevant-subcommands.md)

* [Microsoft RSAT Tools](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/Microsoft-RSAT-Tools.md)

* [Microsoft MMC](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/Microsoft-MMC.md)

* [Microsoft ADSI Edit](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/Microsoft-ADSI-Edit.md)

* [Microsoft LDP](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/Microsoft-LDP.md)

* [Microsoft Regedit](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/Microsoft-Regedit.md)

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

* [systemctl](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/systemctl.md)

* [smbcontrol (including relevant message types)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smbcontrol(including-relevant-message-types).md)

* [smbstatus](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smbstatus.md)

* [tdbbackup](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/tdbbackup.md)

* [tdbrestore](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/tdbrestore.md)

* [samba-tool domain backup (including subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-tool-domain-backup(including-subcommands).md)

* [Virtual Machine Generation Identifier](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/Virtual-Machine-Generation-Identifier.md)

* [Virtual Machine Snapshots](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/Virtual-Machine-Snapshots.md)


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

* [smb.conf:](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smb.conf.md)

* [log level](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/log-level.md)

* [debuglevel](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/debuglevel.md)

* [/var/log/samba/](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/varlogsamba.md)

* [smbpasswd](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smbpasswd.md)

* [pdbedit](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/pdbedit.md)

* [registry.tdb](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/registry.tdb.md)

* [secrets.tdb](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/secrets.tdb.md)

* [tdbdump](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/tdbdump.md)

* [tdbtool](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/tdbtool.md)

* [ldbsearch](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/ldbsearch.md)

* [ldbmodify](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/ldbmodify.md)

* [ldbedit](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/ldbedit.md)

* [ldbadd](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/ldbadd.md)

* [ldbdel](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/ldbdel.md)

* [LDIF](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/LDIF.md)

* [samba-tool dbcheck](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-tool-dbcheck.md)

* [samba-tool domain backup (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-tool-domain-backup(including-relevant-subcommands).md)

* [rpcclient](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/rpcclient.md)

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

* [smb.conf:](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smb.conf.md)

* [server role](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/server-role.md)

* [log level](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/log-level.md)

* [samba-tool domain (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-tool-domain(including-relevant-subcommands).md)

* [samba-tool fsmo (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-tool-fsmo(including-relevant-subcommands).md)

* [samba-tool drs (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-tool-drs(including-relevant-subcommands).md)

* [samba-tool sites (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/rsync.conf.md)

* [rsync](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/rsync.md)

* [rsync.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/rsync.conf.md)

* [/var/lib/samba/sysvol](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/varlibsambasysvol.md)

* [robocopy](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/robocopy.md)

* [ntpd.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/ntpd.conf.md)

* [ntpsigndsocket](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/ntpsigndsocket.md)

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

* [smb.conf:](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smb.conf.md)

* [dns forwarder](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/dns-forwarder.md)

* [allow dns updates](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/allow-dns-updates.md)

* [multicst dns register](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/multicst-dns-register.md)

* [samba-tool dns (with subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-tool-dns-(with-subcommands).md)

* [samba_dnsupdate](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba_dnsupdate.md)

* [dig](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/dig.md)

* [host](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/host.md)

* [/etc/resolv.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/etcresolv.conf.md)


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

* [samba-tool user (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-tool-user-(including-relevant-subcommands).md)

* [samba-tool group (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-tool-group(including-relevant-subcommands).md)

* [samba-tool domain passwordsettings](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-tool-domain-passwordsettings.md)

* [samba-tool domain exportkeytab](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-tool-domain-exportkeytab.md)

* [samba-tool spn (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-tool-spn-(including-relevant-subcommands).md)

* [smbpasswd](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smbpasswd.md)

* [pdbedit](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/pdbedit.md)

* [kinit](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/kinit.md)

* [klist](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/klist.md)


### 302.4 Samba Domain Membership * [weight: 4]

#### Description: Candidates should be able to join a Samba server into an existing Active Directory domain and authorize domain users to use the server. This includes installing and configuring the Winbind service.

Key Knowledge Areas:

Join Samba to an existing AD domain
Configure Winbind service, including ID mapping
Understand and configure Winbind ID mapping, including various mapping backends
Configure PAM and NSS to use Winbind
The following is a partial list of the used files, terms and utilities:

* [smb.conf:](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smb.conf.md)

* [security](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/security.md)

* [server role](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/server-role.md)

* [realm](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/realm.md)

* [workgroup](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/workgroup.md)

* [idmap config](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/idmap-config.md)

* [winbind enumerate users](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/winbind-enumerate-users.md)

* [winbind enumerate groups](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/winbind-enumerate-groups.md)

* [winbind offline logon](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/winbind-offline-logon.md)

* [winbind separator](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/winbind-separator.md)

* [template shell](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/template-shell.md)

* [template homedir](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/template-homedir.md)

* [allow trusted domains](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/allow-trusted-domains.md)

* [idmap_ad](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/idmap_ad.md)

* [idmap_autorid](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/idmap_autorid.md)

* [idmap_ldap](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/idmap_ldap.md)

* [idmap_rfc2307](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/idmap_rfc2307.md)

* [idmap_rid](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/idmap_rid.md)

* [idmap_tdb](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/idmap_tdb.md)

* [idmap_tdb2]()

* [net ads (including relevant subcommands)]()

* [/etc/nsswitch.conf]()

* [/etc/pam.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/etcpam.conf.md)

* [/etc/pam.d/](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/etcpam.d.md)

* [libnss_winbind](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/libnss_winbind.md)

* [libpam_winbind](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/libpam_winbind.md)

* [getent](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/getent.md)

* [wbinfo](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/wbinfo.md)


### 302.5 Samba Local User Management * [weight: 2]

#### Description: Candidates should be able to create and manage local user accounts on a stand alone Samba server.

Key Knowledge Areas:

Setup a local password database
Perform password synchronization
Knowledge of different passdb backends
Convert between Samba passdb backends
The following is a partial list of the used files, terms and utilities:

* [smb.conf:](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smb.conf.md)

* [passdb backend](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/passdb-backend.md)

* [/etc/passwd](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/etcpasswd.md)

* [/etc/group]()

* [pam_smbpass.so](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/pam_smbpass.so.md)

* [smbpasswd](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smbpasswd.md)

* [pdbedit](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/pdbedit.md)


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

* [smb.conf:](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smb.conf.md)

* [path](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/path.md)

* [browsable](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/browsable.md)

* [writable / write ok / read only](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/writable-write-ok-read-only.md)

* [valid users](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/valid%20users.md)

* [invalid users](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/invalid-users.md)

* [read list](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/read-list.md)

* [write list](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/write-list.md)

* [guest ok](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/guest-ok.md)

* [hosts allow / allow hosts](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/hosts-allow-allow-hosts.md)

* [hosts deny / deny hosts](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/hosts-deny-deny-hosts.md)

* [copy](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/copy.md)

* [hide unreadable](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/hide-unreadable.md)

* [hide unwritable files](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/hide-unwritable-files.md)

* [hide dot files](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/hide-dot-files.md)

* [hide special files](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/hide-special-files.md)

* [veto files](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/veto-files.md)

* [delete veto files](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/delete-veto-files.md)

* [(homes)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/(homes).md)

* [(IPC$)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/(IPC%24).md)

* [smbcquotas](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smbcquotas.md)


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

* [smb.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smb.conf.md)

* [create mask / create mode](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/create-mask-create-mode.md)

* [directory mask / directory mode](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/directory-mask-directory-mode.md)

* [force create mode](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/force-create-mode.md)

* [force directory mode](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/force-directory-mode.md)

* [force user](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/force-user.md)

* [force group / group](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/force-group-group.md)

* [profile acls](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/profile-acls.md)

* [inherit acls](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/inherit-acls.md)

* [map acl inherit](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/map-acl-inherit.md)

* [store dos attributes](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/store-dos-attributes.md)

* [vfs objects](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/vfs-objects.md)

* [smb encrypt](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smb-encrypt.md)

* [chown](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/chown.md)

* [chmod](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/chmod.md)

* [getfacl](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/getfacl.md)

* [setfacl](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/setfacl.md)

* [getfattr](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/getfattr.md)

* [smbcacls](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smbcacls.md)

* [sharesec](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/sharesec.md)

* [SeDiskOperatorPrivilege](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/SeDiskOperatorPrivilege.md)

* [vfs_acl_xattr](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/vfs_acl_xattr.md)

* [vfs_acl_tdb](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/vfs_acl_tdb.md)

* [samba-tool ntacl (including subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-tool-ntacl(including-subcommands).md)



### 303.3 DFS Share Configuration * [weight: 1]

#### Description: Candidates should be able to create and manage DFS shares in Samba.

Key Knowledge Areas:

Understand DFS
Configure DFS shares
The following is a partial list of the used files, terms and utilities:

* [smb.conf:](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smb.conf.md)

* [host msdfs](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/host%20msdfs.md)

* [msdfs root](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/msdfs-root.md)

* [msdfs proxy](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/msdfs-proxy.md)

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

* [smb.conf:](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smb.conf.md)

* [printing](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/printing.md)

* [printable / print ok](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/printable-print-ok.md)

* [printcap name / printcap](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/printcap-name-printcap.md)

* [spoolss: architecture = Windows x64](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/spoolss-architecture%3DWindowsx64.md)

* [(printers)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/(printers).md)

* [(print$)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/(print%24).md)

* [CUPS](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/CUPS.md)

* [cupsd.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/cupsd.conf.md)

* [/var/spool/samba/](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/varspoolsamba.md)

* [smbspool](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smbspool.md)

* [rpcclient (to execute topic-related commands (enumdrivers, enumprinters, setdriver))](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/rpcclient(to-execute-topic-related-commands(enumdrivers%2Cenumprinters%2Csetdriver)).md)

* [net (included topic-related subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/net(included-topic-related-subcommands).md)

* [SePrintOperatorPrivilege](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/SePrintOperatorPrivilege.md)


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

* [/etc/pam.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/etcpam.conf.md)

* [/etc/pam.d/](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/etcpam.d.md)

* [/etc/nsswitch.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/etcnsswitch.conf.md)

* [/etc/login.defs](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/etclogin.defs.md)

* [pam_ldap.so](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/pam_ldap.so.md)

* [ldap.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/ldap.conf.md)

* [pam_krb5.so](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/pam_krb5.so.md)

* [pam_cracklib.so](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/pam_cracklib.so.md)

* [pam_tally2.so](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/pam_tally2.so.md)

* [pam_faillock.so](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/pam_faillock.so.md)

* [pam_mkhomedir.so](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/pam_mkhomedir.so.md)

* [chage](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/chage.md)

* [faillog](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/faillog.md)

* [sssd](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/sssd.md)

* [sssd.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/sssd.conf.md)

* [sss_override](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/sss_override.md)

* [sss_cache](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/sss_cache.md)

* [sss_debuglevel](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/sss_debuglevel.md)

* [sss_user* and sss_group*](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/sss_user-and-sss_group.md)

* [/var/lib/sss/db/](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/varlibsssdb.md)

* [krb5.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/krb5.conf.md)

* [kinit](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/kinit.md)

* [klist](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/klist.md)

* [kdestroy](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/kdestroy.md)



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

* [smb.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smb.conf.md)

* [smbclient (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smbclient(including-relevant-subcommands).md)

* [mount](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/mount.md)

* [mount.cifs](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/mount.cifs.md)

* [/etc/fstab](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/etcfstab.md)

* [pam_mount.so](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/pam_mount.so.md)

* [pam_mount.conf.xml](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/pam_mount.conf.xml.md)

* [cifscreds](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/cifscreds.md)

* [getcifsacl](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/getcifsacl.md)

* [setcifsacl](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/setcifsacl.md)

* [smbcquotas](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smbcquotas.md)

* [cifsiostat](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/cifsiostat.md)

* [smbget](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smbget.md)

* [smbtar](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smbtar.md)

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

* [smb.conf:](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/smb.conf.md)

* [logon path](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/logon-path.md)

* [logon script](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/logon-script.md)

* [net (Windows command; including all relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/net(Windows-command%3Bincluding-all-relevant-subcommands).md)

* [samba-tool gpo (including all relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/samba-tool-gpo(including-all-relevant-subcommands).md)

* [gpupdate (Windows command)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/gpupdate(Windows-command).md)

* [rdesktop](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/rdesktop.md)

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

* [ipa-server-install](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/ipa-server-install.md)

* [ipa-replica-prepare](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/ipa-replica-prepare.md)

* [ipa-replica-install](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/ipa-replica-install.md)

* [ipa-client-install](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/ipa-client-install.md)

* [ipactl](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/ipactl.md)

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

* [ipa (including relevant user-*, stageuser-* and group-* and idview-* subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/ipa(including-relevant-user-%2Cstageuser-and-group-and-idview-subcommands).md)

* [ipa (including relevant host-*, hostgroup-*, service-* and getkeytab subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/ipa(including-relevant%20host-%2C%20hostgroup-%2C%20service-and-getkeytab-subcommands).md)

* [ipa (including relevant permission-*, privilege-*, and role-* subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/ipa(including-relevant-permission-%2Cprivilege-%2Cand-role-subcommands).md)

* [ipctl](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/ipctl.md)

* [ipa-advice](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/ipa-advice.md)



### 305.3 FreeIPA Active Directory Integration * [weight: 2]

#### Description: Candidates should be able to set up a cross-forest trust between a FreeIPA and an Active Directory domain.

Key Knowledge Areas:

Understand and set up FreeIPA and Active Directory integration using Kerberos cross-realm trusts
Configure ID ranges in FreeIPA
Understand and manage external non-POSIX groups in FreeIPA
Awareness of Microsoft Privilege Attribute Certificates and how they are handled by FreeIPA
Awareness of replication based FreeIPA and Active Directory integration
The following is a partial list of the used files, terms and utilities:

* [ipa-adtrust-install](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/ipa-adtrust-install.md)
* [ipa (including relevant trust-*, idrange-* and group-* subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/ipa(including-relevant-trust-%2Cidrange-and-group-subcommands).md)


### 305.4 Network File System * [weight: 3]

#### Description: Candidates should be able to use NFSv4. This includes understanding ID mapping, NFSv4 ACLs and Kerberos authentication for NFS.

Key Knowledge Areas:

Understand major NFSv4 features
Configure and manage an NFSv4 server and clients
Understand and use the NFSv4 pseudo file system
Understand and use NFSv4 ACLs
Use Kerberos for for NFSv4 authentication
The following is a partial list of the used files, terms and utilities:

* [exportfs](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/exportfs.md)

* [/etc/exports](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/etcexports.md)

* [/etc/idmapd.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/etcidmapd.conf.md)

* [nfs4_editfacl](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/nfs4_editfacl.md)

* [nfs4_getfacl](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/nfs4_getfacl.md)

* [nfs4_setfacl](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/nfs4_setfacl.md)

* [mount (including common NFS mount options)](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/mount(including-common-NFS-mount-options).md)

* [/etc/fstab](../TXT%20FILES/File-systems-Cocepts/LPIC3-300/etcfstab.md)
