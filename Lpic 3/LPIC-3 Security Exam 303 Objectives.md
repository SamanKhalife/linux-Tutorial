# LPIC-3 Security Exam 303 Objectives
## Topic 331: Cryptography
### 331.1 X.509 Certificates and Public Key Infrastructures (weight: 5)
#### Description:	Candidates should understand X.509 certificates and public key infrastructures. They should know how to configure and use OpenSSL to implement certification authorities and issue SSL certificates for various purposes.

Key Knowledge Areas:

Understand X.509 certificates, X.509 certificate lifecycle, X.509 certificate fields and X.509v3 certificate extensions
Understand trust chains and public key infrastructures, including certificate transparency
Generate and manage public and private keys
Create, operate and secure a certification authority
Request, sign and manage server and client certificates
Revoke certificates and certification authorities
Basic feature knowledge of Let’s Encrypt, ACME and certbot
Basic feature knowledge of CFSSL
Partial list of the used files, terms and utilities:

* [openssl (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/openssl%20(including%20relevant%20subcommands).md)

* [OpenSSL configuration](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/OpenSSL%20configuration.md)

* [PEM, DER, PKCS](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/PEM%2C%20DER%2C%20PKCS.md)

* [CSR](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/CSR.md)

* [CRL](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/CRL.md)

* [OCSP](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/OCSP.md)
 

### 331.2 X.509 Certificates for Encryption, Signing and Authentication 
#### Description:	Candidates should be able to use X.509 certificates for both server and client authentication. This includes implementing user and server authentication for Apache HTTPD. The version of Apache HTTPD covered is 2.4 or higher.

Key Knowledge Areas:

Understand SSL, TLS, including protocol versions and ciphers
Configure Apache HTTPD with mod_ssl to provide HTTPS service, including SNI and HSTS
Configure Apache HTTPD with mod_ssl to serve certificate chains and adjust the cipher configuration (no cipher-specific knowledge)
Configure Apache HTTPD with mod_ssl to authenticate users using certificates
Configure Apache HTTPD with mod_ssl to provide OCSP stapling
Use OpenSSL for SSL/TLS client and server tests
Partial list of the used files, terms and utilities:

* [httpd.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/httpd.conf.md)

* [mod_ssl](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/mod_ssl.md)

* [openssl (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/openssl%20(including%20relevant%20subcommands).md)

### 331.3 Encrypted File Systems (weight: 3)
#### Description:	Candidates should be able to set up and configure encrypted file systems.
Key Knowledge Areas:

Understand block device and file system encryption
Use dm-crypt with LUKS1 to encrypt block devices
Use eCryptfs to encrypt file systems, including home directories and PAM integration
Awareness of plain dm-crypt
Awareness of LUKS2 features
Conceptual understanding of Clevis for LUKS devices and Clevis PINs for TMP2 and Network Bound Disk Encryption (NBDE)/Tang
The following is a partial list of the used files, terms and utilities:

* [cryptsetup (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/cryptsetup%20(including%20relevant%20subcommands).md)

* [cryptmount](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/cryptmount.md)

* [/etc/crypttab](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/-etc-crypttab.md)

* [ecryptfsd](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/ecryptfsd.md)

* [ecryptfs-* commands](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/ecryptfs-%20commands.md)

* [mount.ecryptfs, umount.ecryptfs](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/mount.ecryptfs%2C%20umount.ecryptfs.md)

* [pam_ecryptfs](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/pam_ecryptfs.md)

### 331.4 DNS and Cryptography (weight: 5)
#### Description:	Candidates should have experience and knowledge of cryptography in the context of DNS and its implementation using BIND. The version of BIND covered is 9.7 or higher.

Key Knowledge Areas:

Understand the concepts of DNS, zones and resource records
Understand DNSSEC, including key signing keys, zone signing keys and relevant DNS records such as DS, DNSKEY, RRSIG, NSEC, NSEC3
and NSEC3PARAM

Configure and troubleshoot BIND as an authoritative name server serving DNSSEC secured zones
Manage DNSSEC signed zones, including key generation, key rollover and re-signing of zones
Configure BIND as an recursive name server that performs DNSSEC validation on behalf of its clients
Understand CAA and DANE, including relevant DNS records such as CAA and TLSA
Use CAA and DANE to publish X.509 certificate and certificate authority information in DNS
Use TSIG for secure communication with BIND
Awareness of DNS over TLS and DNS over HTTPS
Awareness of Multicast DNS
Partial list of the used files, terms and utilities:

* [named.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/named.conf.md)

* [dnssec-keygen](../TXT%20FILES/dnssec-signkey.md)

* [dnssec-signzone](../TXT%20FILES/File-systems-Cocepts/LPIC2-202/dnssec-signzone.md)

* [dnssec-settime](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/dnssec-settime.md)

* [dnssec-dsfromkey](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/dnssec-dsfromkey.md)

* [rndc (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/rndc%20(including%20relevant%20subcommands).md)

* [dig](../TXT%20FILES/dig.md)

* [delv](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/delv.md)

* [openssl (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/openssl%20(including%20relevant%20subcommands).md)
 

## Topic 332: Host Security
### 332.1 Host Hardening (weight: 5)
#### Description:	Candidates should be able to secure computers running Linux against common threats.
Key Knowledge Areas:

Configure BIOS and boot loader (GRUB 2) security
Disable unused software and services
Understand and drop unnecessary capabilities for specific systemd units and the entire system
Understand and configure Address Space Layout Randomization (ASLR), Data Execution Prevention (DEP) and Exec-Shield
Black and white list USB devices attached to a computer using USBGuard
Create an SSH CA, create SSH certificates for host and user keys using the CA and configure OpenSSH to use SSH certificates
Work with chroot environments
Use systemd units to limit the system calls and capabilities available to a process
Use systemd units to start processes with limited or no access to specific files and devices
Use systemd units to start processes with dedicated temporary and /dev directories and without network access
Understand the implications of Linux Meltdown and Spectre mitigations and enable/disable the mitigations
Awareness of polkit
Awareness of the security advantages of virtualization and containerization
The following is a partial list of the used files, terms and utilities:

* [grub.cfg](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/grub.cfg.md)

* [systemctl](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/systemctl.md)

* [getcap](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/getcap.md)

* [setcap](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/setcap.md)

* [capsh](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/capsh.md)

* [sysctl](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/sysctl.md)

* [/etc/sysctl.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/-etc-sysctl.conf.md)

* [/etc/usbguard/usbguard-daemon.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/-etc-usbguard-usbguard-daemon.conf.md)

* [/etc/usbguard/rules.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/-etc-usbguard-rules.conf.md)

* [usbguard](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/usbguard.md)

* [ssh-keygen](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/ssh-keygen.md)

* [/etc/ssh/](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/-etc-ssh.md)

* [~/.ssh/](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/~-.ssh-.md)

* [/etc/ssh/sshd_config](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/-etc-ssh-sshd_config-.md)

* [chroot](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/chroot.md)

### 332.2 Host Intrusion Detection (weight: 5)
#### Description:	Candidates should be familiar with the use and configuration of common host intrusion detection software. This includes managing the Linux Audit system and verifying a system’s integrity.
Key Knowledge Areas:

Use and configure the Linux Audit system
Use chkrootkit
Use and configure rkhunter, including updates
Use Linux Malware Detect
Automate host scans using cron
Use RPM and DPKG package management tools to verify the integrity of installed files
Configure and use AIDE, including rule management
Awareness of OpenSCAP
Partial list of the used files, terms and utilities:

* [auditd](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/auditd.md)

* [auditctl](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/auditctl.md)

* [ausearch, aureport](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/ausearch%2C%20aureport.md)

* [auditd.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/auditd.conf.md)

* [audit.rules](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/audit.rules.md)

* [pam_tty_audit.so](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/pam_tty_audit.so.md)

* [chkrootkit](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/chkrootkit.md)

* [rkhunter](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/rkhunter.md)

* [/etc/rkhunter.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/-etc-rkhunter.conf.md)

* [maldet](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/maldet.md)

* [conf.maldet](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/conf.maldet.md)

* [rpm](../TXT%20FILES/rpm.md)

* [dpkg](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/dpkg.md)

* [aide](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/aide.md)

* [/etc/aide/aide.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/-etc-aide-aide.conf.md)

### 332.3 Resource Control (weight: 3)
#### Description:	Candidates should be able to restrict the resources services and programs can consume.
Key Knowledge Areas:

Understand and configure ulimits
Understand cgroups, including classes, limits and accounting
Manage cgroups and process cgroup association
Understand systemd slices, scopes and services
Use systemd units to limit the system resources processes can consume
Awareness of cgmanager and libcgroup utilities
Partial list of the used files, terms and utilities:

* [ulimit](utorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC3-303/ulimit.md)

* [/etc/security/limits.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/-etc-security-limits.conf.md)

* [pam_limits.so](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/pam_limits.so.md)

* [/sys/fs/group/](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/-sys-fs-group-.md)

* [/proc/cgroups](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/-proc-cgroups.md)

* [systemd-cgls](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/systemd-cgls.md)

* [systemd-cgtop](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/systemd-cgtop.md)

## Topic 333: Access Control
### 333.1 Discretionary Access Control (weight: 3)
#### Description:	Candidates should understand discretionary access control (DAC) and know how to implement it using access control lists (ACL). Additionally, candidates are required to understand and know how to use extended attributes.
Key Knowledge Areas:

Understand and manage file ownership and permissions, including SetUID and SetGID bits
Understand and manage access control lists
Understand and manage extended attributes and attribute classes
Partial list of the used files, terms and utilities:

* [getfacl](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/getfacl.md)

* [setfacl](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/setfacl.md)

* [getfattr](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/getfattr.md)

* [setfattr](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/setfattr.md)
 

### 333.2 Mandatory Access Control (weight: 5)
#### Description:	Candidates should be familiar with mandatory access control (MAC) systems for Linux. Specifically, candidates should have a thorough knowledge of SELinux. Also, candidates should be aware of other mandatory access control systems for Linux. This includes major features of these systems but not configuration and use.
Key Knowledge Areas:

Understand the concepts of type enforcement, role based access control, mandatory access control and discretionary access control
Configure, manage and use SELinux
Awareness of AppArmor and Smack
Partial list of the used files, terms and utilities:

* [getenforce](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/getenforce.md)

* [setenforce](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/setenforce.md)

* [selinuxenabled](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/selinuxenabled.md)

* [getsebool](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/getsebool.md)

* [setsebool](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/setsebool.md)

* [togglesebool](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/togglesebool.md)

* [fixfiles](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/fixfiles.md)

* [restorecon](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/restorecon.md)

* [setfiles](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/setfiles.md)

* [newrole](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/newrole.md)

* [setcon](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/setcon.md)

* [c](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/c.md)

* [chcon](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/chcon.md)

* [semanage](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/semanage.md)

* [sestatus](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/sestatus.md)

* [seinfo](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/seinfo.md)

* [apol](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/apol.md)

* [seaudit](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/seaudit.md)

* [audit2why](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/audit2why.md)

* [audit2allow](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/audit2allow.md)

* [/etc/selinux/*](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/-etc-selinux-.md)

## Topic 334: Network Security
### 334.1 Network Hardening (weight: 4)
#### Description:	Candidates should be able to secure networks against common threats. This includes analyzing network traffic of specific nodes and protocols.
Key Knowledge Areas:

Understand wireless networks security mechanisms
Configure FreeRADIUS to authenticate network nodes
Use Wireshark and tcpdump to analyze network traffic, including filters and statistics
Use Kismet to analyze wireless networks and capture wireless network traffic
Identify and deal with rogue router advertisements and DHCP messages
Awareness of aircrack-ng and bettercap
The following is a partial list of the used files, terms and utilities:

* [radiusd](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/radiusd.md)

* [radmin](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/radmin.md)

* [radtest](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/radtest.md)

* [radclient](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/radclient.md)

* [radlast](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/radlast.md)

* [radwho](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/radwho.md)

* [radiusd.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/radiusd.conf.md)

* [/etc/raddb/*](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/-etc-raddb-.md)

* [wireshark](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/wireshark.md)

* [tshark](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/tshark.md)

* [tcpdump](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/tcpdump.md)

* [kismet](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/kismet.md)

* [ndpmon](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/ndpmon.md)
 

### 334.2 Network Intrusion Detection (weight: 4)
#### Description:	Candidates should be familiar with the use and configuration of network security scanning, network monitoring and network intrusion detection software. This includes updating and maintaining the security scanners.
Key Knowledge Areas:

Implement bandwidth usage monitoring
Configure and use Snort, including rule management
Configure and use OpenVAS, including NASL
Partial list of the used files, terms and utilities:

* [ntop](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/ntop.md)

* [snort](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/snort.md)

* [snort-stat](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/snort-stat.md)

* [pulledpork.pl](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/pulledpork.pl.md)

* [/etc/snort/*](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/-etc-snort-.md)

* [openvas-adduser](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/openvas-adduser.md)

* [openvas-rmuser](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/openvas-rmuser.md)

* [openvas-nvt-sync](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/openvas-nvt-sync.md)

* [openvassd](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/openvassd.md)

* [openvas-mkcert](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/openvas-mkcert.md)

* [openvas-feed-update](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/openvas-feed-update.md)

* [/etc/openvas/*](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/-etc-openvas-.md)

### 334.3 Packet Filtering (weight: 5)
#### Description:	Candidates should be familiar with the use and configuration of the netfilter Linux packet filter.
Key Knowledge Areas:

Understand common firewall architectures, including DMZ
Understand and use iptables and ip6tables, including standard modules, tests and targets
Implement packet filtering for IPv4 and IPv6
Implement connection tracking and network address translation
Manage IP sets and use them in netfilter rules
Awareness of nftables and nft
Awareness of ebtables
Awareness of conntrackd
Partial list of the used files, terms and utilities:

* [iptables](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/iptables.md)

* [ip6tables](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/ip6tables.md)

* [iptables-save](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/iptables-save.md)

* [iptables-restore](../TXT%20FILES/iptables-restore.md)

* [ip6tables-save](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/ip6tables-save.md)

* [ip6tables-restore](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/ip6tables-restore.md)

* [ipset](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/ipset.md)

### 334.4 Virtual Private Networks (weight: 4)
#### Description:	Candidates should be familiar with the use of OpenVPN, IPsec and WireGuard to set up remote access and site to site VPNs.
Key Knowledge Areas:

Understand the principles of bridged and routed VPNs
Understand the principles and major differences of the OpenVPN, IPsec, IKEv2 and WireGuard protocols
Configure and operate OpenVPN servers and clients
Configure and operate IPsec servers and clients using strongSwan
Configure and operate WireGuard servers and clients
Awareness of L2TP
Partial list of the used files, terms and utilities:

* [/etc/openvpn/](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/-etc-openvpn-.md)

* [openvpn](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/openvpn.md)

* [/etc/strongswan.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/-etc-strongswan.conf.md)

* [/etc/strongswan.d/](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/-etc-strongswan.d-.md)

* [/etc/swanctl/swanctl.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/-etc-swanctl-swanctl.conf.md)

* [/etc/swanctl/](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/-etc-swanctl-.md)

* [swanctl](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/swanctl.md)

* [/etc/wireguard/](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/-etc-wireguard-.md)

* [wg](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/wg.md)

* [wg-quick](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/wg-quick.md)

* [ip](../TXT%20FILES/ip.md)

## Topic 335: Threats and Vulnerability Assessment
### 335.1 Common Security Vulnerabilities and Threats (weight: 2)
#### Description:	Candidates should understand the principle of major types of security vulnerabilities and threats.
Key Knowledge Areas:

Conceptual understanding of threats against individual nodes
Conceptual understanding of threats against networks
Conceptual understanding of threats against application
Conceptual understanding of threats against credentials and confidentiality
Conceptual understanding of honeypots
The following is a partial list of the used files, terms and utilities:

* [Trojans](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/Trojans.mdv)

* [Viruses](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/Rootkits.md)

* [Rootkits](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/Rootkits.md)

* [Keylogger](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/Keylogger.md)

* [DoS and DDoS](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/DoS%20and%20DDoS.md)

* [Man in the Middle](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/Man%20in%20the%20Middle.md)

* [ARP and NDP forgery](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/ARP%20and%20NDP%20forgery.md)

* [Rogue Access Points, Routers and DHCP servers](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/Rogue%20Access%20Points%2C%20Routers%20and%20DHCP%20servers.md)

* [Link layer address and IP address spoofing](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/Link%20layer%20address%20and%20IP%20address%20spoofing.md)

* [Buffer Overflows](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/Buffer%20Overflows.md)

* [SQL and Code Injections](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/SQL%20and%20Code%20Injections.md)

* [Cross Site Scripting](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/Cross%20Site%20Scripting.md)

* [Cross Site Request Forgery](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/Cross%20Site%20Request%20Forgery.md)

* [Privilege escalation](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/Privilege%20escalation.md)

* [Brute Force Attacks](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/Brute%20Force%20Attacks.md)

* [Rainbow tables](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/Rainbow%20tables.md)

* [Phishing](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/Phishing.md)

* [Social Engineering](../TXT%20FILES/File-systems-Cocepts/LPIC3-303/Social%20Engineering.md)

### 335.2 Penetration Testing (weight: 3)
#### Description:	Candidates understand the concepts of penetration testing, including an understand of commonly used penetration testing tools. 
Furthermore, candidates should be able to use nmap to verify the effectiveness of network security measures.
Key Knowledge Areas:

Understand the concepts of penetration testing and ethical hacking
Understand legal implications of penetration testing
Understand the phases of penetration tests, such as active and passive information gathering, enumeration, gaining access, privilege escalation, access maintenance, covering tracks
Understand the architecture and components of Metasploit, including Metasploit module types and how Metasploit integrates various security tools
Use nmap to scan networks and hosts, including different scan methods, version scans and operating system recognition
Understand the concepts of Nmap Scripting Engine and execute existing scripts
Awareness of Kali Linux, Armitage and the Social Engineer Toolkit (SET)
Partial list of the used files, terms and utilities:

* [nmap](../TXT%20FILES/File-systems-Cocepts/LPIC2-201/nmap.md)
