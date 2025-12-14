package spotify

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/topvennie/sortifyr/pkg/concurrent"
	"github.com/topvennie/sortifyr/pkg/storage"
	"github.com/topvennie/sortifyr/pkg/utils"
)

func uriToID(uri string) string {
	parts := strings.Split(uri, ":")
	if len(parts) != 3 {
		return ""
	}

	return parts[2]
}

type syncUserDataStruct[T any] struct {
	DB             []T                 // All the items the user has in our db
	API            []T                 // All the items the user has according to spotify
	Equal          func(T, T) bool     // Test 2 items for equality
	Get            func(T) (*T, error) // Get a db item
	Create         func(*T) error      // Create a new db item
	CreateUserLink func(T) error       // Link a database item to the user
	DeleteUserLink func(T) error       // Delete an item
}

func syncUserData[T any](s syncUserDataStruct[T]) error {
	// Go over every entry in the api list
	// to find missing entries
	for i := range s.API {
		if _, ok := utils.SliceFind(s.DB, func(t T) bool { return s.Equal(t, s.API[i]) }); ok {
			// The user already has this entry
			continue
		}

		// The user doesn not have this item yet
		// Do we have it in our database?
		t, err := s.Get(s.API[i])
		if err != nil {
			return err
		}
		if t == nil {
			// We don't have it in our database yet
			// Let's create it so we get an id
			t = &s.API[i]
			if err := s.Create(t); err != nil {
				return err
			}

			// Add the item to the db list for the deletion loop
			s.DB = append(s.DB, *t)
		}

		// We now have an id for the db item
		// Let's link it to the user
		if err := s.CreateUserLink(*t); err != nil {
			return err
		}
	}

	// Do the same but let's look for
	// items that need to be deleted
	for i := range s.DB {
		if _, ok := utils.SliceFind(s.API, func(t T) bool { return s.Equal(t, s.DB[i]) }); !ok {
			// Item is in our db but not in the spotify api list
			// So we have to delete it
			if err := s.DeleteUserLink(s.DB[i]); err != nil {
				return err
			}
		}
	}

	return nil
}

type syncCoverStruct struct {
	CoverURL string
	CoverID  string
	Update   func(string) error
}

func (c *client) syncCover(ctx context.Context, s []syncCoverStruct) error {
	wg := concurrent.NewLimitedWaitGroup(12)

	var mu sync.Mutex
	var errs []error

	for _, item := range s {
		if item.CoverURL == "" {
			continue
		}

		wg.Go(func() {
			cover, err := c.api.ImageGet(ctx, item.CoverURL)
			if err != nil {
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
				return
			}
			if len(cover) == 0 {
				return
			}

			oldCover := []byte{}
			if item.CoverID != "" {
				oldCover, err = storage.S.Get(item.CoverID)
				if err != nil {
					mu.Lock()
					errs = append(errs, fmt.Errorf("get cover for %+v | %w", item, err))
					mu.Unlock()
					return
				}
			}

			if bytes.Equal(cover, oldCover) {
				return
			}

			newID := uuid.NewString()
			if err := storage.S.Set(newID, cover, 0); err != nil {
				mu.Lock()
				errs = append(errs, fmt.Errorf("store new cover %+v | %w", item, err))
				mu.Unlock()
				return
			}

			if err := item.Update(newID); err != nil {
				_ = storage.S.Delete(newID) // nolint:errcheck // Too bad if it fails
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
				return
			}

			storage.S.Delete(item.CoverID) // nolint:errcheck // Too bad if it fails
		})
	}

	wg.Wait()

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
