package gorddo

import "os"

func TextFile(file_location string) (*RDD, err) {
	file, err := os.Open(file_location)
	if err != nil {
		return nil, err
	}

	return &RDD{
		Data: file,
	}
}
