# Google Cloud Functions Sample

Google Cloud FunctionsのGo1.11ランタイム向けのサンプルコードです。<br>
Google Cloud Storageへの保存をトリガーに関数を呼び出し画像を縮小加工します(画像はJPEG画像を対象にしています)。

### 前提

Google Cloud Platformのアカウントが必要です。
デプロイするため、gcloudコマンドがローカル環境にインストールされている事が前提です。

### 環境設定

1. コンポーネントの更新とベータ機能のインストール

```bash
$ gcloud components update
$ gcloud components install beta
```

2. Go1.11のModule機能を有効にします。

```bash
$ export GO111MODULE=on
```

3. アップロード先のバケットを作成します。

```bash
$ gsutil mb -l us-east1 gs://images-sample/
```

4. 加工された画像の保存用バケットを作成します。

```bash
$ gsutil mb -l us-east1 gs://images-sample-thum/
```

### デプロイ

```bash
$ gcloud functions deploy google-cloud-functions-sample-go 
    --entry-point OnStorageFinalize
    --runtime go111
    --trigger-event=google.storage.object.finalize
    --trigger-resource images-sample
```

### 実行方法

以下のコマンドを実行するか、GCPコンソール上から`images-sample` バケットにアップロードすることで実行されます。

```bash
$ gsutil cp [original file name] gs://images-sample/
```

`images-sample-thum` バケットに加工された画像ファルが保存されています。GCPコンソール上から参照、もしくは以下のコマンドでダウンロードできます。

```bash
$ gsutil cp gs://images-sample-thum/thum_[original file name].jpg .
```
