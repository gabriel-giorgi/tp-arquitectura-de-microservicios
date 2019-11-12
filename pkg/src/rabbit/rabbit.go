package rabbit

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"loyalty_go/pkg/src/domains/userProfile"
	"loyalty_go/pkg/src/repositories/catalogRepo"
	"loyalty_go/pkg/src/repositories/userProfileRepo"
)

var channel *amqp.Channel

type message struct {
	Type    string `json:"type"`
	Message userProfile.NotificationLevelUp `json:"message"`
}

func InitRabbit(){
	ch, _ := getChannel()
	err := ch.ExchangeDeclare(
		"loyalty", // name
		"direct",      // type
		false,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		panic(err)
	}
	q, err := ch.QueueDeclare(
		"loyalty",    // name
		false, // durable
		false, // delete when usused
		false,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	ch.QueueBind("loyalty", "loyalty" , "loyalty", false , nil)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			calculateExperience(d)
			d.Ack(true)
		}
	}()

	log.Printf(" [*] Waiting for msgs.")
	<-forever
}
/**
 * @api {direct} /loyalty/loyalty_notification Notificacion de LevelUp
 * @apiGroup RabbitMQ POST
 *
 * @apiDescription Loyalty envia un mensaje de catalog para notificarle que un usuario subio de nivel
 *
 * @apiExample {json} Mensaje
 *    {
 *      	"currentLevel"   :  "{nivel actual del usuario}",
 *          "currentDiscount":  "{descuento actual del usuario}",
 *    }
 *
**/
func SendLevelUpNotification (notification userProfile.NotificationLevelUp) error{
		send := message{
		Type:    "loyalty_notification",
		Message : notification,
	}

		chanel, err := getChannel()
		if err != nil {
		channel = nil
		return err
	}

		err = chanel.ExchangeDeclare(
		"loyalty",   // name
		"direct", // type
		false,    // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

	_, err = chanel.QueueDeclare(
		"loyalty_notification",    // name
		false, // durable
		false, // delete when usused
		false,  // exclusive
		false, // no-wait
		nil,   // arguments
	)

		body, err := json.Marshal(send)
		if err != nil {
		return err
	}

		err = chanel.Publish(
		"loyalty", // exchange
		"loyalty_notification",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
		Body: []byte(body),
	})
		if err != nil {
		channel = nil
		return err
	}

		log.Output(1, "Rabbit levelUP enviado")
		return nil
	}

/**
 * @api {direct} /loyalty/loyalty  Actualizar Profile
 * @apiGroup RabbitMQ GET
 *
 * @apiDescription Loyalty escucha un mensaje de Cart para actualizar el profile de un user luego del checkout
 * de un carrito.
 *
 * @apiExample {json} Mensaje
 *    {
 *      "cartId"  : "{cartId}",
 *      "userId"  : "{userId}",
 *      "articles": [
 *      "id"	  : "{articleId}",
 *      "quantity": "{value}"
 *          }, ...
 *     ]"
 *    }
 *
**/
func calculateExperience(msg amqp.Delivery) {
	fmt.Println(string(msg.Body))
	cartMsg := userProfile.CartResponse{}
	err := json.Unmarshal(msg.Body , &cartMsg)
	if err != nil {
		panic("Error during the unmarshall ")
	}
	cRepo := catalogRepo.NewRepo()
	totalSpent := 0.0
	for _, art := range cartMsg.Message.Articles{
		article := cRepo.GetArticle(art.ID)
		totalSpent += article.Price * float64(art.Quantity)
	}
	userService := userProfile.NewService(userProfileRepo.NewRepo())
	hasLevelUp :=userService.UpdateProfile(cartMsg.Message.UserID, totalSpent )
	if hasLevelUp {
		profile :=userService.GetProfile(cartMsg.Message.UserID)
		notification := userProfile.NotificationLevelUp{
			CurrentLevel:     profile.UserLevel,
			CurrentDiscount:  profile.CurrentDiscount,
		}
		SendLevelUpNotification(notification)
	}
}
func getChannel() (*amqp.Channel, error) {
	if channel == nil {
		conn, err := amqp.Dial("amqp://localhost")
		if err != nil {
			return nil, err
		}
		ch, err := conn.Channel()
		if err != nil {
			return nil, err
		}
		channel = ch
	}
	return channel, nil
}
