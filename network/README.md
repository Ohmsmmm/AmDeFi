
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
  - ลง npm
```
npm install
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
  - เข้าไปที่ api-AmDeFi เพื่อ run api ของ orgที่1 โดยคำสั่ง
```
npm run start1
``` 
  - orgที่2 โดยคำสั่ง
```
npm run start2
``` 
  - orgที่3 โดยคำสั่ง
```
npm run start3
``` 
# วิธี run event listener 
```
npm run eventlistener 
```