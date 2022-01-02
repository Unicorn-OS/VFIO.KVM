/* Debian
https://ubuntu.com/blog/lxd-on-debian-using-snapd
*/

func install(){
  sudo snap install lxd

  # https://help.ubuntu.com/lts/serverguide/lxd.html.en
  sudo adduser $USER lxd
  newgrp lxd

  # Init
  lxd init
}

func reactivate(){
  /*
  https://blog.simos.info/how-to-initialize-lxd-again/
  */
}
