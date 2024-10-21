# LPIC-2 Exam 202 Objectives
## Topic 207: Domain Name Server
### 207.1 Basic DNS server configuration (Weight:3)
#### Description	Candidates should be able to configure BIND to function as an authoritative and as a recursive, caching-only DNS server. This objective includes the ability to manage a running server and configuring logging.
Key Knowledge Areas:

BIND 9.x configuration files, terms and utilities
Defining the location of the BIND zone files in BIND configuration files
Reloading modified configuration and zone files
Awareness of dnsmasq, djbdns and PowerDNS as alternate name servers
The following is a partial list of the used files, terms and utilities:

* [/etc/named.conf](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/-etc-named.conf.md)

* [/var/named/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/-var-named-.md)

* [rndc](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/rndc.md)

* [named-checkconf](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/named-checkconf.md)

* [kill](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/kill.md)

* [host](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/host.md)

* [dig](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/dig.md)
 

### 207.2 Create and maintain DNS zones (Weight:3)
#### Description	Candidates should be able to create a zone file for a forward or reverse zone and hints for root level servers. This objective includes setting appropriate values for records, adding hosts in zones and adding zones to the DNS. A candidate should also be able to delegate zones to another DNS server.
Key Knowledge Areas:

BIND 9 configuration files, terms and utilities
Utilities to request information from the DNS server
Layout, content and file location of the BIND zone files
Various methods to add a new host in the zone files, including reverse zones
The following is a partial list of the used files, terms and utilities:

* [/var/named/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/-var-named-.md)

* [zone file syntax](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/zone%20file%20syntax.md)

* [resource record formats](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/resource%20record%20formats.md)

* [named-checkzone](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/named-checkzone.md)

* [named-compilezone](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/named-compilezone.md)

* [masterfile-format](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/masterfile-format.md)

* [dig](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/dig.md)

* [nslookup](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/nslookup.md)

* [host](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/host.md)

### 207.3 Securing a DNS server (Weight:2)
#### Description	Candidates should be able to configure a DNS server to run as a non-root user and run in a chroot jail. This objective includes secure exchange of data between DNS servers.
Key Knowledge Areas:

BIND 9 configuration files
Configuring BIND to run in a chroot jail
Split configuration of BIND using the forwarders statement
Configuring and using transaction signatures (TSIG)
Awareness of DNSSEC and basic tools
Awareness of DANE and related records
The following is a partial list of the used files, terms and utilities:

* [/etc/named.conf](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/-etc-named.conf.md)

* [/etc/passwd](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/-etc-passwd.md)

* [DNSSEC](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/DNSSEC.md)
 
* [dnssec-keygen](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/dnssec-keygen.md)

* [dnssec-signzone](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/dnssec-signzone.md)

## Topic 208: HTTP Services
### 208.1 Basic Apache configuration (Weight:4)
#### Description	Candidates should be able to install and configure a web server. This objective includes monitoring the server’s load and performance, restricting client user access, configuring support for scripting languages as modules and setting up client user authentication. Also included is configuring server options to restrict usage of resources. Candidates should be able to configure a web server to use virtual hosts and customise file access.
Key Knowledge Areas:

Apache 2.4 configuration files, terms and utilities
Apache log files configuration and content
Access restriction methods and files
mod_perl and PHP configuration
Client user authentication files and utilities
Configuration of maximum requests, minimum and maximum servers and clients
Apache 2.4 virtual host implementation (with and without dedicated IP addresses)
Using redirect statements in Apache’s configuration files to customise file access
The following is a partial list of the used files, terms and utilities:


* [access logs and error logs](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/access%20logs%20and%20error%20logs.md)

* [.htaccess](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/.htaccess.md)

* [httpd.conf](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/httpd.conf.md)

* [mod_auth_basic, mod_authz_host and mod_access_compat](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/mod_auth_basic%2C%20mod_authz_host%20and%20mod_access_compat.md)

* [htpasswd](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/htpasswd.md)

* [AuthUserFile, AuthGroupFile](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/AuthUserFile%2C%20AuthGroupFile.md)

* [apachectl, apache2ctl](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/apachectl%2C%20apache2ctl.md)

* [httpd, apache2](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/httpd%2C%20apache2.md)

### 208.2 Apache configuration for HTTPS (Weight:3)
#### Description	Candidates should be able to configure a web server to provide HTTPS.
Key Knowledge Areas:

