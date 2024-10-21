## Topic 101: System Architecture

### 101.1 Determine and configure hardware settings

#### Weight	2

#### Description	Candidates should be able to determine and configure fundamental system hardware.
Key Knowledge Areas:

Enable and disable integrated peripherals.
Differentiate between the various types of mass storage devices.
Determine hardware resources for devices.
Tools and utilities to list various hardware information (e.g. lsusb, lspci, etc.).
Tools and utilities to manipulate USB devices.
Conceptual understanding of sysfs, udev and dbus.
The following is a partial list of the used files, terms and utilities:

* [/sys/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/-sys-.md)

* [/proc/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/-proc-.md)

* [/dev/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/-dev-.md)

* [modprobe](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/modprobe.md)

* [lsmod](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/lsmod.md)

* [lspci](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/lspci.md)

* [lsusb](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/lsusb.md)
 

### 101.2 Boot the system

#### Weight	3

#### Description Candidates should be able to guide the system through the booting process.
Key Knowledge Areas:

Provide common commands to the boot loader and options to the kernel at boot time.
Demonstrate knowledge of the boot sequence from BIOS/UEFI to boot completion.
Understanding of SysVinit and systemd.
Awareness of Upstart.
Check boot events in the log files.
The following is a partial list of the used files, terms and utilities:

* [dmesg](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/dmesg.md)

* [journalctl](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/journalctl.md)

* [BIOS](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/BIOS.md)

* [UEFI](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/UEFI.md)

* [bootloader](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/bootloader.md)

* [kernel](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/kernel.md)

* [initramfs](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/initramfs.md)

* [init](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/init.md)

* [SysVinit](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/SysVinit.md)

* [systemd](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/systemd.md)
 

### 101.3 Change runlevels / boot targets and shutdown or reboot system

##### Weight	3

#### Description	Candidates should be able to manage the SysVinit runlevel or systemd boot target of the system. This objective includes changing to single user mode, shutdown or rebooting the system. Candidates should be able to alert users before switching runlevels / boot targets and properly terminate processes. This objective also includes setting the default SysVinit runlevel or systemd boot target. It also includes awareness of Upstart as an alternative to SysVinit or systemd.
Key Knowledge Areas:

Set the default runlevel or boot target.
Change between runlevels / boot targets including single user mode.
Shutdown and reboot from the command line.
Alert users before switching runlevels / boot targets or other major system events.
Properly terminate processes.
Awareness of acpid.
The following is a partial list of the used files, terms and utilities:

* [/etc/inittab](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-102/-etc-inittab.md)

* [shutdown](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/shutdown.md)

* [init](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/init.md)

* [/etc/init.d/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/-etc-init.d-.md)

* [telinit](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/telinit.md)

* [systemd](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/systemd.md)

* [systemctl](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/systemctl.md)

* [/etc/systemd/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/-etc-systemd-.md)

* [/usr/lib/systemd/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/-user-lib-systemd-.md)

* [wall](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/wall.md)
 

## Topic 102: Linux Installation and Package Management

### 102.1 Design hard disk layout

#### Weight	2

#### Description	Candidates should be able to design a disk partitioning scheme for a Linux system.
Key Knowledge Areas:

Allocate filesystems and swap space to separate partitions or disks.
Tailor the design to the intended use of the system.
Ensure the /boot partition conforms to the hardware architecture requirements for booting.
Knowledge of basic features of LVM.
The following is a partial list of the used files, terms and utilities:

* [/ (root) filesystem  ](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/-root.md)

* [/var filesystem](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/-var.md)

* [/home filesystem](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/-home.md)

* [/boot filesystem](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/-boot.md)

* [EFI System Partition (ESP)](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/EFI%20System%20Partition.md)

* [swap space](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/swap%20space.md)

* [mount points](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/mount%20points.md)

* [partitions](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/partitions.md)
 

### 102.2 Install a boot manager

#### Weight	2

#### Description	Candidates should be able to select, install and configure a boot manager.
Key Knowledge Areas:

Providing alternative boot locations and backup boot options.
Install and configure a boot loader such as GRUB Legacy.
Perform basic configuration changes for GRUB 2.
Interact with the boot loader.
The following is a partial list of the used files, terms and utilities:

* [menu.lst, grub.cfg and grub.conf](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/menu.lst%2C%20grub.cfg%20and%20grub.conf.md)

* [grub-install](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/grub-install.md)

* [grub-mkconfig](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/grub-mkconfig.md)

* [MBR](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/MBR.md)
 

### 102.3 Manage shared libraries

#### Weight	1

