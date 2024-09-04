#!/bin/bash
bash_path=$(cd "$(dirname "$0")";pwd)
source $bash_path/base.config

# 获取当前机器ip
get_localip() {
    ipaddr=$(ip addr | awk '/^[0-9]+: / {}; /inet.*global/ {print gensub(/(.*)\/(.*)/, "\\1", "g", $2)}' | grep $ip_segment)
    echo "$ipaddr"
}

# 关闭防火墙
set_ufw_config() {
    echo -e "---------------------\033[31m close ufw \033[0m---------------------"
    ufw disable
    swapoff -a
    echo -e "---------------------\033[31m close ufw success \033[0m---------------------"
}

# 网络设置
set_net_config() {
    echo -e "---------------------\033[31m set net config \033[0m---------------------"
    grep "net" /etc/sysctl.d/k8s.conf
    if [[ $? -eq 0 ]];then
        echo -e "---------------------\033[31m net config seted \033[0m---------------------"
    else
        cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
            net.bridge.bridge-nf-call-ip6tables = 1 
            net.bridge.bridge-nf-call-iptables = 1 
            net.ipv4.ip_forward = 1 
            vm.swappiness = 0
EOF
        modprobe br_netfilter
        sysctl -p /etc/sysctl.d/k8s.conf
        echo -e "---------------------\033[31m net config set success \033[0m---------------------"
fi

    echo -e "---------------------\033[31m open ipvs \033[0m---------------------"
    apt-get -y install ipset ipvsadm
    grep "modprobe" /etc/sysconfig/modules/ipvs.modules
    if [[ $? -eq 0 ]];then
        echo -e "---------------------\033[31m ipvs opened \033[0m---------------------"
    else
        mkdir -p /etc/sysconfig/modules/
        cat > /etc/sysconfig/modules/ipvs.modules <<EOF 
        #!/bin/bash 
        modprobe -- ip_vs 
        modprobe -- ip_vs_rr 
        modprobe -- ip_vs_wrr 
        modprobe -- ip_vs_sh 
        modprobe -- nf_conntrack_ipv4 
EOF
        chmod 755 /etc/sysconfig/modules/ipvs.modules && sh /etc/sysconfig/modules/ipvs.modules && lsmod | grep -e ip_vs -e nf_conntrack_ipv4
        echo -e "---------------------\033[31m ipvs open success \033[0m---------------------"
    fi
}

# 关闭swap分区
close_swap() {
    echo -e "---------------------\033[31m close swap \033[0m---------------------"
    grep 'vm.swappiness=0' /etc/sysctl.conf
    if [[ $? -eq 0 ]];then
        echo -e "---------------------\033[31m swappiness closed \033[0m---------------------"
    else
        /sbin/swapoff -a
        sed -i '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab
        echo "vm.swappiness=0" >> /etc/sysctl.conf
        /sbin/sysctl -p
        echo "---------------------\033[31m swappiness close success \033[0m---------------------"
    fi
}

# 安装docker
install_docker() {
    echo -e "---------------------\033[31m install docker \033[0m---------------------"
    test -d /etc/docker
    if [[ $? -eq 0 ]];then
        echo -e "---------------------\033[31m docker installed \033[0m---------------------"
    else
        set -e
        sudo apt-get install -y \
            apt-transport-https \
            ca-certificates \
            curl \
            git \
            gnupg-agent \
            software-properties-common

        curl -fsSL http://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | sudo apt-key add -
        arch=`dpkg --print-architecture`
        sudo add-apt-repository \
            "deb [arch=$arch] http://mirrors.aliyun.com/docker-ce/linux/ubuntu \
            $(lsb_release -cs) \
            stable"

        sudo apt-get update
        case "$(lsb_release -r --short)" in
            *16.04* ) sudo apt-get install -y docker-ce=5:19.03.13~3-0~ubuntu-xenial docker-ce-cli=5:19.03.13~3-0~ubuntu-xenial containerd.io ;;
            *18.04* ) sudo apt-get install -y docker-ce=5:19.03.13~3-0~ubuntu-bionic docker-ce-cli=5:19.03.13~3-0~ubuntu-bionic containerd.io ;;
            *20.04* ) sudo apt-get install -y docker-ce=5:19.03.13~3-0~ubuntu-focal docker-ce-cli=5:19.03.13~3-0~ubuntu-focal containerd.io ;;
            * ) exit ;;
            esac

        tee /etc/docker/daemon.json <<EOF
{
    "registry-mirrors": [ "https://docker.m.daocloud.io","https://docker.udayun.com", "https://noohub.ru", "https://huecker.io","https://dockerhub.timeweb.cloud"],
    "exec-opts": ["native.cgroupdriver=systemd"]
}
EOF
        sudo systemctl daemon-reload
        sudo systemctl restart docker
        echo -e "---------------------\033[31m docker install success \033[0m---------------------"
    fi
}

# 安装nvidia docker
install_nvidia_docker() {
    set -e
    echo -e "---------------------\033[31m install nvidia docker \033[0m---------------------"
    distribution=$(. /etc/os-release;echo $ID$VERSION_ID)
    curl -s -L https://nvidia.github.io/nvidia-docker/gpgkey | sudo apt-key add -
    curl -s -L https://nvidia.github.io/nvidia-docker/$distribution/nvidia-docker.list | sudo tee /etc/apt/sources.list.d/nvidia-docker.list
    sudo apt-get update && sudo apt-get install -y nvidia-docker2
    tee /etc/docker/daemon.json <<EOF
{
    "registry-mirrors": ["https://6kx4zyno.mirror.aliyuncs.com"],
    "exec-opts": ["native.cgroupdriver=systemd"],
    "default-runtime": "nvidia",
    "runtimes": {
        "nvidia": {
            "path": "/usr/bin/nvidia-container-runtime",
            "runtimeArgs": []
        }
    }
}
EOF
    sudo systemctl daemon-reload
    sudo systemctl restart docker
    echo -e "---------------------\033[31m nvidia docker install success \033[0m---------------------"
}

