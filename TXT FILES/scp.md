# scp

The `scp` command in Linux is used to securely copy files between two hosts over a network. It is a versatile command that can be used to copy files between local and remote hosts, and between two remote hosts.

The `scp` command is used in the following syntax:

```
scp [options] [source] [destination]
```

The `options` are as follows:

* `-r`: Copies directories recursively.
* `-p`: Preserves file timestamps.
* `-q`: Quiet mode.
* `-i`: Specifies the SSH key file.
* `-h`: Displays help.

The `source` is the file or directory to copy.

The `destination` is the file or directory to copy to.

For example, to copy the file `myfile.txt` from the local host to the remote host `example.com`, you would use the following command:

```
scp myfile.txt example.com:/home/username/
```

This command will copy the file `myfile.txt` from the local host to the directory `/home/username/` on the remote host `example.com`.

To copy the directory `mydir` from the local host to the remote host `example.com`, you would use the following command:

```
scp -r mydir example.com:/home/username/
```

This command will copy the directory `mydir` and all of its contents from the local host to the directory `/home/username/` on the remote host `example.com`.

The `scp` command is a powerful tool that can be used to securely copy files between two hosts over a network. It is supported by most Linux distributions.

Here are some of the benefits of using `scp`:

* It is a secure way to copy files.
* It is supported by most Linux distributions.
* It is a free and open-source software.

Here are some of the drawbacks of using `scp`:

* It can be slow to copy large files.
* It can be difficult to troubleshoot if there are problems.
* The output of the command can be difficult to interpret.

The `scp` command is a powerful tool that can be used to securely copy files between two hosts over a network. However, it is important to use it carefully and to understand the potential risks before you use it.

# help 

```
usage: scp [-346ABCOpqRrsTv] [-c cipher] [-D sftp_server_path] [-F ssh_config]
           [-i identity_file] [-J destination] [-l limit]
           [-o ssh_option] [-P port] [-S program] source ... target

```
