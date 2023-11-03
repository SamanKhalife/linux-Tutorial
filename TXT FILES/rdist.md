# rdist

The `rdist` command in Linux is used to distribute files and directories to multiple hosts. It is a powerful tool that can be used to synchronize files and directories across a network.

The `rdist` command is used in the following syntax:

```
rdist [options] [source] [target]
```

The `source` is the directory or file to distribute.

The `target` is the directory or file to copy the source to.

The `options` can be used to specify the following:

* `-r` : Recursively distribute the source to the target.
* `-u` : Update the target with the latest version of the source.
* `-d` : Delete the target before distributing the source.
* `-c` : Create the target directory if it does not exist.

For example, to distribute the directory `/home/user/data` to the directory `/opt/data` on all hosts in the `mynetwork` network, you would run the following command:

```
rdist -r -u /home/user/data /opt/data mynetwork
```

This command will recursively distribute the directory `/home/user/data` to the directory `/opt/data` on all hosts in the `mynetwork` network. The latest version of the files in `/home/user/data` will be copied to `/opt/data`.

To update the directory `/opt/data` with the latest version of the files in `/home/user/data`, you would run the following command:

```
rdist -u /home/user/data /opt/data
```

This command will update the directory `/opt/data` with the latest version of the files in `/home/user/data`.

To delete the directory `/opt/data` before distributing the source, you would run the following command:

```
rdist -r -c -u /home/user/data /opt/data
```

This command will recursively distribute the directory `/home/user/data` to the directory `/opt/data`. The directory `/opt/data` will be deleted before the distribution starts.

To create the directory `/opt/data` if it does not exist, you would run the following command:

```
rdist -r -c /home/user/data /opt/data
```

This command will recursively distribute the directory `/home/user/data` to the directory `/opt/data`. If the directory `/opt/data` does not exist, it will be created.

The `rdist` command is a powerful tool that can be used to synchronize files and directories across a network. It is a versatile command that can be used to distribute files and directories to multiple hosts.



# help 

```
rdist: invalid option -- '-'
Usage: rdist [-cDFnv] [-A <num>] [-a <num>] [-d var=value]
        [-f distfile] [-l <msgopt>] [-L <msgopt>] [-M <maxproc>]
        [-m host] [-o <distopts>] [-p <rdistd-cmd>] [-P <rsh-path>]
        [-t <timeout>] [target ...]
OR:    rdist [-cDFnv] -c source [...] machine[:dest]
OR:    rdist -V

The values for <distopts> are:
        chknfs,chkreadonly,chksym,compare,follow,ignlnks,nochkgroup,nochkmode,nochkowner,nodescend,noexec,numchkgroup,numchkowner,quiet,remove,savetargets,sparse,verify,whole,younger

Where <msgopt> is of form
        <facility1>=<type1>,<type2>,...:<facility2>=<type1>,<type2>...
Valid <facility> names: stdout file syslog notify
Valid <type> names: change info notice nerror ferror warning verbose all debug
```
