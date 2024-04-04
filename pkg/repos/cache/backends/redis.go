package backends

import (
	"context"
	"encoding/json"
	"grid/pkg/db/redis"
	"grid/pkg/env"
	"grid/pkg/models"
	"grid/pkg/repos/cache/provider"
	"grid/pkg/utils"
)

type Redis struct {
	*redis.Client
}

const (
	//TODO: move this to the env
	AircraftLocationsKey = "aircrafts-locations"
)

func (backend Redis) Initialize(vars *env.Vars) (provider.Provider, error) {
	args := *vars
	args.RedisDB = vars.RedisDBCache

	concrete := Redis{}
	client, err := redis.NewWithEnv(utils.Ref(args))

	if err != nil {
		return concrete, err
	}

	concrete.Client = client

	return concrete, nil
}

func (backend Redis) GetAircraftLocation(key string) (*models.LocationEvent, error) {
	cmd := backend.HGet(context.Background(), AircraftLocationsKey, key)
	err := cmd.Err()

	if err != nil {
		return nil, err
	}

	result := &models.LocationEvent{}
	value := cmd.Val()

	if err := json.Unmarshal([]byte(value), result); err != nil {
		return nil, err
	}

	return result, nil
}

func (backend Redis) GetAircraftsLocations() (map[string]*models.LocationEvent, error) {
	cmd := backend.HGetAll(context.Background(), AircraftLocationsKey)
	err := cmd.Err()

	if err != nil {
		return nil, err
	}

	result := map[string]*models.LocationEvent{}
	value := cmd.Val()

	// TODO: rethink this...
	for k, v := range value {
		event := &models.LocationEvent{}

		if err := json.Unmarshal([]byte(v), event); err != nil {
			return nil, err
		}

		result[k] = event
	}

	return result, nil
}

func (backend Redis) UpsertAircraftLocation(key string, v *models.LocationEvent) error {
	marshalled, err := json.Marshal(v)

	if err != nil {
		return err
	}

	cmd := backend.HSet(context.Background(), AircraftLocationsKey, key, marshalled)
	return cmd.Err()
}
