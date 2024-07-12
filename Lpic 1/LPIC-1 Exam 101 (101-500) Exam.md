## 1- Which SysV init configuration file should be modified to disable the ctrl-alt-delete key combination?
- A. /etc/keys
- B. /proc/keys 
- C. /etc/inittab 
- D. /proc/inittab 
- E. /etc/reboot

## 2- Which of the following information is stored within the BIOS? (Choose TWO correct answers.)
- A. Boot device order
- B. Linux kernel version
- C. Timezone
- D. Hardware configuration - E. The system's hostname

## 3- Which of the following commands reboots the system when using SysV init? (Choose TWO correct answers.)
- A. shutdown -r now
- B. shutdown -r "rebooting"
- C. telinit 6
- D. telinit 0
- E. shutdown -k now "rebooting"

## 4- Which of the following are init systems used within Linux systems? (Choose THREE correct answers.)
- A. startd
- B. systemd 
- C. Upstart 
- D. SysInit 
- E. SysV init

## 5- SIMULATION
Which file in the /proc filesystem lists parameters passed from the bootloader to the kernel? (Specify the file name only without any path.)

## 6- What information can the lspci command display about the system hardware? (Choose THREE correct answers.)
- A. Device IRQ settings
- B. PCI bus speed
- C. System battery type
- D. Device vendor identification - E. Ethernet MAC address

## 7- Which of the following commands brings a system running SysV init into a state in which it is safe to perform maintenance tasks? (Choose TWO correct answers.)
- A. shutdown -R 1 now 
- B. shutdown -single now 
- C. init 1
- D. telinit 1
- E. runlevel 1


## 8- What is the first program that is usually started, at boot time, by the Linux kernel when using SysV init?
- A. /lib/init.so
- B. /sbin/init
- C. /etc/r- C.d/rcinit
- D. /proc/sys/kernel/init 
- E. /boot/init

## 9- SIMULATION
Which command will display messages from the kernel that were output during the normal boot sequence?


ANSWERS:
-------------------------------------------------------------------------------------------
1	Correct Answer: C

2	Correct Answer: AD

3	Correct Answer: AC

4	Correct Answer: BCE Section: System Architecture Explanation

5	Correct Answer: cmdline -or- /proc/cmdline

6	Correct Answer: ABD Section: System Architecture Explanation

7	Correct Answer: CD

8	Correct Answer: B

9	Correct Answer: dmesg -or- /bin/dmesg Section: System Architecture Explanation


## 10- Which of the following commands will write a message to the terminals of all logged in users?
- A. bcast 
- B. mesg 
- C. print 
- D. wall 
- E. yell

## 11- Which of the following kernel parameters instructs the kernel to suppress most boot messages?
- A. silent
- B. verbose=0 
- C. nomesg 
- D. quiet

## 12- Which of the following options for the kernel's command line changes the systemd boot target to rescu- E.target instead of the default target?
- A. system- D.target=rescu- E.target 
- B. system- D.runlevel=rescu- E.target 
- C. system- D.service=rescu- E.target 
- D. system- D.default=rescu- E.target 
- E. system- D.unit=rescu- E.target

## 13- After modifying GNU GRUB's configuration file, which command must be run for the changes to take effect?
- A. kill -HUP $(pidof grub) 
- B. grub-install
- C. grub
- D. No action is required

## 14- Which of the following commands is used to update the list of available packages when using dpkg based package management?
- A. apt-get update
- B. apt-get upgrade
- C. apt-cache update 
- D. apt-get refresh
- E. apt-cache upgrade

## 15- Which of the following commands can be used to download the RPM package kernel without installing it?
- A. yum download --no-install kernel 
- B. yumdownloader kernel
- C. rpm --download --package kernel
- D. rpmdownload kernel


## 16- When using rpm --verify to check files created during the installation of RPM packages, which of the following information is taken into consideration? (Choose THREE correct answers.)
- A. Timestamps
- B. MD5 checksums 
- C. Inodes
- D. File sizes
- E. GnuPG signatures

## 17- Which of the following is correct when talking about mount points?
- A. Every existing directory can be used as a mount point.
- B. Only empty directories can be used as a mount point.
- C. Directories need to have the SetUID flag set to be used as a mount point.
- D. Files within a directory are deleted when the directory is used as a mount point.

## 18- Which function key is used to start Safe Mode in Windows NT?
- A. F10
- B. F8
- C. F6
- D. Windows NT does not support Safe Mode

## 19- Which of the following environment variables overrides or extends the list of directories holding shared libraries?
- A. LD_LOAD_PATH
- B. LD_LIB_PATH
- C. LD_LIBRARY_PATH 
- D. LD_SHARE_PATH 
- E. LD_RUN_PATH

ANSWERS:
-------------------------------------------------------------------------------------------
10	Correct Answer: D

11	Correct Answer: D

12	Correct Answer: E

13	Correct Answer: D

14	Correct Answer: A

15	Correct Answer: B

16	Correct Answer: ABD

17	Correct Answer: A

18	Correct Answer: D

19	Correct Answer: C

## 20- SIMULATION
Which world-writable directory should be placed on a separate partition in order to prevent users from being able to fill up the / filesystem? (Specify the full path to the directory.)

## 21- Which RPM command will output the name of the package which supplied the file /etc/exports?
- A. rpm -F /etc/exports 
- B. rpm -qf /etc/exports 
- C. rpm -Kl /etc/exports 
- D. rpm -qp /etc/exports 
- E. rpm -qi /etc/exports

## 22- SIMULATION
In which directory must definition files be placed to add additional repositories to yum?

## 23- SIMULATION
What is the name of the main configuration file for GNU GRUB? (Specify the file name only without any path.)


## 24- When removing a package, which of the following dpkg options will completely remove the files including configuration files?
- A. --clean 
- B. --delete 
- C. --purge 
- D. remove


## 25- Which file should be edited to select the network locations from which Debian installation package files are loaded?
- A. /etc/dpkg/dpkg.cfg 
- B. /etc/apt/apt.conf
- C. /etc/apt/apt.conf.d 
- D. /etc/apt/sources.list 
- E. /etc/dpkg/dselect.cfg


## 26- SIMULATION
Which option to the yum command will update the entire system? (Specify ONLY the option name without any additional parameters.)

## 27- SIMULATION
Which command will disable swapping on a device? (Specify ONLY the command without any path or parameters.)

## 28- SIMULATION
Which Debian package management tool asks the configuration ##s for a specific already installed package just as if the package were being installed for the first time? (Specify ONLY the command without any path or parameters.)

