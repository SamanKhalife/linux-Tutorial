# WWID, WWN, LUN numbers

Understanding WWID, WWN, and LUN numbers is crucial for managing storage devices and their connections in a computing environment. Here’s a detailed explanation of these terms and their relevance:

### WWID (World Wide Identifier)

**WWID** stands for **World Wide Identifier**, which is a unique identifier assigned to a storage device. It helps in uniquely identifying a storage device across different systems and environments.

#### Characteristics
- **Unique**: Each WWID is unique to a specific storage device.
- **Persistence**: It remains consistent even if the device is moved to a different port or system.
- **Format**: Typically a hexadecimal string.

#### Usage
- **Device Identification**: Used to identify and manage storage devices in SAN (Storage Area Network) environments.
- **Configuration**: Helps in configuring and managing multipath setups.

#### Example
A typical WWID might look like this: `3600508b1001c21a6`.

### WWN (World Wide Name)

**WWN** stands for **World Wide Name**, which is a unique identifier assigned to devices in a storage area network (SAN) or Fibre Channel network. It helps in identifying storage devices, hosts, and switches.

#### Characteristics
- **Unique**: Each WWN is unique to the device or node.
- **Format**: Often presented in hexadecimal format, and may vary in length (e.g., 64-bit or 128-bit).
- **Types**:
  - **N_Port WWN**: Assigned to nodes in a Fibre Channel network.
  - **P_Port WWN**: Assigned to ports on Fibre Channel switches.

#### Usage
- **Device Identification**: Used for identifying and managing devices within a SAN or Fibre Channel network.
- **Configuration**: Helps in zoning and mapping devices in SAN environments.

#### Example
A typical WWN might look like this: `20:00:00:25:b5:fc:00:01`.

### LUN (Logical Unit Number)

**LUN** stands for **Logical Unit Number**. It is an identifier used in storage networks to differentiate between various logical units (i.e., virtual storage devices or partitions) on a physical storage device.

#### Characteristics
- **Identification**: Identifies specific logical units or partitions on a storage device.
- **Mapping**: Used in conjunction with WWNs to map logical units to physical devices.
- **Format**: Typically a numeric value.

#### Usage
- **Device Access**: Used by operating systems and applications to access specific logical units within a storage device.
- **Configuration**: Helps in configuring and managing storage resources and access paths.

#### Example
A LUN might be represented as `0`, `1`, `2`, etc., indicating different logical units or partitions.

### Summary

- **WWID**: Unique identifier for a storage device, ensuring consistent identification across different systems.
- **WWN**: Unique identifier for devices in SAN or Fibre Channel networks, used for managing and identifying devices.
- **LUN**: Identifies specific logical units or partitions within a storage device, used for accessing and managing storage resources.

### Example Use Cases

1. **Multipath Configuration**: When configuring multipath storage, WWID and WWN are used to ensure that the correct paths to storage devices are configured and managed.
2. **SAN Zoning**: In a Fibre Channel SAN, WWNs are used to zone devices, ensuring proper communication between devices and hosts.
3. **Logical Unit Management**: LUN numbers are used to map and access specific logical units on a storage device, allowing for efficient storage management.

Understanding these identifiers and their roles in storage management is essential for system administrators and storage engineers to effectively configure, manage, and troubleshoot storage environments.
