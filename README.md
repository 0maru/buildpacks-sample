# buildpacks-sample

## Usage

### Pack CLI のインストール

```
brew install buildpacks/tap/pack
```

## node

### ローカルでビルド

デフォルトではbuild を行わないので、ビルドするようオプションを追加する
```
pack build --builder=gcr.io/buildpacks/builder nextjs
```

### ローカルで起動

```
docker run -it -ePORT=8000 -p8000:8000 nextjs
```