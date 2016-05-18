package gorddo

import "io"

type RDD struct {
	Data   io.Reader
	filter *func(io.Reader) (io.Reader, err)
}

func (r *RDD) Read(p []byte) (n int, err error) {

}

func (r *RDD) Filter(t func(io.Reader) (io.Reader, err)) (*RDD, error) {
	return &RDD{
		Data:   r,
		filter: t,
	}
}
