## virtualization 

there are currently two approaches to virtualization:

1. Hypervisors - to one degree or another - control host system hardware to provide
each guest operating system the resources it needs. Guest machines are run as
system processes, but with virtualized access to hardware resources. AWS servers,
for instance, are built on the open source Xen hypervisor technology. Other
important hypervisor platforms include VMware ESXi, KVM, and Microsoft’s
Hyper-V.

3. Containers are extremely lightweight virtual servers that, rather than running as
full operating systems, share the underlying kernel of their host OS. Containers can be built from plain-text scripts, created and launched in
seconds, and easily and reliably shared across networks. The best known container
technology right now is probably Docker. The Linux Container project (LXC) that
we’ll be working with in this chapter was Docker’s original inspiration.
