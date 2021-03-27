// sample
package repository

import (
	"app/src/domain"
	"time"
)

type CandleRepository interface {
	Find1sByTime(Time time.Time) (*domain.Btc1sCandle, error)
	Find1mByTime(Time time.Time) (*domain.Btc1mCandle, error)
	Find1hByTime(Time time.Time) (*domain.Btc1hCandle, error)

	Get1s(Limit int) ([]domain.Btc1sCandle, error)
	Get1m(Limit int) ([]domain.Btc1mCandle, error)
	Get1h(Limit int) ([]domain.Btc1hCandle, error)

	Create1s(*domain.BtcCandle) (isCreated bool, err error)
	Create1m(*domain.BtcCandle) (isCreated bool, err error)
	Create1h(*domain.BtcCandle) (isCreated bool, err error)

	Save1s(*domain.Btc1sCandle) (isCreated bool, err error)
	Save1m(*domain.Btc1mCandle) (isCreated bool, err error)
	Save1h(*domain.Btc1hCandle) (isCreated bool, err error)
}
