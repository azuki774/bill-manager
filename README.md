# [WIP] bill-manager

## Description
- 公共料金の明細やクラウドの請求書を、スクレイピングで取りに行く。
- データを溜め込んで、APIとしてデータを提供する
- Rest API の家計簿サーバや、他クライアントと連携する

## Target
- この辺を使ってみる
    - gRPC
    - Twitter API v2
    - 各種kubernetesリソース
    - ORM

## Progress
- 全体のアーキテクチャを決める
- 公共料金の明細の一部として、電力消費量を取得して毎日出力する←今ここ
