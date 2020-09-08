package main

type shareFacade struct {
	twitter *codeTwit
	google  *googlize
	reddit  *reddiator
}

func (s *shareFacade) newShareFacade(u, t, st string) {
	s.twitter = &codeTwit{
		status: st,
		url:    u,
	}

	s.google = &googlize{
		url: u,
	}
	s.reddit = &reddiator{
		title: t,
		url:   u,
	}
}

func (s *shareFacade) share() {
	s.twitter.tweet()
	s.google.share()
	s.reddit.reddit()
}
