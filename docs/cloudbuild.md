## CloudBuild
+ CI

###

+ 鍵の準備
  ```
  gcloud kms keyrings create finport-keyring \
    --location=global
  gcloud kms keys create finport-key \
    --location=global \
    --keyring=finport-keyring \
    --purpose=encryption
  ```

+ CloudBuildからアクセスできるようにする.
  + 下記だとなんかうまくいかなかったので, consoleから叩いた.
  ```
  gcloud kms keys add-iam-policy-binding \
      finport-key --location=global --keyring=finport-keyring \
      --member=serviceAccount:664154218711@cloudbuild.gserviceaccount.com \
      --role=roles/cloudkms.cryptoKeyEncrypterDecrypter
  ```

+ 環境変数
  ```
  echo -n $CODECOV_TOKEN | gcloud kms encrypt \
    --plaintext-file=- \
    --ciphertext-file=- \
    --location=global \
    --keyring=finport-keyring \
    --key=finport-key | base64
  ```

### cloudbuildに関して思うこと.
+ Secretの持ち方がめんどくさすぎる
+ cloud-builders-communityのimage, どこかにホスティングしてほしい.
  + 自分でビルドして, 自分のプロジェクトのGCRにもつとかだるい.
+ 各stepでmount先を指定できないのはなんで?
  + GOPATHどうしたらいいの感