# 安装enflame docker
install_enflame_docker() {
    set -e
    echo -e "---------------------\033[31m install enflame docker \033[0m---------------------"
    OS_TYPE=$(grep -e '^ID=' /etc/os-release | tr -d 'ID="')

    if [ "${OS_TYPE}" == "ubuntu" ]; then
        dpkg -i ${enflame_docker_device_plugin_path}docker-device-plugin_*.deb
    elif [ "${OS_TYPE}" == "tlinux" ] || [ "${OS_TYPE}" == "centos" ] || [ "${OS_TYPE}" == "rhel" ] ||  [ "${OS_TYPE}" == "tencentos" ] ; then
        rpm -ivh ${enflame_docker_device_plugin_path}docker-device-plugin_*.rpm
    else
        echo "${OS_TYPE} not supported" ; exit 0
    fi
    tee /etc/docker/daemon.json <<EOF
{
    "default-runtime": "enflame",
    "runtimes": {
        "enflame": {
            "path": "/usr/bin/enflame-container-runtime",
            "runtimeArgs": []
        }
    },
    "registry-mirrors": ["https://mirror.ccs.tencentyun.com", "https://docker.mirrors.ustc.edu.cn", "https://6kx4zyno.mirror.aliyuncs.com"],
    "insecure-registries": ["127.0.0.1/8", "artifact.enflame.cn", "192.168.202.110:5000", "192.168.202.74:5000"],
    "max-concurrent-downloads": 10,
    "log-driver": "json-file",
    "log-level": "warn",
    "log-opts": {
        "max-size": "30m",
        "max-file": "3"
    },
    "default-shm-size": "1G",
    "default-ulimits": {
         "memlock": { "name":"memlock", "soft":  -1, "hard": -1 },
         "stack"  : { "name":"stack", "soft": 67108864, "hard": 67108864 },
         "nofile": {"name": "nofile","soft": 65536, "hard": 65536}
    },
    "data-root": "/var/lib/docker",
    "exec-opts": ["native.cgroupdriver=systemd"]
}
EOF
    sudo systemctl daemon-reload
    sudo systemctl restart docker
    echo -e "---------------------\033[31m enflame docker install success \033[0m---------------------"
}

# 安装k8s工具
set_repo() {
    set -e
    tee -a /etc/apt/sources.list <<EOF
# kubeadm及kubernetes组件安装源
deb https://mirrors.aliyun.com/kubernetes/apt kubernetes-xenial main
EOF
    cd $bash_path
    cat apt-key.gpg | sudo apt-key add -
    echo -e "---------------------\033[31m install kubelet、kubeadm、kubectl \033[0m---------------------"
    sudo apt-get update -y && apt-get install -y --allow-unauthenticated kubelet=${kube_version}-00 kubeadm=${kube_version}-00 kubectl=${kube_version}-00
    cat <<EOF | sudo tee /etc/sysconfig/kubelet
        KUBELET_EXTRA_ARGS="--cgroup-driver=systemd"
EOF
    systemctl enable kubelet
    echo -e "---------------------\033[31m install kubelet、kubeadm、kubectl success \033[0m---------------------"
}

# pull镜像
install_k8s() {
    set -e
    echo -e "---------------------\033[31m pull images \033[0m---------------------"
	images=(kube-scheduler:${k8s_version}
			kube-proxy:${k8s_version}
			kube-controller-manager:${k8s_version}
			kube-apiserver:${k8s_version}
			pause:3.1
			etcd:3.3.15-0)
	for imagename in ${images[@]}; do
        docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/$imagename
        docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/$imagename k8s.gcr.io/$imagename
        docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/$imagename
	done
	docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:1.6.2
	docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:1.6.2 k8s.gcr.io/coredns:1.6.2
	docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:1.6.2
    echo -e "---------------------\033[31m pull images success \033[0m---------------------"
}

# nvidia gpu节点打标签
nvidia_gpu_label() {
    kubectl label nodes `hostname` hardware-type=NVIDIAGPU
}

# enflame gcu节点打标签
enflame_gcu_label() {
    set -e
    echo -e "---------------------\033[31m enflame gcu label \033[0m---------------------"
    kubectl label nodes `hostname` hardware-type=ENFLAMEGCU
    echo -e "---------------------\033[31m enflame gcu label success \033[0m---------------------"
}

# cambricon mlu节点打标签
cambricon_mlu_label() {
    set -e
    echo -e "---------------------\033[31m cambricon mlu label \033[0m---------------------"
    kubectl label nodes `hostname` hardware-type=CAMBRICONMLU
    echo -e "---------------------\033[31m cambricon mlu label success \033[0m---------------------"
}

# xilinx fpga节点打标签
xilinx_fpga_label() {
    set -e
    echo -e "---------------------\033[31m xilinx fpga label \033[0m---------------------"
    kubectl label nodes `hostname` hardware-type=XILINXFPGA
    echo -e "---------------------\033[31m xilinx fpga label success \033[0m---------------------"
}

# huawei a910节点打标签
huawei_a910_label() {
    kubectl label nodes `hostname` hardware-type=ASCENDNPU
}
