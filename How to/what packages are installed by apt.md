To see what packages are installed by the `apt` package manager on a Debian-based Linux system, you can use the `dpkg` command or the `apt` command. Here's how you can do it:

1. Using `dpkg`:

   Open a terminal and run the following command:

   ```bash
   dpkg --list
   ```

   This will list all installed packages on your system. You can also pipe the output to a pager like `less` for easier navigation:

   ```bash
   dpkg --list | less
   ```

   To search for a specific package, you can use `grep`. For example, if you want to find packages related to Python:

   ```bash
   dpkg --list | grep python
   ```

2. Using `apt`:

   Another way to list installed packages is to use `apt` with the `list` option:

   ```bash
   apt list --installed
   ```

   This command will display a list of all installed packages.

Keep in mind that you may need superuser privileges (using `sudo`) to run these commands, as listing installed packages typically requires administrative access.