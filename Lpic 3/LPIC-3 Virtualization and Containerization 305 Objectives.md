# LPIC-3 Virtualization and Containerization 305 Objectives
## Topic 351: Full Virtualization
### 351.1 Virtualization Concepts and Theory (weight: 6)

#### Description:	Candidates should know and understand the general concepts, theory and terminology of virtualization. This includes Xen, QEMU and libvirt terminology.
Key Knowledge Areas:

Understand virtualization terminology
Understand the pros and cons of virtualization
Understand the various variations of Hypervisors and Virtual Machine Monitors
Understand the major aspects of migrating physical to virtual machines
Understand the major aspects of migrating virtual machines between host systems
Understand the features and implications of virtualization for a virtual machine, such as snapshotting, pausing, cloning and resource limits
Awareness of oVirt, Proxmox, systemd-machined and VirtualBox
Awareness of Open vSwitch
The following is a partial list of the used files, terms and utilities:

[Hypervisor]()

[Hardware Virtual Machine (HVM)]()

[Paravirtualization (PV)]()

[Emulation and Simulation]()

[CPU flags]()

[/proc/cpuinfo]()

[Migration (P2V, V2V)]()


### 351.2 Xen (weight: 3)

#### Description:	Candidates should be able to install, configure, maintain, migrate and troubleshoot Xen installations. The focus is on Xen version 4.x.
Key Knowledge Areas:

Understand architecture of Xen, including networking and storage
Basic configuration of Xen nodes and domains
Basic management of Xen nodes and domains
Basic troubleshooting of Xen installations
Awareness of XAPI
Awareness of XenStore
Awareness of Xen Boot Parameters
Awareness of the xm utility
The following is a partial list of the used files, terms and utilities:

[Domain0 (Dom0), DomainU (DomU)]()

[PV-DomU, HVM-DomU]()

[/etc/xen/]()

[xl]()

[xl.cfg]()

[xl.conf]()

[xentop]()

### 351.3 QEMU (weight: 4)

#### Description:	Candidates should be able to install, configure, maintain, migrate and troubleshoot QEMU installations.
Key Knowledge Areas:

Understand the architecture of QEMU, including KVM, networking and storage
Start QEMU instances from the command line
Manage snapshots using the QEMU monitor
Install the QEMU Guest Agent and VirtIO device drivers
Troubleshoot QEMU installations, including networking and storage
Awareness of important QEMU configuration parameters
The following is a partial list of the used files, terms and utilities:

[Kernel modules: kvm, kvm-intel and kvm-amd]()

[/dev/kvm]()

[QEMU monitor]()

[qemu]()

[qemu-system-x86_64]()

[ip]()

[brctl]()

[tunctl]()

### 351.4 Libvirt Virtual Machine Management (weight: 9)

#### Description:	Candidates should be able to manage virtualization hosts and virtual machines (‘libvirt domains’) using libvirt and related tools.
Key Knowledge Areas:

Understand the architecture of libvirt
Manage libvirt connections and nodes
Create and manage QEMU and Xen domains, including snapshots
Manage and analyze resource consumption of domains
Create and manage storage pools and volumes
Create and manage virtual networks
Migrate domains between nodes
Understand how libvirt interacts with Xen and QEMU
Understand how libvirt interacts with network services such as dnsmasq and radvd
Understand libvirt XML configuration files
Awareness of virtlogd and virtlockd
The following is a partial list of the used files, terms and utilities:

[libvirtd]()

[/etc/libvirt/]()

[virsh (including relevant subcommands)]()


### 351.5 Virtual Machine Disk Image Management (weight: 3)

#### Description:	Candidates should be able to manage virtual machines disk images. This includes converting disk images between various formats and hypervisors and accessing data stored within an image.
Key Knowledge Areas:

Understand features of various virtual disk image formats, such as raw images, qcow2 and VMDK
Manage virtual machine disk images using qemu-img
Mount partitions and access files contained in virtual machine disk images using libguestfish
Copy physical disk content to a virtual machine disk image
Migrate disk content between various virtual machine disk image formats
Awareness of Open Virtualization Format (OVF)
The following is a partial list of the used files, terms and utilities:

[qemu-img]()

[guestfish (including relevant subcommands)]()

[guestmount]()

[guestumount]()

[virt-cat]()

[virt-copy-in]()

[virt-copy-out]()

[virt-diff]()

[virt-inspector]()

[virt-filesystems]()

[virt-rescue]()

[virt-df]()

[virt-resize]()

[virt-sparsify]()

[virt-p2v]()

[virt-p2v-make-disk]()

[virt-v2v]()

[virt-sysprep]()


## Topic 352: Container Virtualization
### 352.1 Container Virtualization Concepts (weight: 7)

#### Description:	Candidates should understand the concept of container virtualization. This includes understanding the Linux components used to implement container virtualization as well as using standard Linux tools to troubleshoot these components.
Key Knowledge Areas:

