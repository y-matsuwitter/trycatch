package trycatch

type TryCatch struct {
	errChan chan interface{}
}

func (t TryCatch) Try(block func()) TryCatch {
	t.errChan = make(chan interface{})
	go func() {
		defer func() {
			err := recover()
			t.errChan <- err
		}()
		block()
	}()
	return t
}

func (t TryCatch) Catch(block func(err error)) TryCatch {
	err := <-t.errChan
	if err != nil {
		block(err.(error))
	}
	return t
}
