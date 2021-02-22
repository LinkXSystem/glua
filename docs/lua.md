# LUA 命令

## 基本命令

- 编译

```shell
luac -o core.luac hello-world.lua
```

- 查看二进制文件

```shell
luac -l core.luac
```

- 基于 XXD 查看

```shell
xxd -u -g 1 core.luac
```
