//go:build ghcr
// +build ghcr

package image

import (
	"context"
	"testing"

	"github.com/Masterminds/semver"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"github.com/akuity/kargo/internal/logging"
)

// All test cases in this file are integration tests that rely on ghcr.io. You
// may get rate-limited executing these tests, so they're disabled by default.

func TestSelectTagGHCR(t *testing.T) {
	const kargoRepo = "ghcr.io/akuity/kargo"
	const platform = "linux/amd64"

	ctx := context.Background()
	logger := logging.LoggerFromContext(ctx)
	logger.Logger.SetLevel(log.TraceLevel)
	ctx = logging.ContextWithLogger(ctx, logger)

	t.Run("digest strategy", func(t *testing.T) {
		const constraint = "v0.1.0"
		s, err := NewTagSelector(
			kargoRepo,
			TagSelectionStrategyDigest,
			&TagSelectorOptions{
				Constraint: constraint,
			},
		)
		require.NoError(t, err)

		tag, err := s.SelectTag(ctx)
		require.NoError(t, err)

		require.NotNil(t, tag)
		require.Equal(t, constraint, tag.Name)
		require.NotEmpty(t, tag.Digest)
		require.NotNil(t, tag.CreatedAt)
	})

	t.Run("digest strategy with platform constraint", func(t *testing.T) {
		const constraint = "v0.1.0"
		s, err := NewTagSelector(
			kargoRepo,
			TagSelectionStrategyDigest,
			&TagSelectorOptions{
				Constraint: constraint,
				Platform:   platform,
			},
		)
		require.NoError(t, err)

		tag, err := s.SelectTag(ctx)
		require.NoError(t, err)

		require.NotNil(t, tag)
		require.Equal(t, constraint, tag.Name)
		require.NotEmpty(t, tag.Digest)
		require.NotNil(t, tag.CreatedAt)
	})

	t.Run("lexical strategy", func(t *testing.T) {
		s, err := NewTagSelector(
			kargoRepo,
			TagSelectionStrategyLexical,
			nil,
		)
		require.NoError(t, err)

		tag, err := s.SelectTag(ctx)
		require.NoError(t, err)

		require.NotNil(t, tag)
		require.NotEmpty(t, tag.Digest)
		require.NotNil(t, tag.CreatedAt)
	})

	t.Run("lexical strategy with platform constraint", func(t *testing.T) {
		s, err := NewTagSelector(
			kargoRepo,
			TagSelectionStrategyLexical,
			&TagSelectorOptions{
				Platform: platform,
			},
		)
		require.NoError(t, err)

		tag, err := s.SelectTag(ctx)
		require.NoError(t, err)

		require.NotNil(t, tag)
		require.NotEmpty(t, tag.Digest)
		require.NotNil(t, tag.CreatedAt)
	})

	t.Run("newest build strategy", func(t *testing.T) {
		s, err := NewTagSelector(
			kargoRepo,
			TagSelectionStrategyNewestBuild,
			&TagSelectorOptions{
				AllowRegex: `^v0.1.0-rc.2\d$`,
			},
		)
		require.NoError(t, err)

		tag, err := s.SelectTag(ctx)
		require.NoError(t, err)

		require.NotNil(t, tag)
		require.Equal(t, "v0.1.0-rc.24", tag.Name)
		require.NotEmpty(t, tag.Digest)
		require.NotNil(t, tag.CreatedAt)
	})

	t.Run("newest build strategy with platform constraint", func(t *testing.T) {
		s, err := NewTagSelector(
			kargoRepo,
			TagSelectionStrategyNewestBuild,
			&TagSelectorOptions{
				AllowRegex: `^v0.1.0-rc.2\d$`,
				Platform:   platform,
			},
		)
		require.NoError(t, err)

		tag, err := s.SelectTag(ctx)
		require.NoError(t, err)

		require.NotNil(t, tag)
		require.Equal(t, "v0.1.0-rc.24", tag.Name)
		require.NotEmpty(t, tag.Digest)
		require.NotNil(t, tag.CreatedAt)
	})

	t.Run("semver strategy", func(t *testing.T) {
		s, err := NewTagSelector(
			kargoRepo,
			TagSelectionStrategySemVer,
			nil,
		)
		require.NoError(t, err)

		tag, err := s.SelectTag(ctx)
		require.NoError(t, err)

		require.NotNil(t, tag)
		semVer, err := semver.NewVersion(tag.Name)
		require.NoError(t, err)
		min := semver.MustParse("0.1.0")
		require.True(t, semVer.GreaterThan(min) || semVer.Equal(min))
		require.True(t, semVer.LessThan(semver.MustParse("1.0.0")))
		require.NotEmpty(t, tag.Digest)
		require.NotNil(t, tag.CreatedAt)
	})

	t.Run("semver strategy with platform constraint", func(t *testing.T) {
		s, err := NewTagSelector(
			kargoRepo,
			TagSelectionStrategySemVer,
			&TagSelectorOptions{
				Platform: platform,
			},
		)
		require.NoError(t, err)

		tag, err := s.SelectTag(ctx)
		require.NoError(t, err)

		require.NotNil(t, tag)
		semVer, err := semver.NewVersion(tag.Name)
		require.NoError(t, err)
		min := semver.MustParse("0.1.0")
		require.True(t, semVer.GreaterThan(min) || semVer.Equal(min))
		require.True(t, semVer.LessThan(semver.MustParse("1.0.0")))
		require.NotEmpty(t, tag.Digest)
		require.NotNil(t, tag.CreatedAt)
	})

	t.Run("tolerance for non-image references", func(t *testing.T) {
		// Image lists or indices may contain non-image references. These could be
		// to things like attestations. This test verifies that we ignore such
		// references to avoid parsing errors.
		const tagName = "unknown"
		s, err := NewTagSelector(
			"ghcr.io/akuity/kargo-test",
			TagSelectionStrategyDigest,
			&TagSelectorOptions{
				Constraint: tagName,
			},
		)
		require.NoError(t, err)

		tag, err := s.SelectTag(ctx)
		require.NoError(t, err)

		require.NotNil(t, tag)
		require.Equal(t, tagName, tag.Name)
		require.NotEmpty(t, tag.Digest)
		require.NotNil(t, tag.CreatedAt)
	})

	t.Run("nothing found", func(t *testing.T) {
		s, err := NewTagSelector(
			kargoRepo,
			TagSelectionStrategyDigest,
			&TagSelectorOptions{
				Constraint: "v0.1.0",
				// Nothing will match this
				Platform: "linux/made-up-arch",
			},
		)
		require.NoError(t, err)

		tag, err := s.SelectTag(ctx)
		require.NoError(t, err)
		require.Nil(t, tag)
	})
}
