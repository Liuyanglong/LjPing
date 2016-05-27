#LjPing

作用：部署于服务器，从cmdb中获取当前管辖的所有server，并对其一一发送icmp包，从而检验当前网络环境的大致健康状态，当交换机/路由器调整或更新时，可快速判断网络情况。

使用说明：
```
# LjPing -h
Usage of ./LjPing:
  -con=100: 并发数
  -timeout=30: ping 超时时间
```