#### Description	Candidates should be able to determine the shared libraries that executable programs depend on and install them when necessary.

Key Knowledge Areas:

Identify shared libraries.
Identify the typical locations of system libraries.
Load shared libraries.
The following is a partial list of the used files, terms and utilities:

* [ldd](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/ldd.md)

* [ldconfig](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/ldconfig.md)

* [/etc/ld.so.conf](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/-etc-ld.so.conf.md)

* [LD_LIBRARY_PATH](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/LD_LIBRARY_PATH.md)
 

### 102.4 Use Debian package management

#### Weight	3

#### Description	Candidates should be able to perform package management using the Debian package tools.
Key Knowledge Areas:

Install, upgrade and uninstall Debian binary packages.
Find packages containing specific files or libraries which may or may not be installed.
Obtain package information like version, content, dependencies, package integrity and installation status (whether or not the package is installed).
Awareness of apt.
The following is a partial list of the used files, terms and utilities:

* [/etc/apt/sources.list](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/-etc-apt-source.list.md)

* [dpkg](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/dpkg.md)

* [dpkg-reconfigure](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/dpkg-reconfigure.md)

* [apt-get](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/apt-get.md)

* [apt-cache](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/apt-cache.md)
 

### 102.5 Use RPM and YUM package management

#### Weight	3

#### Description	Candidates should be able to perform package management using RPM, YUM and Zypper.
Key Knowledge Areas:

Install, re-install, upgrade and remove packages using RPM, YUM and Zypper.
Obtain information on RPM packages such as version, status, dependencies, integrity and signatures.
Determine what files a package provides, as well as find which package a specific file comes from.
Awareness of dnf.
The following is a partial list of the used files, terms and utilities:

* [rpm](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/rpm.md)

* [rpm2cpio](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/rpm2cpio.md)

* [/etc/yum.conf](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/-etc-yum.conf.md)

* [/etc/yum.repos.d/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/-etc-yum.repos.d.md)

* [yum](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/yum.md)

* [zypper](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/zypper.md)
 

### 102.6 Linux as a virtualization guest

#### Weight	1

#### Description	Candidates should understand the implications of virtualization and cloud computing on a Linux guest system.
Key Knowledge Areas:

Understand the general concept of virtual machines and containers
Understand common elements virtual machines in an IaaS cloud, such as computing instances, block storage and networking
Understand unique properties of a Linux system which have to changed when a system is cloned or used as a template
Understand how system images are used to deploy virtual machines, cloud instances and containers
Understand Linux extensions which integrate Linux with a virtualization product
Awareness of cloud-init
The following is a partial list of the used files, terms and utilities:

* [Virtual machine](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/Virtual%20machine.md)

* [Linux container](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/Linux%20container.md)

* [Application container](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/Application%20container.md)

* [Guest drivers](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/Guest%20drivers.md)

* [SSH host keys](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/SSH%20host%20keys.md)

* [D-Bus machine id](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/D-Bus%20machine%20id.md)
 

## Topic 103: GNU and Unix Commands

### 103.1 Work on the command line

#### Weight	4

#### Description	Candidates should be able to interact with shells and commands using the command line. The objective assumes the Bash shell.
Key Knowledge Areas:

Use single shell commands and one line command sequences to perform basic tasks on the command line.
Use and modify the shell environment including defining, referencing and exporting environment variables.
Use and edit command history.
Invoke commands inside and outside the defined path.
The following is a partial list of the used files, terms and utilities:

* [bash](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/bash.md)

* [echo](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/echo.md)

* [env](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/env.md)

* [export](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/export.md)

* [pwd](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/pwd.md)

* [set](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/set.md)

* [unset](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/unset.md)

* [type](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/type.md)

* [which](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/which.md)

* [man](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/man.md)

* [uname](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/uname.md)

* [history](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/history.md)

* [.bash_history](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/.bash_history.md)

* [Quoting](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/Quoting.md)
 

### 103.2 Process text streams using filters

#### Weight	2

#### Description	Candidates should should be able to apply filters to text streams.
Key Knowledge Areas:

Send text files and output streams through text utility filters to modify the output using standard UNIX commands found in the GNU textutils package.
The following is a partial list of the used files, terms and utilities:

* [bzcat](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/bzcat.md)

* [cat](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/cat.md)

* [cut](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/cut.md)

* [head](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/head.md)

* [less](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/less.md)

* [md5sum]()

* [nl](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/nl.md)

* [od](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/od.md)

* [paste](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/paste.md)

* [sed](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/sed.md)

* [sha256sum](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/sha256sum.md)

