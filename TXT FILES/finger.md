# finger

The finger command in Linux is a command-line utility that can be used to get information about users on a system. It can be used to display the user's name, real name, login name, shell, and other information.

The user argument is the name of the user that you want to get information about.

The finger command is a powerful tool that can be used to get information about users on a system. It can be used to troubleshoot problems with users, to manage users, and to learn more about users.


For example, the following command will get information about the user root:

`finger root`

The finger command will output the following information:

```
Login: root
Name: root
Directory: /root
Shell: /bin/bash
On since Wed Jun 22 14:39:33 2022

No mail.

Last login Wed Jun 22 14:39:33 2022 from 192.168.1.1
```


# help 

```
finger [options] user

Get information about a user.

Options:

-l, --long   Display more information about the user.
-p, --plan   Display the user's plan.
-s, --silent   Suppress the display of the user's real name.
-h, --help   Display this help message.
```



## breakdown

```
-l, --long: This option tells finger to display more information about the user, including the user's home directory, login shell, and last login.
-p, --plan: This option tells finger to display the user's plan, which is a message that the user can set to be displayed when other users finger them.
-s, --silent: This option tells finger to suppress the display of the user's real name.
-h, --help: This option displays the help message for the finger command.
```
