# pxelinux.cfg/
The `pxelinux.cfg/` directory is a key component in a PXE (Preboot Execution Environment) boot setup. It contains configuration files that dictate how PXE clients (computers booting over the network) behave during the boot process. Here’s an overview of `pxelinux.cfg/` and its purpose:

### Purpose of `pxelinux.cfg/`

1. **Configuration Storage:**
   - `pxelinux.cfg/` is a directory typically located on the TFTP (Trivial File Transfer Protocol) server used for PXE booting.
   - It stores configuration files that determine which boot options are available to PXE clients and how those options are configured.

2. **File Naming Convention:**
   - Configuration files within `pxelinux.cfg/` follow a specific naming convention:
     - The filename can be based on either the MAC address of the PXE client’s network interface card (NIC) or the IP address of the client.
     - Alternatively, a default configuration file named `default` can be used to define global settings or fallback options for all PXE clients.

3. **Configuration Options:**
   - Each configuration file in `pxelinux.cfg/` specifies various boot parameters, such as:
     - **Kernel Image (`KERNEL`):** Specifies the kernel file (`vmlinuz`) that the client should load.
     - **Initial RAM Disk (`INITRD`):** Specifies the initial RAM disk file (`initrd.img`) used during boot.
     - **Kernel Arguments (`APPEND`):** Provides additional arguments to be passed to the kernel during boot, such as root filesystem location, network settings, and other boot parameters.
     - **Labels (`LABEL`):** Defines different boot options or configurations within the same configuration file.

4. **Default Configuration (`default` file):**
   - The `default` file within `pxelinux.cfg/` is used when no specific configuration file matches the client’s MAC address or IP address.
   - It can define a default boot option or menu structure that applies to all PXE clients unless overridden by a specific configuration file.

5. **Customization:**
   - Administrators can customize `pxelinux.cfg/` to tailor boot options based on the requirements of different systems or network segments.
   - This customization allows for flexibility in deploying different operating systems, configurations, or rescue environments over the network.

### Example Configuration

Here’s an example of what a `pxelinux.cfg/` directory structure might look like:

```plaintext
pxelinux.cfg/
├── default
├── 01-23-45-67-89-ab-cd
├── 192.168.1.100
└── 192.168.1.101
```

- **`default`:** Contains the default boot configuration for all PXE clients.
- **`01-23-45-67-89-ab-cd`:** Configuration file based on the MAC address of a specific client.
- **`192.168.1.100` and `192.168.1.101`:** Configuration files based on the IP addresses of specific clients.

### Configuration File Example

```plaintext
# Example configuration for a specific client (01-23-45-67-89-ab-cd)
LABEL linux
    KERNEL vmlinuz
    APPEND initrd=initrd.img root=/dev/nfs nfsroot=192.168.1.10:/path/to/nfs/root ip=dhcp
```

In this example:
- `LABEL linux`: Defines a label for the boot option.
- `KERNEL vmlinuz`: Specifies the kernel file to load.
- `APPEND initrd=initrd.img root=/dev/nfs nfsroot=192.168.1.10:/path/to/nfs/root ip=dhcp`: Appends additional boot parameters passed to the kernel.

### Conclusion

`pxelinux.cfg/` plays a crucial role in PXE boot environments by storing configuration files that determine how PXE clients boot over the network. Understanding and properly configuring `pxelinux.cfg/` allows administrators to manage and deploy operating systems, recovery tools, and other utilities efficiently across networked computers.
