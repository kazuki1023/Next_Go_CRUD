## 作成物
- Next.jsとGoを使ったWebアプリケーション
- Quiz管理アプリケーション
 - CRUD処理

## やりたいこと
- 

### 参考サイト
[Next.js × Go のWebアプリをDocker上で作成してみませんか？ 〜環境構築編〜](https://qiita.com/takakou/items/a01af0515f49e90bd05c)

### 環境構築について
1. 以下のコマンドを実行してください
```
docker compose up -d
```
2. 下記URlにアクセスして、アプリケーションが起動しているか確認してください

[フロントエンド](http://localhost:3000)
[バックエンド](http://localhost:8080)

### goのパッケージについて
1. goのパッケージを追加する場合は、以下のコマンドを実行してください
```
docker compose exec backend sh
```

2. コンテナ内でgoのパッケージを追加する
