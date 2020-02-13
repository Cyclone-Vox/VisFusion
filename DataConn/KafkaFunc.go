package DataConn

import (
	"github.com/Shopify/sarama"
	"log"
)

var SaramaChanByManual =make(chan *sarama.ProducerMessage,100)
var SaramaChanByRandom =make(chan *sarama.ProducerMessage,100)

var KafkaWaitReload =make(chan struct{})



func KafkaEst(){
	var ProducerManualCloseChan=make(chan struct{})
	var ProducerRandomCloseChan=make(chan struct{})
	for {

		AsyncProducerSelectByManual(CfgLoad.Conf.ZooUrl,ProducerManualCloseChan)
		AsyncProducerSelectByRandom(CfgLoad.Conf.ZooUrl,ProducerRandomCloseChan)

		KafkaWaitReload<- struct{}{}

		ProducerManualCloseChan<- struct{}{}
		ProducerRandomCloseChan<- struct{}{}
	}


}


func AsyncProducerSelectByManual(addr []string,signals chan struct{})  {

	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewManualPartitioner
	producer, err := sarama.NewAsyncProducer(addr, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	//signals := make(chan os.Signal, 1)
	//signal.Notify(signals, os.Interrupt)
	var enqueued, errors int

ProducerLoop:


	for {

		select {
		case producer.Input() <- <-SaramaChanByManual:
			enqueued++
		case err := <-producer.Errors():
			log.Println("Failed to produce message", err)
			errors++
		case <-signals:
			break ProducerLoop
		//default:
		//	log.Println("Erwin Schrodinger‘s Cat")
		}
	}

	log.Printf("Enqueued: %d; errors: %d\n", enqueued, errors)
}

func AsyncProducerSelectByRandom(addr []string,signals chan struct{})  {
	//config := sarama.NewConfig()
	//config.Producer.Partitioner = sarama.NewRandomPartitioner
	producer, err := sarama.NewAsyncProducer(addr, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	//signals := make(chan os.Signal, 1)
	//signal.Notify(signals, os.Interrupt)
	var enqueued, errors int

ProducerLoop:


	for {

		select {
		case producer.Input() <- <-SaramaChanByRandom:
			enqueued++
		case err := <-producer.Errors():
			log.Println("Failed to produce message", err)
			errors++
		case <-signals:
			break ProducerLoop
		//default:
		//	log.Println("Erwin Schrodinger‘s Cat")
		}
	}

	log.Printf("Enqueued: %d; errors: %d\n", enqueued, errors)
}


//func KafkaCreatTopic(addr string,topicName string){
//	config := sarama.NewConfig()
//	config.Producer.Return.Successes = true
//
//	//broker:=sarama.NewBroker(addr)
//	//broker.CreateTopics()
//	admin, err := sarama.NewClusterAdmin([]string{addr}, config)
//	fmt.Println(err)
//	err=admin.CreateTopic(topicName,&sarama.TopicDetail{},false)
//	fmt.Println(err)
//}
