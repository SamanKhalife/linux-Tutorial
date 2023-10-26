# lftpget 

lftpget is a shell script that is used to download files using the lftp command. It is a simple and convenient way to download files from the internet, and it supports a variety of protocols, including FTP, HTTP, and HTTPS.

The syntax of the lftpget command is as follows:

```
lftpget [options] URL [URL...]
```

The `options` argument specifies additional options for lftpget. The most common options are as follows:

* `-c`: Continue a previous download.
* `-d`: Debug output.
* `-v`: Verbose messages.

For example, the following command downloads the file `index.html` from the website `www.example.com`:

```
lftpget http://www.example.com/index.html
```

The lftpget command is a useful tool for downloading files from the internet. It is simple to use and supports a variety of protocols, making it a versatile option for downloading files.

Here are some additional things to keep in mind about the lftpget command:

* The lftpget command must be run as root if you are downloading files to a directory that requires root privileges.
* The lftpget command does not support resuming downloads that were interrupted by a network error.
* The lftpget command does not support downloading files from BitTorrent or other peer-to-peer networks.

It is important to be aware of these limitations when using the lftpget command, so that you do not accidentally download files to a directory that you do not have permissions to access, or so that you do not lose your download if the network connection is interrupted.



# help 

```

```
