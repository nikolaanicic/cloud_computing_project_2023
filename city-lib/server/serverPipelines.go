package server

import "rac_oblak_proj/base_server/pipeline"

func (s *CityLibServer) getAllPipeline() *pipeline.Pipeline {
	p := pipeline.New("/books/getAll", s.handleGetAllBooksRequest)

	p.RegisterMiddleware(s.Auth, s.Session, s.GetMethodAllowed)

	return p
}

func (s *CityLibServer) insertBookPipeline() *pipeline.Pipeline {
	p := pipeline.New("/books/insert", s.handleInsertBookRequest)

	p.RegisterMiddleware(s.Auth, s.Session, s.PostMethodAllowed)

	return p
}

func (s *CityLibServer) loginUserPipeline() *pipeline.Pipeline {
	p := pipeline.New("/users/login", s.handleUserLogin)

	p.RegisterMiddleware(s.PostMethodAllowed)

	return p
}

func (s *CityLibServer) rentBookPipeline() *pipeline.Pipeline {
	p := pipeline.New("/books/rent", s.handleRentBook)

	p.RegisterMiddleware(s.Auth, s.Session, s.PostMethodAllowed)

	return p
}

func (s *CityLibServer) returnBookPipeline() *pipeline.Pipeline {
	p := pipeline.New("/books/return", s.handleReturnBook)

	p.RegisterMiddleware(s.Auth, s.Session, s.PostMethodAllowed)

	return p
}

func (c *CityLibServer) getPipelines() []*pipeline.Pipeline {
	return []*pipeline.Pipeline{
		c.getAllPipeline(),
		c.insertBookPipeline(),
		c.loginUserPipeline(),
		c.rentBookPipeline(),
		c.returnBookPipeline(),
	}
}
