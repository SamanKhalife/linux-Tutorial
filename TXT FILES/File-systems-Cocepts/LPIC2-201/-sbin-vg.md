# LVM Volume Group (vg) 

commands related to managing Volume Groups (VGs) in LVM (Logical Volume Manager). Here's an overview of commands related to managing Volume Groups (`vg`) in LVM:

Volume Groups (VGs) in LVM manage one or more physical volumes (PVs) and provide logical storage units for Logical Volumes (LVs).

### Basic Commands

1. **Creating a Volume Group**:
   ```bash
   vgcreate [options] vg_name /dev/sdX /dev/sdY ...
   ```
   Creates a new volume group named `vg_name` using physical volumes `/dev/sdX`, `/dev/sdY`, etc.

2. **Displaying Volume Group Information**:
   ```bash
   vgdisplay [vg_name]
   ```
   Displays information about all volume groups or a specific `vg_name`.

3. **Adding Physical Volumes to a Volume Group**:
   ```bash
   vgextend vg_name /dev/sdZ
   ```
   Adds `/dev/sdZ` to the existing volume group `vg_name`.

4. **Removing Physical Volumes from a Volume Group**:
   ```bash
   vgreduce vg_name /dev/sdZ
   ```
   Removes `/dev/sdZ` from the volume group `vg_name`.

5. **Renaming a Volume Group**:
   ```bash
   vgrename old_vg_name new_vg_name
   ```
   Renames `old_vg_name` to `new_vg_name`.

6. **Removing a Volume Group**:
   ```bash
   vgremove vg_name
   ```
   Removes the volume group `vg_name`.

7. **Resizing a Volume Group**:
   ```bash
   vgextend vg_name /dev/sdX
   ```
   Increases the size of the volume group `vg_name` by adding `/dev/sdX`.

### Advanced Commands

1. **Moving Physical Extents (PEs)**:
   ```bash
   pvmove /dev/sdX
   ```
   Moves allocated physical extents from `/dev/sdX` to other physical volumes in the volume group.

2. **Changing Volume Group Attributes**:
   ```bash
   vgchange -a y vg_name
   ```
   Activates (mounts) a volume group (`-a y`) or deactivates (`-a n`) it.

3. **Listing Physical Volumes in a Volume Group**:
   ```bash
   pvdisplay -m
   ```
   Lists detailed information about physical volumes and their mappings to logical volumes.

### Conclusion

The `vg` commands in LVM (`vgcreate`, `vgdisplay`, `vgextend`, etc.) are essential for managing volume groups and the logical storage they provide. They allow administrators to efficiently allocate and manage disk space, resize volumes, and ensure data integrity through advanced features like mirroring and striping. For more detailed information and options, refer to the respective manual pages (`man vgcreate`, `man vgdisplay`, etc.) or the LVM documentation.
