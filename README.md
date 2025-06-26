# kfc - Kai's First Container ![Go](https://img.shields.io/badge/Go-1.24-blue) ![License: MIT](https://img.shields.io/badge/license-MIT-green)

📘 [中文](./README-zh-CN.md)

This project is my personal implementation and learning record of building a container runtime, inspired by the excellent [mydocker project](https://github.com/xianlubird/mydocker). It serves both as a coding exercise and documentation of my understanding of container fundamentals.

---

## 📑 Table of Contents

- [kfc - Kai's First Container  ](#kfc---kais-first-container--)
  - [📑 Table of Contents](#-table-of-contents)
  - [📚 Background](#-background)
  - [🛠️ Improvements over the original mydocker](#️-improvements-over-the-original-mydocker)
  - [🧱 Features (in progress)](#-features-in-progress)
  - [💻 Environment](#-environment)
  - [🚀 Usage](#-usage)
  - [🙏 Acknowledgements](#-acknowledgements)
  - [💬 Feedback](#-feedback)

---

## 📚 Background

This repository is part of my journey in learning container technologies by following the book _"自己动手写Docker"_ , which is based on the `mydocker` project.

As of **2025/6/26**, I have studied up to **Chapter 3, Section 3**.

---

## 🛠️ Improvements over the original mydocker

In addition to the original features of `mydocker`, I have made the following changes:

1. ✅ Switched from **cgroup v1** to **cgroup v2** for modern resource control compatibility.
2. 📝 I am gradually adding my **learning notes** to this repository, documenting my personal understanding of the book and the underlying technologies.

---

## 🧱 Features (in progress)

- [x] Namespace isolation (`UTS`, `PID`, `NS`, `IPC`, `NET`)
- [x] Cgroup v2-based resource limitation
- [ ] Filesystem mount and `chroot` isolation
- [ ] Container image management
- [ ] Volume mounting
- [ ] Network bridge support
- [ ] Container lifecycle management (start, stop, remove)

---

## 💻 Environment

My local setup:

- OS: WSL2 with Linux 5.15.133
- Architecture: x86_64
- Go: 1.24.4
- Cgroup: v2 enabled

To verify your environment:

```bash
uname -a
go version
mount | grep cgroup
```

Example output:

```
xzk@yosemite:~$ uname -a
Linux yosemite 5.15.133.1-microsoft-standard-WSL2 #2 SMP Tue Oct 10 00:21:20 CST 2023 x86_64 x86_64 x86_64 GNU/Linux
xzk@yosemite:~$ go version
go version go1.24.4 linux/amd64
xzk@yosemite:~$ mount | grep cgroup
cgroup2 on /sys/fs/cgroup type cgroup2 (rw,nosuid,nodev,noexec,relatime,nsdelegate)
```

---

## 🚀 Usage

Clone and run:
```
git clone https://github.com/ZhongkaiXu/kfc.git
cd kfc
go build .
sudo ./kfc run -tty -m 256m -cpuset 0 -cpushare 1000 /bin/bash
```

This runs a container with:
- 256MB memory
- CPU core 0
- CPU share with 1000
- An interactive bash shell inside the container

> ⚠️ Root Permission is required due to usage of namespace and cgroup.

## 🙏 Acknowledgements

- [mydocker project](https://github.com/xianlubird/mydocker) - the origin and inspiration.
- _"自己动手写Docker"_ - an excellent book for understanding container internals.

## 💬 Feedback

Due to my limited experience, this project may contain mistakes or suboptimal code.
If you notice any issues or have suggestions, I sincerely welcome corrections or advice.

📧 Feel free to reach out via email: kaikaixu2003@163.com, Peace!