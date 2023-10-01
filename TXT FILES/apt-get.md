# apt-get

The apt-get command is a command-line utility that can be used to install, remove, and update software packages on Debian-based Linux distributions. It is a powerful tool that can be used to manage the software on your system.

The apt-get command is used as follows:

`apt-get [options] command`

command: This is the command that you want to run. Some of the most commonly used apt-get commands are:
- install: This command is used to install software packages.
- remove: This command is used to remove software packages.
- update: This command is used to update the list of available software packages.
- upgrade: This command is used to upgrade installed software packages to the latest version.




# help

```
Usage: apt-get [options] command
       apt-get [options] install|remove pkg1 [pkg2 ...]
       apt-get [options] source pkg1 [pkg2 ...]

apt-get is a command line interface for retrieval of packages
and information about them from authenticated sources and
for installation, upgrade and removal of packages together
with their dependencies.

Most used commands:
  update - Retrieve new lists of packages
  upgrade - Perform an upgrade
  install - Install new packages (pkg is libc6 not libc6.deb)
  reinstall - Reinstall packages (pkg is libc6 not libc6.deb)
  remove - Remove packages
  purge - Remove packages and config files
  autoremove - Remove automatically all unused packages
  dist-upgrade - Distribution upgrade, see apt-get(8)
  dselect-upgrade - Follow dselect selections
  build-dep - Configure build-dependencies for source packages
  satisfy - Satisfy dependency strings
  clean - Erase downloaded archive files
  autoclean - Erase old downloaded archive files
  check - Verify that there are no broken dependencies
  source - Download source archives
  download - Download the binary package into the current directory
  changelog - Download and display the changelog for the given package
```
