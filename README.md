# クチコミBOT

## What?

- とある企業の評判、クチコミをスクレイピングし、 LINE でお伝えする BOT
- スクレイピング先のサイトに存在している企業なら、どの企業でもクチコミを BOT 化することが可能

## How?

- コンテナ化しているため、実行環境は問わない

### 実行手順

1. `docker build -t kuchikomi_bot .`
2. `docker run -e LINE_ACCESS_TOKEN={YOUR_ACCESS_TOKEN} kuchikomi_bot`

### 開発環境

- Go 1.12