## clone branch

## activate venv

ptest デイレクトリには入らす。

`. ptest/bin/activate`

hello good after noon

# Github コマンド チートシート

基本

```
branch -a ブランチ一覧
checkout <branch name>ブランチの切り替え
checkout -b <branc name> ブランチの作成 & 切り替え
push <origin> <branch>
pull <origin> <branch>
rmeote -a リモートとして登録されているものを一覧する
```

作業開始前の更新 & 確認

```
git pull origin <branch name>
(pull がうまくいかなければ merge)
git fetch --all 
git merge
git branch -a
```

リモートブランチをローカルに持ってくる

```
// ローカルの更新
git fetch --all 
git pull

//リモートをローカルにコピー
git checkout -b <コピーするローカルブランチ名> <origin>/<リモートブランチ名>
```

ローカルブランチをリモートにpush

```
git push -u <origin> <localBranch>
```

wisun_send.py

## Gateway 起動時

python 実行時にコマンドライン引数を入れると url


# Go ドキュメント

go_serial/docs

https://github.com/golang-standards/project-layout/blob/master/README_ja.md

に準拠
