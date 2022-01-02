---
os: manjaro-xfce-19.0.2-200311-linux54.iso
pci-slot: 03:00.0, 03:00.1
card: gtx_710
debug: 'cat /var/log/libvirt/qemu/*.log'
solution:

'First PCI slot was not working for me! Because it is in 'iommu group 1' and the root-port is preventing passthrough!
But the last slot works, it's in 'group 12' by itself, and can be passed through!'

error:
'First slot throws this error when iommu groups 1 is incorrect:
  qemu-system-x86_64: vfio_region_write(0000:01:00.0:region1+0x17b35e, 0x0,1) failed: Device or resource busy
'
