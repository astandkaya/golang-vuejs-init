# golang-vuejs-test  
  
## 構築手順  
  
1. envコピー  
  `cp env.example .env`  
  
1. コンテナのビルドと起動  
  `$ docker compose build --no-cache && docker compose up -d`  
  
1. node_modulesインストール（時間掛かる）  
  `$ docker compose exec go_vue_client sh -c 'cd .. && vue create app'`  
  選択：Merge -> Default ([Vue 3] babel, eslint)  
  
1. compose.ymlから`command: sh -c "npm run serve"`のコメントアウトを外す  
  
1. コンテナを再起動  
  `$ docker compose restart`  
  
1. clientの起動確認（時間掛かる）  
  http://localhost:8798  
  
## serverの開発の進め方  
  
### ホットリロード効かない  
compose.ymlから`command: sh -c "air -c .air.toml"`のコメントアウトを外せばairが起動するが、WSL2ではホットリロードが効かない  
  
### 変更時  
下記コマンドを変更時に実行  
`$ docker compose exec go_vue_server sh -c 'air -c .air.toml'`  
  