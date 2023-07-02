package load_balancer

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
	"os"
)

type Proxy struct {
	Port string `json:"port"`
}

type Backend struct {
	URL    string `json:"url"`
	IsDead bool
	mutex  sync.RWMutex
}

type Config struct {
	Proxy    Proxy     `json:"proxy"`
	Backends []Backend `json:"backends"`
}

func (backend *Backend) SetDead(b bool) {
	backend.mutex.Lock()
	backend.IsDead = b
	backend.mutex.Unlock()
}

func (backend *Backend) GetIsDead() bool {
	backend.mutex.RLock()
	isAlive := backend.IsDead
	backend.mutex.RUnlock()
	return isAlive
}

var mutex sync.Mutex
var idx int = 0
var cfg Config

func lbHandler(w http.ResponseWriter, r *http.Request) {
	maxLen := len(cfg.Backends)

	// Round Robin
	mutex.Lock()
	currentBackend := cfg.Backends[idx%maxLen]
	if currentBackend.GetIsDead() {
		idx++
	}

	targetURL, err := url.Parse(cfg.Backends[idx%maxLen].URL)
	if err != nil {
		log.Println(err.Error())
	}
	idx++
	mutex.Unlock()
	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)
	reverseProxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, e error) {
		log.Printf("%v is dead", targetURL)
		currentBackend.SetDead(true)
		lbHandler(w, r)
	}
	reverseProxy.ServeHTTP(w, r)
}

func isAlive(url *url.URL) bool {
	conn, err := net.DialTimeout("tcp", url.Host, time.Minute*1)
	if err != nil {
		log.Printf("Unreachable tp %v, error:%v", url.Host, err.Error())
		return false
	}
	defer conn.Close()
	return true
}

func HealthCheck() {
	t := time.NewTicker(time.Minute * 1)
	for {
		select {
		case <-t.C:
			for _, backend := range cfg.Backends {
				pingURL, err := url.Parse(backend.URL)
				if err != nil {
					log.Fatal(err.Error())
				}
				isAlive := isAlive(pingURL)
				backend.SetDead(!isAlive)
				msg := "alive"
				if !isAlive {
					msg = "dead"
				}
				log.Printf("%v checked %v by healthcheck", backend.URL, msg)
			}
		}
	}
}

func LoadBalancer() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filePath := wd + "/pkg/load_balancer/config.json"

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	json.Unmarshal(data, &cfg)

	go HealthCheck()

	s := http.Server{
		Addr:    ":" + cfg.Proxy.Port,
		Handler: http.HandlerFunc(lbHandler),
	}
	if err = s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
