/*
Works with LVM!

gh: http://libguestfs.org/
https://serverascode.com/2018/06/26/using-cloud-images.html
*/

path=/dev/stor/fs

Ansible=~/.uni/code/os.Hyperion/Hyper/alter/Ansible

func install(){
  ansible = '''
  sudo apt install libguestfs-tools -y
  sudo usermod -aG kvm $USER
  '''
}

func initImg(){
  sudo chmod 755 $path
}

func info(){
  qemu-img info $path
}

func passUser(){
virt-customize -a $path --password me:password:$pass
}

func passRoot(){
virt-customize -a $path --root-password password:$pass
}

func createUser(){
  func simple_works(){
  virt-customize -a $path --run-command 'useradd me'
  }

  func noninteractive(){
  virt-customize -a $path --run-command 'adduser --disabled-password --gecos "" me'
  }
}

func fstransform(){
/*
* https://fedoramagazine.org/transform-file-systems-in-linux/, http://www.linux-magazine.com/Online/Features/Converting-Filesystems-with-Fstransform
*/
  init(){
  virt-customize -a $path --run-command 'add-apt-repository -y universe'
  virt-customize -a $path --run-command 'apt update'
  virt-customize -a $path --network --install fstransform
  }

  fsInfo(){
  virt-customize -a $path --run-command 'blkid /dev/sda1'
  }

  toBtrfs(){
  # Pre
  virt-customize LTS-a $path --network --install btrfs-progs

  virt-customize -a $path --run-command 'btrfs-convert /dev/sda1'
  }

  toXfs(){
  virt-customize -a $path --network --install xfsprogs
  fstransform /dev/sda1 xfs
  }
}


func info(){
qemu-img info $path
}

func sshSetup(){
virt-customize -a $path --copy-in $Ansible/ssh/sshd_config:/etc/ssh/

  func injectKey(){
  // https://www.cyberciti.biz/faq/how-to-add-ssh-public-key-to-qcow2-linux-cloud-images-using-virt-sysprep/
    inUser(){
    virt-customize -a $path --mkdir $HOME/.ssh/
    virt-customize -a $path --ssh-inject $USER:file:$HOME/.ssh/id_rsa.pub
    }

    inRoot(){
    virt-customize -a $path --mkdir /root/.ssh/
    //virt-customize -a $path --ssh-inject root:file:$HOME/.ssh/id_rsa.pub
    virt-customize -a $path --ssh-inject root:file:$me/.ssh/id_rsa.pub
    }
  }
  regenHostKey(){
  /* https://www.cloudvps.com/helpcenter/knowledgebase/linux-vps-configuration/regenerating-ssh-host-keys
  */
  virt-customize -a $path --run-command "ssh-keygen -f /etc/ssh/ssh_host_rsa_key -N '' -t rsa"
  }
}

func Ansible(){
  func network(){
  // https://netplan.io/examples
  net=single.yaml
  net=single-ubu1804.yaml

  virt-customize -a $path --copy-in $Ansible/netplan/$net:/etc/netplan/
  }

  func prepAnsible(){
    virt-customize -a $path --install git
    virt-customize -a $path --firstboot-install git #python
  }
}
