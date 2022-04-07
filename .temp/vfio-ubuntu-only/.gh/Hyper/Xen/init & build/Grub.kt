Grub(){
echo '''GRUB_CMDLINE_XEN="dom0_mem=1024M"''' | sudo tee -a /etc/default/grub

sudo sed -i 's/GRUB_DEFAULT=0/GRUB_DEFAULT=2/' /etc/default/grub

sudo sed -i 's/CMDLINE_LINUX_DEFAULT="/CMDLINE_LINUX_DEFAULT="intel_iommu=on /' /etc/default/grub

sudo update-grub
}
