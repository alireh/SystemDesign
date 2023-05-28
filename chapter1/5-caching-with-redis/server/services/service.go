package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"

	"strconv" // new import

	"github.com/gorilla/mux" // new import

	"context"

	mdl "caching-with-redis/models"

	"github.com/go-redis/redis/v8"
)

var dbconn *gorm.DB

type Response struct {
	Data    []mdl.Post `json:"data"`
	Message string     `json:"message"`
}

func SetDB(db *gorm.DB) {
	dbconn = db
	var post = mdl.GetPost()
	dbconn.AutoMigrate(&post)
}

var ctx = context.Background()

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if pong := client.Ping(ctx); pong.String() != "ping: PONG" {
		log.Println("-------------Error connection redis ----------:", pong)
	} else {
		var resp Response
		postKey := fmt.Sprintf("%s%d", "post_", id)
		val, err := client.Get(ctx, postKey).Result()
		if err == nil && val != "null" {
			model := mdl.Post{}
			b1 := []byte(val)
			//b1 := []byte(`{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"ali","age":10}`)
			json.Unmarshal(b1, &model)
			resp.Data = append(resp.Data, model)
			resp.Message = "SUCCESS"
			json.NewEncoder(w).Encode(&resp)
		} else {
			var post = mdl.GetPost()
			err := dbconn.Where("id = ?", id).Find(&post).Error

			if err == nil {

				log.Println(post)
				resp.Data = append(resp.Data, post)
				resp.Message = "SUCCESS"
				json.NewEncoder(w).Encode(&resp)

				_, err1 := client.Ping(ctx).Result()
				b, _ := json.Marshal(resp.Data[0])
				err1 = client.Set(ctx, postKey, string(b), 0).Err()
				if err1 != nil {
					log.Panic(err1)
				}

			} else {
				log.Println(err)
				http.Error(w, err.Error(), 400)
			}
		}
	}

}
