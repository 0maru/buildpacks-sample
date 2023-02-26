# buildpacks-sample

## Usage

### Pack CLI のインストール

```
brew install buildpacks/tap/pack
```

## node

### ローカルでビルド

#### Next
Next.js のビルドのときに`npm run build` は実行されないので、package.json のscript に`"gcp-build": "next build"` を追加する必要がある

```
pack build --builder=gcr.io/buildpacks/builder --path ./node/node-16-nextjs-13 nextjs
```

#### Go

```
pack build --builder=gcr.io/buildpacks/builder --path ./go/server go-server
```


### ローカルで起動

```
docker run -it -ePORT=8000 -p8000:8000 nextjs
```