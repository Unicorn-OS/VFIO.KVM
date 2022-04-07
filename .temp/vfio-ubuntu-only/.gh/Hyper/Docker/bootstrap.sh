Debian(){
  # from Official Docs
  if which docker > /dev/null ;then
    version
  else
    echo "installing docker"
    sudo apt-get update
    sudo apt-get install \
        apt-transport-https \
        ca-certificates \
        curl \
        gnupg2 \
        software-properties-common -y
    curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -
    sudo apt-key fingerprint 0EBFCD88
    sudo add-apt-repository \
       "deb [arch=amd64] https://download.docker.com/linux/debian \
       $(lsb_release -cs) \
       stable"
    sudo apt-get update
    sudo apt-get install docker-ce docker-ce-cli containerd.io -y
    postinstall
  fi
}

postinstall(){
  sudo groupadd docker
  sudo usermod -aG docker $USER
  echo 'Postinstall done'
  # docs.docker.com/install/linux/linux-postinstall/
}

try(){
  docker run --rm hello-world
  docker rmi hello-world
}


convert_MXLinux_to_Debian(){
  # Fixes MX
  if [ -f "/etc/lsb-release_original" ]; then
    sudo touch /etc/apt/sources.list

    sudo cp /etc/lsb-release /etc/lsb-release_original
    sudo sed -i 's/MX/Debian/' /etc/lsb-release
    sudo sed -i 's/Continuum/stretch/' /etc/lsb-release
  fi
}


distro_Install(){
  Distro=`lsb_release --id`
  case $Distro in
    'Distributor ID:	Debian')
      Debian
      ;;
    'Distributor ID:	MX')
      echo Converting
      convert_MXLinux_to_Debian
      Debian
      ;;
  esac
}

version(){
  docker --version
}

main(){

  case $1 in
    t | test )
      try
      ;;
    v )
      version
      ;;
    *)
      distro_Install
      ;;
  esac
}

main $@
