

# demo gif





![ex](README.assets/thoreau.gif)





----











![image-20220512214959982](README.assets/image-20220512214959982.png)









# hello thoreau









# sshserver


从私有仓库https://github.com/yimtun/sshServer 整理来  核心功能还未完成迁移 不是sshServer的全部功能


# buid env

```
go  version
go version go1.17.10 linux/amd64
```

# repo path

```
/root/go-code/
```


# build

```
cd /root/go-code/
git clone https://github.com/yimtun/thoreau
cd thoreau
mkdir ./pluginBin
go build -o pluginBin/    thoreau.go
go build -o pluginBin/    controller/controller.go 
```


```
ls /root/go-code/thoreau/pluginBin/
controller	thoreau
```


# install plugin


```
cd /root/go-code/thoreau
git clone https://github.com/yimtun/cf
cd /root/go-code/thoreau/cf
go build cf.go
cp cf ../pluginBin
```

```
cd /root/go-code/thoreau
rm -rf cf
```


```
touch  /root/go-code/thoreau/pluginBin/a.json
```


```
cat /root/go-code/thoreau/pluginBin/a.json
```

```
{
  "resources": [
    {
      "user_name": "tom",
      "pass_word":"tom@321"
    },
    {
      "user_name": "cat",
      "pass_word":"cat@123"
    }
  ]
}
```



```
cd /root/go-code/thoreau/pluginBin
./thoreau
```




# test on linux  server 


```
ssh  tom@127.0.0.1 -p 2022 
```

## press ctrl + c   to quit 



