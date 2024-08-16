package main

import (
        "context"
        "log"
        // "os"
        // "strings"
        "time"

        amqp "github.com/rabbitmq/amqp091-go"
)

var Data string = `{
	   "id": 3689596,
	   "title": "بسته ۳ عددی کنسرو ماهی تون در روغن گالکسی ۱۸۰ گرمی",
	   "subtitle": "۱۸۰ گرم",
	   "description": "",
	   "content": null,
	   "max_order_cap": 1,
	   "price": 262500,
	   "discount_percent": 30,
	   "discounted_price": 183750,
	   "has_alternative": false,
	   "images": [
		   {
			   "image": "https://api.snapp.market/media/cache/product-variation_image2/uploads/images/vendors/users/app/20210818-177823-1.jpg",
			   "thumb": "https://api.snapp.market/media/cache/product_variation_transparent_image/20210818-177823-1.png"
		   },
		   {
			   "image": "https://api.snapp.market/media/cache/product-variation_image2/uploads/images/vendors/users/app/20210818-177823-2.jpg",
			   "thumb": "https://api.snapp.market/media/cache/product-variation_image_thumbnail/uploads/images/vendors/users/app/20210818-177823-2.jpg"
		   },
		   {
			   "image": "https://api.snapp.market/media/cache/product-variation_image2/uploads/images/vendors/users/app/20210818-177823-3.jpg",
			   "thumb": "https://api.snapp.market/media/cache/product-variation_image_thumbnail/uploads/images/vendors/users/app/20210818-177823-3.jpg"
		   }
	   ],
	   "brand": {
		   "id": 1138,
		   "title": "گالکسی",
		   "slug": "galaxy",
		   "english_title": "Galaxy"
	   },
	   "review_count": null,
	   "rating_value": 0,
	   "html_description": "",
	   "meta_description": "",
	   "meta_keywords": "",
	   "needs_server_approval": false,
	   "tags": [],
	   "coupons": [],
	   "badges": [
		   {
			   "title": "بسته ۳ عددی",
			   "color": "#2347fb",
			   "icon": "box"
		   }
	   ],
	   "pureTitle": "کنسرو ماهی تون در روغن گالکسی"
   }`

func failOnError(err error, msg string) {
        if err != nil {
                log.Panicf("%s: %s", msg, err)
        }
}

func main() {
        conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
        failOnError(err, "Failed to connect to RabbitMQ")
        defer conn.Close()

        ch, err := conn.Channel()
        failOnError(err, "Failed to open a channel")
        defer ch.Close()

        q, err := ch.QueueDeclare(
                "task_queue", // name
                true,         // durable
                false,        // delete when unused
                false,        // exclusive
                false,        // no-wait
                nil,          // arguments
        )
        failOnError(err, "Failed to declare a queue")

        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()

        body := Data
        err = ch.PublishWithContext(ctx,
                "",           // exchange
                q.Name,       // routing key
                false,        // mandatory
                false,
                amqp.Publishing{
                        DeliveryMode: amqp.Persistent,
                        ContentType:  "text/plain",
                        Body:         []byte(body),
                })
        failOnError(err, "Failed to publish a message")
        log.Printf(" [x] Sent %s", body)
}

// func bodyFrom(args []string) string {
//         var s string
//         if (len(args) < 2) || os.Args[1] == "" {
//                 s = "hello"
//         } else {
//                 s = strings.Join(args[1:], " ")
//         }
//         return s
// }


