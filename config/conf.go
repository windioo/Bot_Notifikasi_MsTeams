package config
import(
"os"
"github.com/joho/godotenv"
"log"
"fmt"
)

func Parameter() string{
	err:= godotenv.Load("config.env")

	if  err != nil {
    log.Panic("Error loading .env file")
  }

		Token:= os.Getenv("Token")
		CHAT_ID:= os.Getenv("CHAT_ID")

 		var Pesan = os.Args[1:]
		  fmt.Println(Pesan)
	  var Urls = "https://api.telegram.org/bot" + Token +
		 "/sendMessage?chat_id=" + CHAT_ID + "&text=" + Pesan[0]
  return Urls
}
