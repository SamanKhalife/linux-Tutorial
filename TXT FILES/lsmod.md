# lsmod

The `lsmod` command in Linux is used to display information about all the currently loaded kernel modules. Kernel modules are pieces of code that can be loaded and unloaded into the kernel upon demand, allowing the kernel to extend its capabilities without the need to reboot the system. 

### Understanding `lsmod`

The `lsmod` command provides a snapshot of the status of modules in the Linux kernel. It formats the contents of `/proc/modules`, which is the file that contains information about the modules that are currently loaded into the kernel.

#### Basic Usage

To display the currently loaded kernel modules, simply run:
```sh
lsmod
```

The output typically includes three columns:
1. **Module**: The name of the loaded module.
2. **Size**: The size of the module in bytes.
3. **Used by**: The number of instances that refer to this module and a list of referring modules.

#### Example Output

Here’s an example of what the `lsmod` command might return:
```sh
Module                  Size  Used by
nf_conntrack_netbios_ns 16384  0
nf_conntrack_broadcast  16384  1 nf_conntrack_netbios_ns
nf_nat_ftp              20480  0
nf_nat                 45056  1 nf_nat_ftp
nf_conntrack_ftp        24576  1 nf_nat_ftp
nf_conntrack          139264  5 nf_conntrack_ftp,nf_conntrack_broadcast,nf_conntrack_netbios_ns,nf_nat
libcrc32c               16384  2 nf_nat,nf_conntrack
```

#### Explanation of the Output

- **Module**: The name of the module. For example, `nf_conntrack_netbios_ns`.
- **Size**: The size of the module. For example, `16384` bytes.
- **Used by**: The number of instances using this module and other modules dependent on it. For example, `nf_conntrack_broadcast` uses `nf_conntrack_netbios_ns`.

### Practical Examples

1. **List All Loaded Modules**:
   ```sh
   lsmod
   ```
   This command lists all the currently loaded modules with their size and dependencies.

2. **Check if a Specific Module is Loaded**:
   If you want to check if a specific module, say `ext4`, is loaded, you can filter the output:
   ```sh
   lsmod | grep ext4
   ```
   If the module is loaded, you will see an output line corresponding to it. If not, there will be no output.

3. **Interpreting Dependencies**:
   To understand which modules depend on another module, look at the `Used by` column. For example:
   ```sh
   nf_nat_ftp              20480  0
   nf_nat                 45056  1 nf_nat_ftp
   ```
   This shows that `nf_nat_ftp` is used by `nf_nat`.

### Related Commands

1. **modprobe**: 
   Used to add or remove modules from the Linux kernel.
   - To load a module: 
     ```sh
     sudo modprobe module_name
     ```
   - To remove a module:
     ```sh
     sudo modprobe -r module_name
     ```

2. **insmod**: 
   Used to insert a module into the kernel.
   ```sh
   sudo insmod /path/to/module.ko
   ```

3. **rmmod**: 
   Used to remove a module from the kernel.
   ```sh
   sudo rmmod module_name
   ```

4. **modinfo**: 
   Displays information about a kernel module.
   ```sh
   modinfo module_name
   ```

### Conclusion

The `lsmod` command is a powerful tool for viewing the status of kernel modules in your system. It is essential for troubleshooting and managing the kernel’s functionality. Understanding how to use `lsmod` along with related commands like `modprobe`, `insmod`, and `rmmod` allows you to effectively manage kernel modules and optimize your Linux system.

Feel free to ask for more details or examples about other Linux commands or concepts!