## 29- Which of the following commands overwrites the bootloader located on /dev/sda without overwriting the partition table or any data following it?
- A. dd if=/dev/zero of=/dev/sda bs=512
- B. dd if=/dev/zero of=/dev/sda bs=512 count=1 
- C. dd if=/dev/zero of=/dev/sda bs=440 count=1
- D. dd if=/dev/zero of=/dev/sda bs=440


ANSWERS:
-------------------------------------------------------------------------------------------
20	Correct Answer: /tmp -or- tmp -or- /var/tmp -or- /tmp/ -or- /var/tmp/ Section: Linux Installation and Package Management Explanatio

21	Correct Answer: B

22	Correct Answer: /etc/yum.repos.d -or- /etc/yum.repos.d/ -or- yum.repos.d -or- yum.repos.d/

23	Correct Answer: menu.lst -or- gru- B.conf -or- gru- B.cfg Section: Linux Installation and Package Management Explanatio

24	Correct Answer: C

25	Correct Answer: D

26	Correct Answer: update -or- upgrade

27	Correct Answer: swapoff -or- /sbin/swapoff

28	Correct Answer: dpkg-reconfigure

29	Correct Answer: C

## 30- Which of the following commands can be used to create a USB storage media from a disk image?
- A. gdisk 
- B. dd
- C. cc
- D. fdisk 
- E. mount

## 31- In Bash, inserting 1>&2 after a command redirects
 - A. standard error to standard input.
 - B. standard input to standard error. 
 - C. standard output to standard error.
 - D. standard error to standard output.
 - E. standard output to standard input.


## 32- What command will generate a list of user names from /etc/passwd along with their login shell?
- A. column -s : 1,7 /etc/passwd 
- B. chop -c 1,7 /etc/passwd
- C. colrm 1,7 /etc/passwd
- D. cut -d: -f1,7 /etc/passwd

## 33- In a nested directory structure, which find command line option would be used to restrict the command to searching down a particular number of subdirectories?
- A. -dirmax
- B. -maxdepth 
- C. -maxlevels 
- D. -n
- E. -s

## 34- Which of the following statements is correct regarding the command foo 1> bar?
- A. The stdout from the command foo is appended to the file bar.
- B. The stdout from the command foo overwrites the file bar.
- C. The command foo receives its stdin from the file bar.
- D. The command foo receives its stdin from the stdout of the command bar. 
- E. The stderr from the command foo is saved to the file bar.


## 35- Which of the following commands kills the process with the PID 123 but allows the process to "clean up" before exiting?
- A. kill -PIPE 123 
- B. kill -KILL 123 
- C. kill -STOP 123 
- D. kill -TERM 123


## 36- SIMULATION
Which signal is missing from the following command that is commonly used to instruct a daemon to reinitialize itself, including reading configuration files?


## 37- What is the maximum niceness value that a regular user can assign to a process with the nice command when executing a new process?
- A. 9 
- B. 19 
- C. 49 
- D. 99


## 38- Immediately after deleting 3 lines of text in vi and moving the cursor to a different line, which single character command will insert the deleted content below the current line?
- A. i (lowercase) 
- B. P (uppercase) 
- C. p (lowercase) 
- D. U (uppercase) 
- E. u (lowercase)


## 39- A user accidentally created the subdirectory \dir in his home directory. Which of the following commands will remove that directory?
- A. rmdir '~/\dir' 
- B. rmdir "~/\dir" 
- C. rmdir ~/'dir' 
- D. rmdir ~/\dir
- E. rmdir ~/\\dir

ANSWERS:
-------------------------------------------------------------------------------------------
30	Correct Answer: B

31	Correct Answer: C

32	Correct Answer: D

33	Correct Answer: B

34	Correct Answer: B

35	Correct Answer: D

36	Correct Answer: HUP -or- SIGHUP -or- 1 Section: GNU and Unix Commands Explanation

37	Correct Answer: B

38	Correct Answer: C

39	Correct Answer: E


## 40- In compliance with the FHS, in which of the directories are man pages found?

- A. /usr/share/man 
- B. /opt/man
- C. /usr/doc/
- D. /var/pkg/man 
- E. /var/man

## 41- Which of the following commands will send output from the program myapp to both standard output (stdout) and the file file1.log?
- A. cat < myapp | cat > file1.log 
- B. myapp 0>&1 | cat > file1.log 
- C. myapp | cat > file1.log
- D. myapp | tee file1.log
- E. tee myapp file1.log



## 42- What is the purpose of the Bash built-in export command?
- A. It allows disks to be mounted remotely.
- B. It runs a command as a process in a subshell.
- C. It makes the command history available to subshells.
- D. It sets up environment variables for applications.
- E. It shares NFS partitions for use by other systems on the network.

## 43- What is the output of the following command? echo "Hello World" | tr -d aieou
- A. Hello World 
- B. eoo
- C. Hll Wrld
- D. eoo Hll Wrld


## 44- Which of the following characters can be combined with a separator string in order to read from the current input source until the separator string, which is on a separate line and without any trailing spaces, is reached?
- A. << 
- B. <| 
- C. !< 
- D. &<

## 45- Which of the following commands will NOT update the modify timestamp on the file /tmp/myfil- E.txt?
- A. file /tmp/myfil- E.txt
- B. echo "Hello" >/tmp/myfil- E.txt
- C. sed -ie "s/1/2/" /tmp/myfil- E.txt
- D. echo -n "Hello" >>/tmp/myfil- E.txt 
- E. touch /tmp/myfil- E.txt

## 46- What is the default nice level when a process is started using the nice command?
- A. -10 
- B. 10 
- C. 20 
- D. 0


## 47- What is the default action of the split command on an input file?
- A. It will break the file into new files of 1,024 byte pieces each.
- B. It will break the file into new files of 1,000 line pieces each.
- C. It will break the file into new files of 1,024 kilobyte pieces each.
- D. It will break the file into new files that are no more than 5% of the size of the original fil- E.


## 48- What is the difference between the i and a commands of the vi editor?
- A. i (interactive) requires the user to explicitly switch between vi modes whereas a (automatic) switches modes automatically.
- B. i (insert) inserts text before the current cursor position whereas a (append) inserts text after the cursor.
- C. i (independent rows) starts every new line at the first character whereas a (aligned rows) keeps the indentation of the previous lin- E. 
- D. i (interrupt) temporarily suspends editing of a file to the background whereas a (abort) terminates editing.


## 49- SIMULATION
Which command displays a list of all background tasks running in the current shell? (Specify ONLY the command without any path or parameters.)


