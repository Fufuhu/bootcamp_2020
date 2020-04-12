# Ansibleを使ったnginxのインストール

https://github.com/Fufuhu/ansible_nginx にアクセスして、
リポジトリをcloneしてください。(forkしてcloneでも可)

今回解説するコードはひと通り、リポジトリ内部に含んでいます。

## Ansibleの主要な概念

Ansibleを構成する主要な概念としては以下の3つです。

1. インベントリ
2. プレイブック
3. タスク

### インベントリ

インベントリは作業対象となるノードとの接続情報を記載したファイルです。ターゲットとなるノードのIPアドレスやFQDNを記述します。
`inventory/inventory.ini`ファイルを確認してみましょう。

```ini
[nginx]
192.168.56.101
```

今回の場合はセットアップ対象のノードは一台のみなので、ごくシンプルになっています。
`nginx`のグループ内に、`192.168.56.101`のノードが定義されているだけです。

例えば、これとは別に[Apache Tomcat](https://tomcat.apache.org/)のノードを準備したい場合は以下のようになるでしょう。

```ini
[nginx]
192.168.56.101

[tomcat]
192.168.56.108
```

さらに、nginxを複数のサーバにセットアップしたい場合は以下のように記述します。

```ini
[nginx]
192.168.56.101
192.168.56.103
192.168.56.104

[tomcat]
192.168.56.108
```

インベントリについては、他にも様々な形の設定が可能なので、
[公式ドキュメント](https://docs.ansible.com/ansible/2.9/network/getting_started/first_inventory.html)などを参照しつつ、学んでいくと良いでしょう。

### プレイブックとタスク

プレイブック(Playbook)は設定したい事柄について順序立てて定義したファイルです。
プレイブックは1つ以上のタスクから構成されます。

![](Playbook_task.png)

タスクはどのホストグループに対して、何を設定したいかを定義する内容になります。


ではPlaybookの中身を見ていきましょう。

#### プレイブックの中身の解説

`playbook/main.yml`は、以下のように記述されています。

```yaml
---
- hosts: nginx
  roles:
    - nginx
```

ここでは、`inventory.ini`の`nginx`グループ内のホストに、
nginxロール向けのタスクを適用するといった記述になっています。

では、nginxロール向けのタスクを見ていきましょう。
`playbook/roles/nginx/tasks/main.yml`を見てみます。

```yaml
---
- import_tasks: install_nginx.yml
```

`import_task`で`install_nginx.yml`をタスクをインポートしています。
では、`install_nginx.yml`をみてみると以下のようになっています。

```yaml
---
# nginxをインストールする
- name: Install nginx
  become: yes
  apt:
    name:
      - nginx
    state: present
# サーバが起動した際にnginxのプロセスが起動するようsystemdを設定
- name: Automatically start nginx when the server starts
  become: yes
  systemd:
    name: nginx
    enabled: yes
# nginxが停止している場合は、起動
- name: Start nginx
  become: yes
  systemd:
    name: nginx
    state: started
```

全体を解説しているとキリがないので、一部のみ解説します。

```yaml
- name: Install nginx
  become: yes
  apt:
    name:
      - nginx
    state: present
```

上記の`Install nginx`タスクでは、 Ansibleのaptモジュールを使って、
nginxパッケージが導入されている状態(=存在する状態, present)にします。

## 実際にデプロイしてみる

`ansible-playbook`コマンドを使ってデプロイします。
`-i`オプションでインベントリを定義したファイルを指定し、
適用したいプレイブックを指定すると、プレイブックが実行され
nginxのインストールが行われます。

```console
$ ansible-playbook -i inventory/inventory.ini playbook/main.yml

PLAY [all] ***********************************************************************************

TASK [Gathering Facts] ***********************************************************************
ok: [192.168.56.101]

TASK [nginx : Install nginx] *****************************************************************
ok: [192.168.56.101]

TASK [nginx : Automatically start nginx when the server starts] ******************************
ok: [192.168.56.101]

TASK [nginx : Start nginx] *******************************************************************
ok: [192.168.56.101]

PLAY RECAP ***********************************************************************************
192.168.56.101             : ok=4    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   

(local) RyomanoMacBook-Pro:ansible_nginx fujiwara$ ansible-playbook -i inventory/inventory.ini playbook/main.yml 

PLAY [nginx] *********************************************************************************

TASK [Gathering Facts] ***********************************************************************
ok: [192.168.56.101]

TASK [nginx : Install nginx] *****************************************************************
ok: [192.168.56.101]

TASK [nginx : Automatically start nginx when the server starts] ******************************
ok: [192.168.56.101]

TASK [nginx : Start nginx] *******************************************************************
ok: [192.168.56.101]

PLAY RECAP ***********************************************************************************
192.168.56.101             : ok=4    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   

(local) RyomanoMacBook-Pro:ansible_nginx fujiwara$ ansible-playbook -i inventory/inventory.ini playbook/main.yml 

PLAY [nginx] *********************************************************************************

TASK [Gathering Facts] ***********************************************************************
ok: [192.168.56.101]

TASK [nginx : Install nginx] *****************************************************************
ok: [192.168.56.101]

TASK [nginx : Automatically start nginx when the server starts] ******************************
ok: [192.168.56.101]

TASK [nginx : Start nginx] *******************************************************************
ok: [192.168.56.101]

PLAY RECAP ***********************************************************************************
192.168.56.101             : ok=4    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   

(local) RyomanoMacBook-Pro:ansible_nginx fujiwara$ ansible-playbook -i inventory/inventory.ini playbook/main.yml 

PLAY [nginx] *********************************************************************************

TASK [Gathering Facts] ***********************************************************************
ok: [192.168.56.101]

TASK [nginx : Install nginx] *****************************************************************
ok: [192.168.56.101]

TASK [nginx : Automatically start nginx when the server starts] ******************************
ok: [192.168.56.101]

TASK [nginx : Start nginx] *******************************************************************
ok: [192.168.56.101]

PLAY RECAP ***********************************************************************************
192.168.56.101             : ok=4    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
```

nginxのインストールが完了したので、インベントリの`nginx`グループに含まれている
192.168.56.101にwebブラウザでアクセスしてみると、nginxがインストールされ、
動作していることが確認できます。

![](2020-04-12-20-23-23.png)