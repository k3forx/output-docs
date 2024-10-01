# Elasticsearch

## 1. Introduction

- document
- フィールド: ドキュメントのデータ (RDBのカラムに対応する)

## 2. Getting Started

### 12. Understanding the basic architecture

- node
  - データの一部を保存する論理的なブロック？
  - 同じマシンの上で何個もnodeを立ち上げることができる
- cluster
  - nodeのまとまり
  - cluster同士は一般的には独立している
- document
  - データの単位
  - JSONオブジェクト
  - ESが内部で使用するメタデータと一緒に保存される
  - `_source` に保存される
- index
  - documentの論理的なグループ
  - documentを検索するときに指定する

### 13. Inspecting the cluster

## 3. Managing Documents

## 4. Mapping & Analysis
