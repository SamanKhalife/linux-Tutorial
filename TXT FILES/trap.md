# trap

The trap command in Linux is used to define what action should be taken when a certain signal is received. Signals are software interrupts that can be used to notify a process that something has happened, such as a keyboard interrupt or a termination signal.

For example, to trap the SIGINT signal and run the command echo "Received SIGINT", you would use the following command:

trap "echo Received SIGINT" SIGINT

# help 

```
trap: trap [-lp] [[arg] signal_spec ...]

Define a signal handler.

Options:

-l, --list            List all of the currently defined traps.
-p, --print           Print the trap definition for a signal.
-h, --help            Show this help message.
```



## breakdown

```
-l, --list: This option lists all of the currently defined traps.
-p, --print: This option prints the trap definition for a signal.
-h, --help: This option shows this help message.
```