ANSWERS:
-------------------------------------------------------------------------------------------
40	Correct Answer: A

41	Correct Answer: D

42	Correct Answer: D

43	Correct Answer: C

44	Correct Answer: A

45	Correct Answer: A

46	Correct Answer: B

47	Correct Answer: B

48	Correct Answer: B

49	Correct Answer: jobs

## 50- Which of the following commands moves and resumes in the background the last stopped shell job?
- A. run 
- B. bg
- C. fg 
- D. back


## 51- What is the effect of the egrep command when the -v option is used?
- A. It enables color to highlight matching parts.
- B. It only outputs non-matching lines.
- C. It shows the command's version information.
- D. It changes the output order showing the last matching line first.



## 52- What does the ? symbol within regular expressions represent?
- A. Match the preceding qualifier one or more times. 
- B. Match the preceding qualifier zero or more times. 
- C. Match the preceding qualifier zero or one times. 
- D. Match a literal ? character.



## 53- In the vi editor, how can commands such as moving the cursor or copying lines into the buffer be issued multiple times or applied to multiple rows? 
- A. By using the command :repeat followed by the number and the comman- D.
- B. By specifying the number right in front of a command such as 4l or 2yj.
- C. By selecting all affected lines using the shift and cursor keys before applying the comman- D.
- D. By issuing a command such as :set repetition=4 which repeats every subsequent command 4 times.


## 54- Which of the following files, located in the user home directory, is used to store the Bash history?
- A. .bash_history 
- B. .bash_histfile 
- C. .history
- D. .bashrc_history 
- E. .history_bash


## 55- SIMULATION
Which Bash environment variable defines in which file the user history is stored when exiting a Bash process? (Specify ONLY the variable nam- E.)


## 56- Which of the following commands displays the contents of a gzip compressed tar archive?
- A. gzip archiv- E.tgz | tar xvf - 
- B. tar ztf archiv- E.tgz
- C. gzip -d archiv- E.tgz | tar tvf - 
- D. tar cf archiv- E.tgz

## 57- Which grep command will print only the lines that do not end with a / in the file foo?
- A. grep '/$' foo 
- B. grep '/#' foo 
- C. grep -v '/$' foo 
- D. grep -v '/#' foo

## 58- Which of the following commands is used to change options and positional parameters for a running Bash?
- A. history 
- B. set
- C. bashconf 
- D. setsh
- E. envsetup



## 59- Which of the following commands replaces each occurrence of 'bob' in the file letter with 'Bob' and writes the result to the file newletter?
- A. sed '/bob/Bob' letter > newletter
- B. sed s/bob/Bob/ letter < newletter 
- C. sed 's/bob/Bob' letter > newletter
- D. sed 's/bob/Bob/g' letter > newletter 
- E. sed 's/bob, Bob/' letter > newletter

ANSWERS:
-------------------------------------------------------------------------------------------
50	Correct Answer: B

51	Correct Answer: B

52	Correct Answer: C

53	Correct Answer: B

54	Correct Answer: A

55	Correct Answer: HISTFILE

56	Correct Answer: B

57	Correct Answer: C

58	Correct Answer: B

59	Correct Answer: D

## 60- From a Bash shell, which of the following commands directly executes the instruction from the file /usr/local/bin/runm- E.sh without starting a subshell? (Please select TWO answers.)
- A. source /usr/local/bin/runm- E.sh 
- B. . /usr/local/bin/runm- E.sh
- C. /bin/bash /usr/local/bin/runm- E.sh 
- D. /usr/local/bin/runm- E.sh
- E. run /usr/local/bin/runm- E.sh

## 61- Regarding the command:
nice -5 /usr/bin/prog
Which of the following statements is correct?
- A. /usr/bin/prog is executed with a nice level of -5. 
- B. /usr/bin/prog is executed with a nice level of 5.
- C. /usr/bin/prog is executed with a priority of -5. 
- D. /usr/bin/prog is executed with a priority of 5.



## 62- Which shell command is used to continue background execution of a suspended command?
- A. &
- B. bg 
- C. cont 
- D. exec 
- E. :&

## 63- Which of the following shell redirections will write standard output and standard error output to a file named filename?

- A. 2>&1 >filename 
- B. >filename 2>&1 
- C. 1>&2>filename 
- D. >>filename
- E. 1&2>filename

## 64- In the vi editor, which of the following commands will copy the current line into the vi buffer?
- A. c 
- B. cc 
- C. 1c 
- D. yy 
- E. 1y



## 65- Which of the following sequences in the vi editor saves the opened document and exits the editor? (Choose TWO correct answers.)
- A. esc ZZ 
- B. ctrl :w! 
- C. esc zz 
- D. esc :wq! 
- E. ctrl XX

## 66- When starting a program with the nice command without any additional parameters, which nice level is set for the resulting process?
- A. -10 
- B. 0 
- C. 10 
- D. 20


## 67- Which of the following commands will reduce all consecutive spaces down to a single space?
- A. tr '\s' ' ' < - A.txt > - B.txt 
- B. tr -c ' ' < - A.txt > - B.txt
- C. tr -d ' ' < - A.txt > - B.txt 
- D. tr -r ' ' '\n' < - A.txt > - B.txt 
- E. tr -s ' ' < - A.txt > - B.txt


## 68- Which character, added to the end of a command, runs that command in the background as a child process of the current shell?
- A. ! 
- B. + 
- C. & 
- D. % 
- E. #

## 69- Which of the following commands will print the last 10 lines of a text file to the standard output?
- A. cat -n 10 filename 
- B. dump -n 10 filename 
- C. head -n 10 filename 
- D. tail -n 10 filename


ANSWERS:
-------------------------------------------------------------------------------------------
60	Correct Answer: AB

61	Correct Answer: B

62	Correct Answer: B

63	Correct Answer: B

64	Correct Answer: D

65	Correct Answer: AD

66	Correct Answer: C

67	Correct Answer: E

68	Correct Answer: C

69	Correct Answer: D

## 70- Which of the following commands prints a list of usernames (first column) and their primary group (fourth column) from the /etc/passwd file?
- A. fmt -f 1,4 /etc/passwd
- B. split -c 1,4 /etc/passwd 
- C. cut -d : -f 1,4 /etc/passwd 
- D. paste -f 1,4 /etc/passwd


## 71- Which of the following signals is sent to a process when the key combination CTRL+C is pressed on the keyboard?
- A. SIGTERM 
- B. SIGINT 
- C. SIGSTOP 
- D. SIGKILL


