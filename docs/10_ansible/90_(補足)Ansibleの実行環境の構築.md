# (補足)Ansibleの実行環境構築

+ Homebrewの導入
+ Homebrewを使ったpyenvの導入
+ pyenvを使ったpythonバージョンの切り替え
+ venvを使った独立したPython実行環境の準備

## Homebrewの導入

macOS向けのパッケージマネージャです。
https://brew.sh/ にアクセスし、ページ記載のコマンドを実行することで導入できます。

```console
$ /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
```

## Homebrewを使ったpyenvの導入

HomebrewをつかったPythonのバージョンを切り替えるためのツールである、
pyenvを導入します。

```console
$ brew install pyenv
```

インストールが完了したら`~/.bash_profile`ファイルに以下の内容を追記します。
(bash以外のシェルを使っている人はすみません)

```bash
export PYENV_ROOT="$HOME/.pyenv"
export PATH="$PYENV_ROOT/bin:$PATH"
eval "$(pyenv init -)
```

`~/.bash_profile`がない人はファイルを作成の上導入してください。
設定が完了したら有効化します。

```console
$ source ~/.bash_profile
```

## pyenvを使ったpythonバージョンの切り替え

pyenvを使って複数バージョンのpythonを導入することができます。
今回は、Pythonの3.7.7を使うので、以下のコマンドでPython 3.7.7をインストールします。　

```console
$ pyenv install 3.7.7 
```

このバージョンのpythonを特定ディレクトリ配下で利用する場合は、
当該ディレクトリに移動後に以下のコマンドを実行します。

```console
$ pyenv local 3.7.7
```

こうすることで`.python-version`ファイルが作成され、
当該ディレクトリに移動した際にバージョン3.7.7のPythonが
自動的に利用されるようになります。

なお、どのディレクトリでもこのバージョンのPythonを使いたい!といった場合は以下のコマンドを実行します。

```console
$ pyenv glocal 3.7.7
```

### venvを使ったPython実行環境の準備

Pythonのバージョンを切り替えても様々なプロジェクトで利用していると、
プロジェクト毎のパッケージのバージョンを管理したくなるはずです。

その場合はPythonのvenvモジュールを使って**Pythonパッケージをリポジトリ毎に独立して管理**してあげましょう。



```console
$ python -m venv local
```

すると`local`ディレクトリが作成されるので、以下のコマンドを実行します。

```console
source local/bin/activate
```

すると、コンソールのプロンプト表示が少し変わります。

+ 変更前のプロンプト表示例
```console
$
```

+ 変更後のプロンプト表示例
```
(local) $
```

なお、`local`ディレクトリ内部に含まれるのは`pip`などの実行可能ファイルや、
導入済みのパッケージ本体が入るので、`.gitignore`から除外する方が良い。


```text
local/
```