SSL configuration files, tools and utilities
Generate a server private key and CSR for a commercial CA
Generate a self-signed Certificate
Install the key and certificate, including intermediate CAs
Configure Virtual Hosting using SNI
Awareness of the issues with Virtual Hosting and use of SSL
Security issues in SSL use, disable insecure protocols and ciphers
The following is a partial list of the used files, terms and utilities:


* [Apache2 configuration files](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/Apache2%20configuration%20files.md)

* [/etc/ssl/, /etc/pki/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/-etc-ssl-%2C%20-etc-pki-.md)

* [openssl, CA.pl](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/openssl%2C%20CA.pl.md)

* [SSLEngine, SSLCertificateKeyFile, SSLCertificateFile](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/SSLEngine%2C%20SSLCertificateKeyFile%2C%20SSLCertificateFile.md)

* [SSLCACertificateFile, SSLCACertificatePath](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/SSLCACertificateFile%2C%20SSLCACertificatePath.md)

* [SSLProtocol, SSLCipherSuite, ServerTokens, ServerSignature, TraceEnable](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/SSLProtocol%2C%20SSLCipherSuite%2C%20ServerTokens%2C%20ServerSignature%2C%20TraceEnable.md)
 
### 208.3 Implementing Squid as a caching proxy (Weight:2)
#### Description	Candidates should be able to install and configure a proxy server, including access policies, authentication and resource usage.
Key Knowledge Areas:

Squid 3.x configuration files, terms and utilities
Access restriction methods
Client user authentication methods
Layout and content of ACL in the Squid configuration files
The following is a partial list of the used files, terms and utilities:

* [squid.conf](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/squid.conf.md)

* [acl](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/acl.md)

* [http_access](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/http_access.md)

### 208.4 Implementing Nginx as a web server and a reverse proxy (Weight:2)
#### Description	Candidates should be able to install and configure a reverse proxy server, Nginx. Basic configuration of Nginx as a HTTP server is included.
Key Knowledge Areas:

Nginx
Reverse Proxy
Basic Web Server
The following is a partial list of the used files, terms and utilities:

* [/etc/nginx/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/-etc-nginx.md)

* [nginx](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/nginx.md)
 

## Topic 209: File Sharing
### 209.1 Samba Server Configuration (Weight:)	5
#### Description	Candidates should be able to set up a Samba server for various clients. This objective includes setting up Samba as a standalone server as well as integrating Samba as a member in an Active Directory. Furthermore, the configuration of simple CIFS and printer shares is covered. Also covered is a configuring a Linux client to use a Samba server. Troubleshooting installations is also tested.
Key Knowledge Areas:

Samba 4 documentation
Samba 4 configuration files
Samba 4 tools and utilities and daemons
Mounting CIFS shares on Linux
Mapping Windows user names to Linux user names
User-Level, Share-Level and AD security
The following is a partial list of the used files, terms and utilities:

* [smbd, nmbd, winbindd](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/smbd%2C%20nmbd%2C%20winbindd.md)

* [smbcontrol, smbstatus, testparm, smbpasswd, nmblookup](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/smbcontrol%2C%20smbstatus%2C%20testparm%2C%20smbpasswd%2C%20nmblookup.md)

* [samba-tool](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/samba-tool.md)

* [net](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/net.md)

* [smbclient](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/smbclient.md)

* [mount.cifs](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/mount.cifs.md)

* [/etc/samba/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/-etc-samba-.md)

* [/var/log/samba/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/-var-log-samba-.md)

### 209.2 NFS Server Configuration (Weight:3)
#### Description	Candidates should be able to export filesystems using NFS. This objective includes access restrictions, mounting an NFS filesystem on a client and securing NFS.

Key Knowledge Areas:

NFS version 3 configuration files
NFS tools and utilities
Access restrictions to certain hosts and/or subnets
Mount options on server and client
TCP Wrappers
Awareness of NFSv4
The following is a partial list of the used files, terms and utilities:

* [/etc/exports](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/-etc-exports.md)

* [exportfs](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/exportfs.md)

* [showmount](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/showmount.md)

* [nfsstat](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/nfsstat.md)

* [/proc/mounts](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/-proc-mounts.md)

* [/etc/fstab](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/-etc-fstab.md)

* [rpcinfo](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/rpcinfo.md)

* [mountd](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/mountd.md)

* [portmapper](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/portmapper.md)

## Topic 210: Network Client Management
### 210.1 DHCP configuration (Weight:2)
#### Description	Candidates should be able to configure a DHCP server. This objective includes setting default and per client options, adding static hosts and BOOTP hosts. Also included is configuring a DHCP relay agent and maintaining the DHCP server.
Key Knowledge Areas:

