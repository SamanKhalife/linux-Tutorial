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

* [Hypervisor](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/hypervisor.md)

* [Hardware Virtual Machine (HVM)](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/Hardware-Virtual-Machine(HVM).md)

* [Paravirtualization (PV)](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/Paravirtualization(PV).md)

* [Emulation and Simulation](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/Emulation-and-Simulation.md)

* [CPU flags](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/CPU-flags.md)

* [/proc/cpuinfo](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/proccpuinfo.md)

* [Migration (P2V, V2V)](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/Migration(P2V%2C%20V2V).md)

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

* [Domain0 (Dom0), DomainU (DomU)](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/Domain0(Dom0)%2CDomainU%20(DomU).md)

* [PV-DomU, HVM-DomU](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/PV-DomU%2CHVM-DomU.md)

* [/etc/xen/](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/etc-xen-.md)

* [xl](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/xl.md)

* [xl.cfg](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/xl.cfg.md)

* [xl.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/xl.conf.md)

* [xentop](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/xentop.md)

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

* [Kernel modules: kvm, kvm-intel and kvm-amd](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/Kernel-modules-kvm%2Ckvm-intel.and.kvm-amd.md)

* [/dev/kvm](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/-dev-kvm-.md)

* [QEMU monitor](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/QEMU-monitor.md)

* [qemu](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/qemu.md)

* [qemu-system-x86_64]()

* [ip](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/ip.md)

* [brctl](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/brctl.md)

* [tunctl](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/tunctl.md)

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

* [libvirtd](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/libvirtd.md)

* [/etc/libvirt/](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/-etc-libvirt.md)

* [virsh (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/virsh(including-relevant-subcommands).md)


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

* [qemu-img](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/qemu-img.md)

* [guestfish (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/guestfish(including-relevant-subcommands).md)

* [guestmount](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/guestmount.md)

* [guestumount](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/guestumount.md)

* [virt-cat]()

* [virt-copy-in](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/virt-copy-in.md)

* [virt-copy-out](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/virt-copy-out.md)

* [virt-diff](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/virt-diff.md)

* [virt-inspector](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/virt-inspector.md)

* [virt-filesystems](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/virt-filesystems.md)

* [virt-rescue](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/virt-rescue.md)

* [virt-df](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/virt-df.md)

* [virt-resize](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/virt-resize.md)

* [virt-sparsify](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/virt-sparsify.md)

* [virt-p2v](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/virt-p2v.md)

* [virt-p2v-make-disk](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/virt-p2v-make-disk.md)

* [virt-v2v](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/virt-v2v.md)

* [virt-sysprep](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/virt-sysprep.md)


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

* [nsenter](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/nsenter.md)

* [unshare](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/unshare.md)

* [ip (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/ip(including-relevant-subcommands).md)

* [capsh](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/capsh.md)

* [/sys/fs/cgroups](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/-sys-fs-cgroups.md)

* [/proc/*[0-9]+/ns](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/-proc-%5B0-9%5D%2B-ns.md)

* [/proc/*[0-9]+/status](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/-proc%20-%5B0-9%5D%2B-status.md)


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

* [lxd](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/lxd.md)

* [lxc (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/lxc(including%20relevant%20subcommands).md)


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

* [dockerd](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/dockerd.md)

* [/etc/docker/daemon.json](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/-etc-docker-daemon.json.md)

* [/var/lib/docker/]()

* [docker](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/-var-lib-docker-.md)

* [Dockerfile](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/Dockerfile.md)


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

* [IaaS, PaaS, SaaS](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/IaaS%2CPaaS%2CSaaS.md)

* [OpenStack](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/OpenStack.md)

* [Terraform](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/Terraform.md)


### 353.2 Packer (weight: 2)
#### Description:	Candidates should be able to use Packer to create system images. This includes running Packer in various public and private cloud environments as well as building container images for LXC/LXD.
Key Knowledge Areas:

Understand the functionality and features of Packer
Create and maintain template files
Build images from template files using different builders
Partial list of the used files, terms and utilities:

* [packer](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/packer.md)


### 353.3 cloud-init (weight: 3)
#### Description:	Candidates should able to use cloud-init to configure virtual machines created from standardized images. This includes adjusting virtual machines to match their available hardware resources, specifically, disk space and volumes. Additionally, candidates should be able to configure instances to allow secure SSH logins and install a specific set of software packages. Furthermore, candidates should be able to create new system images with cloud-init support.
Key Knowledge Areas:

Understanding the features and concepts of cloud-init, including user-data, initializing and configuring cloud-init
Use cloud-init to create, resize and mount file systems, configure user accounts, including login credentials such as SSH keys and install software packages from the distribution’s repository
Integrate cloud-init into system images
Use config drive datasource for testing
Partial list of the used files, terms and utilities:

* [cloud-init](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/cloud-init.md)

* [user-data](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/user-data.md)

* [/var/lib/cloud/](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/-var-lib-cloud.md)

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

* [vagrant](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/vagrant.md)

* [Vagrantfile](../TXT%20FILES/File-systems-Cocepts/LPIC3-305/Vagrantfile.md)
