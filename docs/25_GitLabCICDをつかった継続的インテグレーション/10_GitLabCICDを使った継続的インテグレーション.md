# GitLab CI/CDを使った継続的インテグレーション

GitLabには、[GitLab CI/CD](https://docs.gitlab.com/ee/ci/)と呼ばれる機能があり、
CIおよびCDのための機能が提供されています。

ここでは、[30_docker-composeを使ったテスト](../20_dockerを使ったローカルテスト/30_docker-composeを使ったテスト.md)で利用していたリポジトリを使ってGitLabを使った継続的インテグレーションを体験します。

## GitHubリポジトリとGitLabリポジトリの連携

GitLabにはGithub上のリポジトリの連携機能があり、
今回はこの機能を利用してGitLabのCI機能のみを利用します。

アカウントがない場合もGitHubアカウントを使ってログイン可能です。
`Sign in with`の`GitHub`を使ってください。

![](2020-04-14-12-52-25.png)

事前にGitHubのパーソナルアクセストークンを取得しておいてください。

以降はGithubのリポジトリをGitLabにミラーリングし、
GitLab CI/CDのパイプラインを動かすまでの流れになります。

![](2020-04-14-12-54-12.png)

最初に`New project`を選択します。

`CI/CD for external repo`のタブを選択して、
`Connect repositories from`の`GitHub`を選択します。

![](2020-04-14-15-20-28.png)

GitHubのパーソナルアクセストークンが要求されるので、
あらかじめ確保されていたものを入力します。

![](2020-04-14-12-59-04.png)

GitHubの認証を行うと、自身の保有するGitHubリポジトリが表示されます。
目的のリポジトリを検索して、`Connect`ボタンをクリックします。

![](2020-04-14-13-00-19.png)

リポジトリのミラーリングがスケジュールされます(ごく短時間で完了します)。

![](2020-04-14-13-01-27.png)

`CI/CD`->`Pipelines`をクリックするとパイプラインの実行結果一覧画面に遷移します。

![](2020-04-14-13-20-37.png)

初期時点では実行結果が空なので、`Run Pipeline`をクリックしてブランチを指定して実行してみましょう。

![](2020-04-14-13-25-41.png)

しばらく待つと実行結果が一覧に表示されます。

![](2020-04-14-14-24-37.png)

パイプラインの実行結果の詳細を表示すると、パイプライン内のジョブの詳細も確認できます。

![](2020-04-14-14-24-19.png)

パイプライン内でコンテナイメージをビルドしてレジストリにプッシュしているので、
GitLab Container Registry(GitLab内部のDockerイメージのレジストリ機能)を確認してみます。

![](2020-04-14-14-49-55.png)

サーバサイドのアプリケーションとクライアントサイドのアプリケーションの両方が作成されています。

![](2020-04-14-14-50-09.png)

ここまでGitLab CI/CDを使ったパイプライン実行についての手順を確認しました。
では、以降の課題にチャレンジしてみましょう。

## 簡単な解説

GitLab CI/CDのパイプライン設定はリポジトリの`.gitlab-ci.yml`に記載されています。
詳細は[公式ドキュメント](https://docs.gitlab.com/ee/ci/)に譲るとして、今回利用しているパイプライン構成については概要を説明しておきます。

```yaml
stages:
  - build
  - test
```

まず、冒頭の`stages`はパイプラインのステージを表しています。
パイプラインの中に含まれるジョブをステージを使ってグループ化することができます。
ここでは、パイプラインをコンテナイメージをビルドする`build`ステージ、
ビルドしたコンテナイメージをテストする`test`ステージに分けています。

では、`build`ステージに所属するジョブ(`todoserver_image_build`)をみてみましょう。

```yaml
todoserver_image_build:
  stage: build
  image: docker:latest
  services:
    - docker:dind
  before_script:
    - docker login -u gitlab-ci-token -p $CI_JOB_TOKEN $CI_REGISTRY
  script:
    - docker build -t ${CI_REGISTRY_IMAGE}/todo/server:${CI_COMMIT_SHA} server
    - docker push ${CI_REGISTRY_IMAGE}/todo/server:${CI_COMMIT_SHA}
  tags:
    - docker
```

`services`に`docker:dind`を指定することでDocker-in-dockerを使ったコンテナイメージのビルドが可能になります。

`before_script`内であらかじめレジストリにログインして、コンテナイメージをプッシュできるようにしておきます。

`script`内では、`docker`の`build`コマンド、`push`コマンドを使ってイメージのビルドと
レジストリへのイメージのプッシュを行なっています。

最後に`test`ステージに所属するジョブ(`todoserver_test`)をみていきましょう。

```yaml
todoserver_test:
  stage: test
  services:
    - name: mysql:5.7
      alias: todo-mysql
      command:
        - mysqld 
        - --character-set-server=utf8 
        - --collation-server=utf8_unicode_ci

  variables:
    # MySQL関連の設定
    MYSQL_USER: test #ユーザ
    MYSQL_PASSWORD: test #ユーザパスワード
    MYSQL_ROOT_PASSWORD: test  #MySQLのrootパスワード
    MYSQL_DATABASE: test #データベース名
    MYSQL_HOST: todo-mysql #データベースホスト名
    # 動作環境設定
    TODO_SERVER_ENVIRONMENT: production
    # 利用するDjangoの設定ファイル
    DJANGO_SETTINGS_MODULE: sampleapp.settings_mysql

  image: ${CI_REGISTRY_IMAGE}/todo/server:${CI_COMMIT_SHA}
  script:
    - |
      cd server/sampleapp
      python manage.py test -v 2  --noinput
```

`services`内で、`todo-mysql`の名前をつけた`mysql:5.7`のイメージを起動しています。
環境変数としてMYSQLに与える環境変数などを設定しつつ、サーバサイドアプリケーションで
テストを実行する際に必要となる環境変数をセットしています。

`script`では、単純に`python manage.py test`を実行してテスト結果を取得しています。

## 課題

1. GitLab CI/CDの出力からサーバサイドアプリケーションのテスト結果を見れるようにしましょう。
   + [30_docker-composeを使ったテスト](../20_dockerを使ったローカルテスト/30_docker-composeを使ったテスト.md)の課題記載と同様に[jazzband/django-nose](https://github.com/jazzband/django-nose)または、[xmlrunner/unittest-xml-reporting](https://github.com/xmlrunner/unittest-xml-reporting)を使ってみてください。
   + (参考) [JUnit test reports](https://docs.gitlab.com/ee/ci/junit_test_reports.html)
2. (応用)テストのカバレッジを取得できるようにしてみましょう。
   + [Test coverage parsing](https://docs.gitlab.com/ee/ci/pipelines/settings.html#test-coverage-parsing)
