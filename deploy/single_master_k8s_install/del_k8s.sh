#!/bin/bash
bash_path=$(cd "$(dirname "$0")";pwd)
source $bash_path/base.config

if [[ "$(whoami)" != "root" ]]; then
	echo "please run this script as root ." >&2
	exit 1
fi

echo -e "---------------------\033[31m del k8s components and docker \033[0m---------------------"

del_kube() {
  echo -e "---------------------\033[31m del k8s components \033[0m---------------------"
  kubeadm reset -f
  modprobe -r ipip
  rm -rf ~/.kube
  rm -rf /etc/kubernetes
  rm -rf /etc/systemd/system/kubelet.service.d
  rm -rf /etc/systemd/system/kubelet.service
  rm -rf /usr/bin/kube*
  rm -rf /etc/cni
  rm -rf /opt/cni
  rm -rf /var/lib/etcd
  rm -rf /var/etcd
  apt-get remove -y kubelet=${kube_version}-00 kubeadm=${kube_version}-00 kubectl=${kube_version}-00
  rm -rf /etc/yum.repos.d/kubernetes.repo
  rm -rf /opt/kube/bin/kubectl
  rm -rf /opt/kube/bin/kubelet
  rm -rf /opt/kube/bin/kube-proxy
  echo -e "---------------------\033[31m k8s components del success \033[0m---------------------"
}

del_docker() {
  echo -e "---------------------\033[31m del docker \033[0m---------------------"
  apt-get -y  autoremove docker docker-ce docker-engine docker.io containerd runc
  dpkg -l |grep ^rc|awk '{print $2}' |sudo xargs dpkg -P
  apt-get autoremove docker-ce-*
  rm -rf /etc/systemd/system/docker.service.d
  rm -rf /var/lib/docker
  rm -rf /etc/docker
  rm -rf /usr/bin/docker
  rm -rf /opt/kube/bin/docker
  echo -e "---------------------\033[31m docker del success \033[0m---------------------"
}

main() {
  del_kube
  del_docker
}

main
echo -e "---------------------\033[31m k8s components and docker del success \033[0m---------------------"