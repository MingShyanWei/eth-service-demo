package cache

import (
	model "eth-service-demo/models"

	"log"
	"time"
)

var Blocks []model.Block

func UpdateBlocksCashe() {
	// Get all blocks
	var block model.Block
	var err error
	Blocks, err = block.ListBlocks(-1)
	if err != nil {
		log.Println(err)
	}
	log.Println(Blocks[0:10])

	for {
		max, min, _ := block.GetMaxMinBlockId()
		log.Println(max, min)

		log.Println(max, Blocks[0].Num, len(Blocks))

		if max > Blocks[0].Num {
			log.Printf("Update blocks...")

			newBlocks, err := block.ListBetweenBlocks(Blocks[0].Num, max)
			if err != nil {
				log.Println(err)
			}
			log.Println(newBlocks)

			Blocks = append(newBlocks, Blocks...)
		}

		time.Sleep(10 * time.Second)
	}
}
