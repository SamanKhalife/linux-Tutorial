# quotaon

The `quotaon` command is a Linux command that is used to enable disk quotas on a file system. Disk quotas are limits that can be placed on the amount of disk space that a user or group can use on a file system.

The syntax for the `quotaon` command is:

```
quotaon [options] [filesystem]
```

The options are:

* `-a` or `--all`: Enables quotas on all file systems.
* `-g` or `--group`: Enables quotas for the specified group.
* `-u` or `--user`: Enables quotas for the specified user.
* `-v` or `--verbose`: Prints verbose output.

If no options are specified, the `quotaon` command will enable quotas on the default file system.

For example, to enable quotas for all file systems on the system, you would use the following command:

```
quotaon -a
```

This would enable quotas on all file systems, including the root file system.

To enable quotas for the user `saman` on the file system `/dev/sda1`, you would use the following command:

```
quotaon -u saman /dev/sda1
```

This would enable quotas for the user `saman` on the file system `/dev/sda1`.

The `quotaon` command is a useful tool for enabling disk quotas on your system. It can be used to enable quotas for users or groups who are using the system, or to enable quotas for file systems that are being used.

Here are some of the benefits of using the `quotaon` command:

* It can help to prevent users from exceeding their disk quota limits.
* It can help to ensure that all users have fair access to disk space.
* It can help to identify and fix problems with the disk quotas.
* It can help to improve the performance of the file system.

If you are using disk quotas on your system, you should make sure to use the `quotaon` command to enable them. This will help to prevent users from exceeding their disk quota limits and ensure that all users have fair access to disk space.





# help 

```
quotaon: Usage:
        quotaon [-guPvp] [-F quotaformat] [-x state] -a
        quotaon [-guPvp] [-F quotaformat] [-x state] filesys ...

-a, --all                turn quotas on for all filesystems
-f, --off                turn quotas off
-u, --user               operate on user quotas
-g, --group              operate on group quotas
-P, --project            operate on project quotas
-p, --print-state        print whether quotas are on or off
-x, --xfs-command=cmd    perform XFS quota command
-F, --format=formatname  operate on specific quota format
-v, --verbose            print more messages
-h, --help               display this help text and exit
-V, --version            display version information and exit
```
