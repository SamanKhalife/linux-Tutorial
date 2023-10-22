# cd

The cd utility is one of the system utilities that allows you to navigate to a specific directory.

	this command changes the directory to the home dir from any where
	
	root@Saman:~# cd /home/

	if you what to go in dir in another dir you can

	root@Saman:~# cd DIR1/DIR2/DIR3/

	this command is the wayback dir

	root@Saman:~# cd ..

	for home dir from anywhere

	root@Saman:~# cd ~
```
 help

    Options:
      -L        force symbolic links to be followed: resolve symbolic
                links in DIR after processing instances of `..'
      -P        use the physical directory structure without following
                symbolic links: resolve symbolic links in DIR before
                processing instances of `..'
      -e        if the -P option is supplied, and the current working
                directory cannot be determined successfully, exit with
                a non-zero status
      -@        on systems that support it, present a file with extended
                attributes as a directory containing the file attribute
```
