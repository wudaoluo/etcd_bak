# etcd_bak
etcd 数据备份恢复

### 设计原则
    1.先同步所有数据到本地临时目录,同步成功,在remame成配置目录
    2.为每个 key 创建一个指纹
    3.watch监控对 key操作
    4.定时和 etcd 核对指纹(没有完成)