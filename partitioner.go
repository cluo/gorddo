package gorddo

import (
	"flag"
	"io"
	"log"
	"os"
)

const (
	DEFAULT_BLOCKSIZE       = 32 * 1024 * 1024
	DEFAULT_PARTITION_COUNT = 0
)

var (
	blockSize = flag.Int64("blocksize",
		DEFAULT_BLOCKSIZE,
		"desired size of each partition in bytes")
	numOfPartitions = flag.Int64("num_of_partitions",
		DEFAULT_PARTITION_COUNT,
		"desired number of partitions")
)

type Partition struct {
	data []byte
}

func NewPartition(data []byte) *Partition {
	return &Partition{
		data: data,
	}
}

type Partitioner struct {
	blocksize      int64
	min_partitions int64

	partitions []*Partition
}

func NewPartitioner(blocksize, min_partitions int64) *Partitioner {
	return &Partitioner{
		blocksize:      blocksize,
		min_partitions: min_partitions,
	}
}

func (p *Partitioner) Partition(file_path string) error {
	file, err := os.Open(file_path)
	if err != nil {
		return err
	}

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	file_size := stat.Size()

	remainder := file_size % p.blocksize
	partition_count := file_size / p.blocksize
	if remainder > 0 {
		partition_count += 1
	}

	if partition_count < p.min_partitions &&
		p.min_partitions != DEFAULT_PARTITION_COUNT {

		partition_count = p.min_partitions
	}

	log.Printf("Creating %s partitions", partition_count)

	partition_size := file_size / partition_count

	buffer := make([]byte, partition_size)

	for {
		bytes_read, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}

		if bytes_read == 0 {
			break
		}

		partition, remaining_bytes, err := NewPartition(buffer)
		if err != nil {
			return err
		}

		p.partitions = append(p.partitions, partition)

		_, err = file.Seek(-len(remaining_bytes), 1)
		if err != nil {
			return err
		}
	}
	return nil
}
