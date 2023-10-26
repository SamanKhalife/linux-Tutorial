# gpgsplit

gpgsplit is a command-line tool that splits an OpenPGP message into individual packets. This can be useful for distributing large messages or for storing messages in a file system.

The syntax for the gpgsplit command is:

```
gpgsplit [options] [files]
```

The `options` that you can use with the gpgsplit command include:

* `-v`, `--verbose`: Verbose output.
* `-p`, `--prefix`: Prepend filenames with STRING.
* `--uncompress`: Uncompress a packet.
* `--secret-to-public`: Convert secret keys to public keys.
* `--no-split`: Write to stdout and don't actually split.

The `files` are the OpenPGP messages that you want to split. If no files are specified, the message is read from standard input.

The gpgsplit command will create a new file for each packet. The filenames will be prefixed with the string that you specify with the `-p` option. If you do not specify a prefix, the filenames will be the numbers of the packets.

For example, to split the OpenPGP message in the file `message.asc` into 100 packets, you would use the following command:

```
gpgsplit -p 100 message.asc
```

This would create 100 files, numbered from 001 to 100. Each file would contain one packet from the original message.

gpgsplit is a useful tool for splitting OpenPGP messages. It can be used to distribute large messages or for storing messages in a file system.




# help 

```

```
