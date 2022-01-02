/*
https://www.google.com/search?q=virt-builder+lvm

https://lukas.zapletalovi.com/2018/02/hidden-gem-of-fedora-virt-builder.html
*/

vol=uni1
pool=hd
size=100G
pass=u

// vm path
path=/dev/stor/${vol}

thin_lv(){
  lvcreate -V $size -T stor/${pool}-pool -n $vol
}

create_vm(){
  // must run with sudo
  ubuntu(){
    virt-builder ubuntu-18.04 --output $path --hostname $vol --root-password password:$pass
  }

  centos(){
    virt-builder centos-7.4 --output /dev/vg_virt/centos \
      --root-password password:redhat --hostname centos.lan
  }
}

delete(){
  sudo lvremove $path -y
}