## 72- What happens after issuing the command vi without any additional parameters?
- A. vi starts and loads the last file used and moves the cursor to the position where vi was when it last exite- D. - B. vi starts and requires the user to explicitly either create a new or load an existing fil- E.
- C. vi exits with an error message as it cannot be invoked without a file name to operate on.
- D. vi starts in command mode and opens a new empty fil- E.
- E. vi starts and opens a new file which is filled with the content of the vi buffer if the buffer contains text.


## 73- Which of the following command sets the Bash variable named TEST with the content FOO?
- A. set TEST="FOO" 
- B. TEST = "FOO" 
- C. var TEST="FOO" 
- D. TEST="FOO"


## 74- Which variable defines the directories in which a Bash shell searches for executable commands?
- A. BASHEXEC 
- B. BASHRC 
- C. PATH
- D. EXECPATH 
- E. PATHRC

## 75- Which of the following commands determines the type of a file by using a definition database file which contains information about all common file types?
- A. magic 
- B. type
- C. file
- D. pmagic 
- E. hash


## 76- SIMULATION
Which command is used in a Linux environment to create a new directory? (Specify ONLY the command without any path or parameters.)

## 77- Which of the following commands prints all files and directories within the /tmp directory or its subdirectories which are also owned by the user root? (Choose TWO correct answers.)
- A. find /tmp -uid root -print 
- B. find -path /tmp -uid root 
- C. find /tmp -user root -print 
- D. find /tmp -user root
- E. find -path /tmp -user root print


## 78- When running the command
sed -e "s/a/b/" /tmp/file >/tmp/file
While /tmp/file contains data, why is /tmp/file empty afterwards?
- A. The file order is incorrect. The destination file must be mentioned before the command to ensure redirection.
- B. The command sed did not match anything in that file therefore the output is empty.
- C. When the shell establishes the redirection it overwrites the target file before the redirected command starts and opens it for reading. 
- D. Redirection for shell commands do not work using the > character. It only works using the | character instea- D.

## 79- When given the following command lin- E. echo "foo bar" | tee bar | cat
Which of the following output is created?
- A. cat
- B. foo bar 
- C. tee bar 
- D. bar
- E. foo

ANSWERS:
-------------------------------------------------------------------------------------------
70	Correct Answer: C

71	Correct Answer: B

72	Correct Answer: D

73	Correct Answer: D

74	Correct Answer: C

75	Correct Answer: C

76	Correct Answer: mkdir -or- /usr/bin/mkdir Section: GNU and Unix Commands Explanation

77	Correct Answer: CD

78	Correct Answer: C

79	Correct Answer: B

## 80- Which of the following commands can be used to determine how long the system has been running? (Choose TWO correct answers.)
- A. uptime
- B. up
- C. top
- D. uname -u 
- E. time up


## 81- Which of the following are valid stream redirection operators within Bash? (Choose THREE correct answers.)
- A. < 
- B. <<< 
- C. > 
- D. >>> 
- E. %>

## 82- After successfully creating a hard link called bar to the ordinary file foo, foo is deleted from the filesystem. Which of the following describes the resulting situation?
- A. foo and bar would both be remove- D.
- B. foo would be removed while bar would remain accessibl- E.
- C. foo would be remove- D. bar would still exist but would be unusabl- E. 
- D. Both foo and bar would remain accessibl- E.
- E. The user is prompted whether bar should be removed, too.

## 83- After moving data to a new filesystem, how can the former path of the data be kept intact in order to avoid reconfiguration of existing applications? (Choose TWO correct answers.)
- A. By creating an ACL redirection from the old to the new path of the dat- A. 
- B. By creating a hard link from the old to the new path of the dat- A.
- C. By creating a symbolic link from the old to the new path of the dat- A.
- D. By running the command touch on the old path.
- E. By mounting the new filesystem on the original path of the dat- A.

## 84- Which of the following commands changes the ownership of fil- E.txt to the user dan and the group staff?
- A. chown dan/staff fil- E.txt
- B. chown dan:staff fil- E.txt
- C. chown -u dan -g staff fil- E.txt
- D. chown dan -g staff fil- E.txt


## 85- Which of the following commands makes /bin/foo executable by everyone but writable only by its owner?
- A. chmod u=rwx,go=rx /bin/foo 
- B. chmod o+rwx,a+rx /bin/foo 
- C. chmod 577 /bin/foo
- D. chmod 775 /bin/foo


## 86- Which of the following commands can be used to search for the executable file foo when it has been placed in a directory not included in $PATH?
- A. apropos 
- B. which 
- C. find
- D. query
- E. whereis


## 87- What does the command mount -a do?
- A. It ensures that all file systems listed with the option noauto in /etc/fstab are mounte- D.
- B. It shows all mounted file systems that have been automatically mounte- D.
- C. It opens an editor with root privileges and loads /etc/fstab for editing.
- D. It ensures that all file systems listed with the option auto in /etc/fstab are mounte- D.
- E. It ensures that all file systems listed in /etc/fstab are mounted regardless of their options.

## 88- Which of the following settings for umask ensures that new files have the default permissions -rw-r----- ?
- A. 0017
- B. 0640
- C. 0038
- D. 0027

## 89- Which of the following is the device file name for the second partition on the only SCSI drive?
- A. /dev/hda1 
- B. /dev/sda2 
- C. /dev/sd0a2 
- D. /dev/sd1p2

## 90- In order to display all currently mounted filesystems, which of the following commands could be used? (Choose TWO correct answers.)
- A. cat /proc/self/mounts 
- B. free
- C. mount
- D. lsmounts
- E. cat /proc/filesystems

ANSWERS:
-------------------------------------------------------------------------------------------
80	Correct Answer: AC

81	Correct Answer: ABC

82	Correct Answer: B

83	Correct Answer: CE

84	Correct Answer: B

85	Correct Answer: A

86	Correct Answer: C

87	Correct Answer: D

88	Correct Answer: D

89	Correct Answer: B


## 91- Which of the following commands can be used to locate programs and their corresponding man pages and configuration files?
- A. dirname 
- B. which
- C. basename 
- D. query
- E. whereis



## 92- Which of the following commands changes the number of days before the ext3 filesystem on /dev/sda1 has to run through a full filesystem check while booting?
- A. tune2fs -d 200 /dev/sda1
- B. tune2fs -c 200 /dev/sda1
- C. tune2fs -i 200 /dev/sda1
- D. tune2fs -n 200 /dev/sda1
- E. tune2fs --days 200 /dev/sda1