* [sha512sum](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/sha512sum.md)

* [sort](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/sort.md)

* [split](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/split.md)

* [tail](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/tail.md)

* [tr](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/tr.md)

* [uniq](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/uniq.md)

* [wc](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/wc.md)

* [xzcat](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/xzcat.md)

* [zcat](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/zcat.md)
 

### 103.3 Perform basic file management

#### Weight	4

#### Description	Candidates should be able to use the basic Linux commands to manage files and directories.
Key Knowledge Areas:

Copy, move and remove files and directories individually.
Copy multiple files and directories recursively.
Remove files and directories recursively.
Use simple and advanced wildcard specifications in commands.
Using find to locate and act on files based on type, size, or time.
Usage of tar, cpio and dd.
The following is a partial list of the used files, terms and utilities:

* [cp](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/cp.md)

* [find](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/find.md)

* [mkdir](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/mkdir.md)

* [mv](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/mv.md)

* [ls](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/ls.md)

* [rm](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/rm.md)

* [rmdir](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/rmdir.md)

* [touch](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/touch.md)

* [tar](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/tar.md)

* [cpio](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/cpio.md)

* [dd](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/dd.md)

* [file](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/file.md)

* [gzip](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/gzip.md)

* [gunzip](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/gunzip.md)

* [bzip2](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/bzip2.md)

* [bunzip2](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/bunzip2.md)

* [xz](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/xz.md)

* [unxz](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/unxz.md)

* [file globbing](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/file%20globbing.md)
 

### 103.4 Use streams, pipes and redirects

#### Weight	4

#### Description	Candidates should be able to redirect streams and connect them in order to efficiently process textual data. Tasks include redirecting standard input, standard output and standard error, piping the output of one command to the input of another command, using the output of one command as arguments to another command and sending output to both stdout and a file.
Key Knowledge Areas:

Redirecting standard input, standard output and standard error.
Pipe the output of one command to the input of another command.
Use the output of one command as arguments to another command.
Send output to both stdout and a file.
The following is a partial list of the used files, terms and utilities:

* [tee](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/tee.md)

* [xargs](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/xargs.md)

 

### 103.5 Create, monitor and kill processes

#### Weight	4

#### Description	Candidates should be able to perform basic process management.
Key Knowledge Areas:

Run jobs in the foreground and background.
Signal a program to continue running after logout.
Monitor active processes.
Select and sort processes for display.
Send signals to processes.
The following is a partial list of the used files, terms and utilities:

* [bg](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/bg.md)

* [fg](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/fg.md)

* [jobs](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/jobs.md)

* [kill](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/kill.md)

* [nohup](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/nohup.md)

* [ps](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/ps.md)

* [top](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/top.md)

* [free](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/free.md)

* [uptime](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/uptime.md)

* [pgrep](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/pgrep.md)

* [pkill](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/pkill.md)

* [killall](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/killall.md)

* [watch](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/watch.md)

* [screen](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/screen.md)

* [tmux](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/tmux.md)
 

### 103.6 Modify process execution priorities

#### Weight	2

#### Description	Candidates should should be able to manage process execution priorities.
Key Knowledge Areas:

Know the default priority of a job that is created.
Run a program with higher or lower priority than the default.
Change the priority of a running process.
The following is a partial list of the used files, terms and utilities:

* [nice](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/nice.md)

* [ps](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/ps.md)

* [renice](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/renice.md)

* [top](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/top.md)
 

### 103.7 Search text files using regular expressions

#### Weight	3

#### Description	Candidates should be able to manipulate files and text data using regular expressions. This objective includes creating simple regular expressions containing several notational elements as well as understanding the differences between basic and extended regular expressions. It also includes using regular expression tools to perform searches through a filesystem or file content.
Key Knowledge Areas:

Create simple regular expressions containing several notational elements.
Understand the differences between basic and extended regular expressions.
Understand the concepts of special characters, character classes, quantifiers and anchors.
Use regular expression tools to perform searches through a filesystem or file content.
Use regular expressions to delete, change and substitute text.
The following is a partial list of the used files, terms and utilities:

* [grep](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/grep.md)

* [egrep](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/egrep.md)

* [fgrep](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/fgrep.md)

* [sed](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/sed.md)

* [regex(7)](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/regex(7).md)
 

### 103.8 Basic file editing

#### Weight	3

#### Description	Candidates should be able to edit text files using vi. This objective includes vi navigation, vi modes, inserting, editing, deleting, copying and finding text. It also includes awareness of other common editors and setting the default editor.
Key Knowledge Areas:

