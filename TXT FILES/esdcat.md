# esdcat

eject is a command-line utility that is used to eject removable media, such as optical discs and USB drives. It is typically used when you are finished using the media and want to remove it from your computer.

The `esdcat` command is used in the following syntax:

```
esdcat [options] file
```



For example, to eject the optical disc in the drive /dev/cdrom, you would use the following command:

`eject /dev/cdrom`

To eject the USB drive in the port /dev/sdb, you would use the following command:

`eject /dev/sdb`

To force the eject of the optical disc in the drive /dev/cdrom, you would use the following command:

`eject -f /dev/cdrom`

To specify a timeout of 10 seconds before the eject command will time out, you would use the following command:

`eject -t 10 /dev/cdrom`


The `file` is the ESD file that you want to display.

The options can be used to specify the following:

* `-l` : List the contents of the ESD file.
* `-x` : Extract the contents of the ESD file to a directory.
* `-v` : Be more verbose in the output of esdcat.


# help 

```
eject [options] [device]

Eject removable media.

Options:

-f, --force        Force eject, even if medium is in use.
-t, --timeout=SECS  Timeout after SECS seconds, 0 waits forever.
-h, --help          Show this help message.

```