## 93- Which type of filesystem is created by mkfs when it is executed with the block device name only and without any additional parameters?
- A. ext2 
- B. ext3 
- C. ext4 
- D. XFS 
- E. VFAT

## 94- How many fields are in a syntactically correct line of /etc/fstab?
- A. 3 
- B. 4 
- C. 5 
- D. 6 
- E. 7


## 95- SIMULATION
Which command is used to create and initialize the files used to store quota information? (Specify ONLY the command without any path or parameters.)


## 96- Which of the following file permissions belong to a symbolic link?
- A. -rwxrwxrwx 
- B. +rwxrwxrwx 
- C. lrwxrwxrwx 
- D. srwxrwxrwx

## 97- Creating a hard link to an ordinary file returns an error. What could be the reason for this?
- A. The source file is hidden.
- B. The source file is read-only.
- C. The source file is a shell script.
- D. The source file is already a hard link.
- E. The source and the target are on different filesystems.

## 98- Which of the following commands creates an ext3 filesystem on /dev/sdb1? (Choose TWO correct answers.)
- A. /sbin/mke2fs -j /dev/sdb1
- B. /sbin/mkfs -t ext3 /dev/sdb1 
- C. /sbin/mkfs -c ext3 /dev/sdb1 
- D. /sbin/mke3fs -j /dev/sdb1

## 99- Which of the following commands will change the quota for a specific user? 
- A. edquota
- B. repquota 
- C. quota -e
- D. quota

## 100- Which utility would be used to change how often a filesystem check is performed on an ext2 filesystem without losing any data stored on that filesystem?
- A. mod2fs 
- B. fsck
- C. tune2fs 
- D. mke2fs 
- E. fixe2fs

## 101- Which of the following Linux filesystems preallocates a fixed number of inodes at the filesystem's make/creation time and does NOT generate them as needed? (Choose TWO correct answers.)
- A. ext3 
- B. JFS 
- C. ext2 
- D. XFS 
- E. procfs

## 102- What is the purpose of the Filesystem Hierarchy Standard?
- A. It is a security model used to ensure files are organized according to their permissions and accessibility. 
- B. It provides unified tools to create, maintain and manage multiple filesystems in a common way.
- C. It defines a common internal structure of inodes for all compliant filesystems.
- D. It is a distribution neutral description of locations of files and directories.

ANSWERS:
-------------------------------------------------------------------------------------------
90	Correct Answer: AC

91	Correct Answer: E

92	Correct Answer: C

93	Correct Answer: A

94	Correct Answer: D

95	Correct Answer: quotacheck

96	Correct Answer: C

97	Correct Answer: E

98	Correct Answer: AB

99	Correct Answer: A

100	Correct Answer: C

101	Correct Answer: AC

102	Correct Answer: D


Exam A

## 1- Which of the following information is stored within the BIOS? (Choose TWO correct answers.)

- A.	Boot device order
- B.	Linux kernel version
- C.	Timezone
- D.	Hardware configuration
- E.	The system's hostname


## 2- Which of the following commands reboots the system when using SysV init? (Choose TWO correct answers.)

- A.	shutdown -r now
- B.	shutdown -r "rebooting"
- C.	telinit 6
- D.	telinit 0
- E.	shutdown -k now "rebooting"


## 3- Which of the following are init systems used within Linux systems? (Choose THREE correct answers.)

- A.	startd
- B.	systemd
- C.	Upstart
- D.	SysInit
- E.	SysV init

## 4- SIMULATION
Which file in the /proc filesystem lists parameters passed from the bootloader to the kernel? (Specify the file name only without any path.)

## 5- What information can the lspci command display about the system hardware? (Choose THREE correct answers.)
- A.	Device IRQ settings
- B.	PCI bus speed
- C.	System battery type
- D.	Device vendor identification
- E.	Ethernet MAC address

## 6- Which of the following commands brings a system running SysV init into a state in which it is safe to perform maintenance tasks? (Choose TWO correct answers.)
- A.	shutdown -R 1 now
- B.	shutdown -single now
- C.	init 1
- D.	telinit 1
- E.	runlevel 1


## 7- What is the first program that is usually started, at boot time, by the Linux kernel when using SysV init?
- A.	/lib/init.so
- B.	/sbin/init
- C.	/etc/r- C.d/rcinit
- D.	/proc/sys/kernel/init
- E.	/boot/init

 
## 8- SIMULATION
Which command will display messages from the kernel that were output during the normal boot sequence?

## 9- Which of the following commands will write a message to the terminals of all logged in users?
- A.	bcast
- B.	mesg
- C.	print
- D.	wall
- E.	yell


## 10- Which of the following kernel parameters instructs the kernel to suppress most boot messages?
- A.	silent
- B.	verbose=0
- C.	nomesg
- D.	quiet

ANSWERS:
-------------------------------------------------------------------------------------------
1	Correct Answer: AD

2	Correct Answer: AC

3	Correct Answer: BCE Section: System Architecture Explanation

4	Correct Answer: cmdline, /proc/cmdline Section: System Architecture Explanation

5	Correct Answer: ABD Section: System Architecture Explanation

6	Correct Answer: CD

7	Correct Answer: B

8	Correct Answer: dmesg, /bin/dmesg Section: System Architecture Explanation

9	Correct Answer: D

10	Correct Answer: D


## 11- Which of the following options for the kernel's command line changes the systemd boot target to rescu- E.target instead of the default target?
- A.	system- D.target=rescu- E.target
- B.	system- D.runlevel=rescu- E.target
- C.	system- D.service=rescu- E.target
- D.	system- D.default=rescu- E.target
- E.	system- D.unit=rescu- E.target

 
## 12- After modifying GNU GRUB's configuration file, which command must be run for the changes to take effect?
- A.	kill -HUP $(pidof grub)
- B.	grub-install
- C.	grub
- D.	No action is required


## 13- Which of the following commands is used to update the list of available packages when using dpkg based package management?
- A.	apt-get update
- B.	apt-get upgrade
- C.	apt-cache update
- D.	apt-get refresh
- E.	apt-cache upgrade


## 14- Which of the following commands lists the dependencies of a given dpkg package?
- A.	apt-cache depends-on package
- B.	apt-cache dependencies package
- C.	apt-cache depends package
- D.	apt-cache requires package


## 15- Which of the following options is used in a GRUB Legacy configuration file to define the amount of time that the GRUB menu will be shown to the user?
- A.	hidemenu
- B.	splash
- C.	timeout
- D.	showmenu


