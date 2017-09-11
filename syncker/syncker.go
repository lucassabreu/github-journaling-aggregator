package syncker

type Syncker struct {
	ch chan bool
}

func New() Syncker {
	return Syncker{
		ch: make(chan bool),
	}
}

func (s *Syncker) Sync(f func()) {
	defer func() {
		<-s.ch
	}()
	s.ch <- true
	f()
}
