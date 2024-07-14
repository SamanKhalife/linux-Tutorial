# /etc/usbguard/rules.conf

The `/etc/usbguard/rules.conf` file is used to define USB authorization policies in USBGuard, specifying which USB devices are allowed or blocked based on various criteria such as device identifiers, device classes, and more. This file plays a crucial role in controlling USB device access and enhancing system security by preventing unauthorized devices from connecting to the system.

### Overview of `/etc/usbguard/rules.conf`

#### Purpose

The `rules.conf` file allows administrators to:
- Define rules for permitting or blocking USB devices based on specific attributes.
- Specify exceptions or special handling for certain devices.
- Ensure that only authorized USB devices can interact with the system.

### Basic Format

The general format of the `/etc/usbguard/rules.conf` file is:

```plaintext
# Comment
<allow|block> <device_policy> [if <condition>]
```

- `<allow|block>`: Specifies whether to allow (`allow`) or block (`block`) the device.
- `<device_policy>`: Describes the device policy, which can include device identifiers, classes, or other attributes.
- `[if <condition>]`: Optional condition to apply the rule only if certain criteria are met.

### Example Entries and Their Descriptions

Here are some common examples of entries in the `rules.conf` file:

1. **Allowing a Specific Device**

   Allow a device based on its USB vendor and product IDs:
   ```plaintext
   allow id 1234:5678
   ```

2. **Blocking Devices by Class**

   Block all USB mass storage devices:
   ```plaintext
   block class 08:*:*  # USB mass storage class
   ```

3. **Conditional Rules**

   Allow a device only if it is plugged into a specific USB port:
   ```plaintext
   allow id 1234:5678 if port 1-2
   ```

4. **Device Properties**

   Allow a device based on its serial number:
   ```plaintext
   allow id 1234:5678 serial "ABC12345"
   ```

### Commonly Used Attributes

- **id vendor_id:product_id**: Allows or blocks devices based on their vendor and product IDs.
- **class class_code:subclass_code:protocol_code**: Allows or blocks devices based on their USB class, subclass, and protocol.
- **serial "serial_number"**: Allows or blocks devices based on their serial number.
- **port port_number[-port_number]**: Applies the rule only if the device is plugged into the specified USB port(s).

### Example Configuration

Here is an example of a complete `/etc/usbguard/rules.conf` file:

```plaintext
# USBGuard rules configuration file

# Allow a specific USB keyboard
allow id 046d:c31c

# Block USB mass storage devices
block class 08:*:*

# Allow devices with specific serial numbers
allow id 1234:5678 serial "ABC12345"
allow id 7890:ABCD serial "XYZ98765"
```

### Applying Changes

After modifying the `/etc/usbguard/rules.conf` file, you need to reload the USBGuard rules to apply the changes:

```bash
sudo systemctl restart usbguard
```

This command restarts the USBGuard service, reloading the rules defined in `/etc/usbguard/rules.conf`.

### Conclusion

The `/etc/usbguard/rules.conf` file is essential for configuring USBGuard rules to control which USB devices are allowed or blocked on your system. By defining specific policies in this file, administrators can enforce security measures to prevent unauthorized devices from accessing sensitive data or compromising system integrity.
