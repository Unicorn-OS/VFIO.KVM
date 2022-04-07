/* Nested
Hyper-V Works!
https://ladipro.wordpress.com/2017/02/24/running-hyperv-in-kvm-guest/

Example
http://alexander.holbreich.org/qemu-kvm-introduction/
https://wiki.gentoo.org/wiki/QEMU/Options
*/

func install(){
sudo su
qemu-system-x86_64 -hda /media/me/uni/disk/hyper-kvm.qcow2 -cdrom /home/me/stor/Artifact/uOS/RS5X64.OFF19.ENU.MAR2019.iso -boot d \
-cpu host,hv_relaxed,hv_spinlocks=0x1fff,hv_vapic,hv_time,+vmx \
-enable-kvm \
-smp 4 -m 6G
}

func store(){
  -hda /media/me/uni/disk/hyper-kvm.qcow2 \
}

func run(){
  // Works!
sudo su
qemu-system-x86_64 -hda /media/me/uni/disk/hyper-kvm.qcow2 \
-cpu host,hv_relaxed,hv_spinlocks=0x1fff,hv_vapic,hv_time,+vmx \
-enable-kvm \
-smp 4 -m 6G
}

func alt(){
  /* -machine */
  -enable-kvm
  else
  -machine type=q35,accel=kvm \\

  /* CPU *?
  -cpu host \

  /* GPU */
  -display gtk,gl=on -vga qxl \
  else
  -display gtk,gl=on -vga std \
}

func feature(){
  /* https://wiki.archlinux.org/index.php/QEMU

  Virtual Machine CPU Passthrough, Topology, Pinning, etc Qemu KVM Libvirt Linux Bitcoin Mining Crypto
  https://youtu.be/DONBXAqGQZI
  */

  func gl(){
  // https://wiki.archlinux.org/index.php/QEMU#Graphic_card
  -vga virtio \
  -display gtk,gl=on \
  }
}