## 16- What can the Logical Volume Manager (LVM) be used for? (Choose THREE correct answers.)
- A.	To create RAID 9 arrays.
- B.	To dynamically change the size of logical volumes.
- C.	To encrypt logical volumes.
- D.	To create snapshots.
- E.	To dynamically create or delete logical volumes.


## 17- Which of the following commands updates the linker cache of shared libraries?
- A.	mkcache
- B.	soconfig
- C.	mkldconfig
- D.	lddconfig
- E.	ldconfig


## 18- Which of the following commands lists all currently installed packages when using RPM package management?
- A.  yum --query --all
- B.	yum --list --installed
- C.	rpm --query --all
- D.	rpm --list installed


## 19- Which of the following commands can be used to download the RPM package kernel without installing it?
- A.	yum download --no-install kernel
- B.	yumdownloader kernel
- C.	rpm --download --package kernel
- D.	rpmdownload kernel

## 20- When using rpm --verify to check files created during the installation of RPM packages, which of the following information is taken into consideration? (Choose THREE correct answers.)
- A.	Timestamps
- B.	MD5 checksums
- C.	Inodes
- D.	File sizes
- E.	GnuPG signatures



ANSWERS:
-------------------------------------------------------------------------------------------
11	Correct Answer: E

12	Correct Answer: D

13	Correct Answer: A

14	Correct Answer: C

15	Correct Answer: C

16	Correct Answer: BDE

17	Correct Answer: E

18	Correct Answer: C

19	Correct Answer: B

20	Correct Answer: ABD



## 21- Which RPM command will output the name of the package which supplied the file /etc/exports?
- A.	rpm -F /etc/exports
- B.	rpm -qf /etc/exports
- C.	rpm -Kl /etc/exports
- D.	rpm -qp /etc/exports
- E.	rpm -qi /etc/exports


## 22- SIMULATION
In which directory must definition files be placed to add additional repositories to yum?

## 23- SIMULATION
What is the name of the main configuration file for GNU GRUB? (Specify the file name only without any path.)

## 24- When removing a package, which of the following dpkg options will completely remove the files including configuration files?
- A.	--clean
- B.	--delete
- C.	--purge
- D.	remove


## 25- Which file should be edited to select the network locations from which Debian installation package files are loaded?
- A.	/etc/dpkg/dpkg.cfg
- B.	/etc/apt/apt.conf
- C.	/etc/apt/apt.conf.d
- D.	/etc/apt/sources.list
- E.	/etc/dpkg/dselect.cfg


## 26- SIMULATION
Which option to the yum command will update the entire system? (Specify ONLY the option name without any additional parameters.)

## 27- SIMULATION
Which command will disable swapping on a device? (Specify ONLY the command without any path or parameters.)

## 28- SIMULATION
Which Debian package management tool asks the configuration ##s for a specific already installed package just as if the package were being installed for the first time? (Specify ONLY the command without any path or parameters.)

## 29- Which of the following commands overwrites the bootloader located on /dev/sda without overwriting the partition table or any data following it?
- A.	dd if=/dev/zero of=/dev/sda bs=512
- B.	dd if=/dev/zero of=/dev/sda bs=512 count=1
- C.	dd if=/dev/zero of=/dev/sda bs=440 count=1
- D.	dd if=/dev/zero of=/dev/sda bs=440


## 30- Which of the following commands can be used to create a USB storage media from a disk image?
- A.	gdisk
- B.	dd
- C.	cc
- D.	fdisk
- E.	mount

ANSWERS:
-------------------------------------------------------------------------------------------
21	Correct Answer: B

22	Correct Answer: /etc/yum.repos.d, /etc/yum.repos.d/, yum.repos.d, yum.repos.d/

23	Correct Answer: menu.lst, gru- B.conf, gru- B.cfg

24	Correct Answer: C

25	Correct Answer: D

26	Correct Answer: update, upgrade

27	Correct Answer: swapoff, /sbin/swapoff

28	Correct Answer: dpkg-reconfigure

29	Correct Answer: C

30	Correct Answer: B

## 31- In Bash, inserting 1>&2 after a command redirects
- A.	standard error to standard input.
- B.	standard input to standard error.
- C.	standard output to standard error.
- D.	standard error to standard output.
- E.	standard output to standard input.


## 32- What command will generate a list of user names from /etc/passwd along with their login shell?
- A.	column -s : 1,7 /etc/passwd
- B.	chop -c 1,7 /etc/passwd
- C.	colrm 1,7 /etc/passwd
- D.	cut -d: -f1,7 /etc/passwd

## 33- In a nested directory structure, which find command line option would be used to restrict the command to searching down a particular number of subdirectories?
- A.	-dirmax
- B.	-maxdepth
- C.	-maxlevels
- D.	-n
- E.	-s

## 34- Which of the following statements is correct regarding the command foo 1> bar?
- A.	The stdout from the command foo is appended to the file bar.
- B.	The stdout from the command foo overwrites the file bar.
- C.	The command foo receives its stdin from the file bar.
- D.	The command foo receives its stdin from the stdout of the command bar.
- E.	The stderr from the command foo is saved to the file bar.

 
## 35- Which of the following commands kills the process with the PID 123 but allows the process to "clean up" before exiting?
- A.	kill -PIPE 123
- B.	kill -KILL 123
- C.	kill -STOP 123
- D.	kill -TERM 123


## 36- SIMULATION
Which signal is missing from the following command that is commonly used to instruct a daemon to reinitialize itself, including reading configuration files?

killall -s 	daemon

## 37- What is the maximum niceness value that a regular user can assign to a process with the nice command when executing a new process?
- A.	9
- B.	19
- C.	49
- D.	99


## 38- Immediately after deleting 3 lines of text in vi and moving the cursor to a different line, which single character command will insert the deleted content below the current line?
- A.	i (lowercase)
- B.	P (uppercase)
- C.	p (lowercase)
- D.	U (uppercase)
- E.	u (lowercase)


## 39- A user accidentally created the subdirectory \dir in his home directory. Which of the following commands will remove that directory?
- A.	rmdir '~/\dir'
- B.	rmdir "~/\dir"
- C.	rmdir ~/'dir'
- D.	rmdir ~/\dir
- E.	rmdir ~/\\dir


## 40- In compliance with the FHS, in which of the directories are man pages found?
- A.	/usr/share/man
- B.	/opt/man
- C.	/usr/doc/
- D.	/var/pkg/man
- E.	/var/man

ANSWERS:
-------------------------------------------------------------------------------------------
31	Correct Answer: C

32	Correct Answer: D

33	Correct Answer: B

34	Correct Answer: B

35	Correct Answer: D