Navigate a document using vi.
Understand and use vi modes.
Insert, edit, delete, copy and find text in vi.
Awareness of Emacs, nano and vim.
Configure the standard editor.
The following is a partial list of the used files, terms and utilities:

* [vi](url)

* [/, ?](url)

* [h,j,k,l](url)

* [i, o, a](url)

* [d, p, y, dd, yy](url)

* [ZZ, :w!, :q!](url)

* [EDITOR](url)
 

## Topic 104: Devices, Linux Filesystems, Filesystem Hierarchy Standard

### 104.1 Create partitions and filesystems

#### Weight	2

#### Description	Candidates should be able to configure disk partitions and then create filesystems on media such as hard disks. This includes the handling of swap partitions.
Key Knowledge Areas:

Manage MBR and GPT partition tables
Use various mkfs commands to create various filesystems such as:

* [ext2/ext3/ext4](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/ext2-ext3-ext4.md)

* [XFS](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/XFS.md)

* [VFAT](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/VFAT.md)

* [exFAT](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/exFAT.md)

Basic feature knowledge of Btrfs, including multi-device filesystems, compression and subvolumes.

The following is a partial list of the used files, terms and utilities:

* [fdisk](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/fdisk.md)

* [gdisk](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/gdisk.md)

* [parted](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/parted.md)

* [mkfs](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/mkfs.md)

* [mkswap](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/mkswap.md)
 

### 104.2 Maintain the integrity of filesystems

#### Weight	2

#### Description	Candidates should be able to maintain a standard filesystem, as well as the extra data associated with a journaling filesystem.
Key Knowledge Areas:

Verify the integrity of filesystems.
Monitor free space and inodes.
Repair simple filesystem problems.
The following is a partial list of the used files, terms and utilities:

* [du](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/du.md)

* [df](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/df.md)

* [fsck](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/fsck.md)

* [e2fsck](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/e2fsck.md)

* [mke2fs](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/mke2fs.md)

* [tune2fs](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/tune2fs.md)

* [xfs_repair](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/xfs_repair.md)

* [xfs_fsr](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/xfs_fsr.md)

* [xfs_db](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/xfs_db.md)
 

### 104.3 Control mounting and unmounting of filesystems

#### Weight	3

#### Description	Candidates should be able to configure the mounting of a filesystem.
Key Knowledge Areas:

Manually mount and unmount filesystems.
Configure filesystem mounting on bootup.
Configure user mountable removable filesystems.
Use of labels and UUIDs for identifying and mounting file systems.
Awareness of systemd mount units.
The following is a partial list of the used files, terms and utilities:

* [/etc/fstab](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/-etc-fstab.md)

* [/media/](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/-media-.md)

* [mount](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/mount.md)

* [umount](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/umount.md)

* [blkid](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC2-201/blkid.md)

* [lsblk](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/lsblk.md)
 

### 104.4 Removed

### 104.5 Manage file permissions and ownership

#### Weight	3

#### Description	Candidates should be able to control file access through the proper use of permissions and ownerships.
Key Knowledge Areas:

Manage access permissions on regular and special files as well as directories.
Use access modes such as suid, sgid and the sticky bit to maintain security.
Know how to change the file creation mask.
Use the group field to grant file access to group members.
The following is a partial list of the used files, terms and utilities:

* [chmod](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/chmod.md)

* [umask](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/umask.md)

* [chown](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/chown.md)

* [chgrp](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/chgrp.md)
 

### 104.6 Create and change hard and symbolic links

#### Weight	2

#### Description	Candidates should be able to create and manage hard and symbolic links to a file.
Key Knowledge Areas:

Create links.
Identify hard and/or soft links.
Copying versus linking files.
Use links to support system administration tasks.
The following is a partial list of the used files, terms and utilities:

* [ln](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/ln.md)

* [ls](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/ls.md)
 

### 104.7 Find system files and place files in the correct location

#### Weight	2

#### Description	Candidates should be thoroughly familiar with the Filesystem Hierarchy Standard (FHS), including typical file locations and directory classifications.
Key Knowledge Areas:

Understand the correct locations of files under the FHS.
Find files and commands on a Linux system.
Know the location and purpose of important file and directories as defined in the FHS.
The following is a partial list of the used files, terms and utilities:

* [find](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/find.md)

* [locate](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/locate.md)

* [updatedb](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/updatedb.md)

* [whereis](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/File-systems-Cocepts/LPIC1-101/whereis.md)

* [which](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/which.md)

* [type](https://github.com/SamanKhalife/linux-Tutorial/blob/main/TXT%20FILES/type.md)
