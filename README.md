## JPEG

拡張子: `.jpg`, `.jpeg`

- **Quality**
  指定できる値: `0~100`
  値が大きいほど高画質かつファイルサイズが大きくなる。

## JPEG LI

拡張子: `.jpg`, `.jpeg`

- **Quality**
  指定できる値: `0~100`
  値が大きいほど高画質かつファイルサイズが大きくなる。
- **ChromaSubsampling**
  指定できる値: `image.YCbCrSubsampleRatio444`, `image.YCbCrSubsampleRatio440`, `image.YCbCrSubsampleRatio422`, `image.YCbCrSubsampleRatio420`
  人の目は「明るさ（輝度）」には敏感だが、「色の違い（色差）」には鈍い。
  そのため、色だけを間引いても人の目にはほとんど違いがわからない。
  - 基本用語
    - `Luma`: 明るさの情報（輝度）
    - `Chroma`: 色の情報（色差）
  - サンプリング種類
    - `4:4:4`: 無圧縮 色も明るさも間引かない
    - `4:4:0`: 縦方向に chroma を 1/2 に間引く（あまり使われることはない）
    - `4:2:2`: 横方向に Chroma を 1/2 に間引く
    - `4:2:0`: 横方向に Chroma を 1/2 に間引き、縦方向にも 1/2 に間引く
    - `4:1:1`: 横方向に Chroma を 1/4 に間引く（古い方式）
- **ProgressiveLevel**
  指定できる値: `0`, `1`, `2`
  値が大きいほどより細かくステップしていく。
  - 基本用語
    - `Baseline`: 画像をすべて読み込んでから表示させる
    - `Progressive`: 画像を読み込んだ時点で表示させる（最初はぼやけた状態から、だんだんクッキリしていく）
  - メリット
    - 圧縮率が良くなることがある
    - 最初の表示が早い
  - デメリット
    - エンコード時間が増える
    - ネット回線が遅いと高画質になるまで時間がかかる
- **OptimizeCoding**
  指定できる値: `true`, `false`
  画質を変えずにファイルサイズを小さくできるが、エンコード時間が長くなる
  - 仕組み
    JPEG では画像データを量子化した後にハフマン符号化で圧縮してる。
    この機能を有効化すると、画像の内容に最適さされたハフマン符号を使って圧縮する。
- **AdaptiveQuantization**
  指定できる値: `true`, `false`
  同じファイルサイズでも画質が向上するが、エンコード時間が長くなる
  - Adaptive Quantization (適応量子化) とは？
    - 画像の領域ごとに量子化の度合を変える
    - 視覚的に目立ちにくい部分は強く圧縮する
    - 重要な部分（エッジや顔など）は圧縮を弱くする
- **StandardQuantTables**
  指定できる値: `true`, `false`
  JPEG 標準の Annex K にある量子化テーブルを使うかどうか
  違いは分からぬ
- **FancyDownsampling**
  指定できる値: `true`, `false`
  通常のサンプリングでは、単純に chroma を間引くだけだが、この機能を使うと、周囲の画素を加味して、より自然になるように補完する。
- **DCTMethod**
  指定できる値: `jpegli.DCTFloat`, `jpegli.DCTISlow`, `jpegli.DCTIFast`
  JPEG 圧縮時に使われる離散コサイン変換のアルゴリズム
  - `jpegli.DCTFloat`: 最も高精度だが、処理に時間がかかる
  - `jpegli.DCTISlow`: 一般的なアルゴリズム
  - `jpegli.DCTIFast`: より高速だが、エンコード後の画質が若干劣化する

## JPEG XL

拡張子: `.jxl`

- **Quality**
  指定できる値: `0~100`
- **Effort**
  指定できる値: `1~10`
  値が大きいほど圧縮率が上がるが、エンコードに時間がかかる

## PNG

拡張子: `.png`

- **CompressionLevel**
  指定できる値: `-3~0` or `png.BestCompression`, `png.NoCompression`, `png.BestSpeed`, `png.DefaultCompression`
  PNG は「可逆圧縮」のため画質劣化なし。画像によっては圧縮率の差がそこまで大きくない。
  - `-3`, `png.BestCompression`: 圧縮率が高く、エンコードに時間がかかる
  - `-2`, `pn.noCompression`: 圧縮を行わない「無圧縮」
  - `-1`, `png.BestSpeed`: 圧縮率が低く、エンコード時間が短い
  - `0`, `png.DefaultCompression`: 基本的な圧縮。圧縮率と速度のバランスが良い。
- **BufferPool**
  PNG 圧縮時に使うメモリ（バッファ）を再利用する仕組み
  複数の画像を並列に処理する場合のパフォーマンスが向上する

## WebP

拡張子: `.webp`

- **Quality**
  指定できる値: `0~100`
- **Lossless**
  指定できる値: `true`, `false`
  ロスレス（可逆）圧縮をするかどうか。true にすると quality は無視される。
- **Method**
  指定できる値: `0~6`
  圧縮時間と圧縮率のバランスを調整する。値が大きいほど圧縮率が良くなるが時間がかかる。
  ロスレスモードでも有効
- **Exact**
  指定できる値: `true`, `false`
  透明領域をピクセル単位で正確に表現するかどうか。
  有効にすると圧縮効率が下がる場合がある。
  無効の場合には透明部分がずれることがある。

## AVIF

拡張子: `.avif`

- **Quality**
  指定できる値: `0~100`
- **QualityAlpha**
  指定できる値: `0~100`
  透明部分の画質を設定する。
  値が大きいほど高画質になるが、ファイルサイズが大きくなる。
- **Speed**
  指定できる値: `0~10`
  値が小さいほど時間がかかるが、圧縮率が良くなる。
- **ChromaSubsampling**
  指定できる値: `image.YCbCrSubsampleRatio444`, `image.YCbCrSubsampleRatio422`, `image.YCbCrSubsampleRatio420`
