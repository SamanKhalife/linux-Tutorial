# /sys/


The `/sys` directory is an essential part of the Linux filesystem, especially related to the system's hardware and kernel. Here’s a detailed explanation of `/sys`:

### Understanding the `/sys` Directory

The `/sys` directory, also known as sysfs, is a virtual filesystem in Linux that provides information about the kernel, devices, and other system-related attributes. It was introduced in Linux kernel version 2.6. The `/sys` directory is used to export kernel data structures, their attributes, and relationships to userspace. It is mounted automatically by the system at boot time.

#### Key Characteristics of `/sys`:
1. **Virtual Filesystem**: Similar to `/proc`, the files in `/sys` do not exist on disk; they are created dynamically by the kernel.
2. **Hierarchy**: The structure of `/sys` mirrors the internal kernel structures. It provides a hierarchical view of devices and their attributes.
3. **Read and Write**: Unlike `/proc`, which is mostly read-only, many files in `/sys` can be written to, allowing for dynamic configuration of hardware parameters.

#### Key Subdirectories in `/sys`:

1. **/sys/class/**:
   - Contains directories for each class of device. Device classes include `block`, `net`, `tty`, etc.
   - Example: `/sys/class/net` contains directories for each network interface (e.g., `eth0`, `lo`).

2. **/sys/block/**:
   - Contains directories for each block device. Block devices include hard drives and other storage devices.
   - Example: `/sys/block/sda` for the first SCSI disk.

3. **/sys/bus/**:
   - Contains directories for each bus type in the system (e.g., `pci`, `usb`).
   - Example: `/sys/bus/usb` contains directories for USB devices.

4. **/sys/devices/**:
   - Contains directories for all devices, organized by their physical hierarchy.
   - Example: `/sys/devices/pci0000:00` for PCI devices.

5. **/sys/firmware/**:
   - Contains directories for firmware-related information.
   - Example: `/sys/firmware/efi` for EFI firmware variables.

#### Examples of Using `/sys`:

1. **Listing Network Interfaces**:
   ```sh
   ls /sys/class/net/
   ```
   Output might include:
   ```
   eth0  lo  wlan0
   ```

2. **Getting Device Information**:
   ```sh
   cat /sys/class/net/eth0/address
   ```
   Output might be the MAC address of `eth0`:
   ```
   01:23:45:67:89:ab
   ```

3. **Modifying Device Attributes**:
   - Example: Changing the brightness of a laptop screen.
   ```sh
   echo 100 > /sys/class/backlight/acpi_video0/brightness
   ```

#### Practical Use Cases:

1. **Monitoring Hardware Health**:
   - You can read sensor data like CPU temperature:
     ```sh
     cat /sys/class/thermal/thermal_zone0/temp
     ```

2. **Managing Power**:
   - You can manage power settings for devices, like enabling or disabling USB ports:
     ```sh
     echo 'auto' > /sys/bus/usb/devices/usb1/power/control
     ```

3. **Configuring Kernel Parameters**:
   - You can configure various kernel parameters dynamically without needing a reboot:
     ```sh
     echo 1 > /sys/class/net/eth0/device/sriov_numvfs
     ```

### Conclusion

The `/sys` directory is a powerful tool for Linux system administrators, providing a way to query and control kernel and hardware attributes dynamically. Understanding and utilizing `/sys` can greatly enhance your ability to manage and optimize your Linux system.

Feel free to ask for explanations of other objectives or for more examples and details!
