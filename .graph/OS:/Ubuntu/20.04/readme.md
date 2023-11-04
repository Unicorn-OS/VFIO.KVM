# guide
https://mathiashueber.com/pci-passthrough-ubuntu-2004-virtual-machine/

# Breaking changes for older passthrough setups in Ubuntu 20.04
'''Starting with kernel version 5.4, the “vfio-pci” driver is no longer a kernel module, but build-in into the kernel. We have to find a different way instead of using initramfs-tools/modules config files, as I recommended in e.g. the 18.04 guide.'''
