# Microsoft Regedit

**Microsoft Regedit (Registry Editor)** is a graphical tool provided by Windows that allows users to view, search, and modify the Windows Registry. The Windows Registry is a hierarchical database where Windows and many applications store settings and configuration data.

### **Understanding the Windows Registry**
The Windows Registry contains configuration settings for:
- Operating System components (like services, hardware drivers, and kernel)
- Installed applications
- User profiles
- System resources (like memory and device configurations)

The Registry is organized in a tree structure, with five primary keys (also known as "hives") at the top level.

### **Top-Level Hives in the Registry**

1. **HKEY_LOCAL_MACHINE (HKLM)**:
   - Contains system-wide settings (hardware, software, security).
   - Applies to all users on the system.
   - Common subkeys include:
     - **SYSTEM**: Settings for system configuration.
     - **SOFTWARE**: Installed programs and drivers.

2. **HKEY_CURRENT_USER (HKCU)**:
   - Contains settings for the currently logged-in user.
   - Personal preferences, such as desktop settings and application preferences.
   
3. **HKEY_CLASSES_ROOT (HKCR)**:
   - Contains information about registered applications and file associations.
   - Defines what happens when certain file types are opened (which programs are used).
   - Includes file extension mappings (e.g., `.txt`, `.jpg`).

4. **HKEY_USERS (HKU)**:
   - Contains settings for all user profiles on the system.
   - Each user's settings are stored under a unique identifier.

5. **HKEY_CURRENT_CONFIG (HKCC)**:
   - Contains hardware settings that are dynamically generated at boot time.
   - Reflects the current hardware configuration.

### **Accessing the Registry Editor**
To access **Regedit**:
1. Press `Windows + R` to open the **Run** dialog.
2. Type `regedit` and press **Enter**.
3. Confirm the User Account Control (UAC) prompt (if necessary).

### **Navigating in Regedit**
The Registry Editor consists of two main panes:
- **Left Pane (Tree View)**: Displays the hierarchical structure of the registry.
- **Right Pane**: Shows the contents (values) of the selected key in the left pane.

Each key in the registry can contain:
- **Subkeys** (similar to folders).
- **Values** (similar to files).

There are three main types of registry values:
1. **String Value (REG_SZ)**: Stores text.
2. **Binary Value (REG_BINARY)**: Stores binary data.
3. **DWORD (32-bit) Value (REG_DWORD)**: Stores 32-bit integers.

You can also encounter other value types, such as `QWORD (64-bit)`, `Multi-String`, and `Expandable String`.

### **Basic Operations in Regedit**

1. **Create a Key or Value**:
   - Right-click in the left pane to create a new **key**.
   - In the right pane, right-click to create a new **value**.

2. **Modify a Value**:
   - Double-click a value in the right pane to edit it.

3. **Delete a Key or Value**:
   - Right-click the key or value you want to delete, then click **Delete**.

4. **Export/Backup Registry**:
   - You can backup the registry (or specific keys) by selecting **File > Export**. This allows you to restore the registry later if needed.
   - To restore, use **File > Import** to apply previously saved `.reg` files.

### **Examples of Common Registry Paths**

1. **Change Desktop Wallpaper**:
   - Key: `HKEY_CURRENT_USER\Control Panel\Desktop`
   - Value: **Wallpaper**
   - This value stores the path of the desktop wallpaper image.

2. **Disable Windows Defender**:
   - Key: `HKEY_LOCAL_MACHINE\SOFTWARE\Policies\Microsoft\Windows Defender`
   - Value: **DisableAntiSpyware** (set to `1` to disable).

3. **Run Programs at Startup**:
   - Key: `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run`
   - Value: You can add a new value with the path to the program to be run at startup.

### **Precautions when using Regedit**
- **Always backup the registry** before making any changes. Incorrect edits can cause system instability or crashes.
- Be careful when following online guides, especially those that require registry modifications. Make sure you're making changes to the correct keys and values.
- Only modify registry settings if you fully understand their purpose.

