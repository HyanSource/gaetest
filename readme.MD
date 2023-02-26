#### 佈署到GAE

- clone
git clone https://github.com/HyanSource/gaetest
- 安裝
gcloud components install app-engine-go
- 如果有問題輸入下面語法
sudo apt-get install google-cloud-sdk-app-engine-go
- 佈署
gcloud app deploy
- 開啟
gcloud app browse

PS:
runtime的go版本需要注意一下