package tdsd

import (
	"os"
	"path/filepath"
	"strings"
)

type KubeConfig struct {
	Path string
	Name string
}

type Finder struct {
	Configs map[string]KubeConfig

	homeDir       string
	supportedDirs []string
}

func NewFinder() *Finder {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	finder := &Finder{
		homeDir: homeDir,
		Configs: map[string]KubeConfig{},
		supportedDirs: []string{
			homeDir + "/.kube",
		},
	}

	finder.scan()

	return finder
}

func (r *Finder) GetConfig(name string) KubeConfig {
	return r.Configs[name]
}

func (r *Finder) scan() {
	for _, dir := range r.supportedDirs {
		r.readDir(dir)
	}
}

func (r *Finder) readDir(dir string) {
	kubeDir, err := os.Open(dir)

	if err != nil {
		panic(err)
	}

	kubeFiles, err := kubeDir.Readdir(0)

	if err != nil {
		panic(err)
	}

	defer kubeDir.Close()

	for _, file := range kubeFiles {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) == ".yaml" {
			name := strings.TrimSuffix(filepath.Base(file.Name()), filepath.Ext(file.Name()))

			r.Configs[name] = KubeConfig{
				Path: dir + "/" + file.Name(),
				Name: file.Name(),
			}
		}
	}
}
