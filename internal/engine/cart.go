package engine

import (
	"git.agehadev.com/elliebelly/jamboy/internal"
	"go.uber.org/zap"
	"os"
)

type Cart struct {
	Path string
	Data []byte
}

func (c *Cart) Load() error {
	f, err := os.Open(c.Path)

	if err != nil {
		return err
	}

	stat, err := f.Stat()

	if err != nil {
		return err
	}

	c.Data = make([]byte, stat.Size())

	count, err := f.Read(c.Data)

	if err != nil {
		return err
	}

	internal.Logger.Info("loaded cartridge", zap.Int("bytes", count))

	return nil
}
