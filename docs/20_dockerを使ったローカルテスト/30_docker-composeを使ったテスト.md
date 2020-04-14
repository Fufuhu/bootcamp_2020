# docker-composeを使ったテスト

なお、本章の作業は[Docker Desktop for Mac](https://hub.docker.com/editions/community/docker-ce-desktop-mac)のdockerおよびdocker-composeでも実施可能です。

## docker-composeを使ったテストの実行

[Fufuhu/RancherBook-samples](https://github.com/Fufuhu/RancherBook-samples)をフォークしてクローンします。

アプリケーションの構成は上記リポジトリ記載の通りです。
`docker-compose.yaml`が準備されているのでそのままテストを始めることが可能です。

以下のようにコマンドを実行することで、クライアントサイドのテストを実行します。

```console
＄ docker-compose up --build --abort-on-container-exit --exit-code-from todo-client
```
`todo-client`コンテナの終了コードを`docker-compose`コマンドの終了コードとすることで、テスト結果クライアントサイドのテストの成否を判断しています。

```console
todo-client_1  | ?      gitlab.com/fufuhu/ti_rancher_k8s_sampleapp      [no test files]
todo-client_1  | === RUN   TestProtocolWithDefaultValue
todo-client_1  | --- PASS: TestProtocolWithDefaultValue (0.00s)
todo-client_1  | === RUN   TestHostWithDefaultValue
todo-client_1  | --- PASS: TestHostWithDefaultValue (0.00s)
todo-client_1  | === RUN   TestPortWithDefaultValue
todo-client_1  | --- PASS: TestPortWithDefaultValue (0.00s)
todo-client_1  | === RUN   TestProtocolWithOptionOverride
todo-client_1  | --- PASS: TestProtocolWithOptionOverride (0.00s)
todo-client_1  | === RUN   TestHostWithOptionOveride
todo-client_1  | --- PASS: TestHostWithOptionOveride (0.00s)
todo-client_1  | === RUN   TestPortWithOptionOverride
todo-client_1  | --- PASS: TestPortWithOptionOverride (0.00s)
todo-client_1  | === RUN   TestProtocolWithConfigFileOveride
todo-client_1  | --- PASS: TestProtocolWithConfigFileOveride (0.00s)
todo-client_1  | === RUN   TestHostWithConfigFileOverride
todo-client_1  | --- PASS: TestHostWithConfigFileOverride (0.00s)
todo-client_1  | === RUN   TestPortWithConfigFileOverride
todo-client_1  | --- PASS: TestPortWithConfigFileOverride (0.00s)
todo-client_1  | === RUN   TestHostWithConfigOverrideWithFlag
todo-client_1  | --- PASS: TestHostWithConfigOverrideWithFlag (0.00s)
todo-client_1  | === RUN   TestPortWithConfigOverrideWithFlag
todo-client_1  | --- PASS: TestPortWithConfigOverrideWithFlag (0.00s)
todo-client_1  | === RUN   TestProtocolWithConfigOverrideWithFlag
todo-client_1  | --- PASS: TestProtocolWithConfigOverrideWithFlag (0.00s)
todo-client_1  | === RUN   TestUsernameWithCommandLineOption
todo-client_1  | --- PASS: TestUsernameWithCommandLineOption (0.00s)
todo-client_1  | === RUN   TestUsernameWithoutCommandLineOption
todo-client_1  | 2020/04/13 09:55:28 ユーザ名情報が設定されていません。
todo-client_1  | --- PASS: TestUsernameWithoutCommandLineOption (0.00s)
todo-client_1  | === RUN   TestPasswordWithCommandLineOption
todo-client_1  | --- PASS: TestPasswordWithCommandLineOption (0.00s)
todo-client_1  | === RUN   TestPasswordWithoutCommandLineOption
todo-client_1  | 2020/04/13 09:55:28 パスワード情報が指定されていません。
todo-client_1  | --- PASS: TestPasswordWithoutCommandLineOption (0.00s)
todo-client_1  | === RUN   TestTokenWithConfigFile
todo-client_1  | --- PASS: TestTokenWithConfigFile (0.00s)
todo-client_1  | === RUN   TestTokenWithoutConfigFile
todo-client_1  | 2020/04/13 09:55:28 トークン情報が見つかりません。
todo-client_1  | --- PASS: TestTokenWithoutConfigFile (0.00s)
todo-client_1  | PASS
todo-client_1  | ok     gitlab.com/fufuhu/ti_rancher_k8s_sampleapp/cmd  (cached)
todo-client_1  | === RUN   TestCreateConfigFile
todo-client_1  | --- PASS: TestCreateConfigFile (0.00s)
todo-client_1  | === RUN   TestLogin
todo-client_1  | 2020/04/13 09:55:58 eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6InRlc3RfdXNlciIsImV4cCI6MTU4Njc3MjA1OCwiZW1haWwiOiIifQ.Y-Mqgkl__vZ8URv4hjIzWk43U7_qb7rD1IGXQl8UFpg
todo-client_1  | --- PASS: TestLogin (0.22s)
todo-client_1  | === RUN   TestRequestPing
todo-client_1  | todo-server
todo-client_1  | --- PASS: TestRequestPing (0.00s)
todo-client_1  | === RUN   TestCreateTask
todo-client_1  | 2020/04/13 09:55:58 Task created. TaskID: 1
todo-client_1  | 2020/04/13 09:55:58 Task Title: TODO
todo-client_1  | --- PASS: TestCreateTask (0.11s)
todo-client_1  | === RUN   TestCreateTaskWithWrongStatusCode
todo-client_1  | 2020/04/13 09:55:58 &{401 Unauthorized 401 HTTP/1.1 1 1 map[Allow:[GET, POST, PATCH, DELETE, HEAD, OPTIONS] X-Frame-Options:[SAMEORIGIN] Server:[gunicorn/19.9.0] Www-Authenticate:[JWT realm="api"] Vary:[Accept] Date:[Mon, 13 Apr 2020 09:55:58 GMT] Content-Type:[application/json] Content-Length:[38]] 0xc00020e100 38 [] true false map[] 0xc000114600 <nil>}
todo-client_1  | 2020/04/13 09:55:58 期待したレスポンスステータスコード(201 Created)ではありません。
todo-client_1  | --- PASS: TestCreateTaskWithWrongStatusCode (0.00s)
todo-client_1  | === RUN   TestGetTask
todo-client_1  | 2020/04/13 09:55:59 Task created. TaskID: 2
todo-client_1  | 2020/04/13 09:55:59 Task Title: TODO
todo-client_1  | 2020/04/13 09:55:59 TargetTaskID: 2
todo-client_1  | --- PASS: TestGetTask (0.10s)
todo-client_1  | === RUN   TestGetTaskWithoutTask
todo-client_1  | 2020/04/13 09:55:59 期待したレスポンスステータスコード(200 OK)ではありません。
todo-client_1  | 2020/04/13 09:55:59 指定したタスクのIDに対応するタスクが見つかりません。
todo-client_1  | --- PASS: TestGetTaskWithoutTask (0.09s)
todo-client_1  | === RUN   TestGetTaskWithOtherUsersTask
todo-client_1  | 2020/04/13 09:55:59 Task created. TaskID: 3
todo-client_1  | 2020/04/13 09:55:59 Task Title: TODO
todo-client_1  | 2020/04/13 09:55:59 期待したレスポンスステータスコード(200 OK)ではありません。
todo-client_1  | 2020/04/13 09:55:59 指定したタスクのIDに対応するタスクが見つかりません。
todo-client_1  | --- PASS: TestGetTaskWithOtherUsersTask (0.18s)
todo-client_1  | === RUN   TestGetTaskWithWrongAuthInfo
todo-client_1  | 2020/04/13 09:55:59 期待したレスポンスステータスコード(200 OK)ではありません。
todo-client_1  | --- PASS: TestGetTaskWithWrongAuthInfo (0.14s)
todo-client_1  | === RUN   TestGetTasks
todo-client_1  | 2020/04/13 09:55:59 Task created. TaskID: 4
todo-client_1  | 2020/04/13 09:55:59 Task Title: TODO
todo-client_1  | 2020/04/13 09:55:59 Task created. TaskID: 5
todo-client_1  | 2020/04/13 09:55:59 Task Title: TODO
todo-client_1  | 2020/04/13 09:55:59 Task created. TaskID: 6
todo-client_1  | 2020/04/13 09:55:59 Task Title: TODO
todo-client_1  | 2020/04/13 09:55:59 Task created. TaskID: 7
todo-client_1  | 2020/04/13 09:55:59 Task Title: TODO
todo-client_1  | --- PASS: TestGetTasks (0.19s)
todo-client_1  | === RUN   TestDeleteTask
todo-client_1  | 2020/04/13 09:55:59 Task created. TaskID: 8
todo-client_1  | 2020/04/13 09:55:59 Task Title: TODO
todo-client_1  | 2020/04/13 09:55:59 期待したレスポンスステータスコード(200 OK)ではありません。
todo-client_1  | 2020/04/13 09:55:59 指定したタスクのIDに対応するタスクが見つかりません。
todo-client_1  | --- PASS: TestDeleteTask (0.11s)
todo-client_1  | === RUN   TestDeleteTaskWithoutTask
todo-client_1  | 2020/04/13 09:55:59 削除の為に指定したIDに対応するTaskが見つかりませんでした。
todo-client_1  | --- PASS: TestDeleteTaskWithoutTask (0.09s)
todo-client_1  | === RUN   TestDeleteTaskWithWrongAuth
todo-client_1  | 2020/04/13 09:55:59 指定されたIDに対応するTaskを削除しようとしましたが、想定外のステータスコードが返されました。
todo-client_1  | --- PASS: TestDeleteTaskWithWrongAuth (0.00s)
todo-client_1  | === RUN   TestDeleteTaskWithOtherUsersTask
todo-client_1  | 2020/04/13 09:55:59 Task created. TaskID: 9
todo-client_1  | 2020/04/13 09:55:59 Task Title: TODO
todo-client_1  | 2020/04/13 09:56:00 削除の為に指定したIDに対応するTaskが見つかりませんでした。
todo-client_1  | --- PASS: TestDeleteTaskWithOtherUsersTask (0.18s)
todo-client_1  | === RUN   TestUpdateTask
todo-client_1  | 2020/04/13 09:56:00 Task created. TaskID: 10
todo-client_1  | 2020/04/13 09:56:00 Task Title: TODO
todo-client_1  | 2020/04/13 09:56:00 Task(ID=10) is updated.
todo-client_1  | --- PASS: TestUpdateTask (0.13s)
todo-client_1  | === RUN   TestUpdateTaskWithBadRequest
todo-client_1  | 2020/04/13 09:56:00 Task created. TaskID: 11
todo-client_1  | 2020/04/13 09:56:00 Task Title: TODO
todo-client_1  | 2020/04/13 09:56:00 更新の為に指定したステータス情報が不正です
todo-client_1  | --- PASS: TestUpdateTaskWithBadRequest (0.11s)
todo-client_1  | === RUN   TestUpdateTaskWithoutTask
todo-client_1  | 2020/04/13 09:56:00 期待したレスポンスステータスコード(200 OK)ではありません。
todo-client_1  | 2020/04/13 09:56:00 指定したタスクのIDに対応するタスクが見つかりません。
todo-client_1  | 2020/04/13 09:56:00 指定したタスクのIDに対応するタスクが見つかりません。
todo-client_1  | 2020/04/13 09:56:00 更新の為に指定したIDに対応するTaskが見つかりませんでした。
todo-client_1  | --- PASS: TestUpdateTaskWithoutTask (0.09s)
todo-client_1  | === RUN   TestUpdateTaskWithOtherUsersTask
todo-client_1  | 2020/04/13 09:56:00 Task created. TaskID: 12
todo-client_1  | 2020/04/13 09:56:00 Task Title: TODO
todo-client_1  | 2020/04/13 09:56:00 期待したレスポンスステータスコード(200 OK)ではありません。
todo-client_1  | 2020/04/13 09:56:00 指定したタスクのIDに対応するタスクが見つかりません。
todo-client_1  | 2020/04/13 09:56:00 指定したタスクのIDに対応するタスクが見つかりません。
todo-client_1  | 2020/04/13 09:56:00 更新の為に指定したIDに対応するTaskが見つかりませんでした。
todo-client_1  | --- PASS: TestUpdateTaskWithOtherUsersTask (0.18s)
todo-client_1  | === RUN   TestUpdateTaskWithWrongAuth
todo-client_1  | 2020/04/13 09:56:00 Task created. TaskID: 13
todo-client_1  | 2020/04/13 09:56:00 Task Title: TODO
todo-client_1  | 2020/04/13 09:56:00 期待したレスポンスステータスコード(200 OK)ではありません。
todo-client_1  | 2020/04/13 09:56:00 期待したレスポンスステータスコード(200 OK)ではありません。
todo-client_1  | 2020/04/13 09:56:00 指定されたIDに対応するTaskを更新しようとしましたが、想定外のステータスコードが返されました。
todo-client_1  | --- PASS: TestUpdateTaskWithWrongAuth (0.09s)
todo-client_1  | PASS
todo-client_1  | ok     gitlab.com/fufuhu/ti_rancher_k8s_sampleapp/service      (cached)
todo-client_1  | 
todo-client_1  | 2020/04/13 09:57:34 生成したバイナリからpingサブコマンドを実行します。
todo-client_1  | Using config file: /root/.todo_config.yaml
todo-client_1  | pong
todo-client_1  | 2020/04/13 09:57:34 生成したバイナリからloginサブコマンドを実行します。
todo-client_1  | Using config file: /root/.todo_config.yaml
todo-client_1  | 2020/04/13 09:57:34 認証トークンを取得しました。
todo-client_1  | 2020/04/13 09:57:34 eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6InRlc3RfdXNlciIsImV4cCI6MTU4Njc3MjE1NCwiZW1haWwiOiIifQ.xoxGO6avFegn0aaPfMcNnbz9YwYNcySaQ2b_RsAZpN8
todo-client_1  | 2020/04/13 09:57:34 生成したバイナリからcreateサブコマンドを実行します。
todo-client_1  | Using config file: /root/.todo_config.yaml
todo-client_1  | 2020/04/13 09:57:34 Task created. TaskID: 18
todo-client_1  | 2020/04/13 09:57:34 Task Title: todo
todo-client_1  | ID: 18
todo-client_1  | TITLE: todo
todo-client_1  | DESCRIPTION: 
todo-client_1  | todo-description
todo-client_1  | 2020/04/13 09:57:34 生成したバイナリからgetサブコマンドを実行します。(全タスク取得)
todo-client_1  | Using config file: /root/.todo_config.yaml
todo-client_1  | ID     Title   Status  Description
todo-client_1  | 14     todo    TODO    todo-description
todo-client_1  | 15     todoGet TODO    Get
todo-client_1  | 17     updated RUNNING updated description
todo-client_1  | 18     todo    TODO    todo-description
todo-client_1  | 2020/04/13 09:57:34 生成したバイナリからgetサブコマンドを実行します。(単一タスク取得)
todo-client_1  | 2020/04/13 09:57:34 Task created. TaskID: 19
todo-client_1  | 2020/04/13 09:57:34 Task Title: todoGet
todo-client_1  | Using config file: /root/.todo_config.yaml
todo-client_1  | ID     Title   Status  Description
todo-client_1  | 19     todoGet TODO    Get
todo-client_1  | 2020/04/13 09:57:34 生成したバイナリからdeleteサブコマンドを実行します。
todo-client_1  | 2020/04/13 09:57:34 Task created. TaskID: 20
todo-client_1  | 2020/04/13 09:57:34 Task Title: todoGet
todo-client_1  | Using config file: /root/.todo_config.yaml
todo-client_1  | 2020/04/13 09:57:34 Task(ID=20) is deleted.
todo-client_1  | ID     Title   Description
todo-client_1  | 20     todoGet Get
todo-client_1  | 2020/04/13 09:57:34 生成したバイナリからupdateサブコマンドを実行します。
todo-client_1  | 2020/04/13 09:57:34 Task created. TaskID: 21
todo-client_1  | 2020/04/13 09:57:34 Task Title: todoGet
todo-client_1  | Using config file: /root/.todo_config.yaml
todo-client_1  | ID     Title   Status  Description
todo-client_1  | 21     updated RUNNING updated description
todo-client_1  | 2020/04/13 09:57:34 Task(ID=21) is updated.
todo-client_1  | 2020/04/13 09:57:34 Task(ID=21) is updated.
rancherbook-samples_todo-client_1 exited with code 0
```

個別のコマンドオプションの意味は`docker-compose`コマンドのヘルプを参照してください。

## 追加課題

1. `docker-compose-server.yaml`ファイルを作成してサーバサイドのテストも`docker-compose`で実行できるようにしてみましょう。
2. 1.で実行した結果をJUnitのxmlファイル形式で出力するようコードの追加などを実行してみましょう。
   + [jazzband/django-nose](https://github.com/jazzband/django-nose)または、[xmlrunner/unittest-xml-reporting](https://github.com/xmlrunner/unittest-xml-reporting)を使ってみてください。
3. (応用) 現状の`docker-compose.yaml`では、MySQLの起動に時間がかかるため、テストが失敗してしまうことがあります。テストの失敗を防ぐための手段について考えて実装してみてください。
   + 2.のサーバサイドのテスト用の`docker-compose`ファイルを利用しても構いません。
   + データベースはMySQLのままにしてください(SQLiteにしてしまうのは禁止)