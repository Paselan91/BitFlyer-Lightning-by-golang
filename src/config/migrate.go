package config

import (
	"app/src/domain"
	"github.com/jinzhu/gorm"
	"log"
	// "time"
)

// DBMigrate will create & migrate the tables, then make the some relationships if necessary
func DBMigrate() (*gorm.DB, error) {
	log.Println("Migration start")
	conn, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if conn.HasTable(domain.SignalEvent{}) {
		conn.DropTable(domain.SignalEvent{})
	}
	if conn.HasTable(domain.Btc1sCandle{}) {
		conn.DropTable(domain.Btc1sCandle{})
	}
	if conn.HasTable(domain.Btc1mCandle{}) {
		conn.DropTable(domain.Btc1mCandle{})
	}
	if conn.HasTable(domain.Btc1hCandle{}) {
		conn.DropTable(domain.Btc1hCandle{})
	}

	conn.AutoMigrate(domain.SignalEvent{})
	conn.AutoMigrate(domain.Btc1sCandle{})
	conn.AutoMigrate(domain.Btc1mCandle{})
	conn.AutoMigrate(domain.Btc1hCandle{})

	log.Println("Migration has been processed")
	return conn, nil
}

func Seeds() (*gorm.DB, error) {
	conn, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// testBtc1sCandle := &domain.Btc1sCandle{
	// 	BtcCandle: domain.BtcCandle{
	// 		Time:   time.Now(),
	// 		Open:   1.234,
	// 		Close:  1.234,
	// 		High:   1.234,
	// 		Low:    1.234,
	// 		Volume: 1.234,
	// 	},
	// }
	// testBtc1mCandle := &domain.Btc1mCandle{
	// 	BtcCandle: domain.BtcCandle{
	// 		Time:   time.Now(),
	// 		Open:   2.234,
	// 		Close:  2.234,
	// 		High:   2.234,
	// 		Low:    2.234,
	// 		Volume: 2.234,
	// 	},
	// }
	// testBtc1hCandle := &domain.Btc1hCandle{
	// 	BtcCandle: domain.BtcCandle{
	// 		Time:   time.Now(),
	// 		Open:   3.234,
	// 		Close:  3.234,
	// 		High:   3.234,
	// 		Low:    3.234,
	// 		Volume: 3.234,
	// 	},
	// }

	// if err := conn.Debug().Create(testBtc1sCandle).Create(testBtc1mCandle).Create(testBtc1hCandle).Error; err != nil {
	// 	return nil, err
	// }
	return nil, err
}