Understand the concepts of system and application container
Understand and analyze kernel namespaces
Understand and analyze control groups
Understand and analyze capabilities
Understand the role of seccomp, SELinux and AppArmor for container virtualization
Understand how LXC and Docker leverage namespaces, cgroups, capabilities, seccomp and MAC
Understand the principle of runc
Understand the principle of CRI-O and containerd
Awareness of the OCI runtime and image specifications
Awareness of the Kubernetes Container Runtime Interface (CRI)
Awareness of podman, buildah and skopeo
Awareness of other container virtualization approaches in Linux and other free operating systems, such as rkt, OpenVZ, systemd-nspawn or BSD Jails
The following is a partial list of the used files, terms and utilities:

[nsenter]()

[unshare]()

[ip (including relevant subcommands)]()

[capsh]()

[/sys/fs/cgroups]()

[/proc/[0-9]+/ns]()

[/proc/[0-9]+/status]()


### 352.2 LXC (weight: 6)
#### Description:	Candidates should be able to use system containers using LXC and LXD. The version of LXC covered is 3.0 or higher.
Key Knowledge Areas:

Understand the architecture of LXC and LXD
Manage LXC containers based on existing images using LXD, including networking and storage
Configure LXC container properties
Limit LXC container resource usage
Use LXD profiles
Understand LXC images
Awareness of traditional LXC tools
Partial list of the used files, terms and utilities:

[lxd]()

[lxc (including relevant subcommands)]()


### 352.3 Docker (weight: 9)
#### Description:	Candidate should be able to manage Docker nodes and Docker containers. This include understand the architecture of Docker as well as understanding how Docker interacts with the node’s Linux system.
Key Knowledge Areas:

Understand the architecture and components of Docker
Manage Docker containers by using images from a Docker registry
Understand and manage images and volumes for Docker containers
Understand and manage logging for Docker containers
Understand and manage networking for Docker
Use Dockerfiles to create container images
Run a Docker registry using the registry Docker image
Partial list of the used files, terms and utilities:

[dockerd]()

[/etc/docker/daemon.json]()

[/var/lib/docker/]()

[docker]()

[Dockerfile]()


### 352.4 Container Orchestration Platforms (weight: 3)

#### Description:	Candidates should understand the importance of container orchestration and the key concepts Docker Swarm and Kubernetes provide to implement container orchestration.
Key Knowledge Areas:

Understand the relevance of container orchestration
Understand the key concepts of Docker Compose and Docker Swarm
Understand the key concepts of Kubernetes and Helm
Awareness of OpenShift, Rancher and Mesosphere DC/OS



## Topic 353: VM Deployment and Provisioning
### 353.1 Cloud Management Tools (weight: 2)

#### Description:	Candidates should understand common offerings in public clouds and have basic feature knowledge of commonly available cloud management tools.
Key Knowledge Areas:

Understand common offerings in public clouds
Basic feature knowledge of OpenStack
Basic feature knowledge of Terraform
Awareness of CloudStack, Eucalyptus and OpenNebula
Partial list of the used files, terms and utilities:

[IaaS, PaaS, SaaS]()

[OpenStack]()

[Terraform]()


### 353.2 Packer (weight: 2)

#### Description:	Candidates should be able to use Packer to create system images. This includes running Packer in various public and private cloud environments as well as building container images for LXC/LXD.
Key Knowledge Areas:

Understand the functionality and features of Packer
Create and maintain template files
Build images from template files using different builders
Partial list of the used files, terms and utilities:

[packer]()


### 353.3 cloud-init (weight: 3)

#### Description:	Candidates should able to use cloud-init to configure virtual machines created from standardized images. This includes adjusting virtual machines to match their available hardware resources, specifically, disk space and volumes. Additionally, candidates should be able to configure instances to allow secure SSH logins and install a specific set of software packages. Furthermore, candidates should be able to create new system images with cloud-init support.
Key Knowledge Areas:

Understanding the features and concepts of cloud-init, including user-data, initializing and configuring cloud-init
Use cloud-init to create, resize and mount file systems, configure user accounts, including login credentials such as SSH keys and install software packages from the distribution’s repository
Integrate cloud-init into system images
Use config drive datasource for testing
Partial list of the used files, terms and utilities:

[cloud-init]()

[user-data]()

[/var/lib/cloud/]()

### 353.4 Vagrant (weight: 3)

#### Description:	Candidate should be able to use Vagrant to manage virtual machines, including provisioning of the virtual machine.
Key Knowledge Areas:

Understand Vagrant architecture and concepts, including storage and networking
Retrieve and use boxes from Atlas
Create and run Vagrantfiles
Access Vagrant virtual machines
Share and synchronize folder between a Vagrant virtual machine and the host system
Understand Vagrant provisioning, i.e. File and Shell provisioners
Understand multi-machine setup
Partial list of the used files, terms and utilities:

[vagrant]()

[Vagrantfile]()
