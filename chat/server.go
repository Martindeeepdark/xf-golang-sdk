package chat

type Server struct {
	appID     string
	apiKey    string
	apiSecret string
	hosturl   string
	cache     map[string]*Session
}

func NewServer(appID, apiKey, apiSecret, hosturl string) *Server {
	svr := &Server{
		appID:     appID,
		apiKey:    apiKey,
		apiSecret: apiSecret,
		hosturl:   hosturl,
		cache:     map[string]*Session{},
	}

	return svr
}

func (s *Server) GetSession(uid string) (*Session, error) {
	session, ok := s.cache[uid]

	if !ok {
		session = NewSession(s, uid)
		s.cache[uid] = session
	}

	return session, nil
}
