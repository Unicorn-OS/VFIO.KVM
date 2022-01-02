init(){

  install_vagrant(){
    echo installing vagrant
    echo 'Debian: Requires user added to sudo first!'

    alt_official_works(){
      # Latest official release!
      ver=2.2.5
      dl=https://releases.hashicorp.com/vagrant/$ver/vagrant_${ver}_x86_64.deb

      cd /tmp
      wget -c $dl
      sudo dpkg -i vagrant_${ver}_x86_64.deb
      # github.com/dotless-de/vagrant-vbguest/issues/292
    }

    alt_repo(){
      # Vagrant from repo. But plugins are prone to Breaking!
      sudo apt install vagrant -y
    }

    kvm(){
      sudo apt install qemu-kvm libvirt-daemon-system -y
    }

    plugins(){
      deps(){
        #sudo apt install gcc -y
        # vagrant-libvirt deps
        sudo apt-get build-dep vagrant ruby-libvirt -y
        sudo apt-get install qemu libvirt-bin ebtables dnsmasq-base -y
        sudo apt-get install libxslt-dev libxml2-dev libvirt-dev zlib1g-dev ruby-dev -y
        }

      deps
      vagrant plugin install vagrant-libvirt
      vagrant plugin install vagrant-cachier
    }

    set_default_provider(){
      if [ -z "$VAGRANT_DEFAULT_PROVIDER" ]; then # if var unset
        echo "export VAGRANT_DEFAULT_PROVIDER=libvirt" >> ~/.bashrc
        echo "export VAGRANT_DEFAULT_PROVIDER=libvirt" >> ~/.profile
      fi
      # www.vagrantup.com/docs/providers/default.html
    }

    if_Debian(){
      sudo apt-get update
      sudo apt-get install -y vagrant-libvirt libvirt-daemon-system qemu-kvm
      sudo apt-get install -y nfs-common nfs-kernel-server ebtables dnsmasq
      sudo apt-get install -y ansible rsync
      sudo usermod -a -G libvirt $USER
      newgrp libvirt
      sudo systemctl restart libvirtd

      sudo usermod -a -G kvm $USER
      newgrp kvm
      sudo rmmod kvm_intel
      sudo rmmod kvm
      sudo modprobe kvm
      sudo modprobe kvm_intel
      # Works! blog.dachary.org/2017/08/15/howto-vagrant-libvirt-provider-on-debian-gnulinux-stretch9/
    }

    # start installing
      alt_official_works
      kvm
      plugins
      set_default_provider
      if_Debian
      echo done init
  }

  if ! which vagrant > /dev/null; then
    install_vagrant
  else
    version
  fi
}

sync_from_UniFS(){
  # Pull a vagrant box from |mirror to save Network!
  vagrant box add
}

try(){

  # up
  mkdir -p ~/.test ; cd ~/.test
  if [ ! -f ~/.test/Vagrantfile ]; then
    vagrant init generic/debian10
  fi
  vagrant up

  # ssh_test
  vagrant ssh -c whoami

  # destroy
  vagrant destroy -f
  rm -r ~/.test

  echo 'test complete'
  # app.vagrantup.com/generic/boxes/debian10
}

uninstall(){
  sudo apt purge vagrant -y
  sudo apt autoremove -y
  echo 'removed vagrant'
}

update(){
  echo todo
}

version(){
  vagrant version	# test
}

help(){
  echo '''
  you can say:
  init              : Installs vagrant and plugins.
  init t | test     : Test vagrant up
  init u | update   : Update
  init v            : Get Version
  init rm           : Uninstall and remove all trace!
  init h | help     : this
  '''
}

main(){
  # Args
  if [ -z "$1" ]; then
    init
  else
    case $1 in
      rm)
        uninstall
        ;;
      t | test)
        try
        ;;
      u)
        update
        ;;
      v)
        version
        ;;
      h | help)
        help
        ;;
      *)
        init
        ;;
    esac
  fi
}

main $@
