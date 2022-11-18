#!/bin/bash
bash_path=$(cd "$(dirname "$0")";pwd)
source $bash_path/comm.sh

if [[ "$(whoami)" != "root" ]]; then
	echo "please run this script as root ." >&2
	exit 1
fi

echo -e "---------------------\033[31m k8s node deploying ,do not close it，\
        The deployment process requires downloading a large number of components, \
        and the deployment time depends on your network environment \033[0m---------------------"

node_type=$1

# node互信
rootssh_trust() {
    set -e
    echo -e "---------------------\033[31m k8s node mutual trust \033[0m---------------------"
    cd $bash_path
    for host in ${hostip[@]}; do
        if [[ `get_localip` != $host ]];then
            if [[ ! -f /root/.ssh/id_rsa.pub ]];then
                echo -e 'init'
                apt -y install expect
                expect ssh_trust_init.exp $root_passwd $host
            else
                echo -e 'add'
                apt -y install expect
                expect ssh_trust_add.exp $root_passwd $host
            fi
            echo -e "---------------------\033[31m "$host" node trust success \033[0m---------------------"
        fi
    done
}

# node加入集群
join_cluster() {
    set -e
    echo -e "---------------------\033[31m node join cluster \033[0m---------------------"
    kubeadm join $masterip:6443 --token $tocken --discovery-token-ca-cert-hash sha256:$sha_value --ignore-preflight-errors=…

    mkdir -p $HOME/.kube
    if [ ! -f "$bash_path/config" ]; then
        echo "$bash_path/config no such file"
    else 
        cp -i $bash_path/config $HOME/.kube/config
        chown $(id -u):$(id -g) $HOME/.kube/config
    fi
    export KUBECONFIG=$HOME/.kube/config
    echo -e "---------------------\033[31m "`hostname`"just cluster success \033[0m---------------------"
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
    # 节点互信
    rootssh_trust
    # 加入集群
    join_cluster

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
}

main
echo -e "---------------------\033[31m node deploy success，you can test 'kubectl get node' on master to check the cluster \033[0m---------------------"
