package main

import(
	"log"
	"time"
	"github.com/d2r2/go-dht"
)

func main() {	
	//di sini menggunakan GPIO 4
	temperature, humidity, retried, err := dht.ReadDHTxxWithRetry(dht.DHT22, 4, false, 10)//jika ingin mengganti jadi dht11 ganti jadi dht.DHT11
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n\n",temperature, humidity, retried)
	for {
		temperature, humidity, retried, err = dht.ReadDHTxxWithRetry(dht.DHT22, 4, false, 10)//jika ingin mengganti jadi dht11 ganti jadi dht.DHT11
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n\n",temperature, humidity, retried)	
		time.Sleep(1 * time.Second)
	}
}

