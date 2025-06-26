# kfc - Kai's First Container ![Go](https://img.shields.io/badge/Go-1.24-blue) ![License: MIT](https://img.shields.io/badge/license-MIT-green)

📘 [English](./README.md)

本项目包含了我在学习[mydocker](https://github.com/xianlubird/mydocker)时的代码实现和学习笔记。

---

## 📑 目录

- [kfc - Kai's First Container  ](#kfc---kais-first-container--)
  - [📑 目录](#-目录)
  - [📚 背景](#-背景)
  - [🛠️ 改进](#️-改进)
  - [🧱 功能](#-功能)
  - [💻 环境](#-环境)
  - [🚀 使用](#-使用)
  - [🙏 鸣谢](#-鸣谢)
  - [💬 反馈和交流](#-反馈和交流)

---

## 📚 背景

跟着书本《自己动手写docker》和对应项目[mydocker](https://github.com/xianlubird/mydocker)，进行动手实践。

截至 **2025/6/25**，我学到了 **第三章第三节**。

---

## 🛠️ 改进

在 `mydocker` 的基础上，我进行了一些修改，主要是因为 `mydocker` 中的一些操作在新的开发环境中有其他的实现。

1. ✅ 将 **cgroup-v1** 换为 **cgroup-v2**。
2. 📒 除了 **AUFS** 之外，了解 **OverlayFS** 等 **UnionFS**。
3. 📝 添加了学习笔记，包含我对技术的理解。

---

## 🧱 功能

- [X] Namespace isolation (`UTS`, `PID`, `NS`, `IPC`, `NET`)
- [X] Cgroup v2-based resource limitation
- [ ] Filesystem mount and `chroot` isolation
- [ ] Container image management
- [ ] Volume mounting
- [ ] Network bridge support
- [ ] Container lifecycle management (start, stop, remove)

---

## 💻 环境

我的本地开发环境：

- OS: WSL2 with Linux 5.15.133
- Architecture: x86_64
- Go: 1.24.4
- Cgroup: v2 enabled

你可以这样检查：

```bash
uname -a
go version
mount | grep cgroup
```

我的输出：

```
xzk@yosemite:~$ uname -a
Linux yosemite 5.15.133.1-microsoft-standard-WSL2 #2 SMP Tue Oct 10 00:21:20 CST 2023 x86_64 x86_64 x86_64 GNU/Linux
xzk@yosemite:~$ go version
go version go1.24.4 linux/amd64
xzk@yosemite:~$ mount | grep cgroup
cgroup2 on /sys/fs/cgroup type cgroup2 (rw,nosuid,nodev,noexec,relatime,nsdelegate)
```

---

## 🚀 使用

拉取项目并使用：

```
git clone https://github.com/ZhongkaiXu/kfc.git
cd kfc
go build .
sudo ./kfc run -tty -m 256m -cpuset 0 -cpushare 1000 /bin/bash
```

上面这个命令代表了：

- 256MB 内存
- CPU核 0
- CPU权重 1000
- 一个内置的TTY

> ⚠️ 因为需要创建namespace核cgroup，所以需要root权限。

## 🙏 鸣谢

- [mydocker](https://github.com/xianlubird/mydocker) - 我学习的项目。
- _"自己动手写Docker"_ - 我学习的书籍。

## 💬 反馈和交流

由于本人目前还是在读学生，只能在课外时间学习，所以学习进度可能会比较慢，希望我能坚持下来。
对于项目中的错误，欢迎交流和讨论！

📧 我的邮箱: kaikaixu2003@163.com。

😁 爱来自北京，Peace！