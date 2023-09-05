# accton


The accton command is a deprecated command in Linux. It was used to control the accounting system, which tracks the resources that a user or process uses. The accton command is no longer necessary, as the accounting system is now controlled by the /etc/acct file.

The OPTION argument is a single option that can be used to control the accounting system. The following options are supported:

`on`: This option turns on the accounting system.
`off`: This option turns off the accounting system.
`ACCOUNTING_FILE`: This option specifies the file that the accounting system will use.

Here are some examples of how to use the accton command:

```
# To turn on the accounting system:
accton on

# To turn off the accounting system:
accton off

# To specify the accounting file:
accton ACCOUNTING_FILE
```

# help 

```
Usage: accton [OPTION] on|off|ACCOUNTING_FILE

          Turns process accounting on or off, or changes the file where this
          info is saved.

          OPTIONS:
          -h, --help       Show help and exit
          -V, --version    Show version and exit

          ARGUMENTS:
          on               Activate process accounting and use default file
          off              Deactivate process accounting
          ACCOUNTING_FILE  Activate (if not active) and save information in
          this file
```

