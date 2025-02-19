# linux Tutorial
**Learning Linux**: We started from scratch to build a comprehensive tutorial covering everything from essential Linux commands and tools to the structure of the Linux kernel, including all the relevant details. Although our primary focus is not on the LPIC exams, we've included each objective along with explanations, a partial list of commands, and even sample exam questions ([**lpic101**](https://github.com/SamanKhalife/linux-Tutorial/tree/main/Lpic%201#lpic-1-exam-101-objectives), [**lpic102**](https://github.com/SamanKhalife/linux-Tutorial/tree/main/Lpic%201#lpic-1-exam-102-objectives), [**lpic201**](https://github.com/SamanKhalife/linux-Tutorial/tree/main/Lpic%202#lpic-2-exam-201-objectives), [**lpic202**](https://github.com/SamanKhalife/linux-Tutorial/tree/main/Lpic%202#lpic-2-exam-202-objectives), [**lpic300**](https://github.com/SamanKhalife/linux-Tutorial/blob/main/Lpic%203/README.md#lpic-3-mixed-environments-300-objectives), [**lpic303**](https://github.com/SamanKhalife/linux-Tutorial/blob/main/Lpic%203/README.md#lpic-3-security-exam-303-objectives), [**lpic305**](https://github.com/SamanKhalife/linux-Tutorial/blob/main/Lpic%203/README.md#lpic-3-virtualization-and-containerization-305-objectives), [**lpic306**](https://github.com/SamanKhalife/linux-Tutorial/blob/main/Lpic%203/README.md#lpic-3-high-availability-and-storage-clusters-306-objectives)).

