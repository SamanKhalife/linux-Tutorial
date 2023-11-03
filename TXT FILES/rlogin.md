# rlogin

The `rlogin` command in Linux is a deprecated command that was used to log in to a remote system from a local system. It is a simple protocol that uses unencrypted passwords. This makes it vulnerable to attacks, such as password sniffing.

Instead of using the `rlogin` command, you should use the `ssh` command to log in to a remote system from a local system. The `ssh` command is a more secure and efficient way to log in to a remote system from a local system. It is also supported by most Linux distributions.

The `rlogin` command is used to allow users to log in to a remote system from a local system. It is a simple protocol that uses unencrypted passwords. This makes it vulnerable to attacks, such as password sniffing.

The `ssh` command is used to log in to a remote system from a local system. It is a more secure protocol that uses encrypted passwords. This makes it much more secure than `rlogin`.

Here are some of the benefits of using `ssh` instead of `rlogin`:

* It is more secure.
* It is more efficient.
* It is supported by most Linux distributions.

Here are some of the drawbacks of using `ssh` instead of `rlogin`:

* It can be more difficult to configure.
* It can be slower than `rlogin` in some cases.

Overall, the `ssh` command is a better choice than `rlogin` for logging in to a remote system from a local system. It is more secure, efficient, and widely supported.

# help 

```
unknown option -- -
usage: ssh [-46AaCfGgKkMNnqsTtVvXxYy] [-B bind_interface]
           [-b bind_address] [-c cipher_spec] [-D [bind_address:]port]
           [-E log_file] [-e escape_char] [-F configfile] [-I pkcs11]
           [-i identity_file] [-J [user@]host[:port]] [-L address]
           [-l login_name] [-m mac_spec] [-O ctl_cmd] [-o option] [-p port]
           [-Q query_option] [-R address] [-S ctl_path] [-W host:port]
           [-w local_tun[:remote_tun]] destination [command [argument ...]]
```