36	Correct Answer: HUP, SIGHUP, 1 Section: GNU and Unix Commands Explanation

37	Correct Answer: B

38	Correct Answer: C

39	Correct Answer: E

40	Correct Answer: A

## 41- Which of the following commands will send output from the program myapp to both standard output (stdout) and the file file1.log?
- A.	cat < myapp | cat > file1.log
- B.	myapp 0>&1 | cat > file1.log
- C.	myapp | cat > file1.log
- D.	myapp | tee file1.log
- E.	tee myapp file1.log

 
## 42- What is the purpose of the Bash built-in export command?
- A.	It allows disks to be mounted remotely.
- B.	It runs a command as a process in a subshell.
- C.	It makes the command history available to subshells.
- D.	It sets up environment variables for applications.
- E.	It shares NFS partitions for use by other systems on the network.


## 43- What is the output of the following command? echo "Hello World" | tr -d aieou
- A.	Hello World
- B.	eoo
- C.	Hll Wrld
- D.	eoo Hll Wrld


## 44- Which of the following characters can be combined with a separator string in order to read from the current input source until the separator string, which is on a separate line and without any trailing spaces, is reached?
- A.	<<
- B.	<|
- C.	!<
- D.	&<


## 45- In the vi editor, how can commands such as moving the cursor or copying lines into the buffer be issued multiple times or applied to multiple rows?
- A.	By using the command :repeat followed by the number and the comman- D.
- B.	By specifying the number right in front of a command such as 4l or 2yj.
- C.	By selecting all affected lines using the shift and cursor keys before applying the comman- D.
- D.	By issuing a command such as :set repetition=4 which repeats every subsequent command 4 times.


## 46- Which of the following files, located in the user home directory, is used to store the Bash history?
- A.	.bash_history
- B.	.bash_histfile
- C.	.history
- D.	.bashrc_history
- E.	.history_bash


## 47- SIMULATION
Which Bash environment variable defines in which file the user history is stored when exiting a Bash process? (Specify ONLY the variable nam- E.)


## 48- Which of the following commands displays the contents of a gzip compressed tar archive?
- A.	gzip archiv- E.tgz | tar xvf -
- B.	tar ztf archiv- E.tgz
- C.	gzip -d archiv- E.tgz | tar tvf -
- D.	tar cf archiv- E.tgz


## 49- Which grep command will print only the lines that do not end with a / in the file foo?
- A.	grep '/$' foo
- B.	grep '/#' foo
- C.	grep -v '/$' foo
- D.	grep -v '/#' foo

 ANSWERS:
-------------------------------------------------------------------------------------------
41	Correct Answer: D

42	Correct Answer: D

43	Correct Answer: C

44	Correct Answer: A

45	Correct Answer: B

46	Correct Answer: A

47	Correct Answer: HISTFILE

48	Correct Answer: B

49	Correct Answer: C

## 50- Which of the following commands is used to change options and positional parameters for a running Bash?
- A.	history
- B.	set
- C.	bashconf
- D.	setsh
- E.	envsetup


## 51- Which of the following commands replaces each occurrence of 'bob' in the file letter with 'Bob' and writes the result to the file newletter?
- A.	sed '/bob/Bob' letter > newletter
- B.	sed s/bob/Bob/ letter < newletter
- C.	sed 's/bob/Bob' letter > newletter
- D.	sed 's/bob/Bob/g' letter > newletter
- E.	sed 's/bob, Bob/' letter > newletter


## 52- From a Bash shell, which of the following commands directly executes the instruction from the file /usr/local/ bin/runm- E.sh without starting a subshell? (Please select TWO answers.)
- A.	source /usr/local/bin/runm- E.sh
- B.	. /usr/local/bin/runm- E.sh
- C.	/bin/bash /usr/local/bin/runm- E.sh
- D.	/usr/local/bin/runm- E.sh
- E.	run /usr/local/bin/runm- E.sh


## 53- Regarding the command: nice -5 /usr/bin/prog
Which of the following statements is correct?

- A.	/usr/bin/prog is executed with a nice level of -5.
- B.	/usr/bin/prog is executed with a nice level of 5.
- C.	/usr/bin/prog is executed with a priority of -5.
- D.	/usr/bin/prog is executed with a priority of 5.


## 54- Which shell command is used to continue background execution of a suspended command?
- A.	&
- B.	bg
- C.	cont
- D.	exec
- E.	:&


## 55- Which of the following shell redirections will write standard output and standard error output to a file named filename?
- A.	2>&1 >filename
- B.	>filename 2>&1
- C.	1>&2>filename
- D.	>>filename
- E.	1&2>filename

## 56- In the vi editor, which of the following commands will copy the current line into the vi buffer?
- A.	c
- B.	cc
- C.	1c
- D.	yy
- E.	1y


## 57- Which of the following sequences in the vi editor saves the opened document and exits the editor? (Choose TWO correct answers.)
- A.	esc ZZ
- B.	ctrl :w!
- C.	esc zz
- D.	esc :wq!
- E.	ctrl XX


## 58- When starting a program with the nice command without any additional parameters, which nice level is set for the resulting process?
- A. -10
- B.	0
- C.	10
- D.	20


## 59- Which of the following commands will reduce all consecutive spaces down to a single space?
- A.	tr '\s' ' ' < - A.txt > - B.txt
- B.	tr -c ' ' < - A.txt > - B.txt
- C.	tr -d ' ' < - A.txt > - B.txt
- D.	tr -r ' ' '\n' < - A.txt > - B.txt
- E.	tr -s ' ' < - A.txt > - B.txt

ANSWERS:
-------------------------------------------------------------------------------------------
50	Correct Answer: B

51	Correct Answer: D

52	Correct Answer: AB

53	Correct Answer: B

54	Correct Answer: B

55	Correct Answer: B

56	Correct Answer: D

57	Correct Answer: AD

58	Correct Answer: C

59	Correct Answer: E

## 60- Which character, added to the end of a command, runs that command in the background as a child process of the current shell?
- A. !
- B. +
- C.	&
- D.	%
- E.	#


## 61- Which of the following commands will print the last 10 lines of a text file to the standard output?
- A.	cat -n 10 filename
- B.	dump -n 10 filename
- C.	head -n 10 filename
- D.	tail -n 10 filename


## 62- Which of the following commands prints a list of usernames (first column) and their primary group (fourth column) from the /etc/passwd file?
- A.	fmt -f 1,4 /etc/passwd
- B.	split -c 1,4 /etc/passwd
- C.	cut -d : -f 1,4 /etc/passwd
- D.	paste -f 1,4 /etc/passwd


