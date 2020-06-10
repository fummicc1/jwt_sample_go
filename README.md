### prepare Public/PrivateKey
```sh
$ cd sec
# generate demo.rsa and demo.rsa.pub
$ ssh-keygen -t rsa
# convert pub to pkcs8 (or pkcs1)
$ ssh-keygen -f demo.rsa.pub -e -m pkcs8 > demo.rsa.pub.pkcs8
```