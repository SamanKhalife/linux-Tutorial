## Topic 200: Capacity Planning
### 200.1 Measure and Troubleshoot Resource Usage (weight: 6)
#### Weight	6
#### Description	Candidates should be able to measure hardware resource and network bandwidth, identify and troubleshoot resource problems.
Key Knowledge Areas:

Measure CPU usage.
Measure memory usage.
Measure disk I/O.
Measure network I/O.
Measure firewalling and routing throughput.
Map client bandwidth usage.
Match / correlate system symptoms with likely problems.
Estimate throughput and identify bottlenecks in a system including networking.
The following is a partial list of the used files, terms and utilities:

iostat
iotop
vmstat
netstat
ss
iptraf
pstree, ps
w
lsof
top
htop
uptime
sar
swap
processes blocked on I/O
blocks in
blocks out
 

### 200.2 Predict Future Resource Needs (weight: 2)
#### Weight	2
#### Description	Candidates should be able to monitor resource usage to predict future resource needs.
Key Knowledge Areas:

Use monitoring and measurement tools to monitor IT infrastructure usage.
Predict capacity break point of a configuration.
Observe growth rate of capacity usage.
Graph the trend of capacity usage.
Awareness of monitoring solutions such as Icinga2, Nagios, collectd, MRTG and Cacti
The following is a partial list of the used files, terms and utilities:

diagnose
predict growth
resource exhaustion
 

## Topic 201: Linux Kernel
### 201.1 Kernel components (weight: 2)
#### Weight	2
#### Description	Candidates should be able to utilise kernel components that are necessary to specific hardware, hardware drivers, system resources and requirements. This objective includes implementing different types of kernel images, understanding stable and longterm kernels and patches, as well as using kernel modules.
Key Knowledge Areas:

Kernel 2.6.x, 3.x and 4.x documentation
The following is a partial list of the used files, terms and utilities:

/usr/src/linux/
/usr/src/linux/Documentation/
zImage
bzImage
xz compression
 

### 201.2 Compiling a Linux kernel (weight: 3)
#### Weight	3
#### Description	Candidates should be able to properly configure a kernel to include or disable specific features of the Linux kernel as necessary. This objective includes compiling and recompiling the Linux kernel as needed, updating and noting changes in a new kernel, creating an initrd image and installing new kernels.
Key Knowledge Areas:

/usr/src/linux/
Kernel Makefiles
Kernel 2.6.x, 3.x and 4.x make targets
Customize the current kernel configuration.
Build a new kernel and appropriate kernel modules.
Install a new kernel and any modules.
Ensure that the boot manager can locate the new kernel and associated files.
Module configuration files
Use DKMS to compile kernel modules.
Awareness of dracut
The following is a partial list of the used files, terms and utilities:

mkinitrd
mkinitramfs
make
make targets (all, config, xconfig, menuconfig, gconfig, oldconfig, mrproper, zImage, bzImage, modules, modules_install, rpm-pkg, binrpm-pkg, deb-pkg)
gzip
bzip2
module tools
/usr/src/linux/.config
/lib/modules/kernel-version/
depmod
dkms
 

### 201.3 Kernel runtime management and troubleshooting (weight: 4)
#### Weight	4
#### Description	Candidates should be able to manage and/or query a 2.6.x, 3.x or 4.x kernel and its loadable modules. Candidates should be able to identify and correct common boot and run time issues. Candidates should understand device detection and management using udev. This objective includes troubleshooting udev rules.
Key Knowledge Areas:

Use command-line utilities to get information about the currently running kernel and kernel modules.
Manually load and unload kernel modules.
Determine when modules can be unloaded.
Determine what parameters a module accepts.
Configure the system to load modules by names other than their file name.
/proc filesystem
Content of /, /boot/ , and /lib/modules/
Tools and utilities to analyse information about the available hardware
udev rules
The following is a partial list of the used files, terms and utilities:

/lib/modules/kernel-version/modules.dep
module configuration files in /etc/
/proc/sys/kernel/
/sbin/depmod
/sbin/rmmod
/sbin/modinfo
/bin/dmesg
/sbin/lspci
/usr/bin/lsdev
/sbin/lsmod
/sbin/modprobe
/sbin/insmod
/bin/uname
/usr/bin/lsusb
/etc/sysctl.conf, /etc/sysctl.d/
/sbin/sysctl
udevmonitor
udevadm monitor
/etc/udev/
 

## Topic 202: System Startup
### 202.1 Customising system startup (weight: 3)
#### Weight	3
#### Description	Candidates should be able to query and modify the behaviour of system services at various targets / run levels. A thorough understanding of the systemd, SysV Init and the Linux boot process is required. This objective includes interacting with systemd targets and SysV init run levels.
Key Knowledge Areas:

