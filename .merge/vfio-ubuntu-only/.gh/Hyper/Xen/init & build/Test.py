def Iommu():
    dmesg | grep Xen\ version
    dmesg | grep -e DMAR -e IOMMU
    find /sys/kernel/iommu_groups/ -type l  #Only works on 16.04
    
def Grub():
    cat /proc/cmdline
    sudo xl dmesg |grep virtualisation

def VFIO():
    xl pci-assignable-add 00:14.0
    xl pci-assignable-list
    sudo lspci -s 03: -v


def Version():
    # https://wiki.xenproject.org/wiki/Xen_Common_Problems#How_can_I_check_the_version_of_Xen_Project_hypervisor.3F
