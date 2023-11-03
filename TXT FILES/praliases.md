# praliases

The `praliases` command in Linux is used to display the current system aliases, one per line, in no particular order. The special internal @:@ alias will be displayed if present.

The syntax for the `praliases` command is as follows:

```
praliases [options]
```

The `options` argument can be used to control the output of the command.

Some of the most useful `praliases` options include:

* `-a`: Display all aliases, including the special internal @:@ alias.
* `-f file`: Read aliases from the specified file instead of the default system aliases file.
* `-k key`: Display aliases that match the specified key.

Here is an example of how to use the `praliases` command to display all aliases:

```
praliases -a
```

This command will display all aliases, including the special internal @:@ alias.

Here is an example of how to use the `praliases` command to display aliases that match the key `root`:

```
praliases -k root
```

This command will display all aliases that match the key `root`.

The `praliases` command is a useful tool for viewing the current system aliases. It can be used to troubleshoot problems with aliases, or to simply see what aliases are defined on the system.

Here are some of the benefits of using the `praliases` command:

* It can be used to troubleshoot problems with aliases.
* It can be used to see what aliases are defined on the system.
* It can be used to create new aliases.
* It can be used to delete aliases.

If you are using aliases on your system, you should make sure to learn how to use the `praliases` command. It is a valuable tool for viewing and managing aliases.



# help 

```
praliases: invalid option -- '-'
usage: praliases [-C cffile] [-f aliasfile] [key ...]
```
