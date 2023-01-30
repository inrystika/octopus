#!/bin/bash
bash_path=$(cd "$(dirname "$0")";pwd)
source $bash_path/comm.sh

if [[ "$(whoami)" != "root" ]]; then
	echo "please run this script as root ." >&2
	exit 1
fi

echo -e "---------------------\033[31m k8s master deploying，do not close it，\
        The deployment process requires downloading a large number of components, \
        and the deployment time depends on your network environment \033[0m---------------------"

node_type=$1

# k8s初始化
init_k8s() {
	set -e
	rm -rf /root/.kube
    rm -rf /var/lib/etcd/*
	kubeadm reset -f
	echo "1" >/proc/sys/net/bridge/bridge-nf-call-iptables
	kubeadm init --image-repository=registry.aliyuncs.com/google_containers --kubernetes-version=$k8s_version --pod-network-cidr=$pod_network_cidr --service-cidr=$service_cidr --apiserver-advertise-address=$masterip --ignore-preflight-errors=Swap
	mkdir -p $HOME/.kube
    cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
	chown $(id -u):$(id -g) $HOME/.kube/config
    export KUBECONFIG=$HOME/.kube/config
}

# calico网络配置
install_calico() {
    set -e
    echo -e "---------------------\033[31m master deploy calico \033[0m---------------------"
    cd $bash_path
    test -f calico.yaml
	kubectl apply -f calico.yaml
    echo -e "---------------------\033[31m master deploy calico success \033[0m---------------------"
}

# master节点token设置
token_shar_value() {
    set -e
    echo -e "---------------------\033[31m master set token \033[0m---------------------"
    cd $bash_path
    /usr/bin/kubeadm token list > $bash_path/token_shar_value.text
    sed -i "s/tocken=/tocken=$(sed -n "2, 1p" token_shar_value.text | awk '{print $1}')/g" $bash_path/base.config
    sed -i "s/sha_value=/sha_value=$(openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //')/g" $bash_path/base.config
        
    rm -rf $bash_path/token_shar_value.text
    echo -e "---------------------\033[31m master token set success \033[0m---------------------"
}

# node节点互信
rootssh_trust() {
    set -e
    echo -e "---------------------\033[31m k8s node mutual trust \033[0m---------------------"
    cd $bash_path
    num=0
    for host in ${hostip[@]}; do
        let num+=1
        if [[ `get_localip` != $host ]];then
            if [[ ! -f /root/.ssh/id_rsa.pub ]];then
                echo 'init'
                apt -y install expect
                expect ssh_trust_init.exp $root_passwd $host
            else
                echo 'add'
                apt -y install expect
                expect ssh_trust_add.exp $root_passwd $host
            fi
            scp $HOME/.kube/config apt-key.gpg base.config comm.sh node_install_k8s.sh del_k8s.sh ssh_trust_init.exp ssh_trust_add.exp root@$host:/root
        fi
        echo -e "---------------------\033[31m "$host" node trust success \033[0m---------------------"
    done
}

check_cluster() {
    kubectl get node 
    kubectl cluster-info
}

main() {
    # 关闭防火墙
    set_ufw_config
    # 网络设置
    set_net_config
    # 关闭swap分区
    close_swap
    # 安装docker
    install_docker
    # 安装nvidia docker
    if [[ $node_type == "nvidia_gpu" ]];then
        install_nvidia_docker
    fi
    # 安装enflame docker
    if [[ $node_type == "enflame_gcu" ]];then
        install_enflame_docker
    fi

    # 安装k8s工具
    set_repo
    # pull镜像
    install_k8s

    # 初始化k8s
    init_k8s
    # 安装calico
    install_calico
    # 设置token
    token_shar_value
    # 节点互信
    rootssh_trust

    # 节点打标签
    if [[ $node_type == "nvidia_gpu" ]];then
        nvidia_gpu_label
    elif [[ $node_type == "huawei_a910" ]];then
        huawei_a910_label
    elif [[ $node_type == "enflame_gcu" ]];then
        enflame_gcu_label
    elif [[ $node_type == "cambricon_mlu" ]];then
        cambricon_mlu_label
    fi
    
    # 验证
    check_cluster
}

main
echo -e "---------------------\033[31m master deploy success, you can test 'kubectl get pod -n kube-system' to check the cluster \033[0m---------------------"
