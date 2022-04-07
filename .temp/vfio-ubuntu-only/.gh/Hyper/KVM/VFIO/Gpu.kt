/*
https://wiki.archlinux.org/index.php/PCI_passthrough_via_OVMF
*/


// Isolating_the_GPU
echo "options vfio-pci ids=10de:13c2,10de:0fbb" \
| sudo tee /etc/modprobe.d/vfio.conf
