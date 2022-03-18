** ДЗ-02. Логирование в ELK
 0. Проверяем порты, что свободные
 - Port 5044, 5601, 8080, 9200
 1. Run **docker-compose up -d**
 2. Check deployed container status 
 **docker ps**
 3. Hit golang services to write some ramdom logs using
 **curl localhost:8080/trigger**
 4. Create index pattern in Kibana  ->  Discover
  **localhost:5601**
 5.  Check our application log in Kibana -> Discover
