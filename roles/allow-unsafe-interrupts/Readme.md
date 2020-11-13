# Src
[prime.](access.redhat.com/documentation/en-us/red_hat_virtualization/4.1/html/installation_guide/appe-configuring_a_hypervisor_host_for_pci_passthrough)


# guide
[Proxmox guide](https://pve.proxmox.com/wiki/Pci_passthrough)


# Solution
on CentOS when you load from Dracut, it doesn't pick up options from /etc/modprobe.d/

To solve it:
Run 'Allow unsafe interrupts allow_unsafe_interrupts=1' Task before, and then run 'dracut regenerate' to build options in Initramfs!


# test
https://wiki.archlinux.org/index.php/Kernel_module#Obtaining_information
'''To list the options that are set for a loaded module:
`systool -v -m module_name`
'''

`systool -v -m vfio_iommu_type1 | grep allow_unsafe_interrupts`
