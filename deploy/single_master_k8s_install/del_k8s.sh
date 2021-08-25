#!/bin/bash
if [[ "$(whoami)" != "root" ]]; then
	echo "please run this script as root ." >&2
	exit 1
fi

echo -e "---------------------\033[31m del k8s components and docker \033[0m---------------------"

del_k8s() {
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
  apt-get remove -y kubelet=1.16.3-00 kubeadm=1.16.3-00 kubectl=1.16.3-00
  rm -rf /etc/yum.repos.d/kubernetes.repo
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
  echo -e "---------------------\033[31m docker del success \033[0m---------------------"
}

main() {
  del_kube
  del_docker
}

main
echo -e "---------------------\033[31m k8s components and docker del success \033[0m---------------------"