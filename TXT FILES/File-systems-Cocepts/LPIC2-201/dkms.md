# dkms

**DKMS** (Dynamic Kernel Module Support) is a framework used to manage the compilation and installation of kernel modules in Linux. It allows kernel modules to be automatically rebuilt and installed when the kernel is updated, which is crucial for maintaining the functionality of modules across different kernel versions. DKMS is particularly useful for handling third-party or custom kernel modules that are not included in the standard kernel distribution.

### Overview

**`dkms`** provides tools and scripts for creating, managing, and removing kernel modules. It ensures that kernel modules remain compatible with new kernel versions without requiring manual intervention.

### Basic Commands and Usage

#### 1. **Adding a New Module**

To add a new module to DKMS, you need to specify the module name, version, and path to the module source code.

```sh
dkms add [module-source-path]
```

- **`[module-source-path]`**: Path to the directory containing the module's source code and a `dkms.conf` file.

**Example:**

```sh
dkms add /usr/src/mymodule-1.0
```

#### 2. **Building a Module**

After adding the module, you need to build it. DKMS will compile the module for the current kernel version.

```sh
dkms build [module-name]/[module-version]
```

- **`[module-name]`**: Name of the module.
- **`[module-version]`**: Version of the module.

**Example:**

```sh
dkms build mymodule/1.0
```

#### 3. **Installing a Module**

Once the module is built, install it for the current kernel.

```sh
dkms install [module-name]/[module-version]
```

**Example:**

```sh
dkms install mymodule/1.0
```

#### 4. **Removing a Module**

To remove a module, you first need to uninstall it and then remove its source code from DKMS.

```sh
dkms remove [module-name]/[module-version] --all
```

**Example:**

```sh
dkms remove mymodule/1.0 --all
```

#### 5. **Listing Installed Modules**

To list all modules managed by DKMS, use:

```sh
dkms status
```

This command displays the status of all DKMS-managed modules, including their version, build status, and the kernels they are installed for.

#### 6. **Uninstalling a Module**

If you need to remove a module, you first uninstall it:

```sh
dkms uninstall [module-name]/[module-version]
```

**Example:**

```sh
dkms uninstall mymodule/1.0
```

#### 7. **Rebuilding Modules**

If you need to rebuild modules (e.g., after a kernel upgrade), use:

```sh
dkms autoinstall
```

This command automatically builds and installs all modules for the new kernel versions.

### Configuration File: `dkms.conf`

The `dkms.conf` file is essential for DKMS to understand how to build and install the module. It should be placed in the module source directory and contains the following sections:

- **`PACKAGE_NAME`**: Name of the module.
- **`PACKAGE_VERSION`**: Version of the module.
- **`CLEAN`**: Commands to clean the build directory.
- **`MAKE`**: Commands to compile the module.
- **`BUILT_MODULE_NAME[0]`**: Name of the module file.
- **`DEST_MODULE_LOCATION[0]`**: Directory where the module should be installed.

**Example `dkms.conf` File:**

```sh
PACKAGE_NAME="mymodule"
PACKAGE_VERSION="1.0"
CLEAN="make clean"
MAKE="make"
BUILT_MODULE_NAME[0]="mymodule"
DEST_MODULE_LOCATION[0]="/updates/dkms"
```

### Common Use Cases

- **Third-Party Drivers**: Managing proprietary drivers like NVIDIA or AMD graphics drivers.
- **Custom Kernel Modules**: Handling custom or experimental kernel modules.
- **Kernel Updates**: Automatically recompiling and installing modules after a kernel update.

### Summary

- **`dkms`** provides a framework for managing kernel modules, making it easier to handle custom and third-party modules across kernel updates.
- **Key Commands**: `dkms add`, `dkms build`, `dkms install`, `dkms remove`, `dkms status`, `dkms uninstall`, and `dkms autoinstall`.
- **Configuration File**: `dkms.conf` defines the build and installation process for modules.

Using DKMS helps ensure that kernel modules remain compatible with your system’s kernel versions and provides a streamlined approach to managing kernel module updates.
