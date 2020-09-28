docker-compose up --build -d
docker exec -it kumparan_kafka_1 kafka-topics.sh --create --topic kumparan --bootstrap-server kafka:9092