DHCP configuration files, terms and utilities
Subnet and dynamically-allocated range setup
Awareness of DHCPv6 and IPv6 Router Advertisements
The following is a partial list of the used files, terms and utilities:

* [dhcpd.conf](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/dhcpd.conf.md)

* [dhcpd.leases](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/dhcpd.leases.md)

* [DHCP Log messages in syslog or systemd journal](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/DHCP%20Log%20messages%20in%20syslog%20or%20systemd%20journal.md)

* [arp](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/arp.md)

* [dhcpd](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/dhcpd.md)

* [radvd](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/radvd.md)

* [radvd.conf](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/radvd.conf.md)

### 210.2 PAM authentication (Weight:3)
#### Description	The candidate should be able to configure PAM to support authentication using various available methods. This includes basic SSSD functionality.
Key Knowledge Areas:

PAM configuration files, terms and utilities
passwd and shadow passwords
Use sssd for LDAP authentication
The following is a partial list of the used files, terms and utilities:

* [/etc/pam.d/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/etcpam.d.md)

* [pam.conf](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/pam.conf.md)

* [nsswitch.conf](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/nsswitch.conf.md)

* [pam_unix, pam_cracklib, pam_limits, pam_listfile, pam_sss](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/pam_unix%2C%20pam_cracklib%2C%20pam_limits%2C%20pam_listfile%2C%20pam_sss.md)

* [sssd.conf](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/sssd.conf.md)

### 210.3 LDAP client usage (Weight:)	2
#### Description	Candidates should be able to perform queries and updates to an LDAP server. Also included is importing and adding items, as well as adding and managing users.
Key Knowledge Areas:

LDAP utilities for data management and queries
Change user passwords
Querying the LDAP directory
The following is a partial list of the used files, terms and utilities:

* [ldapsearch](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/ldapsearch.md)

* [ldappasswd](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/ldappasswd.md)

* [ldapadd](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/ldapadd.md)

* [ldapdelete](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/ldapdelete.md)

### 210.4 Configuring an OpenLDAP server (Weight:4)
#### Description	Candidates should be able to configure a basic OpenLDAP server including knowledge of LDIF format and essential access controls.
Key Knowledge Areas:

OpenLDAP
Directory based configuration
Access Control
Distinguished Names
Changetype Operations
Schemas and Whitepages
Directories
Object IDs, Attributes and Classes
The following is a partial list of the used files, terms and utilities:

* [slapd](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/slapadd.md)

* [slapd-config](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/slapd-config.md)

* [LDIF](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/LDIF.md)

* [slapadd](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/slapadd.md)

* [slapcat](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/slapcat.md)

* [slapindex](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/slapindex.md)

* [/var/lib/ldap/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/varlibldap.md)

* [loglevel](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/loglevel.md)

## Topic 211: E-Mail Services
### 211.1 Using e-mail servers (Weight:4)
#### Description	Candidates should be able to manage an e-mail server, including the configuration of e-mail aliases, e-mail quotas and virtual e-mail domains. This objective includes configuring internal e-mail relays and monitoring e-mail servers.
Key Knowledge Areas:

Configuration files for postfix
Basic TLS configuration for postfix
Basic knowledge of the SMTP protocol
Awareness of sendmail and exim
The following is a partial list of the used files, terms and utilities:

* [Configuration files and commands for postfix](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/Configuration%20files%20and%20commands%20for%20postfix.md)

* [/etc/postfix/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/-etc-postfix-.md)

* [/var/spool/postfix/]()

* [sendmail emulation layer commands](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/sendmail%20emulation%20layer%20commands.md)

* [/etc/aliases](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/-etc-aliases.md)

* [mail-related logs in /var/log/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/mail-related%20logs%20in%20-var-log-.md)

### 211.2 Managing E-Mail Delivery (Weight:2)
#### Description	Candidates should be able to implement client e-mail management software to filter, sort and monitor incoming user e-mail.
Key Knowledge Areas:

Understanding of Sieve functionality, syntax and operators
Use Sieve to filter and sort mail with respect to sender, recipient(s), headers and size
Awareness of procmail
The following is a partial list of the used files, terms and utilities:

* [Conditions and comparison operators](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/Conditions%20and%20comparison%20operators.md)
* [keep, fileinto, redirect, reject, discard, stop](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/keep%2C%20fileinto%2C%20redirect%2C%20reject%2C%20discard%2C%20stop%20.md)
* [Dovecot vacation extension](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/Dovecot%20vacation%20extension.md)

