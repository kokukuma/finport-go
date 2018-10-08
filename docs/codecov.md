## codecov
+ カバレッジを収集して, Gihtubに貼り付けるなにか.
+ CIで実行して取得したcoverrageをアップロードするぽい.

### 使い方
+ codecovでgithubのレポジトリと連携する.
  + tokenが発行される.
  + CODECOV_TOKEN="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
+ アップロードする際, 下記が正しいフォーマットじゃないと上がらなかった.
  + branch=test
  + commit=e306d84481e00e4f3e0fe12627bdc32e8e99df9e
  + slug=kokukuma/finport
