# Vagrantとは

[Vagrant](https://www.vagrantup.com/)とは仮想マシン環境の構築及び管理を行うためのツールです。
Vagrantfileを用いて仮想マシンの構築に必要な定義やプロビジョニングを行います。

## Vagrantの導入方法
Homebrewをインストールし、以下のコマンドを実行してください
```
brew cask install vagrant
```

## Vagrantの実行方法
以下のコマンドを実行し、ディレクトリの移動及び、コマンドの実行をしてください。
```
cd `git rev-parse --show-toplevel`/docs/05_Vagrantを使った環境構築/src/vagrant 
vagrant up
```

色々コンソールが流れてきますが、最後に以下のように表示されれば成功です
```
default: Successfully built ansible
    default: Installing collected packages: ansible
    default: Successfully installed ansible-2.9.7
```

## ログインの確認
上記メッセージが出てセットアップが完了したら以下のコマンドでログインを行います

```
vagrant ssh
```
これでインスタンスにログインできたらvagrantの環境構築は完了です
