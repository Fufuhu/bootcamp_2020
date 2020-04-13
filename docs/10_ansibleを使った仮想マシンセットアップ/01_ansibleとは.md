# Ansibleとは

AnsibleはRedHatの管理する自動化ツールです。
本ハンズオンではAnsibleを使った仮想マシンへのnginxの導入をやってみます。

Ansibleの特徴としてはおもに以下の2つが挙げられます。

1. 宣言型の定義
2. エージェントレス

宣言型の定義とは、**最終的にこうなっていて欲しい構成**をコード定義することで、
自動的にツール側で意図した構成へと設定してくれます。

対義語は手続き型であり、手続き型は**セットアップして欲しい手順**をコード定義します。


参考) [ansible.com](https://www.ansible.com/)

## Ansibleの導入方法

Ansibleはデプロイするためのマシン(今回の場合はmac)にインストールしておけば問題ありません。AnsibleはPythonが必要なのでPythonの動作環境を準備しましょう。

Pythonの実行環境構築については[90_(補足)Ansibleの実行環境の構築](90_(補足)Ansibleの実行環境の構築.md)で解説しています。

基本的にはpyenvとvenvモジュールをうまく組み合わせることで、
単一のマシンの中で複数バージョンのPythonとAnsibleを同居させることが可能です。
プロジェクトによって利用しているAnsibleのバージョンが異なるため、容易に切り替えられるようにしておきましょう。

Pythonの動作環境(とpipコマンド)が導入できていれば、
Ansible自体の導入は難しくはありません。

```console
$ pip install ansible
```

以上でエラーが発生しなければansibleの導入は完了です。

## 本資料で用いているAnsibleなどのバージョン

本資料では、以下のようなパッケージ構成でAnsibleのハンズオンを記載しています。
出力内容からわかる通り、Ansibleのバージョンは`2.9.6`を利用しています。

```console
$ pip freeze
ansible==2.9.6
cffi==1.14.0
cryptography==2.9
Jinja2==2.11.1
MarkupSafe==1.1.1
pycparser==2.20
PyYAML==5.3.1
six==1.14.0
```