### 211.3 Managing Mailbox Access (Weight:2)
#### Description	Candidates should be able to install and configure POP and IMAP daemons.
Key Knowledge Areas:

Dovecot IMAP and POP3 configuration and administration
Basic TLS configuration for Dovecot
Awareness of Courier
The following is a partial list of the used files, terms and utilities:

* [/etc/dovecot/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/-etc-dovecot-.md)

* [dovecot.conf](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/dovecot.conf.md)

* [doveconf](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/doveconf.md)

* [doveadm](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/doveadm.md)

## Topic 212: System Security
### 212.1 Configuring a router (Weight:3)
#### Description	Candidates should be able to configure a system to forward IP packet and perform network address translation (NAT, IP masquerading) and state its significance in protecting a network. This objective includes configuring port redirection, managing filter rules and averting attacks.

Key Knowledge Areas:

iptables and ip6tables configuration files, tools and utilities
Tools, commands and utilities to manage routing tables.
Private address ranges (IPv4) and Unique Local Addresses as well as Link Local Addresses (IPv6)
Port redirection and IP forwarding
List and write filtering and rules that accept or block IP packets based on source or destination protocol, port and address
Save and reload filtering configurations
The following is a partial list of the used files, terms and utilities:

* [/proc/sys/net/ipv4/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/-proc-sys-net-ipv4-.md)

* [/proc/sys/net/ipv6/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/-proc-sys-net-ipv6-.md)

* [/etc/services](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/-etc-services.md)

* [iptables](https://github.com/SamanKhalife/linux-Tutorial/tree/main/Firewalls/Iptables)

* [ip6tables](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/ip6tables.md)

### 212.2 Managing FTP servers (Weight:2)
#### Description	Candidates should be able to configure an FTP server for anonymous downloads and uploads. This objective includes precautions to be taken if anonymous uploads are permitted and configuring user access.
Key Knowledge Areas:

Configuration files, tools and utilities for Pure-FTPd and vsftpd
Awareness of ProFTPd
Understanding of passive vs. active FTP connections
The following is a partial list of the used files, terms and utilities:

* [vsftpd.conf](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/vsftpd.conf.md)

* [important Pure-FTPd command line options](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/important%20Pure-FTPd%20command%20line%20options.md)

### 212.3 Secure shell (SSH) (Weight:4)
#### Description	Candidates should be able to configure and secure an SSH daemon. This objective includes managing keys and configuring SSH for users. Candidates should also be able to forward an application protocol over SSH and manage the SSH login.
Key Knowledge Areas:

OpenSSH configuration files, tools and utilities
Login restrictions for the superuser and the normal users
Managing and using server and client keys to login with and without password
Usage of multiple connections from multiple hosts to guard against loss of connection to remote host following configuration changes
The following is a partial list of the used files, terms and utilities:

* [ssh](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/ssh.md)

* [sshd](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/sshd.md)

* [/etc/ssh/sshd_config](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/-etc-ssh-sshd_config.md)

* [/etc/ssh/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/-etc-ssh.md)

* [Private and public key files](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/Private%20and%20public%20key%20files.md)

* [PermitRootLogin, PubKeyAuthentication, AllowUsers, PasswordAuthentication, Protocol](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/PermitRootLogin%2C%20PubKeyAuthentication%2C%20AllowUsers%2C%20PasswordAuthentication%2C%20Protocol.md)

### 212.4 Security tasks (Weight:3)
#### Description	Candidates should be able to receive security alerts from various sources, install, configure and run intrusion detection systems and apply security patches and bugfixes.
Key Knowledge Areas:

Tools and utilities to scan and test ports on a server
Locations and organisations that report security alerts as Bugtraq, CERT or other sources
Tools and utilities to implement an intrusion detection system (IDS)
Awareness of OpenVAS and Snort
The following is a partial list of the used files, terms and utilities:

* [telnet](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/telnet.md)

* [nmap](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-201/nmap.md)

* [fail2ban](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/fail2ban.md)

* [nc](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-201/nc.md)

* [iptables](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/iptables.md)
 

### 212.5 OpenVPN (Weight:2)
#### Description	Candidates should be able to configure a VPN (Virtual Private Network) and create secure point-to-point or site-to-site connections.
Key Knowledge Areas:

OpenVPN
The following is a partial list of the used files, terms and utilities:

* [/etc/openvpn/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/-etc-openvpn.md)

* [openvpn](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-202/openvpn.md)
 
