// Xen stable release
// From https://github.com/intel/gvt-linux/wiki/GVTd_Setup_Guide

// Prereq
sudo apt-get install -y git libfdt-dev libpixman-1-dev libssl-dev vim socat libsdl1.2-dev libspice-server-dev autoconf libtool xtightvncviewer tightvncserver x11vnc libsdl1.2-dev uuid-runtime uuid uml-utilities bridge-utils python-dev liblzma-dev libc6-dev

// FIX pkgs
sudo apt-get install -y bcc iasl uuid-dev libyajl-dev libpixman-1.dev lib32ncurses5-dev

// Grub
// https://askubuntu.com/questions/148095/how-do-i-set-the-grub-timeout-and-the-grub-default-boot-entry

sudo sed -i 's/GRUB_DEFAULT=0/GRUB_DEFAULT=2/' /etc/default/grub
sudo sed -i 's/GRUB_HIDDEN/#GRUB_HIDDEN/' /etc/default/grub


configure(){
# OVMF
# FIX 1
./configure --libdir=/usr/lib --enable-ovmf
}

/* FIX  xen services not started
https://unix.stackexchange.com/questions/234873/xen-error-xl-list-complains-about-libxl
*/
sudo systemctl enable xencommons
sudo systemctl enable xendomains
sudo systemctl enable xen-watchdog
sudo systemctl enable xendriverdomain
