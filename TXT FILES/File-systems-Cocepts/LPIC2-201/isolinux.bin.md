# ISOLINUX

**ISOLINUX** is a bootloader for Linux systems that is part of the Syslinux Project. It is specifically designed for booting from ISO 9660 (CD-ROM) filesystems. ISOLINUX is commonly used for creating bootable live CDs and DVDs, providing a flexible and easy-to-use environment for launching Linux distributions from optical media.

### Overview of ISOLINUX

**ISOLINUX** operates similarly to other bootloaders in the Syslinux family (such as SYSLINUX for FAT filesystems and PXELINUX for network booting) but is optimized for ISO 9660 filesystems. It uses the El Torito boot specification to boot Linux systems from CD-ROMs.

### Installing ISOLINUX

To create a bootable ISO image with ISOLINUX, follow these steps:

1. **Install the Syslinux Package**:
   - On Debian-based systems:
     ```bash
     sudo apt-get install syslinux
     ```
   - On Red Hat-based systems:
     ```bash
     sudo yum install syslinux
     ```

2. **Prepare the Directory Structure**:
   - Create a directory structure for your ISO image:
     ```bash
     mkdir -p iso/boot/isolinux
     ```

3. **Copy ISOLINUX Files**:
   - Copy the `isolinux.bin` file and the ISOLINUX configuration file (`isolinux.cfg`) to the `iso/boot/isolinux` directory:
     ```bash
     cp /usr/lib/syslinux/isolinux.bin iso/boot/isolinux/
     touch iso/boot/isolinux/isolinux.cfg
     ```

### Configuring ISOLINUX

The ISOLINUX configuration file (`isolinux.cfg`) controls the boot process. Here is an example configuration:

1. **Create the Configuration File**:
   - Open the `isolinux.cfg` file in a text editor:
     ```bash
     nano iso/boot/isolinux/isolinux.cfg
     ```

2. **Add Configuration Entries**:
   - An example configuration might look like this:
     ```ini
     DEFAULT linux
     LABEL linux
         KERNEL /boot/vmlinuz
         APPEND initrd=/boot/initrd.img root=/dev/ram0 rw
     ```

### Creating the Bootable ISO

Once the ISOLINUX bootloader and configuration are in place, you can create the bootable ISO image using the `mkisofs` or `genisoimage` command:

1. **Create the ISO Image**:
   ```bash
   mkisofs -o bootable.iso -b boot/isolinux/isolinux.bin -c boot/isolinux/boot.cat \
           -no-emul-boot -boot-load-size 4 -boot-info-table iso/
   ```

2. **Explanation of Parameters**:
   - `-o bootable.iso`: Specifies the output ISO file.
   - `-b boot/isolinux/isolinux.bin`: Specifies the path to the ISOLINUX bootloader.
   - `-c boot/isolinux/boot.cat`: Specifies the path to the boot catalog file.
   - `-no-emul-boot`: Indicates no emulation mode.
   - `-boot-load-size 4`: Specifies the number of virtual sectors to load (usually 4).
   - `-boot-info-table`: Adds a boot information table to the ISO.
   - `iso/`: Specifies the path to the directory containing the ISO image files.

### Example: Full ISOLINUX Setup

Let’s go through a full example of setting up and creating a bootable ISO image with ISOLINUX:

1. **Install Syslinux Package**:
   ```bash
   sudo apt-get install syslinux
   ```

2. **Prepare Directory Structure**:
   ```bash
   mkdir -p iso/boot/isolinux
   ```

3. **Copy ISOLINUX Files**:
   ```bash
   cp /usr/lib/syslinux/isolinux.bin iso/boot/isolinux/
   touch iso/boot/isolinux/isolinux.cfg
   ```

4. **Configure ISOLINUX**:
   ```bash
   nano iso/boot/isolinux/isolinux.cfg
   ```
   Add the following content:
   ```ini
   DEFAULT linux
   LABEL linux
       KERNEL /boot/vmlinuz
       APPEND initrd=/boot/initrd.img root=/dev/ram0 rw
   ```

5. **Copy Kernel and Initrd**:
   ```bash
   cp /path/to/vmlinuz iso/boot/
   cp /path/to/initrd.img iso/boot/
   ```

6. **Create the ISO Image**:
   ```bash
   mkisofs -o bootable.iso -b boot/isolinux/isolinux.bin -c boot/isolinux/boot.cat \
           -no-emul-boot -boot-load-size 4 -boot-info-table iso/
   ```

### Booting with ISOLINUX

After creating the bootable ISO image, you can burn it to a CD/DVD or use it with virtualization software like VirtualBox or VMware. Booting from the CD/DVD will load the ISOLINUX bootloader, which will then load the kernel and initrd as specified in the `isolinux.cfg` file.

### Advantages of ISOLINUX

- **Specialized for CDs/DVDs**: Optimized for booting from ISO 9660 filesystems, making it ideal for live CDs/DVDs.
- **Simple Configuration**: Easy to set up and configure compared to more complex bootloaders like GRUB.
- **Lightweight**: Minimal resource usage, suitable for live environments and rescue disks.

### Conclusion

ISOLINUX is a powerful and lightweight bootloader for booting Linux systems from CD-ROMs. Its ease of installation and configuration, combined with its specialization for ISO 9660 filesystems, makes it an excellent choice for creating bootable live CDs and DVDs. Understanding how to use and configure ISOLINUX can be a valuable skill for creating custom bootable media and rescue disks.
