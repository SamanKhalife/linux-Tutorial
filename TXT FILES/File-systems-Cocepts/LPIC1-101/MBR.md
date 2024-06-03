# MBR

MBR stands for Master Boot Record, and it is a crucial component of the partitioning scheme used in many legacy BIOS-based systems. The Master Boot Record resides at the very beginning of a storage device, typically the first sector (sector 0), and contains essential information for booting the operating system.

### Purpose of the MBR

1. **Boot Loader Location**: The MBR contains the boot loader code, such as GRUB or NTLDR, which is responsible for loading the operating system kernel into memory during the boot process.

2. **Partition Table**: The MBR also contains the partition table, which defines the layout of the disk and the locations and sizes of the partitions. The MBR partition table supports up to four primary partitions, or three primary partitions and one extended partition.

3. **Boot Signature**: The MBR ends with a two-byte signature (0x55AA) that indicates to the BIOS that the MBR is valid and bootable.

### Structure of the MBR

The MBR has a fixed size of 512 bytes and is divided into several components:

- **Bootstrap Code**: The first 446 bytes of the MBR are reserved for the bootstrap code, which is responsible for locating and loading the boot loader.

- **Partition Table**: The next 64 bytes are used for the partition table, with each partition entry occupying 16 bytes. Each entry contains information such as the partition type, starting sector, and size.

- **Boot Signature**: The final two bytes (511 and 512) of the MBR contain the boot signature (0x55AA), which marks the end of the MBR and indicates that it is bootable.

### Limitations of the MBR

While the MBR partitioning scheme served well for many years, it has several limitations:

- **Limited Partition Support**: The MBR partition table supports up to four primary partitions or three primary partitions and one extended partition. This limitation can be overcome by using extended partitions and logical partitions, but it adds complexity.

- **Disk Size Limit**: MBR disks have a maximum disk size of 2 terabytes (TB) due to the use of 32-bit sector addresses. Larger disks may require the use of GPT (GUID Partition Table) instead.

- **Security Concerns**: The MBR is vulnerable to malware attacks, such as boot sector viruses, which can overwrite or corrupt the MBR code.

### Conclusion

The Master Boot Record (MBR) is a critical component of legacy BIOS-based systems, responsible for booting the operating system and managing disk partitions. While the MBR served well for many years, its limitations led to the development of the GUID Partition Table (GPT) for modern UEFI-based systems. Understanding the MBR's structure and function is essential for system administrators and users involved in disk management and system booting. 
