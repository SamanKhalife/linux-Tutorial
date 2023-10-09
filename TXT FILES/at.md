# at 
The `at` command in Linux is used to schedule a command or script to run at a later time. It is a powerful tool that can be used to automate tasks on your system.

The `at` command is used in the following syntax:

```
at [options] [time] [command]
```

The `time` is the time that you want the command or script to run. The `command` is the command or script that you want to run.

The `options` can be used to specify the following:

* `-f` : Specify a file that contains the command or script to run.
* `-m` : Send a mail notification when the command or script has finished running.
* `-r` : Remove the job from the at queue.

For example, to schedule the command `ls` to run at 10:00 AM, you would run the following command:

```
at 10:00 am ls
```

This command will schedule the command `ls` to run at 10:00 AM.

To schedule the script `/home/user/myscript.sh` to run at 11:00 AM, you would run the following command:

```
at 11:00 am -f /home/user/myscript.sh
```

This command will schedule the script `/home/user/myscript.sh` to run at 11:00 AM.

To send a mail notification when the command or script has finished running, you would run the following command:

```
at 12:00 pm -m ls
```

This command will schedule the command `ls` to run at 12:00 PM and will send a mail notification when the command has finished running.

To remove the job from the at queue, you would run the following command:

```
at -r 12345
```

This command will remove the job with the ID 12345 from the at queue.

The `at` command is a powerful tool that can be used to automate tasks on your system. It is a versatile tool that can be used to schedule commands or scripts to run at any time.

Here are some additional things to note about the `at` command:

* The `at` command can be used to schedule commands or scripts to run at any time.
* The `at` command can be used to schedule commands or scripts to run on a recurring basis.
* The `at` command can be used to schedule commands or scripts to run even if you are not logged in to your system.
* The `at` command is a powerful tool that can be used to automate tasks on your system.




# help 

```

```
