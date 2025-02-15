「Go言語で学ぶYahoo! ID連携の黒帯ハンズオン」
=========

# はじめに

このソースコードは演習を目的として作成しています。  
脆弱な実装や、演習上の理由で実装を簡潔にしているためセキュリティ上の検証が不十分な部分があります。  
そのため実際のサービス等には利用しないでください。利用した際に生じた損害等の責任は負いかねます。  

# 概要

本内容は以下の講義で実施されたハンズオンの演習用のソースコードです。

* Go言語で学ぶYahoo! ID連携の黒帯ハンズオン
  * https://speakerdeck.com/kuralab/20200204-kurobi-hands-on

# コンテンツ

* practice00
  * Go言語環境セットアップ
* practice01
  * Authorizationリクエストの実装
* practice02
  * Tokenリクエストの実装
* practice03
  * UserInfoリクエストの実装
* practice04
  * CSRF対策の実装
* practice05
  * リプレイ攻撃対策の実装
* answer
  * practice00〜05までの模範解答


# 使い方メモ
上に記載のspeakerdeckの資料を参照します。
以下の条件でhttps://e.developer.yahoo.co.jp/からクライアントIDとシークレットを発行しておきます。
- サーバサイド
- コールバックURLには以下を指定しておくこと
  - http://localhost:8080/callback

発行したクライアントIDとシークレットは環境変数から読み込ませるために指定します。
```
export ClientID=xxxxx
export ClientSecret=yyyyy
```

あとは各種practice0Xディレクトリに移動してgo run xxxx.goを実行すると8080ポート上で実行できます。
- http://localhost:8080/

Authorizartion codeフローの実装の完全版コードはanswerディレクトリに存在しています。
