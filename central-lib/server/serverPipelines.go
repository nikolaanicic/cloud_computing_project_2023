package server

import "rac_oblak_proj/base_server/pipeline"

func (s *CentralLibServer) insertUserPipeline() *pipeline.Pipeline {
	p := pipeline.New("/users/insert", s.handleInsertUser)

	p.RegisterMiddleware(s.AllowedHost)

	return p
}

func (s *CentralLibServer) loginUserPipeline() *pipeline.Pipeline {
	p := pipeline.New("/users/login", s.handleUserLogin)

	p.RegisterMiddleware(s.AllowedHost)

	return p
}

func (c *CentralLibServer) getPipelines() []*pipeline.Pipeline {
	return []*pipeline.Pipeline{
		c.insertUserPipeline(),
		c.loginUserPipeline(),
	}
}
