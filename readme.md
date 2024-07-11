## 作成物
- Next.jsとGoを使ったWebアプリケーション
- Quiz管理アプリケーション
 - CRUD処理

## やりたいこと
- 

### 参考サイト
[Next.js × Go のWebアプリをDocker上で作成してみませんか？ 〜環境構築編〜](https://qiita.com/takakou/items/a01af0515f49e90bd05c)
[Next.js × Go のWebアプリをDocker上で作成してみませんか？ 〜DB編〜](https://qiita.com/takakou/items/4f8cd3686c7ec84c141d)

### 環境構築について
1. 以下のコマンドを実行してください
```
docker compose build --no-vache && docker compose up -d
```
2. 下記URlにアクセスして、アプリケーションが起動しているか確認してください

[バックエンド](http://localhost:8080)

3. フロントエンドの環境を構築する場合は、以下のコマンドを実行してください
```
cd frontend
npm install
npm run dev
```

### goのパッケージについて
1. goのパッケージを追加する場合は、以下のコマンドを実行してください
```
docker compose exec backend sh
```

2. コンテナ内でgoのパッケージを追加する

### テーブル確認方法

### これからのtodo
- 環境構築
  - [x] フロントエンドの環境をdockerじゃなくす
- バックエンド
  - [x] postgresqlとの接続を確立する
