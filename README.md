# linux-Tutorial
linux commands and any thing basic you need yo know for using linux it is for my own Knowledge but if you want to use I try to make it very complex and from base so you can learn linux so easy.

<!--if you want to share please share it with the source of it -->
<!-- SAMAN KHALIFE  -->



# Linux Getting Started Overview
- **Kernel operating system**
- **Drivers** 
- **System software** 
- **Application software** 
- **Application software with graphical interface**
<!-- PIC11111111111  -->
![img](https://github.com/SamanKhalife/linux-commands/blob/main/IMAGES/1111111111111111111.png)



- **Application software with graphical interface**
<!-- PIC22222222222  -->
![img](https://github.com/SamanKhalife/linux-commands/blob/main/IMAGES/2222222222222222222.png)


# Linux Boot Process
- There are 6 stages of the Linux boot process
<!-- PIC33333333333  -->
![img](https://github.com/SamanKhalife/linux-commands/blob/main/IMAGES/3333333333333333333.png)
## Stage 1: BIOS (Basic Input/Output System)
- `Bios` stands for basic input/output system
- This is the first step of the boot process, `Bios` doing the POST (Power-on Self-test) job.
- **POST**  is a process of checking the properties.
- When the BIOS process is successful, the BIOS will search and boot an operating system contained in devices such as hard drives, CD/DVD, USB.
- Normally, the BIOS will check the floppy or CD-ROM drives to see if they can boot from them, then the hardware, the order of checking the drives depends on the configuration in the BIOS .
     + If the **BIOS** can't find it `boot device`  it will report `No boot device found`
     + If the operating system `Linux` is installed on the hard disk, it will find the Master Boot Record at the first sector of the first hard drive.

## Stage 2: MBR loading
- MBR (Master Boot Record) is stored in the first sector of a data storage device ( `/dev/sda`, `/dev/hda`)
- MBR is very small, only `512byte`=`1 sector`
- The MBR contains the following information:
     + Primary boot loader code (`446 byte`) : provides information for `boot loaderand` location `boot loaderand` the           hard drive.
     + Partition table information ( `64 byte`) : stores information about `partition`.
     + Magic number ( 2 byte) : used to check MBR , if MBR fails, it will restore.

## Stage 3: GRUB Loader
- After locating the Boot Loader , this step will execute `load` the Boot Loader into memory and read the configuration - information then display the GRUB Boot Menu for the user to choose from.
- If the User does not select the OS, after a preset time, GRUB will load the default Kernel into memory to boot.

## Stage 4: Kernel
- The operating system's kernel will be loaded into **RAM** -> When the kernel works, the first thing will execute the INIT process.

## Stage 5: Run Level (INIT)
- This is the main stage of the process `boot`.

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

- Run level0 : halt - Shutdown system
- Run level1 : single-user mode -No network configuration, start processes and allow loginuser non-root
- Run level2 : multi-user mode -no network configuration, start processes.
- Run level3 : multi-user mode with networking -boots the system normally on the command line interface.
- Run level4 : undefined
- Run level5 : X11- start the system on the graphical interface
- Run level6 : reboot- reboot the system.

##Stage 6: User Prompt
- User login and use.




# Shell
## Concepts
- **Shell** is a program that executes commands from users or applications - the utility requires passing to **Kernel** for processing.
- Activities of **Shell**:
  + 1: Command parsing.
  + 2: Interpreting the command's request.
  + 3: Pass message to **Kernel**.
  + 4: Display the result of the command.
    
### Types of **Shell**
- **sh(the Bourne Shell)**
   + Do not autocomplete file names
   + Do not save the typed statement in memory (`history`)

- **bash (Bourne-Again Shell)**
   + Linux defaults.
   + Support more commands.

- **csh (the C shell)**
   + Written in C . programming language.

### `Shell Prompt (Shell Prompt)`
- Is a character or a character set that always starts at the beginning of any command line.
- Indicates that **shell** is ready to receive commands from the user.
- `$`: ordinary user
- `#`: user `root`




# Things you should know about Linux Partitions
## 1.Hard drives
- As a data storage device, for example in Linux it `isdev/sda`

## 2. Partition
- Partition is a small partition (Logical Partition) divided fromHard drives
- Each Hard driveshas 4 partition, in which partition includes 3 types: Primary partition , Extended partition and Logical Partition
- Primary partition
     +The primary partition can be used to boot the operating system.
- Extended partition
     +Is the remaining data area when we divide into Primary partitions.
     +Extended partition contains Logical Partitions in it. Each drive can only hold 1 Extended partition.
     +Logical partition : These are small partitions in the extended partition , usually used to store data.
3. MBR and GPT
- Information about hard disk partitions will be stored on MBR (Master Boot record) or GPT (GUID Partition table).
- These are two standards for configuring and managing Partitions on hard drives.
- The information stored here includes the location and size of the Partitions.

## Swap file
- Swap is virtual memory concept used on Linux.
- When VPS/Serveroperating, if the system runs out of RAM, it will automatically use part of the hard drive as memory for applications to operate.
- Swap is slower than RAM because Swap is part of the hard drive.




# File Directory
- /home – User personal data
- /boot – Boot files
- /sbin – System binaries
- /dev – Device files
- /sys 
- /etc – Configuration files
- /proc – Process and kernel files
- /var – Variable data files
	-	/tmp – Temporary files
	-	/lock
	-	/log	
- /usr – User binaries and program data
	-	/bin – Binaries
	-	/man
	-	/lib – Shared libraries
	-	/local
	-	/share
- /opt – Optional software
- /root – The home directory of the root
- /media – Mount point for removable media
- /mnt – Mount directory
- /srv – Service data	




# common commands preview
- [View Linux command help information] - commands：`help`, `whatis`, `info`, `which`, `whereis`, `man`
- [Linux file directory management] -commands：`cd`, `ls`, `pwd`, `mkdir`, `rmdir`, `tree`, `touch`, `ln`, `rename`, `stat`, `file`, `chmod`, `chown`, `locate`, `find`, `cp`, `mv`, `rm`
- [Linux file content viewing command] - commands：`cat`, `head`, `tail`, `more`, `less`, `sed`, `vi`, `grep`
- [Linux File Compression and Decompression] -commands：`tar`, `gzip`, `zip`, `unzip`
- [Linux User Management] - commands：`groupadd`, `groupdel`, `groupmod`, `useradd`, `userdel`, `usermod`, `passwd`, `su`, `sudo`
- [Linux System Administration ] - commands：`reboot`, `exit`, `shutdown`, `date`, `mount`, `umount`, `ps`, `kill`, `systemctl`, `service`, `crontab`
- [Linux Network Administration] - commands：`curl`, `wget`, `telnet`, `ip`, `hostname`, `ifconfig`, `route`, `ssh`, `ssh-keygen`, `firewalld`, `iptables`, `host`, `nslookup`, `nc`/`netcat`, `ping`, `traceroute`, `netstat`
- [Linux Hardware Management ] - commands：`df`, `du`, `top`, `free`, `iotop`
- [Linux Software Management ] - commands：`rpm`, `yum`, `apt-get`

## all default Linux commands in bin dir
<details>
<summary>VIEW ALL</summary><br><b>
```
 envsubst                             lslogins                rgrep                              systemd-cryptenroll
 VGAuthService                        eqn                     lsmem                              rlogin                  systemd-delta
'['                                   ex                      lsmod                              rm                      systemd-detect-virt
 aa-enabled                           expand                  lsns                               rmdir                   systemd-escape
 aa-exec                              expiry                  lsof                               rnano                   systemd-hwdb
 aa-features-abi                      expr                    lspci                              routef                  systemd-id128
 add-apt-repository                   factor                  lspgpot                            routel                  systemd-inhibit
 addpart                              faillog                 lspower                            rrsync                  systemd-machine-id-setup
 addr2line                            fallocate               lsusb                              rsh                     systemd-mount
 apport-bug                           false                   lzcat                              rsync                   systemd-notify
 apport-cli                           fgconsole               lzcmp                              rsync-ssl               systemd-path
 apport-collect                       fgrep                   lzdiff                             rtstat                  systemd-repart
 apport-unpack                        file                    lzegrep                            run-one                 systemd-run
 apropos                              finalrd                 lzfgrep                            run-one-constantly      systemd-socket-activate
 apt                                  fincore                 lzgrep                             run-one-until-failure   systemd-stdio-bridge
 apt-add-repository                   find                    lzless                             run-one-until-success   systemd-sysext
 apt-cache                            findmnt                 lzma                               run-parts               systemd-sysusers
 apt-cdrom                            flock                   lzmainfo                           run-this-one            systemd-tmpfiles
 apt-config                           fmt                     lzmore                             runcon                  systemd-tty-ask-password-agent
 apt-extracttemplates                 fold                    mailmail3                          rview                   systemd-umount
 apt-ftparchive                       free                    man                                rvim                    tabs
 apt-get                              ftp                     man-recode                         savelog                 tac
 apt-key                              fuser                   mandb                              sbattach                tail
 apt-mark                             fusermount              manifest                           sbkeysync               tar
 apt-sortpkgs                         fusermount3             manpath                            sbsiglist               taskset
 ar                                   gapplication            mapscrn                            sbsign                  tbl
 arch                                 gawk                    mawk                               sbvarsign               tclsh
 as                                   gdbus                   mcookie                            sbverify                tclsh8.6
 automat-visualize3                   geqn                    md5sum                             scandeps                tcpdump
 awk                                  getconf                 md5sum.textutils                   scp                     tee
 b2sum                                getent                  mdig                               screen                  telnet
 base32                               getkeycodes             mesg                               screendump              telnet.netkit
 base64                               getopt                  migrate-pubring-from-classic-gpg   script                  tempfile
 basename                             gettext                 mk_modmap                          scriptlive              test
 basenc                               gettext.sh              mkdir                              scriptreplay            tic
 bash                                 ginstall-info           mkfifo                             scsi_logging_level      time
 bashbug                              gio                     mknod                              scsi_mandat             timedatectl
 bc                                   gio-querymodules        mksquashfs                         scsi_readcap            timeout
 boltctl                              git                     mktemp                             scsi_ready              tkconch3
 btrfs                                git-receive-pack        mokutil                            scsi_satl               tload
 btrfs-convert                        git-shell               more                               scsi_start              tmux
 btrfs-find-root                      git-upload-archive      mount                              scsi_stop               tnftp
 btrfs-image                          git-upload-pack         mountpoint                         scsi_temperature        toe
 btrfs-map-logical                    glib-compile-schemas    mt                                 sdiff                   top
 btrfs-select-super                   gold                    mt-gnu                             sed                     touch
 btrfsck                              gp-archive              mtr                                select-editor           tput
 btrfstune                            gp-collect-app          mtr-packet                         sensible-browser        tr
 busctl                               gp-display-html         mv                                 sensible-editor         tracepath
 busybox                              gp-display-src          namei                              sensible-pager          trial3
 byobu                                gp-display-text         nano                               seq                     troff
 byobu-config                         gpasswd                 nawk                               setarch                 true
 byobu-ctrl-a                         gpg                     nc                                 setfont                 truncate
 byobu-disable                        gpg-agent               nc.openbsd                         setkeycodes             tset
 byobu-disable-prompt                 gpg-connect-agent       neqn                               setleds                 tsort
 byobu-enable                         gpg-wks-server          netcat                             setlogcons              tty
 byobu-enable-prompt                  gpg-zip                 networkctl                         setmetamode             twist3
 byobu-export                         gpg2                    networkd-dispatcher                setpci                  twistd3
 byobu-janitor                        gpgcompose              newgrp                             setpriv                 tzselect
 byobu-keybindings                    gpgconf                 ngettext                           setsid                  ua
 byobu-launch                         gpgparsemail            nice                               setterm                 ubuntu-advantage
 byobu-launcher                       gpgsm                   nisdomainname                      setupcon                ubuntu-bug
 byobu-launcher-install               gpgsplit                nl                                 sftp                    ubuntu-core-launcher
 byobu-launcher-uninstall             gpgtar                  nm                                 sg                      ubuntu-distro-info
 byobu-layout                         gpgv                    nohup                              sg_bg_ctl               ubuntu-security-status
 byobu-prompt                         gpic                    nproc                              sg_compare_and_write    ucf
 byobu-quiet                          gprof                   nroff                              sg_copy_results         ucfq
 byobu-reconnect-sockets              gprofng                 nsenter                            sg_dd                   ucfr
 byobu-screen                         grep                    nslookup                           sg_decode_sense         uclampset
 byobu-select-backend                 gresource               nstat                              sg_emc_trespass         udevadm
 byobu-select-profile                 groff                   nsupdate                           sg_format               ul
 byobu-select-session                 grog                    ntfs-3g                            sg_get_config           umount
 byobu-shell                          grops                   ntfs-3g.probe                      sg_get_elem_status      uname
 byobu-silent                         grotty                  ntfscat                            sg_get_lba_status       unattended-upgrade
 byobu-status                         groups                  ntfscluster                        sg_ident                unattended-upgrades
 byobu-status-detail                  growpart                ntfscmp                            sg_inq                  uncompress
 byobu-tmux                           grub-editenv            ntfsdecrypt                        sg_logs                 unexpand
 byobu-ugraph                         grub-file               ntfsfallocate                      sg_luns                 unicode_start
 byobu-ulevel                         grub-fstest             ntfsfix                            sg_map                  unicode_stop
 c++filt                              grub-glue-efi           ntfsinfo                           sg_map26                uniq
 c_rehash                             grub-kbdcomp            ntfsls                             sg_modes                unlink
 captoinfo                            grub-menulst2cfg        ntfsmove                           sg_opcodes              unlzma
 cat                                  grub-mkfont             ntfsrecover                        sg_persist              unmkinitramfs
 catman                               grub-mkimage            ntfssecaudit                       sg_prevent              unshare
 cftp3                                grub-mklayout           ntfstruncate                       sg_raw                  unsquashfs
 chage                                grub-mknetdir           ntfsusermap                        sg_rbuf                 unxz
 chardet                              grub-mkpasswd-pbkdf2    ntfswipe                           sg_rdac                 unzstd
 chardetect                           grub-mkrelpath          numfmt                             sg_read                 update-alternatives
 chattr                               grub-mkrescue           objcopy                            sg_read_attr            update-mime-database
 chcon                                grub-mkstandalone       objdump                            sg_read_block_limits    uptime
 chfn                                 grub-mount              od                                 sg_read_buffer          usb-devices
 chgrp                                grub-ntldr-img          oem-getlogs                        sg_read_long            usbhid-dump
 chmod                                grub-render-label       on_ac_power                        sg_readcap              usbreset
 choom                                grub-script-check       openssl                            sg_reassign             users
 chown                                grub-syslinux2cfg       openvt                             sg_referrals            utmpdump
 chrt                                 gsettings               os-prober                          sg_rep_pip              uuidgen
 chsh                                 gtbl                    pager                              sg_rep_zones            uuidparse
 chvt                                 gunzip                  partx                              sg_requests             vcs-run
 ckbcomp                              gzexe                   passwd                             sg_reset                vdir
 ckeygen3                             gzip                    paste                              sg_reset_wp             vi
 cksum                                h2ph                    pastebinit                         sg_rmsn                 view
 clear                                h2xs                    patch                              sg_rtpg                 vigpg
 clear_console                        hardlink                pathchk                            sg_safte                vim
 cloud-id                             hd                      pbget                              sg_sanitize             vim.basic
 cloud-init                           head                    pbput                              sg_sat_identify         vim.tiny
 cloud-init-per                       helpztags               pbputs                             sg_sat_phy_event        vimdiff
 cmp                                  hexdump                 pdb3                               sg_sat_read_gplog       vimtutor
 codepage                             host                    pdb3.10                            sg_sat_set_features     vm-support
 col                                  hostid                  peekfd                             sg_scan                 vmhgfs-fuse
 col1                                 hostname                perl                               sg_seek                 vmstat
 col2                                 hostnamectl             perl5.34-x86_64-linux-gnu          sg_senddiag             vmtoolsd
 col3                                 htop                    perl5.34.0                         sg_ses                  vmware-alias-import
 col4                                 hwe-support-status      perlbug                            sg_ses_microcode        vmware-checkvm
 col5                                 i386                    perldoc                            sg_start                vmware-hgfsclient
 col6                                 iconv                   perlivp                            sg_stpg                 vmware-namespace-cmd
 col7                                 id                      perlthanks                         sg_stream_ctl           vmware-rpctool
 col8                                 info                    pgrep                              sg_sync                 vmware-toolbox-cmd
 col9                                 infobrowser             pic                                sg_test_rwbuf           vmware-vgauth-cmd
 colcrt                               infocmp                 pico                               sg_timestamp            vmware-vmblock-fuse
 colrm                                infotocap               piconv                             sg_turs                 vmware-xferlogs
 column                               install                 pidof                              sg_unmap                w
 comm                                 install-info            pidwait                            sg_verify               wall
 conch3                               instmodsh               pinentry                           sg_vpd                  watch
 corelist                             ionice                  pinentry-curses                    sg_wr_mode              watchgnupg
 cp                                   ip                      ping                               sg_write_buffer         wc
 cpan                                 ipcmk                   ping4                              sg_write_long           wdctl
 cpan5.34-x86_64-linux-gnu            ipcrm                   ping6                              sg_write_same           wget
 cpio                                 ipcs                    pinky                              sg_write_verify         whatis
 crontab                              iptables-xml            pkaction                           sg_write_x              whereis
 csplit                               ischroot                pkcheck                            sg_xcopy                which
 ctail                                iscsiadm                pkcon                              sg_zone                 which.debianutils
 ctstat                               join                    pkexec                             sginfo                  whiptail
 curl                                 journalctl              pkill                              sgm_dd                  who
 cut                                  jq                      pkmon                              sgp_dd                  whoami
 cvtsudoers                           json-patch-jsondiff     pkttyagent                         sh                      wifi-status
 dash                                 json_pp                 pl2pm                              sha1sum                 write
 date                                 jsondiff                pldd                               sha224sum               x86_64
 dbus-cleanup-sockets                 jsonpatch               plymouth                           sha256sum               x86_64-linux-gnu-addr2line
 dbus-daemon                          jsonpointer             pmap                               sha384sum               x86_64-linux-gnu-ar
 dbus-monitor                         jsonschema              pod2html                           sha512sum               x86_64-linux-gnu-as
 dbus-run-session                     kbd_mode                pod2man                            shasum                  x86_64-linux-gnu-c++filt
 dbus-send                            kbdinfo                 pod2text                           showconsolefont         x86_64-linux-gnu-dwp
 dbus-update-activation-environment   kbxutil                 pod2usage                          showkey                 x86_64-linux-gnu-elfedit
 dbus-uuidgen                         keep-one-running        podchecker                         shred                   x86_64-linux-gnu-gold
 dd                                   kernel-install          pollinate                          shuf                    x86_64-linux-gnu-gp-archive
 deallocvt                            keyring                 pr                                 size                    x86_64-linux-gnu-gp-collect-app
 deb-systemd-helper                   kill                    preconv                            skill                   x86_64-linux-gnu-gp-display-html
 deb-systemd-invoke                   killall                 printenv                           slabtop                 x86_64-linux-gnu-gp-display-src
 debconf                              kmod                    printf                             sleep                   x86_64-linux-gnu-gp-display-text
 debconf-apt-progress                 kmodsign                prlimit                            slogin                  x86_64-linux-gnu-gprof
 debconf-communicate                  landscape-sysinfo       pro                                snap                    x86_64-linux-gnu-gprofng
 debconf-copydb                       last                    prove                              snapctl                 x86_64-linux-gnu-ld
 debconf-escape                       lastb                   prtstat                            snapfuse                x86_64-linux-gnu-ld.bfd
 debconf-set-selections               lastlog                 ps                                 snice                   x86_64-linux-gnu-ld.gold
 debconf-show                         lcf                     psfaddtable                        soelim                  x86_64-linux-gnu-nm
 debian-distro-info                   ld                      psfgettable                        sort                    x86_64-linux-gnu-objcopy
 delpart                              ld.bfd                  psfstriptable                      sos                     x86_64-linux-gnu-objdump
 delv                                 ld.gold                 psfxtable                          sos-collector           x86_64-linux-gnu-ranlib
 df                                   ldd                     pslog                              sosreport               x86_64-linux-gnu-readelf
 dh_bash-completion                   less                    pstree                             splain                  x86_64-linux-gnu-size
 diff                                 lessecho                pstree.x11                         split                   x86_64-linux-gnu-strings
 diff3                                lessfile                ptar                               splitfont               x86_64-linux-gnu-strip
 dig                                  lesskey                 ptardiff                           sqfscat                 xargs
 dir                                  lesspipe                ptargrep                           sqfstar                 xauth
 dircolors                            lexgrog                 ptx                                ss                      xdg-user-dir
 dirmngr                              libnetcfg               purge-old-kernels                  ssh                     xdg-user-dirs-update
 dirmngr-client                       link                    pwd                                ssh-add                 xsubpp
 dirname                              linux-boot-prober       pwdx                               ssh-agent               xxd
 distro-info                          linux-check-removal     py3clean                           ssh-argv0               xz
 dmesg                                linux-update-symlinks   py3compile                         ssh-copy-id             xzcat
 dnsdomainname                        linux-version           py3versions                        ssh-import-id           xzcmp
 do-release-upgrade                   linux32                 pybabel                            ssh-import-id-gh        xzdiff
 domainname                           linux64                 pybabel-python3                    ssh-import-id-lp        xzegrep
 dpkg                                 ln                      pydoc3                             ssh-keygen              xzfgrep
 dpkg-deb                             lnstat                  pydoc3.10                          ssh-keyscan             xzgrep
 dpkg-divert                          loadkeys                pygettext3                         stat                    xzless
 dpkg-maintscript-helper              loadunimap              pygettext3.10                      static-sh               xzmore
 dpkg-query                           locale                  pyhtmlizer3                        stdbuf                  yes
 dpkg-realpath                        locale-check            pyserial-miniterm                  strace                  ypdomainname
 dpkg-split                           localectl               pyserial-ports                     strace-log-merge        zcat
 dpkg-statoverride                    localedef               python3                            streamzip               zcmp
 dpkg-trigger                         logger                  python3.10                         strings                 zdiff
 du                                   login                   pzstd                              strip                   zdump
 dumpkeys                             loginctl                ranlib                             stty                    zegrep
 dwp                                  logname                 rbash                              su                      zfgrep
 eatmydata                            look                    rcp                                sudo                    zforce
 ec2metadata                          lowntfs-3g              rdma                               sudoedit                zgrep
 echo                                 ls                      readelf                            sudoreplay              zipdetails
 ed                                   lsattr                  readlink                           sum                     zless
 editor                               lsb_release             realpath                           sync                    zmore
 efibootdump                          lsblk                   red                                systemctl               znew
 efibootmgr                           lscpu                   renice                             systemd                 zstd
 egrep                                lsfd                    rescan-scsi-bus.sh                 systemd-analyze         zstdcat
 eject                                lshw                    reset                              systemd-ask-password    zstdgrep
 elfedit                              lsinitramfs             resizecons                         systemd-cat             zstdless
 enc2xs                               lsipc                   resizepart                         systemd-cgls            zstdmt
 encguess                             lsirq                   resolvectl                         systemd-cgtop
 env                                  lslocks                 rev                                systemd-creds
```
</b>
</details>





























































