Systemd
SysV init
Linux Standard Base Specification (LSB)
The following is a partial list of the used files, terms and utilities:

/usr/lib/systemd/
/etc/systemd/
/run/systemd/
systemctl
systemd-delta
/etc/inittab
/etc/init.d/
/etc/rc.d/
chkconfig
update-rc.d
init and telinit
 

### 202.2 System recovery (weight: 4)
#### Weight	4
#### Description	Candidates should be able to properly manipulate a Linux system during both the boot process and during recovery mode. This objective includes using both the init utility and init-related kernel options. Candidates should be able to determine the cause of errors in loading and usage of bootloaders. GRUB version 2 and GRUB Legacy are the bootloaders of interest. Both BIOS and UEFI systems are covered.
Key Knowledge Areas:

BIOS and UEFI
NVMe booting
GRUB version 2 and Legacy
grub shell
boot loader start and hand off to kernel
kernel loading
hardware initialisation and setup
daemon/service initialisation and setup
Know the different boot loader install locations on a hard disk or removable device.
Overwrite standard boot loader options and using boot loader shells.
Use systemd rescue and emergency modes.
The following is a partial list of the used files, terms and utilities:

mount
fsck
inittab, telinit and init with SysV init
The contents of /boot/, /boot/grub/ and /boot/efi/
EFI System Partition (ESP)
GRUB
grub-install
efibootmgr
UEFI shell
initrd, initramfs
Master boot record
systemctl
 

### 202.3 Alternate Bootloaders (weight: 2)
#### Weight	2
#### Description	Candidates should be aware of other bootloaders and their major features.
Key Knowledge Areas:

SYSLINUX, ISOLINUX, PXELINUX
Understanding of PXE for both BIOS and UEFI
Awareness of systemd-boot and U-Boot
The following is a partial list of the used files, terms and utilities:

syslinux
extlinux
isolinux.bin
isolinux.cfg
isohdpfx.bin
efiboot.img
pxelinux.0
pxelinux.cfg/
uefi/shim.efi
uefi/grubx64.efi
 

## Topic 203: Filesystem and Devices
### 203.1 Operating the Linux filesystem (weight: 4)
#### Weight	4
#### Description	Candidates should be able to properly configure and navigate the standard Linux filesystem. This objective includes configuring and mounting various filesystem types.
Key Knowledge Areas:

The concept of the fstab configuration
Tools and utilities for handling swap partitions and files
Use of UUIDs for identifying and mounting file systems
Understanding of systemd mount units
The following is a partial list of the used files, terms and utilities:

/etc/fstab
/etc/mtab
/proc/mounts
mount and umount
blkid
sync
swapon
swapoff
 

### 203.2 Maintaining a Linux filesystem (weight: 3)
#### Weight	3
#### Description	Candidates should be able to properly maintain a Linux filesystem using system utilities. This objective includes manipulating standard filesystems and monitoring SMART devices.
Key Knowledge Areas:

Tools and utilities to manipulate and ext2, ext3 and ext4
Tools and utilities to perform basic Btrfs operations, including subvolumes and snapshots
Tools and utilities to manipulate XFS
Awareness of ZFS
The following is a partial list of the used files, terms and utilities:

mkfs (mkfs.*)
mkswap
fsck (fsck.*)
tune2fs, dumpe2fs and debugfs
btrfs, btrfs-convert
xfs_info, xfs_check, xfs_repair, xfsdump and xfsrestore
smartd, smartctl
 

### 203.3 Creating and configuring filesystem options (weight: 2)
#### Weight	2
#### Description	Candidates should be able to configure automount filesystems using AutoFS. This objective includes configuring automount for network and device filesystems. Also included is creating filesystems for devices such as CD-ROMs and a basic feature knowledge of encrypted filesystems.
Key Knowledge Areas:

autofs configuration files
Understanding of automount units
UDF and ISO9660 tools and utilities
Awareness of other CD-ROM filesystems (HFS)
Awareness of CD-ROM filesystem extensions (Joliet, Rock Ridge, El Torito)
Basic feature knowledge of data encryption (dm-crypt / LUKS)
The following is a partial list of the used files, terms and utilities:

/etc/auto.master
/etc/auto.[dir]
mkisofs
cryptsetup
 

## Topic 204: Advanced Storage Device Administration
### 204.1 Configuring RAID (weight: 3)
#### Weight	3
#### Description	Candidates should be able to configure and implement software RAID. This objective includes using and configuring RAID 0, 1 and 5.
Key Knowledge Areas:

