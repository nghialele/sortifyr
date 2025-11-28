// nolint:gosec // md5 is not used for non-cryptographic purposes
// The only thing we need from it is speed and to change reliably
package api

import (
	"crypto/md5"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/shareed2k/goth_fiber"
)

const (
	mimeWEBP = "image/webp"
)

func storeInSession(ctx *fiber.Ctx, key string, value any) error {
	session, err := goth_fiber.SessionStore.Get(ctx)
	if err != nil {
		return fmt.Errorf("failed to get session %w", err)
	}

	session.Set(key, value)

	return session.Save()
}

func sendCached(c *fiber.Ctx, img []byte) error {
	etag := fmt.Sprintf(`"%x"`, md5.Sum(img))

	if match := c.Get("If-None-Match"); match != "" {
		if match == etag {
			c.Status(fiber.StatusNotModified)
			return nil
		}
	}

	c.Set("Cache-Control", "public, max-age=3600, stale-while-revalidate=120")
	c.Set("ETag", etag)

	return c.Send(img)
}
