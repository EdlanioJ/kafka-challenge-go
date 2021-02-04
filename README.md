# kafka-challenge-go
kafka challenge on fullstack fullcycle immersion by Code Education

### setup

```bash javascript
# Clone Repository
$ git clone https://github.com/EdlanioJ/kafka-challenge-go.git
$ cd kafka-challenge-go
```
Then create a .env file with the following variables:
```bash javascript
smtpHost=
smtpPort=
smtpUser=
smtpPass=
smtpSender=

kafkaBootstrapServers=
kafkaConsumerGroupId=
kafkaTransactionTopic=
```

##### * I had used [MailTrap](https://mailtrap.io) to send email

### execute:
```bash javascript
docker-compose up -d

docker exec -it [contaner name] bash

to run Consumer run:

$  go run consumer/main.go

to run Producer run:

$  go run producer/main.go
```
