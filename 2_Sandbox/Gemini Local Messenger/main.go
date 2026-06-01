package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

type ClientInfo struct {
	Nickname string
	Platform string
	Addr     string
}

var (
	// Теперь в мапе храним соединение и текущий никнейм пользователя
	clients   = make(map[*websocket.Conn]ClientInfo)
	broadcast = make(chan string)
	mutex     sync.Mutex
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	log.Println("Мессенджер запущен на :8080. Используйте /newnickname [имя] для смены ника.")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal("Ошибка сервера: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Извлекаем данные устройства из заголовков
	userAgent := r.Header.Get("User-Agent")
	platform := getPlatform(userAgent)
	remoteAddr := r.RemoteAddr

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	mutex.Lock()
	clients[ws] = ClientInfo{
		Nickname: "Аноним",
		Platform: platform,
		Addr:     remoteAddr,
	}
	mutex.Unlock()

	for {
		_, msgBytes, err := ws.ReadMessage()
		if err != nil {
			mutex.Lock()
			delete(clients, ws)
			mutex.Unlock()
			break
		}

		msg := string(msgBytes)

		mutex.Lock()
		client := clients[ws]

		// ОБРАБОТКА КОМАНДЫ НИКА
		if strings.HasPrefix(msg, "/newnickname ") {
			newNick := strings.TrimPrefix(msg, "/newnickname ")
			clients[ws] = ClientInfo{Nickname: newNick, Platform: client.Platform, Addr: client.Addr}
			broadcast <- fmt.Sprintf("📢 %s теперь %s", client.Nickname, newNick)
			mutex.Unlock()
			continue
		}

		// АДМИН-ОТЛАДКА В ТЕРМИНАЛ
		log.Printf("[DEBUG] MSG: %s | User: %s | OS: %s | IP: %s",
			msg, client.Nickname, client.Platform, client.Addr)

		mutex.Unlock()

		broadcast <- fmt.Sprintf("[%s]: %s", client.Nickname, msg)
	}
}

func handleMessages() {
	for {
		msg := <-broadcast

		mutex.Lock()
		for client := range clients {
			// Отправляем сообщение обратно всем клиентам как простой текст
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Printf("Ошибка отправки: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}

func getPlatform(ua string) string {
	ua = strings.ToLower(ua)
	if strings.Contains(ua, "android") {
		return "Android"
	} else if strings.Contains(ua, "iphone") || strings.Contains(ua, "ipad") {
		return "iOS"
	} else if strings.Contains(ua, "windows") {
		return "Windows"
	} else if strings.Contains(ua, "macintosh") {
		return "macOS"
	} else if strings.Contains(ua, "linux") {
		return "Linux"
	}
	return "Unknown Device"
}
