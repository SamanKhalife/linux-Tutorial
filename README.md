# linux-Tutorial
linux commands and any thing basic you need yo know for using linux it is for my own Knowledge but if you what to use I try to make it very complex and from base so you can learn linux so easy.

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




# Linux Boot Process
- There are 6 stages of the Linux boot process
<!-- PIC33333333333  -->

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






























































































