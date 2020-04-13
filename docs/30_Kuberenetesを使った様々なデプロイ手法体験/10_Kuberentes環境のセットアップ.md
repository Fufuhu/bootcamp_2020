# Kuberenetes環境のセットアップ

## サーバサイドの動作環境セットアップ

https://github.com/Fufuhu/ansible_k3s_install の手順にしたがって、Kuberenetesディストリビューションの[k3s](https://k3s.io/)の動作環境のセットアップを行ってください。

## クライアントサイドの動作環境セットアップ

操作端末(macOS)上にKubernetesを操作するCLIである、
kubectlの導入と、アクセスに必要な認証情報の配置を行います。

### kubectlの導入

Kubernetesのクラスタを操作するにはkubectlの導入が必要です。
公式ドキュメントの[macOSへkubectlをインストールする](https://kubernetes.io/ja/docs/tasks/tools/install-kubectl/#install-kubectl-on-macos)の項を参考にしつつ導入してください。

### kubeconfigファイルの導入

すでにkubeconfigファイルが導入済みの人は退避させてください。

```console
$ mkdir ~/.kube
```

`.kube`ディレクトリの配下に`config`のファイル名で、
k3sをデプロイしたサーバの`/etc/rancher/k3s/k3s.yaml`を`~/.kube/config`に配置しましょう。

認証情報についてはホスト指定が少し修正が必要なので、その内容をいかに示します。

+ 修正前

```yaml
    server: https://127.0.0.1:6443
```

+ 修正後

```yaml
    server: https://192.168.56.101:6443
```

### 動作確認

```console
$ kubectl get nodes
NAME                  STATUS   ROLES    AGE   VERSION
fujiwara-virtualbox   Ready    master   62m   v1.17.4+k3s1
```
