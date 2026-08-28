# Backpacker Travel Sharing Platform

## 概要

バックパッカー向けの旅行共有プラットフォーム。

自分の旅行を作成・記録・公開し、他のバックパッカーがどのような旅をしているのかを参考にできるサービス。

旅行の行程や予算、旅行記などを共有し、公開されている他人の旅行をコピーして、自分の旅行計画として編集できることを特徴とする。

### 主な機能

* ユーザー登録・ログイン
* プロフィール
* 旅行の作成・編集
* 旅行期間・行程・予算の管理
* 旅行記
* 旅行の公開・非公開
* 公開された旅行の閲覧
* 他人の旅行のコピー・編集

## 技術選定

### Backend

| 項目            | 技術    |
| ------------- | ----- |
| Language      | Go    |
| Web Framework | Gin   |
| ORM           | GORM  |
| Database      | MySQL |
| Hot Reload    | Air   |

### Frontend

| 項目         | 技術         |
| ---------- | ---------- |
| Library    | React      |
| Language   | TypeScript |
| Build Tool | Vite       |

### Infrastructure

| 項目                   | 技術             |
| -------------------- | -------------- |
| Container            | Docker         |
| Container Management | Docker Compose |
| Repository           | GitHub         |
