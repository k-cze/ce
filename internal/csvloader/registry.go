package csvloader

import (
	"ce/internal/csvloader/crypto"
	"path/filepath"
)



type CSVDrivers struct {
	crypto crypto.Interface
}

func (d *CSVDrivers) Crypto() crypto.Interface {
	return d.crypto
}

type Drivers interface {
	Crypto() crypto.Interface
}

func InitAllDrivers() (*CSVDrivers, error) {
	cryptoDriver := &crypto.Driver{}

	processors := []CSVProcessor{
		WrapProcessor(NewProcessor[crypto.CryptoRate](
			filepath.Join(dataDir, crypto.CsvFileName),
			cryptoDriver,
			crypto.ParseCryptoRate,
		)),
	}

	for _, proc := range processors {
		if err := proc.Process(); err != nil {
			return nil, err
		}
	}

	return &CSVDrivers{
		crypto: cryptoDriver,
	}, nil
}
