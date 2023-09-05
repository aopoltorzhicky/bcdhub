package indexer

import (
	"context"
	"sync"

	"github.com/baking-bad/bcdhub/internal/bcd/tezerrors"
	"github.com/baking-bad/bcdhub/internal/config"
	"github.com/dipdup-io/workerpool"
	"github.com/rs/zerolog/log"
)

// CreateIndexers -
func CreateIndexers(ctx context.Context, cfg config.Config, g workerpool.Group) ([]Indexer, error) {
	if err := tezerrors.LoadErrorDescriptions(); err != nil {
		return nil, err
	}

	var (
		mx       sync.Mutex
		wg       sync.WaitGroup
		indexers = make([]Indexer, 0)
	)

	for network, indexerCfg := range cfg.Indexer.Networks {
		wg.Add(1)
		go func(network string, indexerCfg config.IndexerConfig) {
			defer wg.Done()

			if indexerCfg.Periodic != nil {
				periodicIndexer, err := NewPeriodicIndexer(ctx, network, cfg, indexerCfg, g)
				if err != nil {
					log.Err(err).Msg("NewPeriodicIndexer")
					return
				}
				mx.Lock()
				indexers = append(indexers, periodicIndexer)
				mx.Unlock()
			} else {
				bi, err := NewBlockchainIndexer(ctx, cfg, network, indexerCfg)
				if err != nil {
					log.Err(err).Msg("NewBlockchainIndexer")
					return
				}
				mx.Lock()
				indexers = append(indexers, bi)
				mx.Unlock()
			}
		}(network, indexerCfg)
	}

	wg.Wait()

	return indexers, nil
}
