# 升级说明

如果当前已部署章鱼且需要升级到新版本，参考以下升级说明，注意跨多个版本升级时需要参考这之间多个版本的升级说明



## v4.2.0

升级前：

1.sftpgo的存储路径下需要创建目录权限为1000:1000的data、home目录

升级后：

2.执行sql修复数据

```sql
update user set resource_pools='["common-pool"]' where resource_pools is null
update notebook set resource_pool = 'common-pool' where workspace_id = 'default-workspace';
update notebook as n set n.resource_pool = (select w.r_pool_id from workspace as w where n.workspace_id = w.id) where n.workspace_id != 'default-workspace';
update train_job set resource_pool = 'common-pool' where workspace_id = 'default-workspace';
update train_job as n set n.resource_pool = (select w.r_pool_id from workspace as w where n.workspace_id = w.id) where n.workspace_id != 'default-workspace';
update model_deploy set resource_pool = 'common-pool' where workspace_id = 'default-workspace';
update model_deploy as n set n.resource_pool = (select w.r_pool_id from workspace as w where n.workspace_id = w.id) where n.workspace_id != 'default-workspace';
update train_job_template set resource_pool = 'common-pool' where workspace_id = 'default-workspace';
update train_job_template as n set n.resource_pool = (select w.r_pool_id from workspace as w where n.workspace_id = w.id) where n.workspace_id != 'default-workspace';
```



## v4.2.4

升级后:

1.执行sql

```sql
UPDATE  octopus.train_job AS a
JOIN core.jobs AS b 
ON a.id = b.id
SET a.detail = b.detail;

UPDATE  octopus.notebook_job AS a
JOIN core.jobs AS b 
ON a.id = b.id
SET a.detail = b.detail;
```



## v4.2.5

升级后:

1.执行sql

```sql
alter table resourcespec modify price decimal(10,2) null;
alter table notebook_job modify resource_spec_price decimal(10,2) default '0' not null comment '资源规格价格';
alter table model_deploy modify res_spec_price decimal(10,2) null comment '资源规格单价';
```


## v4.2.6

升级前：
1. v4.2.6已经将nvidia设备插件集成到章鱼安装包，需要先将原先的nvidia设备插件卸载
```
wget https://raw.githubusercontent.com/NVIDIA/k8s-device-plugin/v0.9.0/nvidia-device-plugin.yml
kubectl delete -f nvidia-device-plugin.yml
```