

# demo gif





![ex](README.assets/thoreau.gif)





----











![image-20220512214959982](README.assets/image-20220512214959982.png)









# hello thoreau









# sshserver


从私有仓库https://github.com/yimtun/sshServer 整理来  核心功能还未完成迁移 不是sshServer的全部功能



# build

```
cd sshserver
mkdir ./pluginBin
go build -o pluginBin/    thoreau.go
go build -o pluginBin/    controller/controller.go 
```


```
ls pluginBin/
controller	thoreau
```


install plugin


```
cd sshserver
git clone https://github.com/yimtun/cf.git
cd cf
go build cf.go
cp cf ../pluginBin
```

