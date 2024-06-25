package main

import (
	pb "Go11Group/Kupalov-Muhammadjon/lesson44/hometask/proto/translator"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

type TranslatorServer struct {
	pb.UnimplementedTranslatorServer
}

func main() {
	listener, err := net.Listen("tcp", ":50051")

	s := grpc.NewServer()
	pb.RegisterTranslatorServer(s, &TranslatorServer{})

	if err != nil {
		log.Fatal(err)
	}

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal(err)
	}
	
}

func (t *TranslatorServer) Translate(ctx context.Context, words *pb.Request) (*pb.Translation, error) {

    res := []string{}
	for _, word := range words.Words {
		if _, ok := translations[word]; ok {
			res = append(res, translations[word])
		} else {
			res = append(res, "not found")
		}
	}
	fmt.Println(res)

	return &pb.Translation{Translations: res}, nil
}

var translations = map[string]string{
	"apple":     "olma",
	"book":      "kitob",
	"house":     "uy",
	"car":       "mashina",
	"tree":      "daraxt",
	"water":     "suv",
	"fire":      "olov",
	"sky":       "osmon",
	"earth":     "yer",
	"sun":       "quyosh",
	"moon":      "oy",
	"star":      "yulduz",
	"flower":    "gul",
	"mountain":  "tog'",
	"river":     "daryo",
	"lake":      "ko'l",
	"forest":    "o'rmon",
	"animal":    "hayvon",
	"bird":      "qush",
	"fish":      "baliq",
	"bread":     "non",
	"milk":      "sut",
	"cheese":    "pishloq",
	"butter":    "yog'",
	"egg":       "tuxum",
	"meat":      "go'sht",
	"vegetable": "sabzavot",
	"fruit":     "meva",
	"school":    "maktab",
	"teacher":   "o'qituvchi",
	"student":   "talaba",
	"pen":       "ruchka",
	"pencil":    "qalam",
	"notebook":  "daftar",
	"computer":  "kompyuter",
	"phone":     "telefon",
	"window":    "deraza",
	"door":      "eshik",
	"chair":     "stul",
	"table":     "stol",
	"bed":       "karavot",
	"cup":       "fincan",
	"plate":     "likopcha",
	"spoon":     "qoshiq",
	"fork":      "sanchqi",
	"knife":     "pichoq",
	"clothes":   "kiyim",
	"shoe":      "oyoq kiyim",
	"hat":       "shapka",
	"glasses":   "ko'zoynak",
	"watch":     "soat",
	"bag":       "sumka",
	"road":      "yo'l",
	"street":    "ko'cha",
	"city":      "shahar",
	"village":   "qishloq",
	"country":   "mamlakat",
	"world":     "dunyo",
	"family":    "oila",
	"friend":    "do'st",
	"love":      "sevgi",
	"happiness": "baxt",
	"sadness":   "g'am",
	"anger":     "g'azab",
	"fear":      "qo'rquv",
	"surprise":  "hayrat",
	"health":    "salomatlik",
	"medicine":  "dori",
	"doctor":    "shifokor",
	"hospital":  "kasalxona",
	"work":      "ish",
	"money":     "pul",
	"time":      "vaqt",
	"day":       "kun",
	"night":     "tun",
	"morning":   "ertalab",
	"afternoon": "tushdan keyin",
	"evening":   "kechqurun",
	"week":      "hafta",
	"month":     "oy",
	"year":      "yil",
	"spring":    "bahor",
	"summer":    "yoz",
	"autumn":    "kuz",
	"winter":    "qish",
}