Feel free to contribute ([lets-fork-it](https://github.com/login?return_to=%2FSamanKhalife%2Flinux-Tutorial)), and if there's any problem or recommendation, I'd be happy to sort it out if you raise an [issue](https://github.com/SamanKhalife/linux-Tutorial/issues/new/choose).

# Repo Contents (More Than You Imagine 😜)
* [**Complete List Of Commands**](https://github.com/SamanKhalife/linux-commands/blob/main/Commands.md)
* LPIC-Exams: Objectives-and-Samples:
  * [LPIC1](https://github.com/SamanKhalife/linux-Tutorial/tree/main/Lpic%201)
  * [LPIC2](https://github.com/SamanKhalife/linux-Tutorial/tree/main/Lpic%202)
  * [LPIC3](https://github.com/SamanKhalife/linux-Tutorial/tree/main/Lpic%203)
* [Quick review of important commants](https://github.com/SamanKhalife/linux-Tutorial/blob/main/Quick%20review%20for%20commands.md)
* [A Handy Guide for Installing Some Useful Tools and Apps](https://github.com/SamanKhalife/linux-Tutorial/tree/main/Installing%20tools%20and%20apps)
* [Diving Deep into Linux](https://github.com/SamanKhalife/linux-Tutorial?tab=readme-ov-file#common-commands-preview)
* [Working with Cli Editors](https://github.com/SamanKhalife/linux-Tutorial/tree/main/Commadline-Editors)
* [Working with Git Cli](https://github.com/SamanKhalife/linux-Tutorial/tree/main/Git%20Commands)
* [How to Use Tools in Iran(DNS Settings)](https://github.com/SamanKhalife/linux-Tutorial/tree/main?tab=readme-ov-file#if-you-live-in-countries-that-are-embargoed-by-some-companies-and-applications-you-can-use-some-thing-like-dns-changing-for-using-the-apps)
* [Secure Shell(ssh)](https://github.com/SamanKhalife/linux-Tutorial/blob/main/ssh.md)

# Advanced Mode (On)
* [Common Troubleshooting Scenarios](https://www.hipeople.io/interview-questions/linux-troubleshooting-interview-questions)
* [Shell Scripting](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/readme/shell-scripting.md)
* [Linux File Systems and Storage Management](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/readme/file-systems-storage-management.md)
* [Linux Networking](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/readme/linux-networking.md)
* [Ansible Automation](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/readme/ansible-automation.md)
* [eBPF,XDP,DPDK Resources](https://github.com/SamanKhalife/X-Defender/blob/main/docs/resources-xdp-ebpf.md)
  * [Awesome eBPF](https://github.com/zoidyzoidzoid/awesome-ebpf)
  * [eBPF Tools](https://www.brendangregg.com/ebpf.html)
* [Linux Kernel Customization and Compilation](https://docs.rockylinux.org/guides/custom-linux-kernel/)
* [Virtualization in linux](https://github.com/SamanKhalife/linux-Tutorial/blob/main/virtualization-in-linux.md)
* [Backup and Recovery](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/readme/backup-and-recovery.md)
* [Cloud Delivery Models and SLA,SLI,SLO](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/readme/SRE.md)

## Common commands preview

* \[View Linux command help information] - commands：`help`, `whatis`, `info`, `which`, `whereis`, `man`, `stat`
* \[Linux file directory management] - commands：`cd`, `ls`, `pwd`, `mkdir`, `rmdir`, `tree`, `touch`, `ln`, `rename`, `stat`, `file`, `chmod`, `chown`, `locate`, `find`, `cp`, `mv`, `rm`
* \[Linux file content viewing command] - commands：`cat`, `head`, `tail`, `more`, `less`, `sed`, `vi`, `grep`
* \[Linux File Compression and Decompression] - commands：`tar`, `gzip`, `zip`, `unzip`
* \[Linux User Management] - commands：`groupadd`, `groupdel`, `groupmod`, `useradd`, `userdel`, `usermod`, `passwd`, `su`, `sudo`
* \[Linux System Administration ] - commands：`reboot`, `exit`, `shutdown`, `date`, `mount`, `umount`, `ps`, `kill`, `systemctl`, `service`, `crontab`
* \[Linux Network Administration] - commands：`curl`, `wget`, `telnet`, `ip`, `hostname`, `ifconfig`, `route`, `ssh`, `ssh-keygen`, `firewalld`, `iptables`, `host`, `nslookup`, `nc`, `netcat`, `ping`, `traceroute`, `netstat`
* \[Linux Hardware Management ] - commands：`df`, `du`, `top`, `free`, `iotop`
* \[Linux Software Management ] - commands：`rpm`, `yum`, `apt-get`, `dpkg`


"The art of command line" is all about getting good at using the command line interface (CLI) in Unix-like systems, like Linux. If you want to learn more, check out this [Link](https://github.com/jlevy/the-art-of-command-line)

## Getting Started With linux Overview (Hardware Interaction)

* **Kernel operating system**
* **Drivers**
* **System software**
* **Application software**
* **Application software with graphical interface**
<p align="center">
  <img src="https://github.com/SamanKhalife/linux-commands/blob/main/IMAGES/1111111111111111111.png" alt="Image description" width="700" style="max-width: 50%;">
</p>

* **Application software with graphical interface**
<p align="center">
  <img src="https://github.com/SamanKhalife/linux-commands/blob/main/IMAGES/2222222222222222222.png" alt="Image description" width="700" style="max-width: 100%;">
</p>

## More info about Linux kernel 

[Linux kernel Map](https://makelinux.github.io/kernel/map/)

## Linux Boot Process

* There are 6 stages of the Linux boot process

<p align="center">
  <img src="https://github.com/SamanKhalife/linux-commands/blob/main/IMAGES/3333333333333333333.png" alt="Image description" width="700" style="max-width: 100%;">
</p>

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

[gorupmod](/TXT%20FILES/gorupmod.md)

[groupadd](/TXT%20FILES/groupadd.md)

[groupdel](/TXT%20FILES/groupdel.md)

[groups](/TXT%20FILES/groups.md)

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

[wget command](/TXT%20FILES/wget.md)

Installing wget:

```
sudo apt install wget
```

Download a file or webpage using wget:

```
wget URL
```

2.Download files from Linux command line using curl

[curl command](/TXT%20FILES/curl.md)

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

[ping](/TXT%20FILES/ping.md)

[host](/TXT%20FILES/host.md)

[finger](/TXT%20FILES/finger.md)

[traceroute](/TXT%20FILES/traceroute.md)

[netstat](/TXT%20FILES/netstat.md)

[tracepath](/TXT%20FILES/tracepath.md)

[dig](/TXT%20FILES/dig.md)

[hostname](/TXT%20FILES/hostname.md)

[route](/TXT%20FILES/route.md)

[nslookup](/TXT%20FILES/nslookup.md)

## Firewalls

Here are the Best Firewalls for Linux for Effective System Protection(We Mostly use Iptables ,(ufw for regular usages))

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
<p align="center">
  <img src="https://github.com/SamanKhalife/linux-commands/blob/main/IMAGES/4444444444444444444.png" alt="Image description" width="700" style="max-width: 100%;">
</p>

### Packet Flow in Netfilter and General Networking
![Netfilter Packet Flow](https://upload.wikimedia.org/wikipedia/commons/3/37/Netfilter-packet-flow.svg)

### If you live in countries that are embargoed by some companies and applications, you can use some thing like DNS changing for using the apps.

* [403.online](https://403.online/)
* [shecan](https://shecan.ir/)

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

* [certbot](https://github.com/certbot/certbot)
* [acme](https://github.com/acmesh-official/acme.sh)
* [zerossl](https://github.com/zerossl)
* [free-ssl](https://github.com/topics/free-ssl-certificates)

[learn getting free ssl with acme step by step](acme-freessl.md)

## Some of the monitoring tools (Cli) 

#### Viewing Running Processes in Linux

* [top](</TXT%20FILES/top.md>)
* [htop](</TXT%20FILES/htop.md>)
* [ps](</TXT%20FILES/ps.md>)
* [lsof](</TXT%20FILES/lsof.md>)

#### Monitor Network

* [nethogs](<README (1).md>)
* [iptraf-ng](<README (1).md>)
* [netstat](</TXT%20FILES/netstat.md>)
* [iftop](</TXT%20FILES/iftop.md>)
* [(speed test)](<README (1).md>)

#### Monitor Disk Usage

* [df](</TXT%20FILES/df.md>)
* [du](</TXT%20FILES/du.md>)
* [iotop](</TXT%20FILES/iotop.md>)
* [iostat](</TXT%20FILES/iostat.md>)

#### Monitor Memory Usage

* [free(Ram usage)](</TXT%20FILES/free.md>)
* [vmstat](</TXT%20FILES/vmstat.md>)

## Managing Processes in Linux

* [kill](</TXT%20FILES/kill.md>)
* [nice](</TXT%20FILES/nice.md>)

mosltly use for controlling the cpu usage (nice number is from -20 to 19 ).

## How To Add Swap Space

What is Swap ? Swap is a portion of hard drive storage that has been set aside for the operating system to temporarily store data that it can no longer hold in RAM. This lets you increase the amount of information that your server can keep in its working memory, with some caveats.

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
<p align="center">
  <img src="https://github.com/SamanKhalife/linux-commands/blob/main/IMAGES/5555555555555555555.png" alt="Image description" width="700" style="max-width: 100%;">
</p>


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



