# sfto

SFTP stands for Secure File Transfer Protocol. It is a secure version of the File Transfer Protocol (FTP) that uses the SSH (Secure Shell) protocol to encrypt all data transfers. This makes SFTP a very secure way to transfer files between computers.

Here are some examples of how to use the sftp command:

```
# To connect to the remote host `example.com` as the user `username`:
sftp username@example.com

# To connect to the remote host `example.com` on port 2222:
sftp -p 2222 username@example.com

# To connect to the remote host `example.com` and force the use of the 3DES cipher:
sftp -o cipher=3des username@example.com
```

The sftp command is a powerful tool that can be used to transfer files securely between computers. It is a simple command to use, but it can be very effective.


# help 

```
sftp [options] [user@]host[:port]

SFTP client.

Options:

-v, --verbose       verbose mode
-p, --port          port number to connect to
-o, --option        option to set
-F, --config        specify config file
-i, --identity      identity file to use
-b, --batch          batch mode
-l, --login-timeout login timeout
-q, --quiet         quiet mode
-k, --knownhosts     known hosts file
-R, --remote-command command to execute on remote side
-S, --subsystem     subsystem to execute
-h, --help          show this help message

For more information, see the sftp man page.
```
## breakdown 

```
-v: This option tells sftp to be verbose and print out more information about the transfer.
-p: This option tells sftp to use the specified port number.
-o: This option tells sftp to use the specified option. For example, -o cipher=3des will force sftp to use the 3DES cipher.
-F: This option tells sftp to use the specified config file.
-i: This option tells sftp to use the specified identity file.
-b: This option tells sftp to enter batch mode.
-l: This option tells sftp to set the login timeout.
-q: This option tells sftp to be quiet.
-k: This option tells sftp to use the specified known hosts file.
-R: This option tells sftp to execute the specified command on the remote side.
-S: This option tells sftp to execute the specified subsystem.
-h: This option shows this help message.
```
