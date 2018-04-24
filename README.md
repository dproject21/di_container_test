# di_container_test
github.com/fgrosse/goldi を使ってみた

## 実行
```
go run main.go
```

で実行できる。

## Directory構成
- config : DIコンテナの設定yamlが入っている
- lib : goldigenで生成したインスタンスの作成、containerの作成・初期化
- sampleinterface : Interface定義と実装

## インスタンスの変え方
config/types.yml 5行目を
```
factory: NewSampleBar
```
に変えて、

```
goldigen --in config/types.yml --out lib/dependency_injection.go
```
を実行すると、lib/dependency_injection.goが書き換えられる。


## 動き方
lib.CreateContainer()でcontainerの作成を行う。
lib.CreateContainer()では、goldi.TypeRegisterを使ってconfigで定義したインスタンスを作成。Registerに登録。goldi.ContainerにRegisterを納めている。

lib.GetPrinter() で lib.containerからPrinterインスタンスを持ってくる。
p.Print()でlib.containerに納められていたPrinterのPrintが実行される。


