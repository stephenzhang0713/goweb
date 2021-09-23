package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
)

var globalSession *SessionManager
var providers = make(map[string]Provider)

func init() {
	globalSession, _ = NewSessionManager("memory", "sessionId", 3600)
}

type Session interface {
	Set(key, value interface{}) error //设置Session
	Get(key interface{}) interface{}  //获取Session
	Delete(key interface{}) error     //删除Session
	SessionID() string                //当前GetSessionId
}

type Provider interface {
	SessionInit(sessionId string) (Session, error)
	SessionRead(sessionId string) (Session, error)
	SessionDestroy(sessionId string) error
	GarbageCollector(maxLifeTime int64)
}

type SessionManager struct {
	cookieName  string     //cookie的名称
	lock        sync.Mutex //锁，保证并发时数据的安全一致
	provider    Provider   //管理session
	maxLifeTime int64      //超时时间
}

func RegisterProvider(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}

	if _, p := providers[name]; p {
		panic("session: Register provider is existed")
	}
	providers[name] = provider
}

func NewSessionManager(providerName, cookieName string, maxLifetime int64) (*SessionManager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", providerName)
	}
	//返回一个 SessionManager 对象
	return &SessionManager{
		cookieName:  cookieName,
		provider:    provider,
		maxLifeTime: maxLifetime,
	}, nil
}

func (manager *SessionManager) GetSessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *SessionManager) SessionBegin(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sessionId := manager.GetSessionId()
		session, _ = manager.provider.SessionInit(sessionId)
		cookie := http.Cookie{
			Name:     "manager.cookieName",
			Value:    url.QueryEscape(sessionId),
			Path:     "/",
			MaxAge:   int(manager.maxLifeTime),
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
	} else {
		sessionId, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sessionId)
	}
	return session
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("chapter3/session/login.html")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, "aaaa")
	} else {
		sess := globalSession.SessionBegin(w, r)
		sess.Set("phone", r.Form["phone"])
		//http.Redirect(w, r, "/", 302)
	}
}

func main() {
	http.HandleFunc("/login", login)         //设置路由
	err := http.ListenAndServe(":8088", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
