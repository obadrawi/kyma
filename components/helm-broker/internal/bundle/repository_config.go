package bundle

import (
	"path"
	"path/filepath"
	"strings"
)

// RepositoryConfig provides configuration for HTTP Repository.
// Deprecated: Use the implementation defined on ClusterAddonsConfiguration CRD
type RepositoryConfig struct {
	URL string `json:"URL" valid:"required"`
}

// IndexFileName returns name of yaml file with configuration (if not exist return default name)
// Deprecated: Use the implementation defined on ClusterAddonsConfiguration CRD
func (cfg RepositoryConfig) IndexFileName() string {
	if cfg.hasConfigFile() {
		return path.Base(cfg.URL)
	}

	return "index.yaml"
}

// BaseURL returns base url to bundles with trailing slash
// Deprecated: Use the implementation defined on ClusterAddonsConfiguration CRD
func (cfg RepositoryConfig) BaseURL() string {
	if cfg.hasConfigFile() {
		return strings.TrimRight(cfg.URL, cfg.IndexFileName())
	}

	return strings.TrimRight(cfg.URL, "/") + "/"
}

func (cfg RepositoryConfig) hasConfigFile() bool {
	extension := filepath.Ext(path.Base(cfg.URL))

	return extension == ".yaml"
}