## 63- Which of the following signals is sent to a process when the key combination CTRL+C is pressed on the keyboard?
- A.	SIGTERM
- B.	SIGINT
- C.	SIGSTOP
- D.	SIGKILL
 

## 64- What happens after issuing the command vi without any additional parameters?
- A.	vi starts and loads the last file used and moves the cursor to the position where vi was when it last exite- D.
- B.	vi starts and requires the user to explicitly either create a new or load an existing fil- E.
- C.	vi exits with an error message as it cannot be invoked without a file name to operate on.
- D.	vi starts in command mode and opens a new empty fil- E.
- E.	vi starts and opens a new file which is filled with the content of the vi buffer if the buffer contains text.


## 65- Which of the following command sets the Bash variable named TEST with the content FOO?
- A.	set TEST="FOO"
- B.	TEST = "FOO"
- C.	var TEST="FOO"
- D.	TEST="FOO"


## 66- Which variable defines the directories in which a Bash shell searches for executable commands?
- A.	BASHEXEC
- B.	BASHRC
- C.	PATH
- D.	EXECPATH
- E.	PATHRC


## 67- Which of the following commands determines the type of a file by using a definition database file which contains information about all common file types?
- A.	magic
- B.	type
- C.	file
- D.	pmagic
- E.	hash


## 68- Which of the following commands can be used to search for the executable file foo when it has been placed in a directory not included in $PATH?
- A.	apropos
- B.	which
- C.	find
- D.	query
- E.	whereis


## 69- What does the command mount -a do?
- A.	It ensures that all file systems listed with the option noauto in /etc/fstab are mounte- D.
- B.	It shows all mounted file systems that have been automatically mounte- D.
- C.	It opens an editor with root privileges and loads /etc/fstab for editing.
- D.	It ensures that all file systems listed with the option auto in /etc/fstab are mounte- D.
- E.	It ensures that all file systems listed in /etc/fstab are mounted regardless of their options.

ANSWERS:
-------------------------------------------------------------------------------------------
60	Correct Answer: C

61	Correct Answer: D

62	Correct Answer: C

63	Correct Answer: B

64	Correct Answer: D

65	Correct Answer: D

66	Correct Answer: C

67	Correct Answer: C

68	Correct Answer: C

69	Correct Answer: D

## 70- Which of the following settings for umask ensures that new files have the default permissions -rw-r----- ? 
- A. 0017
- B.  0640
- C. 0038
- D. 0027


## 71- Which of the following is the device file name for the second partition on the only SCSI drive?
- A.	/dev/hda1
- B.	/dev/sda2
- C.	/dev/sd0a2
- D.	/dev/sd1p2


## 72- In order to display all currently mounted filesystems, which of the following commands could be used? (Choose TWO correct answers.)
- A.	cat /proc/self/mounts
- B.	free
- C.	mount
- D.	lsmounts
- E.	cat /proc/filesystems


## 73- Which of the following commands can be used to locate programs and their corresponding man pages and configuration files?
- A.	dirname
- B.	which
- C.	basename
- D.	query
- E.	whereis

 
## 74- Which of the following commands changes the number of days before the ext3 filesystem on /dev/sda1 has to run through a full filesystem check while booting?
- A.	tune2fs -d 200 /dev/sda1
- B.	tune2fs -c 200 /dev/sda1
- C.	tune2fs -i 200 /dev/sda1
- D.	tune2fs -n 200 /dev/sda1
- E.	tune2fs --days 200 /dev/sda1


## 75- Which type of filesystem is created by mkfs when it is executed with the block device name only and without any additional parameters?
- A.	ext2
- B.	ext3
- C.	ext4
- D.	XFS
- E.	VFAT


## 76- How many fields are in a syntactically correct line of /etc/fstab?
- A.	3
- B.	4
- C.	5
- D.	6
- E.	7


## 77- SIMULATION
Which command is used to create and initialize the files used to store quota information? (Specify ONLY the command without any path or parameters.)


## 78- Which of the following file permissions belong to a symbolic link?
- A. -rwxrwxrwx
- B. +rwxrwxrwx
- C.	lrwxrwxrwx
- D.	srwxrwxrwx


## 79- Creating a hard link to an ordinary file returns an error. What could be the reason for this?
- A.	The source file is hidden.
- B.	The source file is read-only.
- C.	The source file is a shell script.
- D.	The source file is already a hard link.
- E.	The source and the target are on different filesystems.

## 80- Which of the following commands creates an ext3 filesystem on /dev/sdb1? (Choose TWO correct answers.)
- A.	/sbin/mke2fs -j /dev/sdb1
- B.	/sbin/mkfs -t ext3 /dev/sdb1
- C.	/sbin/mkfs -c ext3 /dev/sdb1
- D.	/sbin/mke3fs -j /dev/sdb1


## 81- Which of the following commands will change the quota for a specific user?
- A.	edquota
- B.	repquota
- C.	quota -e
- D.	quota


## 82- Which utility would be used to change how often a filesystem check is performed on an ext2 filesystem without losing any data stored on that filesystem?
- A.	mod2fs
- B.	fsck
- C.	tune2fs
- D.	mke2fs
- E.	fixe2fs


## 83- Which of the following Linux filesystems preallocates a fixed number of inodes at the filesystem's make/ creation time and does NOT generate them as needed? (Choose TWO correct answers.)
- A.	ext3
- B.	JFS
- C.	ext2
- D.	XFS
- E.	procfs


## 84- What is the purpose of the Filesystem Hierarchy Standard?
- A.	It is a security model used to ensure files are organized according to their permissions and accessibility.
- B.	It provides unified tools to create, maintain and manage multiple filesystems in a common way.
- C.	It defines a common internal structure of inodes for all compliant filesystems.
- D.	It is a distribution neutral description of locations of files and directories.


## 85- SIMULATION
Which umask value will result in the default access permissions of 600 (rw-------) for files and 700 (rwx------) for directories? (Specify only the numerical umask valu- E.)
 
ANSWERS:
-------------------------------------------------------------------------------------------
70	Correct Answer: D

71	Correct Answer: B

72	Correct Answer: AC

73	Correct Answer: E

74	Correct Answer: C

75	Correct Answer: A

76	Correct Answer: D

77	Correct Answer: quotacheck

78	Correct Answer: C

79	Correct Answer: E

80	Correct Answer: AB

81	Correct Answer: A

82	Correct Answer: C

83	Correct Answer: AC

84	Correct Answer: D

85	Correct Answer: 0077, 077
 



