# Octopus Platform

<img src="./logo.png" width="100">

---

[简体中文](./readme.md)

**Octopus** is a one-stop computing fusion platform for multiple computing scenarios. the platform is mainly designed for the needs of computing and resource management in AI, HPC and other scenarios. It provides users with computing power management and use functions for data, algorithms, mirroring, models, and computing power, which is convenient for users to build a one-stop shop Computing environment, realizing calculation.
At the same time, cluster management personnel are provided with functions such as cluster resource management and monitoring, computing task management and monitoring, etc., to facilitate cluster management personnel to operate and analyze the overall system.

**Octopus** is based on the container orchestration platform [Kubernetes](https://kubernetes.io/zh/docs/concepts/overview/what-is-kubernetes) , octopus makes full use of the agility, light weight, and isolation of containers to meet the needs of diverse computing scenarios.

## Features and Scenarios

Octopus has the following characteristics:

- **One-stop Development**, provide users with one-stop AI and HPC computing scenarios development functions, through data management, model development and model training, open up the entire computing link;
- **Easy to manage**, provide a one-stop resource management platform for platform managers, and greatly reduce the management cost of platform managers through visual tools such as resource configuration, monitoring, and authority management and control;
- **Easy to deploy**, octopus supports rapid deployment in [Helm](https://helm.sh), simplifying the complex deployment process;
- **Superior performance**, provide a high-performance distributed computing experience, and ensure the smooth operation of each environment through multiple optimizations. At the same time, through resource scheduling optimization and distributed computing optimization, the efficiency of model training is further improved;
- **Good compatibility**, the platform supports heterogeneous hardware, such as GPU, NPU, FPGA, etc., to meet various hardware cluster deployment needs. It supports multiple deep learning frameworks, such as TensorFlow, Pytorch, PaddlePaddle, etc., and can support new additions through custom mirroring frame.

Octopus is suitable for use in the following scenarios:

- Build a large-scale AI computing platform;
- Hope to share computing resources;
- Hope to complete model training in a unified environment;
- Hope to use the integrated plug-in to assist model training and improve efficiency.

## Get Started

**Octopus**　manages computing resources and optimizes computing tasks for scenarios such as AI and HPC. Decoupling computing hardware and software through mirroring and container technology ([Docker](https://docs.docker.com)) enables easy switching between different computing environments．

Octopus users usually have two different roles:

- **Cluster administrators** are the owners and maintainers of computing resources. The administrator is responsible for the deployment and availability of the cluster.
- **Cluster users** are consumers of cluster computing resources. According to the deployment scenario, cluster users can be machine learning and deep learning researchers, data scientists, laboratory teachers, students, etc.

Octopus provides end-to-end manuals for cluster users and administrators.

### For cluster administrators

Documents related to cluster administrators include the following:

- ***Cluster Deployment Guide***: the main contents provided in this part include: preparation and installation of cluster dependent environment and components, Octopus system deployment guide and follow-up system upgrade instructions to facilitate installation and maintenance. For details, please refer to [here](https://octopus.pcl.ac.cn/docs/deployment/environment) 。

- ***Cluster Management Manual***: This part mainly introduces the operations that the cluster administrator can perform after entering the Octopus management system through the management system page entrance. The main function descriptions include: platform monitoring, resource management, user management, machine time management, data management, algorithm management, development and training management And other functions. For details, please refer to [here](https://octopus.pcl.ac.cn/docs/management/intro) 。

### For cluster users

The main documents related to cluster users are as follows:

- ***User Manual***： this part mainly introduces the operations that cluster users can perform after entering the Octopus system through the Octopus system page entrance. The main function descriptions include: data management, algorithm management, mirroring management, development and training management and other functions. For details, please refer to [here](https://octopus.pcl.ac.cn/docs/manual/intro) 。

## Documentations

For detailed documentation, please refer to [here](https:///octopus.pcl.ac.cn/docs/introduction/intro).

## How to Contribute

For detailed contribution guidelines, please refer to [here](https://octopus.pcl.ac.cn/docs/community/contribution).

## License

[Apache License](https://octopus.pcl.ac.cn/docs/community/LICENSE)