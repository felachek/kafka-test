
# Run Kafka on Apple with M1 chip

###  Download as zip file
[confluent](https://www.confluent.io/get-started/?product=software)
and update in server.properties
folder path kafka/etc/server.properties
```
listeners=PLAINTEXT://0.0.0.0:9092

advertised.listeners=PLAINTEXT://localhost:9092

listener.security.protocol.map=PLAINTEXT:PLAINTEXT,SSL:SSL,SASL_PLAINTEXT:SASL_PLAINTEXT,SASL_SSL:SASL_SSL
```

### Add in .zshrc
```
export CONFLUENT_HOME=/Users/<username>/confluent-6.0.1
export PATH=$PATH:$CONFLUENT_HOME/bin
```

### Start kafka
```
sdk use java 11.0.9.hs-adpt
arch -x86_64 confluent local services start
```

### Send the message
```
kafka-console-producer --topic test --broker-list localhost:9092
```


### Receive the message
```
kafka-console-consumer --topic test --bootstrap-server localhost:9092 --from-beginning
```

### Stop and destroy kafka
```
confluent local services stop
confluent local destroy
```

### Test app
1. Start kafka
2. Open browser to http://localhost:9021
3. Run the app with ```go run .```

