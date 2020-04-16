# ブートキャンプ資料

本資料はブートキャンプの一環として、

1. Vagrantを使った仮想マシンセットアップ
2. Ansibleを使ったOS環境セットアップ
3. Docker(docker-compose)を使ったローカルテスト
4. Kubernetesを使った様々なデプロイ手法

について学ぶためのものです。

内容としては以下を含んでいます。

+ [対象とする環境](00_前提とする環境/01_対象とする環境.md)
  + [(補足)Virtualboxのネットワーク設定](docs/00_前提とする環境/10_(補足)Virtualboxのネットワーク設定.md)
  + [(補足)仮想マシンの設定](docs/00_前提とする環境/20_(補足)仮想マシンの設定.md)
  + なお、仮想マシンの設定についてはVagrantを別途提供しているのでそれを利用しても構いません
    + [Vagrantを使った仮想マシンセットアップ](docs/05_Vagrantを使った環境構築/01_Vagrant環境の構築.md)
+ Ansibleを使ったOS環境セットアップ
  + [Ansibleとは](docs/10_ansibleを使った仮想マシンセットアップ/01_ansibleとは.md)
  + [Ansibleを使ったNginxのインストール](docs/10_ansibleを使った仮想マシンセットアップ/10_ansibleを使ったnginxのインストール.md)
  + [(補足)Ansible実行環境の構築](docs/10_ansibleを使った仮想マシンセットアップ/90_(補足)Ansibleの実行環境の構築.md)
+ Docker(docker-compose)を使ったローカルテスト
  + [Docker環境の構築](docs/20_dockerを使ったローカルテスト/10_docker環境の構築.md)
  + [docker-composeを使ったテスト](docs/20_dockerを使ったローカルテスト/30_docker-composeを使ったテスト.md)
+ Kubernetesをつかった様々なデプロイ手法
  + [Kubernetesの解説(簡易版)](docs/30_Kuberenetesを使った様々なデプロイ手法体験/00_Kubernetesの解説.md)
  + [ローリングデプロイ](docs/30_Kuberenetesを使った様々なデプロイ手法体験/20_ローリングデプロイ.md)
  + [ブルーグリーンデプロイ](docs/30_Kuberenetesを使った様々なデプロイ手法体験/30_ブルーグリーンデプロイ.md)
  + [カナリアリリース](docs/30_Kuberenetesを使った様々なデプロイ手法体験/40_カナリアリリース.md)
