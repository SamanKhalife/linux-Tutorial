# linux-Tutorial (with some SRE's info)

linux commands and any thing basic you need yo know for using linux (for cloud computing) it is for my own Knowledge but if you want to use I try to make it very complex and from base so you can learn linux so easy and work on linux based systems.

# Contents (Not-Only-In-This-Page)
* [Quick-review-important-commants](https://github.com/SamanKhalife/linux-Tutorial/blob/main/Quick%20review%20for%20commands.md)
* [full-list-for-all-Commands](https://github.com/SamanKhalife/linux-commands/blob/main/Commands.md)
* [tree-of-repo](https://github.com/SamanKhalife/linux-Tutorial/blob/main/SUMMARY.md)
* Exams-Objectives-and-samples:
  * [LPIC1](https://github.com/SamanKhalife/linux-Tutorial/tree/main/Lpic%201)
  * [LPIC2](https://github.com/SamanKhalife/linux-Tutorial/tree/main/Lpic%202)
  * [LPIC3](https://github.com/SamanKhalife/linux-Tutorial/tree/main/Lpic%203)
* [Some-How-To-do](https://github.com/SamanKhalife/linux-Tutorial/tree/main/How%20to)
* [Deep-into-linux](https://github.com/SamanKhalife/linux-Tutorial?tab=readme-ov-file#common-commands-preview)

## Common commands preview

* \[View Linux command help information] - commands：`help`, `whatis`, `info`, `which`, `whereis`, `man`
* \[Linux file directory management] -commands：`cd`, `ls`, `pwd`, `mkdir`, `rmdir`, `tree`, `touch`, `ln`, `rename`, `stat`, `file`, `chmod`, `chown`, `locate`, `find`, `cp`, `mv`, `rm`
* \[Linux file content viewing command] - commands：`cat`, `head`, `tail`, `more`, `less`, `sed`, `vi`, `grep`
* \[Linux File Compression and Decompression] -commands：`tar`, `gzip`, `zip`, `unzip`
* \[Linux User Management] - commands：`groupadd`, `groupdel`, `groupmod`, `useradd`, `userdel`, `usermod`, `passwd`, `su`, `sudo`
* \[Linux System Administration ] - commands：`reboot`, `exit`, `shutdown`, `date`, `mount`, `umount`, `ps`, `kill`, `systemctl`, `service`, `crontab`
* \[Linux Network Administration] - commands：`curl`, `wget`, `telnet`, `ip`, `hostname`, `ifconfig`, `route`, `ssh`, `ssh-keygen`, `firewalld`, `iptables`, `host`, `nslookup`, `nc`, `netcat`, `ping`, `traceroute`, `netstat`
* \[Linux Hardware Management ] - commands：`df`, `du`, `top`, `free`, `iotop`
* \[Linux Software Management ] - commands：`rpm`, `yum`, `apt-get`

## Linux Getting Started Overview

* **Kernel operating system**
* **Drivers**
* **System software**
* **Application software**
* **Application software with graphical interface**

![img](https://github.com/SamanKhalife/linux-commands/blob/main/IMAGES/1111111111111111111.png)

* **Application software with graphical interface**

![img](https://github.com/SamanKhalife/linux-commands/blob/main/IMAGES/2222222222222222222.png)

## More info about Linux kernel 

[Linux kernel map](https://makelinux.github.io/kernel/map/)

## Linux Boot Process

* There are 6 stages of the Linux boot process

![img](https://github.com/SamanKhalife/linux-commands/blob/main/IMAGES/3333333333333333333.png)

### Stage 1: BIOS (Basic Input/Output System)

* `Bios` stands for basic input/output system.
* This is the first step of the boot process, `Bios` doing the POST (Power-on Self-test) job.
* **POST** is a process of checking the properties.
* When the BIOS process is successful, the BIOS will search and boot an operating system contained in devices such as hard drives, CD/DVD, USB.
* Normally, the BIOS will check the floppy or CD-ROM drives to see if they can boot from them, then the hardware, the order of checking the drives depends on the configuration in the BIOS .
  * If the **BIOS** can't find it `boot device` it will report `No boot device found`.
  * If the operating system `Linux` is installed on the hard disk, it will find the Master Boot Record at the first sector of the first hard drive.

### Stage 2: MBR loading

* MBR (Master Boot Record) is stored in the first sector of a data storage device ( `/dev/sda`, `/dev/hda`).
* MBR is very small, only `512byte`=`1 sector`.
* The MBR contains the following information:
  * Primary boot loader code (`446 byte`) : provides information for `boot loaderand` location `boot loaderand` the hard drive.
  * Partition table information ( `64 byte`) : stores information about `partition`.
  * Magic number ( 2 byte) : used to check MBR , if MBR fails, it will restore.

### Stage 3: GRUB Loader

* After locating the Boot Loader , this step will execute `load` the Boot Loader into memory and read the configuration - information then display the GRUB Boot Menu for the user to choose from.
* If the User does not select the OS, after a preset time, GRUB will load the default Kernel into memory to boot.

### Stage 4: Kernel

* The operating system's kernel will be loaded into **RAM** -> When the kernel works, the first thing will execute the INIT process.

### Stage 5: Run Level (INIT)

* This is the main stage of the process `boot`.

```
       Table 1. Mapping between runlevels and systemd targets
       ┌───────────────────────────────────────┐
       │Runlevel    │ Target            │      │
       ├───────────────────────────────────────┤
       │0           │ poweroff.target   │      │
       ├───────────────────────────────────────┤
       │1           │ rescue.target     │      │
       ├───────────────────────────────────────┤
       │2, 3, 4     │ multi-user.target │      │
       ├───────────────────────────────────────┤
       │5           │ graphical.target  │      │
       ├───────────────────────────────────────┤
       │6           │   reboot.target   │      │
       └───────────────────────────────────────┘
```

* Run level0 : halt - Shutdown system
* Run level1 : single-user mode -No network configuration, start processes and allow loginuser non-root
* Run level2 : multi-user mode -no network configuration, start processes.
* Run level3 : multi-user mode with networking -boots the system normally on the command line interface.
* Run level4 : undefined
* Run level5 : X11- start the system on the graphical interface
* Run level6 : reboot- reboot the system.

### Stage 6: User Prompt

* User login and use.


[for-deeper-view-of-kernel-click](https://github.com/0xAX/linux-insides/blob/master/SUMMARY.md)

## Shell

### Concepts

* **Shell** is a program that executes commands from users or applications - the utility requires passing to **Kernel** for processing.
* Activities of **Shell**:
  * 1: Command parsing.
  * 2: Interpreting the command's request.
  * 3: Pass message to **Kernel**.
  * 4: Display the result of the command.

#### Types of **Shell**

* **sh(the Bourne Shell)**
  * Do not autocomplete file names
  * Do not save the typed statement in memory (`history`)
* **bash (Bourne-Again Shell)**
  * Linux defaults.
  * Support more commands.
* **csh (the C shell)**
  * Written in C . programming language.

#### Shell Prompt (Shell Prompt)

* Is a character or a character set that always starts at the beginning of any command line.
* Indicates that **shell** is ready to receive commands from the user.
* `$`: ordinary user
* `#`: user `root`

## Things you should know about Linux Partitions

### 1.Hard drives

* As a data storage device, for example in Linux it `isdev/sda`
* `/dev/sda3` is the 3rd primary partition on the first disk
* `/dev/sdb5` is the first logical partition on the second disk
* `/dev/sda7` is the 3rd logical partition of the first physical disk

### 2. Partition

* Partition is a small partition (Logical Partition) divided fromHard drives
* Each Hard driveshas 4 partition, in which partition includes 3 types: Primary partition , Extended partition and Logical Partition
* Primary partition +The primary partition can be used to boot the operating system.
* Extended partition +Is the remaining data area when we divide into Primary partitions. +Extended partition contains Logical Partitions in it. Each drive can only hold 1 Extended partition. +Logical partition : These are small partitions in the extended partition , usually used to store data.

fdisk

```
root@saman:~# fdisk /dev/sda
Welcome to fdisk (util-linux 2.25.1).
Changes will remain in memory only, until you decide to write them.
Be careful before using the write command.


root@saman:~# Command (m for help): p
Disk /dev/sda: 298.1 GiB, 320072933376 bytes, 625142448 sectors
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: dos
Disk identifier: 0x000beca1

Device     Boot     Start       End   Sectors   Size Id Type
/dev/sda1  *         2048  43094015  43091968  20.6G 83 Linux
/dev/sda2        43094016  92078390  48984375  23.4G 83 Linux
/dev/sda3        92080126 625141759 533061634 254.2G  5 Extended
/dev/sda5        92080128 107702271  15622144   7.5G 82 Linux swap / Solaris
/dev/sda6       107704320 625141759 517437440 246.8G 83 Linux
```

parted

```
root@saman:~# sudo parted /dev/sda p
Model: ATA ST320LT000-9VL14 (scsi)
Disk /dev/sda: 320GB
Sector size (logical/physical): 512B/512B
Partition Table: msdos
Disk Flags:

Number  Start   End     Size    Type      File system     Flags
 1      1049kB  22.1GB  22.1GB  primary   ext4            boot
 2      22.1GB  47.1GB  25.1GB  primary   ext4
 3      47.1GB  320GB   273GB   extended
 5      47.1GB  55.1GB  7999MB  logical   linux-swap(v1)
 6      55.1GB  320GB   265GB   logical
```

* LVM (logical volume management)

In many cases, you need to resize your partitions or even install new disks and add them to your current mount points; Increasing the total size. LVM is designed for this. + Physical Volume (PV): A whole drive or a partition. It is better to define partitions and not use whole disks - unpartitioned. + Volume Groups (VG): This is the collection of one or more PVs. OS will see the vg as one big disk. PVs in one VG, can have different sizes or even be on different physical disks. + Logical Volumes (LV): OS will see lvs as partitions. You can format an LV with your OS and use it.

* To find a file with a larger size, you can use the du (Disk Usage) and sort commands. In the command line, run the following

```
du -h /path/to/directory | sort -rh
```

### MBR and GPT

* Information about hard disk partitions will be stored on MBR (Master Boot record) or GPT (GUID Partition table).
* These are two standards for configuring and managing Partitions on hard drives.
* The information stored here includes the location and size of the Partitions.

### Swap file

* Swap is virtual memory concept used on Linux.
* When VPS/Serveroperating, if the system runs out of RAM, it will automatically use part of the hard drive as memory for applications to operate.
* Swap is slower than RAM because Swap is part of the hard drive.
* (mostly we set swap 2_ram but less than 8 GB for hdd and set 1_ram or 2\*ram )

## These are some of the most used bash environment variables

| USER     | The name of the logged-in user                                                                  |
| -------- | ----------------------------------------------------------------------------------------------- |
| PATH     | List of directories to search for commands, colon separated                                     |
| EDITOR   | Default editor                                                                                  |
| HISTFILE | System hostname                                                                                 |
| HOSTNAME | Show Info                                                                                       |
| PS1      | The Prompt! Play with it                                                                        |
| UID      | The numeric user id of the logged-in user                                                       |
| HOME     | The user's home directory                                                                       |
| PWD      | The current working directory                                                                   |
| SHELL    | The name of the shell                                                                           |
| $        | The process id (or PID of the running bash shell (or other) process                             |
| PPID     | The process id of the process that started this process (that is, the id of the parent process) |
| ?        | The exit code of the last command                                                               |

## File Directory

* /home – User personal data
* /boot – Static files of the boot loader
* /sbin – System binaries
* /dev – Device files
* /sys
* /etc – Configuration files
* /proc – Process and kernel files
* /var – Variable data files
  * /tmp – Temporary files
  * /lock
  * /log
* /usr – User binaries and program data
  * /bin – Binaries
  * /man
  * /lib – Shared libraries
  * /local
  * /share
* /opt – Optional software
* /root – The home directory of the root
* /media – Mount point for removable media
* /mnt – Mount point for mounting a filesystem temporarily
* /srv – Service data

### All default Linux commands in bin dir

<details>

<summary>VIEW ALL</summary>

\
for better view you can see the code----------------------------------------------------------------------------------------------------------- envsubst lslogins rgrep systemd-cryptenroll VGAuthService eqn lsmem rlogin systemd-delta '\[' ex lsmod rm systemd-detect-virt aa-enabled expand lsns rmdir systemd-escape aa-exec expiry lsof rnano systemd-hwdb aa-features-abi expr lspci routef systemd-id128 add-apt-repository factor lspgpot routel systemd-inhibit addpart faillog lspower rrsync systemd-machine-id-setup addr2line fallocate lsusb rsh systemd-mount apport-bug false lzcat rsync systemd-notify apport-cli fgconsole lzcmp rsync-ssl systemd-path apport-collect fgrep lzdiff rtstat systemd-repart apport-unpack file lzegrep run-one systemd-run apropos finalrd lzfgrep run-one-constantly systemd-socket-activate apt fincore lzgrep run-one-until-failure systemd-stdio-bridge apt-add-repository find lzless run-one-until-success systemd-sysext apt-cache findmnt lzma run-parts systemd-sysusers apt-cdrom flock lzmainfo run-this-one systemd-tmpfiles apt-config fmt lzmore runcon systemd-tty-ask-password-agent apt-extracttemplates fold mailmail3 rview systemd-umount apt-ftparchive free man rvim tabs apt-get ftp man-recode savelog tac apt-key fuser mandb sbattach tail apt-mark fusermount manifest sbkeysync tar apt-sortpkgs fusermount3 manpath sbsiglist taskset ar gapplication mapscrn sbsign tbl arch gawk mawk sbvarsign tclsh as gdbus mcookie sbverify tclsh8.6 automat-visualize3 geqn md5sum scandeps tcpdump awk getconf md5sum.textutils scp tee b2sum getent mdig screen telnet base32 getkeycodes mesg screendump telnet.netkit base64 getopt migrate-pubring-from-classic-gpg script tempfile basename gettext mk\_modmap scriptlive test basenc gettext.sh mkdir scriptreplay tic bash ginstall-info mkfifo scsi\_logging\_level time bashbug gio mknod scsi\_mandat timedatectl bc gio-querymodules mksquashfs scsi\_readcap timeout boltctl git mktemp scsi\_ready tkconch3 btrfs git-receive-pack mokutil scsi\_satl tload btrfs-convert git-shell more scsi\_start tmux btrfs-find-root git-upload-archive mount scsi\_stop tnftp btrfs-image git-upload-pack mountpoint scsi\_temperature toe btrfs-map-logical glib-compile-schemas mt sdiff top btrfs-select-super gold mt-gnu sed touch btrfsck gp-archive mtr select-editor tput btrfstune gp-collect-app mtr-packet sensible-browser tr busctl gp-display-html mv sensible-editor tracepath busybox gp-display-src namei sensible-pager trial3 byobu gp-display-text nano seq troff byobu-config gpasswd nawk setarch true byobu-ctrl-a gpg nc setfont truncate byobu-disable gpg-agent nc.openbsd setkeycodes tset byobu-disable-prompt gpg-connect-agent neqn setleds tsort byobu-enable gpg-wks-server netcat setlogcons tty byobu-enable-prompt gpg-zip networkctl setmetamode twist3 byobu-export gpg2 networkd-dispatcher setpci twistd3 byobu-janitor gpgcompose newgrp setpriv tzselect byobu-keybindings gpgconf ngettext setsid ua byobu-launch gpgparsemail nice setterm ubuntu-advantage byobu-launcher gpgsm nisdomainname setupcon ubuntu-bug byobu-launcher-install gpgsplit nl sftp ubuntu-core-launcher byobu-launcher-uninstall gpgtar nm sg ubuntu-distro-info byobu-layout gpgv nohup sg\_bg\_ctl ubuntu-security-status byobu-prompt gpic nproc sg\_compare\_and\_write ucf byobu-quiet gprof nroff sg\_copy\_results ucfq byobu-reconnect-sockets gprofng nsenter sg\_dd ucfr byobu-screen grep nslookup sg\_decode\_sense uclampset byobu-select-backend gresource nstat sg\_emc\_trespass udevadm byobu-select-profile groff nsupdate sg\_format ul byobu-select-session grog ntfs-3g sg\_get\_config umount byobu-shell grops ntfs-3g.probe sg\_get\_elem\_status uname byobu-silent grotty ntfscat sg\_get\_lba\_status unattended-upgrade byobu-status groups ntfscluster sg\_ident unattended-upgrades byobu-status-detail growpart ntfscmp sg\_inq uncompress byobu-tmux grub-editenv ntfsdecrypt sg\_logs unexpand byobu-ugraph grub-file ntfsfallocate sg\_luns unicode\_start byobu-ulevel grub-fstest ntfsfix sg\_map unicode\_stop c++filt grub-glue-efi ntfsinfo sg\_map26 uniq c\_rehash grub-kbdcomp ntfsls sg\_modes unlink captoinfo grub-menulst2cfg ntfsmove sg\_opcodes unlzma cat grub-mkfont ntfsrecover sg\_persist unmkinitramfs catman grub-mkimage ntfssecaudit sg\_prevent unshare cftp3 grub-mklayout ntfstruncate sg\_raw unsquashfs chage grub-mknetdir ntfsusermap sg\_rbuf unxz chardet grub-mkpasswd-pbkdf2 ntfswipe sg\_rdac unzstd chardetect grub-mkrelpath numfmt sg\_read update-alternatives chattr grub-mkrescue objcopy sg\_read\_attr update-mime-database chcon grub-mkstandalone objdump sg\_read\_block\_limits uptime chfn grub-mount od sg\_read\_buffer usb-devices chgrp grub-ntldr-img oem-getlogs sg\_read\_long usbhid-dump chmod grub-render-label on\_ac\_power sg\_readcap usbreset choom grub-script-check openssl sg\_reassign users chown grub-syslinux2cfg openvt sg\_referrals utmpdump chrt gsettings os-prober sg\_rep\_pip uuidgen chsh gtbl pager sg\_rep\_zones uuidparse chvt gunzip partx sg\_requests vcs-run ckbcomp gzexe passwd sg\_reset vdir ckeygen3 gzip paste sg\_reset\_wp vi cksum h2ph pastebinit sg\_rmsn view clear h2xs patch sg\_rtpg vigpg clear\_console hardlink pathchk sg\_safte vim cloud-id hd pbget sg\_sanitize vim.basic cloud-init head pbput sg\_sat\_identify vim.tiny cloud-init-per helpztags pbputs sg\_sat\_phy\_event vimdiff cmp hexdump pdb3 sg\_sat\_read\_gplog vimtutor codepage host pdb3.10 sg\_sat\_set\_features vm-support col hostid peekfd sg\_scan vmhgfs-fuse col1 hostname perl sg\_seek vmstat col2 hostnamectl perl5.34-x86\_64-linux-gnu sg\_senddiag vmtoolsd col3 htop perl5.34.0 sg\_ses vmware-alias-import col4 hwe-support-status perlbug sg\_ses\_microcode vmware-checkvm col5 i386 perldoc sg\_start vmware-hgfsclient col6 iconv perlivp sg\_stpg vmware-namespace-cmd col7 id perlthanks sg\_stream\_ctl vmware-rpctool col8 info pgrep sg\_sync vmware-toolbox-cmd col9 infobrowser pic sg\_test\_rwbuf vmware-vgauth-cmd colcrt infocmp pico sg\_timestamp vmware-vmblock-fuse colrm infotocap piconv sg\_turs vmware-xferlogs column install pidof sg\_unmap w comm install-info pidwait sg\_verify wall conch3 instmodsh pinentry sg\_vpd watch corelist ionice pinentry-curses sg\_wr\_mode watchgnupg cp ip ping sg\_write\_buffer wc cpan ipcmk ping4 sg\_write\_long wdctl cpan5.34-x86\_64-linux-gnu ipcrm ping6 sg\_write\_same wget cpio ipcs pinky sg\_write\_verify whatis crontab iptables-xml pkaction sg\_write\_x whereis csplit ischroot pkcheck sg\_xcopy which ctail iscsiadm pkcon sg\_zone which.debianutils ctstat join pkexec sginfo whiptail curl journalctl pkill sgm\_dd who cut jq pkmon sgp\_dd whoami cvtsudoers json-patch-jsondiff pkttyagent sh wifi-status dash json\_pp pl2pm sha1sum write date jsondiff pldd sha224sum x86\_64 dbus-cleanup-sockets jsonpatch plymouth sha256sum x86\_64-linux-gnu-addr2line dbus-daemon jsonpointer pmap sha384sum x86\_64-linux-gnu-ar dbus-monitor jsonschema pod2html sha512sum x86\_64-linux-gnu-as dbus-run-session kbd\_mode pod2man shasum x86\_64-linux-gnu-c++filt dbus-send kbdinfo pod2text showconsolefont x86\_64-linux-gnu-dwp dbus-update-activation-environment kbxutil pod2usage showkey x86\_64-linux-gnu-elfedit dbus-uuidgen keep-one-running podchecker shred x86\_64-linux-gnu-gold dd kernel-install pollinate shuf x86\_64-linux-gnu-gp-archive deallocvt keyring pr size x86\_64-linux-gnu-gp-collect-app deb-systemd-helper kill preconv skill x86\_64-linux-gnu-gp-display-html deb-systemd-invoke killall printenv slabtop x86\_64-linux-gnu-gp-display-src debconf kmod printf sleep x86\_64-linux-gnu-gp-display-text debconf-apt-progress kmodsign prlimit slogin x86\_64-linux-gnu-gprof debconf-communicate landscape-sysinfo pro snap x86\_64-linux-gnu-gprofng debconf-copydb last prove snapctl x86\_64-linux-gnu-ld debconf-escape lastb prtstat snapfuse x86\_64-linux-gnu-ld.bfd debconf-set-selections lastlog ps snice x86\_64-linux-gnu-ld.gold debconf-show lcf psfaddtable soelim x86\_64-linux-gnu-nm debian-distro-info ld psfgettable sort x86\_64-linux-gnu-objcopy delpart ld.bfd psfstriptable sos x86\_64-linux-gnu-objdump delv ld.gold psfxtable sos-collector x86\_64-linux-gnu-ranlib df ldd pslog sosreport x86\_64-linux-gnu-readelf dh\_bash-completion less pstree splain x86\_64-linux-gnu-size diff lessecho pstree.x11 split x86\_64-linux-gnu-strings diff3 lessfile ptar splitfont x86\_64-linux-gnu-strip dig lesskey ptardiff sqfscat xargs dir lesspipe ptargrep sqfstar xauth dircolors lexgrog ptx ss xdg-user-dir dirmngr libnetcfg purge-old-kernels ssh xdg-user-dirs-update dirmngr-client link pwd ssh-add xsubpp dirname linux-boot-prober pwdx ssh-agent xxd distro-info linux-check-removal py3clean ssh-argv0 xz dmesg linux-update-symlinks py3compile ssh-copy-id xzcat dnsdomainname linux-version py3versions ssh-import-id xzcmp do-release-upgrade linux32 pybabel ssh-import-id-gh xzdiff domainname linux64 pybabel-python3 ssh-import-id-lp xzegrep dpkg ln pydoc3 ssh-keygen xzfgrep dpkg-deb lnstat pydoc3.10 ssh-keyscan xzgrep dpkg-divert loadkeys pygettext3 stat xzless dpkg-maintscript-helper loadunimap pygettext3.10 static-sh xzmore dpkg-query locale pyhtmlizer3 stdbuf yes dpkg-realpath locale-check pyserial-miniterm strace ypdomainname dpkg-split localectl pyserial-ports strace-log-merge zcat dpkg-statoverride localedef python3 streamzip zcmp dpkg-trigger logger python3.10 strings zdiff du login pzstd strip zdump dumpkeys loginctl ranlib stty zegrep dwp logname rbash su zfgrep eatmydata look rcp sudo zforce ec2metadata lowntfs-3g rdma sudoedit zgrep echo ls readelf sudoreplay zipdetails ed lsattr readlink sum zless editor lsb\_release realpath sync zmore efibootdump lsblk red systemctl znew efibootmgr lscpu renice systemd zstd egrep lsfd rescan-scsi-bus.sh systemd-analyze zstdcat eject lshw reset systemd-ask-password zstdgrep elfedit lsinitramfs resizecons systemd-cat zstdless enc2xs lsipc resizepart systemd-cgls zstdmt encguess lsirq resolvectl systemd-cgtop env lslocks rev systemd-creds

</details>

## User Permissions and attributes

| command | what it means     | what it does                    | how to use it      |
| ------- | ----------------- | ------------------------------- | ------------------ |
| chmod   | change modifiers  | edits the permissions of a file | chmod +x filename  |
| chattr  | change attributes | edits the attributes of a file  | chattr +i filename |
| chown   | change owner      |                                 |                    |

`chmod` is used to edit the permissions on a file. There are 3 permissions and 3 groups, for a total of 9 permissions. The permissions are:

* read - allows the group to read the file
* write - allows the group to edit the file
* execute - allows the group to run the file as a program

for example when using `ls -la` you see all of the files permissions

```
(user)
 |   
 |     
 |     (others)
 |     |
-rw-r--r-- 1 username groupname date time ...
    |
    |
   (group) 

```

so this file the owner has rw, the group has r and everyone has r

### to make a file executable for all users you would run the following command

`chmod a+x filename`

| character | group name | who it applies to                         |
| --------- | ---------- | ----------------------------------------- |
| u         | user       | the owner of the file                     |
| g         | group      | users who are members of the files' group |
| o         | others     | users not included in u or g              |
| a         | all        | u, g and o                                |

chmod can take either a set of three octal (base 8) numbers to set the permissions across all groups, or list the groups "a+" and the permissions that you want to add or "a-" and the permissions you want to subtract.

The character representations of the permissions are:

| letter | number | binary | permission |
| ------ | ------ | ------ | ---------- |
| r      | 4      | 100    | read       |
| w      | 2      | 010    | write      |
| x      | 1      | 001    | execute    |

So `chmod a+rwx filename` is the same as `chmod 777 filename`

* examples:
* `chmod u+x,g+w filename`
* `chmod u=rwx,g=rw,o=r filename`
* `chmod 764 filename`
* `chmod a+x filename`

The `chown` command changes the owner and group of a file or directory. This command can authorize a user to become the owner of the specified file or change the group to which the file belongs. User can be user or user D, user group can be group name or group id. The filename can be a list of files separated by spaces, and wildcards can be included in the filename.

Only the file owner and superuser can use this command.

## Group Permissions and attributes

[gorupmod](<README (1).md>)

[groupadd](<TXT FILES/groupadd.md>)

[groupdel](<TXT FILES/groupdel.md>)

[groups](<TXT FILES/groups.md>)

## apt command in linux (for installing packages)

apt provides a high-level CLI (Command Line Interface) for the package management system and is intended as an interface for the end user which enables some options better suited for interactive usage by default compared to more specialized APT tools like apt-cache and apt-get.

Syntax :`apt [...COMMANDS] [...PACKAGES]`

update :`apt update`

full-upgrade :`apt full-upgrade`

install :`apt install [...PACKAGES]`

remove :`apt remove [...PACKAGES]`

purge :`apt purge [...PACKAGES]`

search :`apt search [...REGEX]` This command is used when the user wants to search for the given regex term(s) in the list of available packages and display matches. For example, this command can be useful when you want to search for packages having a specific feature.

show :`apt show [...PACKAGES]`

### APT-GET Command in Linux

Basic Syntax :`sudo apt-get [options] [command] [package(s)]`

If you want to download only one specific package, you can do:`apt-get download [...PACKAGES]`

Removing debian packages:`apt-get remove [package(s)]`

### package information with dpkg

If you want to reconfigure a package that is already installed, you can use the:`dpkg-reconfigure tzdata`

dpkg \[OPTIONS] ACTION PACKAGE

<details>

<summary>Switch Description</summary>

\


</details>

## wildcard characters for globbing

Wildcard characters are used in globbing and regular expressions to represent patterns of text. While the exact set of wildcard characters can vary depending on the context and programming language, here are some common wildcard characters and their meanings:

1. **Asterisk (`*`)**: Matches zero or more characters. It is often used as a wildcard for any sequence of characters.
   * Example: `*.txt` matches all files ending with ".txt."
2. **Question Mark (`?`)**: Matches a single character. It is used to represent any character in a specific position.
   * Example: `file?.txt` matches "file1.txt" and "file2.txt."
3. **Square Brackets (`[]`) and Dash (`-`) for Character Ranges**: Square brackets are used to define character classes, and the dash is used to specify character ranges within those classes.
   * `[aeiou]` matches any one vowel.
   * `[0-9]` matches any digit.
   * `[A-Za-z]` matches any uppercase or lowercase letter.
4. **Curly Braces (`{}`) for Alternatives**: Curly braces are used to specify alternative patterns or choices.
   * `{jpg,jpeg,png}` matches either "jpg," "jpeg," or "png."
5. **Exclamation Mark (`!`) or Caret (`^`) (Inside Square Brackets)**: When placed at the beginning of a character class, these symbols negate the class, matching any character not listed in the brackets.
   * `[^0-9]` matches any character that is not a digit.
6. **Period (`.`)**: In regular expressions, the period is a wildcard that matches any single character.
   * Example: `a.c` matches "abc," "adc," etc.
7. **Pipe (`|`) (In Regular Expressions)**: Represents the logical OR operation and is used to specify alternatives.
   * Example: `apple|banana` matches either "apple" or "banana."
8. **Backslash (`\`) (Escape Character)**: In some contexts, the backslash is used to escape special characters, so they are treated as literal characters rather than as wildcards or metacharacters.

These wildcard characters provide flexible ways to match patterns in text and are commonly used in various programming languages, command-line utilities, and regular expressions for tasks such as text searching, text processing, and file manipulation. Keep in mind that the availability and behavior of these characters can vary depending on the specific environment or tool you're using.

[more info wildcard characters for globbing](<How to/wildcard characters for globbing .md>)

## 2 Ways to Download Files From Linux Terminal

how do you download a file from the terminal?

1.Download files from Linux terminal using wget command

[wget command](<TXT FILES/wget.md>)

Installing wget:

```
sudo apt install wget
```

Download a file or webpage using wget:

```
wget URL
```

2.Download files from Linux command line using curl

[curl command](<TXT FILES/curl.md>)

Installing curl:

```
sudo apt install curl
```

Download a file or webpage using curl:

```
curl URL
```

## How to your find your Linux IP address

The following commands will get you the IP address list to find public IP addresses for your machine:

If you enter the command `ifconfig` in the terminal, all information will be displayed. You can also use these commands:

The following commands will get you the IP address list to find public IP addresses for your machine:

* `curl ifconfig.me`
* `curl -4/-6 icanhazip.com`
* `curl ipinfo.io/ip`
* `curl api.ipify.org`
* `curl checkip.dyndns.org`
* `dig +short myip.opendns.com @resolver1.opendns.com`
* `host myip.opendns.com resolver1.opendns.com`
* `curl ident.me`

The following commands will get you the private IP address of your interfaces:

* `ifconfig -a`
* `hostname -I'`
* `ip route get 1.2.3.4'`
* `nmcli -p device show`(best option in my opinion)

## Linux Networking Tools

[ping](<TXT FILES/ping.md>)

[host](<TXT FILES/host.md>)

[finger](<TXT FILES/finger.md>)

[traceroute](<TXT FILES/traceroute.md>)

[netstat](<TXT FILES/netstat.md>)

[tracepath](<TXT FILES/tracepath.md>)

[dig](<TXT FILES/dig.md>)

[hostname](<TXT FILES/hostname.md>)

[route](<TXT FILES/route.md>)

[nslookup](<TXT FILES/nslookup.md>)

## SSH

## DNS

## Firewalls

Here are the Best Firewalls for Linux for Effective System Protection

<details>

<summary>IPFire</summary>

\
Key Features:

Firewall engine and Instruction Prevention system.

Offers default zones with different security policies. For example, DMZ and LAN.

Frequently updated to prevent attack vectors and security vulnerabilities.

Offers Stateful Package Inspection(SPI) firewall built on top of Netfilter.

Provides an intuitive web user interface.

Protects against Denial-of-Service attacks.

It lets users create logging and graphical reports for insights.

It can be installed on hardware devices such as Raspberry Pi.

[IPFire](https://www.ipfire.org/)

</details>

<details>

<summary>PfSense</summary>

\
Key Features:

FreeBSD-based.

Supports a wide variety of hardware.

Clean web interface.

It comes with commercial-grade features.

Supports VPN endpoint and wireless access point configuration.

Outbound and Inbound load balancing.

Real-time information.

[pfsense](https://www.pfsense.org/)

</details>

<details>

<summary>Zenarmor</summary>

\
Key Features:

Web filtering, application control, and cloud threat intelligence.

Auto-block malware/phishing attempts in real-time.

Instantly deploy firewall with minimal setup requirements.

Offers centralized cloud management to manage multiple firewalls.

Improves network visibility with rich analytics and reporting.

[Zenarmor](https://www.zenarmor.com/)

</details>

<details>

<summary>UFW</summary>

\


[UFW help](https://help.ubuntu.com/community/UFW)

</details>

<details>

<summary>IPcop</summary>

\
[IPcop docs](https://www.ipcop.org/docs.html)

</details>

<details>

<summary>Opnsence</summary>

\


[Opnsence](https://opnsense.org/)

</details>

<details>

<summary>Surewall</summary>

\


[Surewall](https://shorewall.org/)

</details>

<details>

<summary>ClearOS</summary>

\


[ClearOS](https://www.clearos.com/)

</details>

<details>

<summary>Arista Edge</summary>

\


[Arista Edge](https://edge.arista.com/)

</details>

## iptables

iptables is a powerful configuration tool for controlling traffic to and from your system. Modern Linux kernels come with a packet filtering framework called Netfilter. Netfilter provides operations such as allow, drop, and modify to control the flow of packets into and out of the system. The user-level command-line tool iptables based on the Netfilter framework provides a powerful firewall configuration function, allowing you to add rules to build firewall policies. The richness and complexity of iptables and its baroque command syntax can be overwhelming.

### main view of iptable command

```
                                          (chains)
         (table)                             |                                  (source)                                                                                                             MASQURADE
            |                            PREROUTING          1                      |                                                             1                    1                             ACCEPT
            |       filter     -A        INPUT               2            udp       |   xxx.xxx.xxx.xxx            xxx.xxx.xxx.xxx                2                    2             -------         REJECT
iptables   -t       nat        -I        FORWARDING          .     -P     tcp      -S   37.16.2.4           -d     185.674.46.4       --sport     .       --dport      .      -m     module   -J     DROP
                    mangle     -D        OUTPUT              .            icmp                               |                                    .                    .             -------   |     DNAT 
                               -U        POSTROUTING         n                                               |                                  65535                65535                     |     SNAT
                                |                            |                                         (destination)                                                                           |
               (insert)(apend)(delete)(update)               |                                                                                                                 [this is the most important thing]
                                                             |                                                                                                                                         
                                                      (line you want) 
```

### chains that packets will go through in iptable

![img](https://github.com/SamanKhalife/linux-Tutorial/blob/main/IMAGES/4444444444444444444.png)

### If you live in countries that are embargoed by some companies and applications, you can use some thing like DNS changing for using the apps.

* [shecan](https://shecan.ir/)
* [hostiran](https://hostiran.net/landing/proxy)
* [electrotm](https://electrotm.org/)
* [begzar](https://begzar.ir/)
* [vanillapp](https://vanillapp.ir/)
* [403.online](https://403.online/)

### If you dont know how to change and set it follow these steps

* connect to your srver with ssh.
* open `/etc/resolv.conf` with some editor like nano or vim.
*   change the in front of namesrver like this:

    ```
    nameserver 172.29.2.100
    nameserver 185.51.200.2
    ```
* save changes and exit.

## How to Check Ubuntu Version Details and Other System Information

type the following command:

```
lsb_release -a
```

Alternate way:

```
cat /etc/lsb-release
```

Another Fun way:

[Neofetch command](Widgets/Fun/Neofetch.md)

## Getting SSL Certificat

you can use any kind tools you know here are some of them:

* [acme](https://github.com/acmesh-official/acme.sh)
* [zerossl](https://github.com/zerossl)
* [free-ssl](https://github.com/topics/free-ssl-certificates)
* [certbot](https://github.com/certbot/certbot)

[learn getting free ssl with acme step by step](acme-freessl.md)

## Some of the server monitoring tools

#### How To View Running Processes in Linux

* [top](<TXT FILES/top.md>)
* [htop](<TXT FILES/htop.md>)
* [ps](<TXT FILES/ps.md>)
* [lsof](<TXT FILES/lsof.md>)

#### How To Monitor Your Network

* [nethogs](<README (1).md>)
* [iptraf-ng](<README (1).md>)
* [netstat](<TXT FILES/netstat.md>)
* [iftop](<TXT FILES/iftop.md>)
* [(speed test)](<README (1).md>)

#### How To Monitor Your Disk Usage

* [df](<TXT FILES/df.md>)
* [du](<TXT FILES/du.md>)
* [iotop](<TXT FILES/iotop.md>)
* [iostat](<TXT FILES/iostat.md>)

#### How To Monitor Your Memory Usage

* [free(Ram usage)](<TXT FILES/free.md>)
* [vmstat](<TXT FILES/vmstat.md>)

## How to to Manage Processes in Linux

* [kill](<TXT FILES/kill.md>)
* [nice](<TXT FILES/nice.md>)

mosltly use for controlling the cpu usage (nice number is from -20 to 19 ).

## Some of the monitoring apps

* [prometheus](https://prometheus.io/)
  * [my docs](<README (1).md>)
* [grafana](https://grafana.com/)
  * [my docs](<README (1).md>)
* [zabbix](https://www.zabbix.com/)
* [paessler](https://www.paessler.com/)
* [nagios](https://www.nagios.org/)
* [elastic.co](https://www.elastic.co/kibana)
* [elastic](https://www.elastic.co/)
* [influxdata](https://www.influxdata.com/)

## How To Add Swap Space

What is Swap????? Swap is a portion of hard drive storage that has been set aside for the operating system to temporarily store data that it can no longer hold in RAM. This lets you increase the amount of information that your server can keep in its working memory, with some caveats.

#### Checking the System for Swap Information

We can see if the system has any configured swap by typing:`sudo swapon --show`.

If you don’t get back any output, this means your system does not have swap space available currently.

You can verify that there is no active swap using the free utility:`free -h`

#### Checking Available Space on the Hard Drive Partition

Before we create our swap file, we’ll check our current disk usage to make sure we have enough space:`df -h`

#### Creating a Swap File

Now that we know our available hard drive space, we can create a swap file on our filesystem. We will allocate a file of the size that we want called swapfile in our root (/) directory.

The best way of creating a swap file is with the fallocate program. This command instantly creates a file of the specified size.

Since the server in our example has 1G of RAM, we will create a 1G file in this guide. Adjust this to meet the needs of your own server:`sudo fallocate -l 1G /swapfile`

We can verify that the correct amount of space was reserved by typing:`ls -lh /swapfile`

#### Enabling the Swap File

Now that we have a file of the correct size available, we need to actually turn this into swap space.

First, we need to lock down the permissions of the file so that only users with root privileges can read the contents. This prevents normal users from being able to access the file, which would have significant security implications.

Make the file only accessible to root by typing:`sudo chmod 600 /swapfile`

Verify the permissions change by typing:`ls -lh /swapfile`

```
Output
-rw------- 1 root root 1.0G jul 16 16:18 /swapfile
```

only the root user has the read and write flags enabled.

We can now mark the file as swap space by typing:`sudo mkswap /swapfile`

After marking the file, we can enable the swap file, allowing our system to start using it:`sudo swapon /swapfile`

Verify that the swap is available by typing:`sudo swapon --show`

```
Output
NAME      TYPE  SIZE USED PRIO
/swapfile file 1024M   0B   -2
```

We can check the output of the free utility again to corroborate our findings:`free -h`

```
Output
              total        used        free      shared  buff/cache   available
Mem:          981Mi       123Mi       644Mi       0.0Ki       213Mi       714Mi
Swap:         1.0Gi          0B       1.0Gi
```

#### Making the Swap File Permanent

Our recent changes have enabled the swap file for the current session. However, if we reboot, the server will not retain the swap settings automatically. We can change this by adding the swap file to our /etc/fstab file.

Back up the /etc/fstab file in case anything goes wrong:`sudo cp /etc/fstab /etc/fstab.bak`

Add the swap file information to the end of your /etc/fstab file by typing:`echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab`

Next we’ll review some settings we can update to tune our swap space.

## Linux as a virtualization guest

topics:Virtual machine/Linux container/Application container/Guest drivers/SSH host keys/D-Bus machine id

To check and see if your host operating system / CPU, supports using hypervisors check for the `vmx` (for Intel CPUs) or `svm` (for AMD CPUs) in your `/proc/cpuinfo` in flags.

Based on your CPU you should have `kvm` or `kvm-amd` kernel modules loaded.

```
lsmod | grep -i kvm
sudo modprobe kvm
```

(If you see hypervisor in your /proc/cpuinfo it means that you are inside a virtualized Linux machine )

#### let's see the 2 types of hypervisors. First type 2, since it's easier to understand.

![img](https://github.com/SamanKhalife/linux-Tutorial/blob/main/IMAGES/5555555555555555555.png)

**Type 2 Hypervisor**

These hypervisors run on a conventional operating system (OS) just as other computer programs do. A guest operating system runs as a process on the host. Type-2 hypervisors abstract guest operating systems from the host operating system.

In other words, a type 2 hypervisor is the software between the guest and host. It completely runs on the host OS and provides virtualization to the guest.

Two of the most famous Type 2 hypervisors are VirtualBox (from Oracle) and VMware.

**Type 1 Hypervisor**

These hypervisors run directly on the host's hardware to control the hardware and manage guest operating systems. For this reason, they are sometimes called bare-metal hypervisors. The first hypervisors, which IBM developed in the 1960s, were native hypervisors. These included the test software SIMMON and the CP/CMS operating system, the predecessor of IBM z/VM.

Some of the most famous Type 1 hypervisors are KVM, Xen & Hyper-V. KVM is built-in since Linux Kernel version 2.6.20.

### Creating a Virtual Machine

First, create the machine itself. We tell the hypervisor this machine how much RAM/disk/CPU/... needs and set a name for our machine. Then we need to install the guest OS. This can be done using:

* Installing from a CD / DVD / ...
* Cloning an existing machine.
* Using Open Virtualization Format (OVF) to move machines between hypervisors. This is a standard format for virtual machine definition and may include several files, in this case, you can archive all of them into one Open Virtualization Archive (OVA) file.
* It is also possible to create Templates that are master copies\* to initiate new machines.

### Guest-specific configs

Some configurations are machine specific. For example, a network card's MAC address should be unique for whole the network. If we are cloning a machine or sometimes creating them from templates, at least we need to change these on each machine before booting them:

* Host Name
* NIC MAC Address
* NIC IP (If not using DHCP)
* Machine ID (delete the `/etc/machine-id` and `/var/lib/dbus/machine-id` and run `dbus-uuidgen --ensure`. These two files might be soft links to each other)
* Encryption Keys like SSH Fingerprints and PGP keys
* HDD UUIDs
* Any other UUIDs on the system

##

##

##

##

## Some info about Clouds and SRE things

##

##

##

## Cloud Delivery Models

#### Infrastructure as a Service (IaaS)

IaaS is the on-demand delivery of computing infrastructure, including operating systems, networking, storage, and other infrastructural components. Acting much like a virtual equivalent to physical servers, IaaS relieves cloud users of the need to buy and maintain physical servers while also providing the flexibility to scale and pay for resources as needed. IaaS is a popular option for businesses that wish to leverage the advantages of the cloud and have system administrators who can oversee the installation, configuration, and management of operating systems, development tools, and other underlying infrastructure that they wish to use. However, IaaS is also used by developers, researchers, and others who wish to customize the underlying infrastructure of their computing environment. Given its flexibility, IaaS can support everything from a company’s computing infrastructure to web hosting to big data analysis.

#### Platform as a Service (PaaS)

PaaS provides a computing platform where the underlying infrastructure (such as the operating system and other software) is installed, configured, and maintained by the provider, allowing users to focus their efforts on developing and deploying apps in a tested and standardized environment. PaaS is commonly used by software developers and developer teams as it cuts down on the complexity of setting up and maintaining computer infrastructure, while also supporting collaboration among distributed teams. PaaS can be a good choice for developers who don’t have the need to customize their underlying infrastructure, or those who want to focus their attention on development rather than DevOps and system administration.

#### Software as a Service (SaaS)

SaaS providers are cloud-based applications that users access on demand from the internet without needing to install or maintain the software. Examples include GitHub, Google Docs, Slack, and Adobe Creative Cloud. SaaS applications are popular among businesses and general users given that they’re often easy to adopt, accessible from any device, and have free, premium, and enterprise versions of their applications. Like PaaS, SaaS abstracts away the underlying infrastructure of the software application so that users are only exposed to the interface they interact with.

![img](https://github.com/SamanKhalife/linux-Tutorial/blob/main/IMAGES/7777777777777777777.png)

## Cloud Environments

#### Public Cloud

The public cloud refers to cloud services (such as virtual machines, storage, or applications) offered publicly by a commercial provider to businesses and individuals. Public cloud resources are hosted on the commercial provider’s hardware, which users access through the internet. They are not always suitable for organizations in highly-regulated industries, such as healthcare or finance, as public cloud environments may not comply with industry regulations regarding customer data.

#### Private Cloud

The private cloud refers to cloud services that are owned and managed by the organization that uses them and available only to the organization’s employees and customers. Private clouds allow organizations to exert greater control over their computing environment and their stored data, which can be necessary for organizations in highly-regulated industries. Private clouds are sometimes seen as more secure than public clouds as they are accessed through private networks and enable the organization to directly oversee their cloud security. Public cloud providers sometimes provide their services as applications that can be installed on private clouds, allowing organizations to keep their infrastructure and data on premise while taking advantage of the public cloud’s latest innovations.

#### Hybrid Cloud and Multicloud

Many organizations use a hybrid cloud environment which combines public and private cloud resources to support the organization’s computing needs while maintaining compliance with industry regulation. Multicloud environments are also common, which entail the use of more than one public cloud provider (for example, combining Amazon Web Services and DigitalOcean).

#### Managed cloud

Managed cloud services refer to cloud computing solutions in which a third-party provider takes on the responsibility of managing various aspects of your cloud infrastructure, applications, and services.

#### Micro cloud

Micro clouds are a new class of infrastructure for on-demand computing at the edge. They differ from the internet-of-things (IoT), which uses thousands of single machines or sensors to gather data, yet they perform computing tasks. Micro clouds reuse proven cloud primitives but with the unattended, autonomous and clustering features that resolve typical edge computing challenges.

#### Multi-cloud

Multi-cloud (also referred to as multi cloud or multicloud) is a concept that refers to using multiple clouds from more than one cloud service provider at the same time. The term is also used to refer to the simultaneous running of bare metal, virtualised and containerised workloads.

## SLA(Service Level agreement) / SLI(Service Level Indicators) / SLO(Service Level Objectives)

![img](https://github.com/SamanKhalife/linux-Tutorial/blob/main/IMAGES/6666666666666666666.png)

If the running time and availability of a system cannot be measured, it is very difficult to maintain and operate the system that is already online, which often causes the maintenance and operation team to continue to be in the state of a fire brigade, and when the root cause of the problem is finally found, it may be the development There is a problem with the code written by the team.

Development teams often don't see "stability" as a potential problem if they can't figure out how to measure runtime and availability. This problem has plagued Google for many years, which is why the SRE principle was developed. One of the motivations.

SRE ensures that everyone knows how to measure reliability and what to do when a service fails. This will be detailed to the point that when a problem occurs, from VP or CxO to every relevant employee within the organization, they all have to do what they should do. What each "person" should do is clearly regulated. SRE will communicate with all relevant personnel to determine Service Level Indicators (SLIs) and Service Level Objectives (SLOs).

SLIs define metrics related to the "response time" of a system, such as response time, throughput per second, requests, etc., and often convert this metric into a ratio or average.

SLOs are a time interval obtained after discussions with relevant personnel. It is expected that SLIs can maintain a certain level of figures, such as "what is the level of SLIs every month", which is more internal indicators.

The video also discusses Service Level Agreements (SLAs), even though it's not a number that SREs care about on a daily basis. As an online service provider, SLA is a commitment to customers to ensure that the percentage of continuous operation of the service is usually "negotiated" with customers, and the downtime per year (or month) should not be less than a few minutes.

The concepts of SLI, SLO, and SLA are very similar to the "everything can be measured" mentioned by DevOps, which is one of the reasons why it is said that class SRE implements DevOps.

#### Uptime

| Availability % | Downtime per year | Downtime per month | Downtime per Week |
| -------------- | ----------------- | ------------------ | ----------------- |
| 90%            | 36.5 days         | 72 hours           | 16.8 hours        |
| 95%            | 18.25 days        | 36 hours           | 8.4 hours         |
| 98%            | 7.30 days         | 14.4 hours         | 3.36 hours        |
| 99%            | 3.65 days         | 7.20 hours         | 1.68 hours        |
| 99.5%          | 1.83 days         | 3.60 hours         | 50.4 minutes      |
| 99.8%          | 17.52 hours       | 86.23 minutes      | 20.16 minutes     |
| 99.9%          | 8.76 hours        | 43.2 minutes       | 10.1 minutes      |
| 99.95%         | 4.38 hours        | 21.56 minutes      | 5.04 minutes      |
| 99.99%         | 52.6 minutes      | 4.32 minutes       | 1.01 minutes      |
| 99.999%        | 5.26 minutes      | 25.9 seconds       | 6.05 seconds      |
| 99.9999%       | 31.5 seconds      | 2.59 seconds       | 0.605 seconds     |

## Mean Time To Recover (MTTR)

MTTR is the average time that a device will take to recover from any failure.

## Mean Time Between Failures (MTBF)

MTBF is the predicted elapsed time between inherent failures of a mechanical or electronic system, during normal system operation. MTBF can be calculated as the arithmetic mean (average) time between failures of a system. The term is used for repairable systems, while mean time to failure (MTTF) denotes the expected time to failure for a non-repairable system.

## Mean Time To Failure (MTTF)

MTTF denotes the expected time to failure for a non-repairable system.


## Critical incident response time (CIRT)

Critical incident response time (CIRT) is a new, significantly more accurate method to evaluate operations performance. CIRT focuses on the incidents that are most likely to impact business by culling noise from incoming signals using the following techniques:

Real business-impacting (or potentially impacting) incidents are very rarely low urgency, so rule out all low-urgency incidents.

Real business-impacting incidents are very rarely (if ever) auto-resolved by monitoring tools without the need for human intervention, so rule out incidents that were not resolved by a human.

Short, bursting, and transient incidents that are resolved within 120 seconds are highly unlikely to be real business-impacting incidents, so rule them out.

Incidents that go unnoticed, tabled, or ignored (not acknowledged, not resolved) for a very long time are rarely business-impacting; rule them out. Note: 

This threshold can be a statistically derived number that is customer-specific (e.g., two standard deviations above the mean) to avoid using an arbitrary number.

Individual, ungrouped incidents generated by separate alerts are not representative of the larger business-impacting incident. Therefore, simulate incident 

groupings with a very conservative threshold, e.g., two minutes, to calculate response time.

## books that you need to read as sre

- [incident-management](https://www.atlassian.com/incident-management/get-the-handbook)
- [Building Secure & Reliable Systems](https://sre.google/books/)
- [The Site Reliability Workbook](https://sre.google/books/)
- [Site Reliability Engineering](https://sre.google/books/)


## Sources I get help from

* Digital ocean
* google youtube chanel
* some open source projects

