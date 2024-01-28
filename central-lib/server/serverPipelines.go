package server

import "rac_oblak_proj/base_server/pipeline"

func (s *CentralLibServer) insertUserPipeline() *pipeline.Pipeline {
	p := pipeline.New("/users/signUp", s.handleUserSignUp)

	p.RegisterMiddleware(s.AllowedHost, s.PostMethodAllowed)

	return p
}

func (s *CentralLibServer) loginUserPipeline() *pipeline.Pipeline {
	p := pipeline.New("/users/login", s.handleUserLogin)

	p.RegisterMiddleware(s.AllowedHost, s.PostMethodAllowed)

	return p
}

func (s *CentralLibServer) rentBookPipeline() *pipeline.Pipeline {
	p := pipeline.New("/books/rent", s.handleRentBook)

	p.RegisterMiddleware(s.AllowedHost, s.PostMethodAllowed)

	return p
}

func (c *CentralLibServer) getPipelines() []*pipeline.Pipeline {
	return []*pipeline.Pipeline{
		c.insertUserPipeline(),
		c.loginUserPipeline(),
		c.rentBookPipeline(),
	}
}
