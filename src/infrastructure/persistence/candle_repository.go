package persistence

import (
	"app/src/domain"
	"app/src/domain/repository"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type CandleRepositoryImpl struct {
	Conn *gorm.DB
}

// CandleRepositoryWithRDB returns initialized CandleRepositoryImpl
func CandleRepositoryWithRDB(conn *gorm.DB) repository.CandleRepository {
	return &CandleRepositoryImpl{Conn: conn}
}

func (r *CandleRepositoryImpl) Find1sByTime(Time time.Time) (*domain.Btc1sCandle, error) {
	Btc1sCandle := &domain.Btc1sCandle{}
	log.Printf("repo Find1sByTime : %v", Time)
	if err := r.Conn.Where("Time = ?", Time).Find(&Btc1sCandle).Error; err != nil {
		return nil, err
	}
	return Btc1sCandle, nil
}
func (r *CandleRepositoryImpl) Find1mByTime(Time time.Time) (*domain.Btc1mCandle, error) {
	Btc1mCandle := &domain.Btc1mCandle{}
	log.Printf("repo Find1mByTime : %v", Time)
	if err := r.Conn.Where("Time = ?", Time).Find(&Btc1mCandle).Error; err != nil {
		return nil, err
	}
	return Btc1mCandle, nil
}
func (r *CandleRepositoryImpl) Find1hByTime(Time time.Time) (*domain.Btc1hCandle, error) {
	Btc1hCandle := &domain.Btc1hCandle{}
	log.Printf("repo Find1hByTime : %v", Time)
	if err := r.Conn.Where("Time = ?", Time).Find(&Btc1hCandle).Error; err != nil {
		return nil, err
	}
	return Btc1hCandle, nil
}
func (r *CandleRepositoryImpl) Get1s(Limit int) ([]domain.Btc1sCandle, error) {
	Btc1sCandles := []domain.Btc1sCandle{}
	if err := r.Conn.Order("time desc").Limit(Limit).Find(&Btc1sCandles).Error; err != nil {
		return nil, err
	}
	return Btc1sCandles, nil
}
func (r *CandleRepositoryImpl) Get1m(Limit int) ([]domain.Btc1mCandle, error) {
	Btc1mCandles := []domain.Btc1mCandle{}
	if err := r.Conn.Order("time desc").Limit(Limit).Find(&Btc1mCandles).Error; err != nil {
		return nil, err
	}
	return Btc1mCandles, nil
}
func (r *CandleRepositoryImpl) Get1h(Limit int) ([]domain.Btc1hCandle, error) {
	Btc1hCandles := []domain.Btc1hCandle{}
	if err := r.Conn.Order("time desc").Limit(Limit).Find(&Btc1hCandles).Error; err != nil {
		return nil, err
	}
	return Btc1hCandles, nil
}

func (r *CandleRepositoryImpl) Create1s(Candle *domain.BtcCandle) (isCreated bool, err error) {
	log.Printf("repo Create1s : %v", Candle)

	Btc1sCandle := domain.Btc1sCandle{BtcCandle: *Candle}
	if err := r.Conn.Create(&Btc1sCandle).Error; err != nil {
		log.Printf("repo Create1s errrr")

		return false, err
	}
	log.Printf("repo Create1s true")
	return true, nil
}

func (r *CandleRepositoryImpl) Create1m(Candle *domain.BtcCandle) (isCreated bool, err error) {
	log.Printf("repo Create1m : %v", Candle)

	Btc1mCandle := domain.Btc1mCandle{BtcCandle: *Candle}
	if err := r.Conn.Create(&Btc1mCandle).Error; err != nil {
		log.Printf("repo Create1m errrr")

		return false, err
	}
	log.Printf("repo Create1m true")
	return true, nil
}
func (r *CandleRepositoryImpl) Create1h(Candle *domain.BtcCandle) (isCreated bool, err error) {
	log.Printf("repo Create1h : %v", Candle)

	Btc1hCandle := domain.Btc1hCandle{BtcCandle: *Candle}
	if err := r.Conn.Create(&Btc1hCandle).Error; err != nil {
		log.Printf("repo Create1h errrr")

		return false, err
	}
	log.Printf("repo Create1h true")
	return true, nil
}

func (r *CandleRepositoryImpl) Save1s(Btc1sCandle *domain.Btc1sCandle) (isSaved bool, err error) {
	log.Printf("repo Save1s : %v", Btc1sCandle)

	// Btc1sCandle := Candle
	if err := r.Conn.Save(&Btc1sCandle).Error; err != nil {
		log.Printf("repo Save1s errrr")

		return false, err
	}
	log.Printf("repo Save1s true")
	return true, nil
}

func (r *CandleRepositoryImpl) Save1m(Btc1mCandle *domain.Btc1mCandle) (isSaved bool, err error) {
	log.Printf("repo Save1m : %v", Btc1mCandle)

	// Btc1mCandle := Candle
	if err := r.Conn.Save(&Btc1mCandle).Error; err != nil {
		log.Printf("repo Save1m errrr")

		return false, err
	}
	log.Printf("repo Save1m true")
	return true, nil
}

func (r *CandleRepositoryImpl) Save1h(Btc1hCandle *domain.Btc1hCandle) (isSaved bool, err error) {
	log.Printf("repo Save1h : %v", Btc1hCandle)

	// Btc1hCandle := Candle
	if err := r.Conn.Save(&Btc1hCandle).Error; err != nil {
		log.Printf("repo Save1h errrr")

		return false, err
	}
	log.Printf("repo Save1h true")
	return true, nil
}
