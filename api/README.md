# วิธีการ clone .

  - เปิด terminal
  - ทำการโคลนแบบ shh โดยพิมพ์    
  ```
  git clone git@gitlab.com:CSBlockchainKmitl/api-AmDeFi.git
  ```
  - ย้ายพาทของงาน
  ```
  git checkout AmDeFi
  ```

# วิธีการ install npm
  - ลง nodemon
```
 npm install nodemon
``` 
  - node ที่ใช่ในการ install คือ v 10.16.0
 ``` 
 nvm install 10.16.0
```
  - วิธีเปลี่ยน v ของ node
```
 nvm use 10.16.0
```

# วิธี run network
  - ให้เข้าไปที่ไฟล์ start โดย
```
 cd network-AmDeFi/start
``` 
  - ถ้ายังไม่มี network ให้ clone ลงมาโดย
```
 gti clone git@gitlab.com:CSBlockchainKmitl/network-AmDeFi.git
```
  - ให้ทำการ start network โดยการพิมพ์ 
```
 ./startFabric.sh
```
# วิธี run api
  - เข้าไปที่ api-AmDeFi เพื่อ run api โดยคำสั่ง
```
 npm run start1
``` 
# วิธี merge

**Git pull dev ลงมาในเครื่อง**
- git pull 

## **Merge verify **
- git add .
- git commit  -m “pull_update”
- git checkout verify
- git merge dev

##**Up date verify**
- git add .
- git commit - m “push_verify”
- git push origin verify

##**เข้าไปที่ dev และpush**
- git checkout dev
- git merge verify
- git add .
- git commit -m “merge_verify”
- git push origin verify