package session

import (
	"github.com/gofiber/fiber/v2"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

type Session struct {
	store        *session.Store
	session      *session.Session
	FiberContext *fiber.Ctx
	defaultExp   time.Duration
}

func (s *Session) Load() {
	var err error

	s.session, err = s.store.Get(s.FiberContext)
	if err != nil {
		panic(err)
	}
}

func (s *Session) Save() {
	s.SetDefaultExpiry()
	if err := s.session.Save(); err != nil {
		panic(err)
	}
}

func (s *Session) Get(key string) interface{} {
	return s.session.Get(key)
}

func (s *Session) GetString(key string) string {
	return s.session.Get(key).(string)
}

func (s *Session) Set(key string, value interface{}) {
	s.session.Set(key, value)
}

func (s *Session) SetString(key string, value string) {
	s.session.Set(key, value)
}

func (s *Session) Keys() []string {
	return s.session.Keys()
}

func (s *Session) Delete(key string) {
	s.session.Delete(key)
}

func (s *Session) Destroy() {
	if err := s.session.Destroy(); err != nil {
		panic(err)
	}
}

func (s *Session) SetExpiry(exp time.Duration) {
	s.session.SetExpiry(exp)
}

func (s *Session) SetDefaultExpiry() {
	s.session.SetExpiry(s.defaultExp)
}

func (s *Session) SetExpiryDay() {
	s.SetExpiry(time.Hour * 24)
}

func (s *Session) SetExpiryWeek() {
	s.SetExpiry(time.Hour * 24 * 7)
}

type WrapFunc func(c *fiber.Ctx, s *Session) error

func (s *Session) Wrap(c *fiber.Ctx, fn WrapFunc) error {
	s.FiberContext = c
	s.Load()
	err := fn(c, s)
	s.Save()
	return err
}

func NewSessionStorage(config ...session.Config) *Session {
	s := &Session{
		defaultExp: time.Hour,
	}

	s.store = session.New(config...)

	return s
}
