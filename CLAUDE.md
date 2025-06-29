# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## プロジェクト概要

`timeins`パッケージは、JSON APIで時刻を秒精度でシリアライズするためのGoライブラリです。標準の`time.Time`をラップしたカスタム`Time`型を提供し、JSON出力を`2006-01-02T15:04:05-07:00`形式で秒精度にフォーマットします。

## コマンド

### 開発コマンド
```bash
# カバレッジ付きでテスト実行
make test

# リント実行（golangci-lint使用）
make lint

# ベンチマーク実行
make benchmark

# 特定のテストを実行
go test -v -run TestFunctionName

# 特定のファイルのテストを実行
go test -v timeins_test.go timeins.go
```

### Goバージョン
- 最小Goバージョン: 1.20
- CI環境でのテスト対象: 1.20.x, 1.21.x, 1.22.x, 1.23.x

## アーキテクチャ

外部依存のないシングルパッケージライブラリ：

- `timeins.go` - カスタム`Time`型のコア実装
- `timeins_test.go` - Time型機能の単体テスト
- `json_test.go` - JSONマーシャリング/アンマーシャリングの統合テスト

主な設計方針：
- `Time`型は`time.Time`を埋め込み、標準の時刻機能を継承
- `MarshalJSON`で秒単位に切り捨ててRFC3339形式でフォーマット
- `UnmarshalJSON`は柔軟性のため様々な時刻形式を受け入れ
- すべてのメソッドは新しいTimeインスタンスを返すことで不変性を保持

## テスト方針

テーブル駆動テストを広範囲に使用。新機能追加時：
1. 関連するテストテーブルにテストケースを追加
2. 正常系と異常系の両方のテストケースをカバー
3. ベンチマークを実行してパフォーマンス低下がないことを確認

## CI/CD

GitHub Actionsがプッシュごとに自動実行：
- マルチプラットフォームテスト（Ubuntu、macOS、Windows）
- 複数のGoバージョンでのテスト
- golangci-lintでのリント
- ベンチマーク実行

ワークフローは`.github/workflows/ci.yml`で定義。

## 開発フロー（GitHub Flow）

このプロジェクトではGitHub Flowを採用しています。

**重要**: ファイルの変更を伴う作業を行う場合は、必ず新しいブランチを作成してから作業を開始してください。mainブランチへの直接的な変更は禁止されています。

### ブランチ命名規則
- **機能追加**: `feature/機能名` （例: `feature/add-timezone-support`）
- **バグ修正**: `fix/修正内容` （例: `fix/parsing-error`）
- **ドキュメント**: `docs/更新内容` （例: `docs/update-readme`）
- **リファクタリング**: `refactor/対象` （例: `refactor/improve-performance`）
- **その他**: `chore/作業内容` （例: `chore/update-dependencies`）

### 開発手順
1. issueを作成または選択
2. **必ず**mainブランチから適切な名前で新しいブランチを作成
  ```bash
  # まずmainブランチに切り替えて最新の状態に更新
  git checkout main
  git pull origin main

  # 新しいブランチを作成
  git checkout -b feature/your-feature-name
  ```
3. 変更を実装し、テストとリントを実行
  ```bash
  make test
  make lint
  ```
4. コミットを作成（コミットメッセージは明確に）
5. GitHubにプッシュしてPull Requestを作成
6. CIがすべてパスすることを確認
7. レビュー後、mainブランチにマージ

### 注意事項
- mainブランチへの直接プッシュは禁止
- どんな小さな変更でも必ずブランチを作成してPRを経由する
- 作業開始前に必ずmainブランチを最新に更新する

### コミットメッセージのフォーマット

**重要**: Claude Codeでコミット作成時は必ずこのフォーマットに従ってください:

```
タイトル (50文字以内)

- 変更内容の詳細説明
- 箇条書きで記載

🤖 Generated with Claude Code
```

**注意**: リンクやCo-Authored-Byは追加しないでください。

### Pull Request 作成時の設定

Claude Code で PR を作成する際は、以下の設定を行ってください:

* **タイトルと本文は日本語で記載する**
* Assignee に作業者を追加（pushしたユーザーとPRの作業者を一致させる）
* label に `Claude Code` を追加
* GitHub Copilot をレビュワーとして追加（Web UIで手動設定）

```bash
gh pr create --title "日本語のタイトル" --body "日本語の本文" --assignee @me --label "Claude Code"
```

**注意**:
- PRのタイトルと説明文は必ず日本語で記載してください
- GitHub CopilotのレビュワーはCLIから設定できないため、PR作成後にWeb UIで手動追加してください
- **PR作成後に追加変更を行った場合は、必ずPRの説明文を更新してください**
  - 新しく追加した機能や修正内容を説明文に反映
  - 変更の経緯や理由を明記
  - レビュワーが変更内容を理解できるよう詳細に記載
