# who

The `who` command is a Linux command that displays a list of users who are currently logged in to the system. It is a versatile command that can be used to quickly get a list of users who are logged in and what they are doing.

The `who` command is used as follows:

```
who
```

The `who` command will display the following information about each user who is logged in:

* Username: The username of the user.
* Terminal: The terminal that the user is logged into.
* PID: The process ID of the user's shell.
* Time: The time that the user logged in.
* Idle: The amount of time that the user has been idle.
* What: The command that the user is currently running.

For example, the following output from the `who` command shows that the user `johndoe` is logged into the terminal `tty1` and has been idle for 2 minutes. The user is currently running the command `vim`.

```
$ who

 13:58:35 up 4 days, 3:42,  1 user,  load average: 0.00, 0.01, 0.05
USER     TTY      FROM              LOGIN@   IDLE   JCPU   PCPU WHAT
johndoe   tty1     pts/0            13:56    2m    0.09s  0.09s vim
```

The `who` command is a valuable tool for system administrators who need to monitor the activity of users on their system. It can also be used by users to see who else is logged in and what they are doing.

Here are some other options that can be used with the `who` command:

* `-a`: This option will display all users who are logged in, including users who are logged in via remote terminals.
* `-l`: This option will display additional information about each user, such as their real name and home directory.
* `-r`: This option will display a list of users who are currently logged out, but who have not yet logged out completely.

The `who` command is a powerful tool that can be used to get information about the users who are logged in to your system. It is a valuable tool for system administrators and users alike.
# help 

```
Usage: who [OPTION]... [ FILE | ARG1 ARG2 ]
Print information about users who are currently logged in.

  -a, --all         same as -b -d --login -p -r -t -T -u
  -b, --boot        time of last system boot
  -d, --dead        print dead processes
  -H, --heading     print line of column headings
      --ips         print ips instead of hostnames. with --lookup,
                    canonicalizes based on stored IP, if available,
                    rather than stored hostname
  -l, --login       print system login processes
      --lookup      attempt to canonicalize hostnames via DNS
  -m                only hostname and user associated with stdin
  -p, --process     print active processes spawned by init
  -q, --count       all login names and number of users logged on
  -r, --runlevel    print current runlevel
  -s, --short       print only name, line, and time (default)
  -t, --time        print last system clock change
  -T, -w, --mesg    add user's message status as +, - or ?
  -u, --users       list users logged in
      --message     same as -T
      --writable    same as -T
      --help     display this help and exit
      --version  output version information and exit

If FILE is not specified, use /var/run/utmp.  /var/log/wtmp as FILE is common.
If ARG1 ARG2 given, -m presumed: 'am i' or 'mom likes' are usual.
```