Software RAID configuration files and utilities
The following is a partial list of the used files, terms and utilities:

mdadm.conf
mdadm
/proc/mdstat
partition type 0xFD
 

### 204.2 Adjusting Storage Device Access (weight: 2)
#### Weight	2
#### Description	Candidates should be able to configure kernel options to support various drives. This objective includes software tools to view & modify hard disk settings including iSCSI devices.
Key Knowledge Areas:

Tools and utilities to configure DMA for IDE devices including ATAPI and SATA
Tools and utilities to configure Solid State Drives including AHCI and NVMe
Tools and utilities to manipulate or analyse system resources (e.g. interrupts)
Awareness of sdparm command and its uses
Tools and utilities for iSCSI
Awareness of SAN, including relevant protocols (AoE, FCoE)
The following is a partial list of the used files, terms and utilities:

hdparm, sdparm
nvme
tune2fs
fstrim
sysctl
/dev/hd*, /dev/sd*, /dev/nvme*
iscsiadm, scsi_id, iscsid and iscsid.conf
WWID, WWN, LUN numbers
 

### 204.3 Logical Volume Manager (weight: 3)
#### Weight	3
#### Description	Candidates should be able to create and remove logical volumes, volume groups, and physical volumes. This objective includes snapshots and resizing logical volumes.
Key Knowledge Areas:

Tools in the LVM suite
Resizing, renaming, creating, and removing logical volumes, volume groups, and physical volumes
Creating and maintaining snapshots
Activating volume groups
The following is a partial list of the used files, terms and utilities:

/sbin/pv*
/sbin/lv*
/sbin/vg*
mount
/dev/mapper/
lvm.conf
 

## Topic 205: Networking Configuration
### 205.1 Basic networking configuration (weight: 3)
#### Weight	3
#### Description	Candidates should be able to configure a network device to be able to connect to a local, wired or wireless, and a wide-area network. This objective includes being able to communicate between various subnets within a single network including both IPv4 and IPv6 networks.
Key Knowledge Areas:

Utilities to configure and manipulate ethernet network interfaces
Configuring basic access to wireless networks
The following is a partial list of the used files, terms and utilities:

ip
ifconfig
route
arp
iw
iwconfig
iwlist
 

### 205.2 Advanced Network Configuration (weight: 4)
#### Weight	4
#### Description	Candidates should be able to configure a network device to implement various network authentication schemes. This objective includes configuring a multi-homed network device and resolving communication problems.
Key Knowledge Areas:

Utilities to manipulate routing tables
Utilities to configure and manipulate ethernet network interfaces
Utilities to analyse the status of the network devices
Utilities to monitor and analyse the TCP/IP traffic
The following is a partial list of the used files, terms and utilities:

ip
ifconfig
route
arp
ss
netstat
lsof
ping, ping6
nc
tcpdump
nmap
 

### 205.3 Troubleshooting network issues (weight: 4)
#### Weight	4
#### Description	Candidates should be able to identify and correct common network setup issues, to include knowledge of locations for basic configuration files and commands.
Key Knowledge Areas:

Location and content of access restriction files
Utilities to configure and manipulate ethernet network interfaces
Utilities to manage routing tables
Utilities to list network states.
Utilities to gain information about the network configuration
Methods of information about the recognised and used hardware devices
System initialisation files and their contents (Systemd and SysV init)
Awareness of NetworkManager and its impact on network configuration
The following is a partial list of the used files, terms and utilities:

ip
ifconfig
route
ss
netstat
/etc/network/, /etc/sysconfig/network-scripts/
ping, ping6
traceroute, traceroute6
mtr
hostname
System log files such as /var/log/syslog, /var/log/messages and the systemd journal
dmesg
/etc/resolv.conf
/etc/hosts
/etc/hostname, /etc/HOSTNAME
/etc/hosts.allow, /etc/hosts.deny
 

## Topic 206: System Maintenance
### 206.1 Make and install programs from source (weight: 2)
#### Weight	2
#### Description	Candidates should be able to build and install an executable program from source. This objective includes being able to unpack a file of sources.
Key Knowledge Areas:

Unpack source code using common compression and archive utilities.
Understand basics of invoking make to compile programs.
Apply parameters to a configure script.
Know where sources are stored by default.
The following is a partial list of the used files, terms and utilities:

/usr/src/
gunzip
gzip
bzip2
xz
tar
configure
make
uname
install
patch
 

### 206.2 Backup operations (weight: 3)
#### Weight	3
#### Description	Candidates should be able to use system tools to back up important system data.
Key Knowledge Areas:

Knowledge about directories that have to be included in backups
