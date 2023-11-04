# sysctl

The sysctl command is used to view and change kernel parameters. Kernel parameters are settings that control the behavior of the Linux kernel.

For example, to view the value of the kernel parameter net.ipv4.tcp_syncookies, you would use the following command:

`sysctl net.ipv4.tcp_syncookies`

To change the value of the kernel parameter net.ipv4.tcp_syncookies to 1, you would use the following command:

`sysctl -w net.ipv4.tcp_syncookies=1`

# help 

```
Usage:
 sysctl [options] [variable[=value] ...]

Options:
  -a, --all            display all variables
  -A                   alias of -a
  -X                   alias of -a
      --deprecated     include deprecated parameters to listing
  -b, --binary         print value without new line
  -e, --ignore         ignore unknown variables errors
  -N, --names          print variable names without values
  -n, --values         print only values of the given variable(s)
  -p, --load[=<file>]  read values from file
  -f                   alias of -p
      --system         read values from all system directories
  -r, --pattern <expression>
                       select setting that match expression
  -q, --quiet          do not echo variable set
  -w, --write          enable writing a value to variable
  -o                   does nothing
  -x                   does nothing
  -d                   alias of -h

 -h, --help     display this help and exit
 -V, --version  output version information and exit

For more details see sysctl(8).
```

## breakdown

```
-a, --all: This option shows all kernel parameters.
-w, --value=VALUE: This option sets the value of a kernel parameter.
-p, --file=FILE: This option reads kernel parameters from a file.
-h, --help: This option shows this help message.
```
