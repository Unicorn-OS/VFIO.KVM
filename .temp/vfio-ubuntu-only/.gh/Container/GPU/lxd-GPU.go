/*
Graph: LXD.gpu
. current: https://blog.simos.info/how-to-easily-run-graphics-accelerated-gui-apps-in-lxd-containers-on-your-ubuntu-desktop/
*/

func init(){

  # if deb package:
    echo "root:$UID:1" | sudo tee -a /etc/subuid /etc/subgid

  # all
  lxc profile create gui
  cat lxdguiprofile.txt | lxc profile edit gui

  # list
  lxc profile list
}

func launch(){
  lxc launch --profile default --profile gui ubuntu:18.10 guitry
}
