package main

import (
	"sync"

	"github.com/HUSTSecLab/criticality_score/pkg/collector/alpine"
	"github.com/HUSTSecLab/criticality_score/pkg/collector/archlinux"
	"github.com/HUSTSecLab/criticality_score/pkg/collector/aur"
	"github.com/HUSTSecLab/criticality_score/pkg/collector/centos"
	"github.com/HUSTSecLab/criticality_score/pkg/collector/debian"
	"github.com/HUSTSecLab/criticality_score/pkg/collector/deepin"
	"github.com/HUSTSecLab/criticality_score/pkg/collector/fedora"
	"github.com/HUSTSecLab/criticality_score/pkg/collector/gentoo"
	"github.com/HUSTSecLab/criticality_score/pkg/collector/homebrew"
	"github.com/HUSTSecLab/criticality_score/pkg/collector/nix"
	"github.com/HUSTSecLab/criticality_score/pkg/collector/ubuntu"
	"github.com/HUSTSecLab/criticality_score/pkg/config"
	"github.com/spf13/pflag"
)

var (
	flagType    = pflag.String("type", "", "type of the distribution")
	flagGenDot  = pflag.String("gendot", "", "output dot file")
	workerCount = pflag.Int("worker", 1, "number of workers")
	batchSize   = pflag.Int("batch", 1000, "batch size")
	downloadDir = pflag.String("downloadDir", "./download", "download directory")
)

func main() {
	config.RegistCommonFlags(pflag.CommandLine)
	config.ParseFlags(pflag.CommandLine)

	if *flagType == "" {
		var wg sync.WaitGroup
		wg.Add(11)

		go func() {
			defer wg.Done()
			archlinux.NewArchLinuxCollector().Collect(*flagGenDot)
		}()
		go func() {
			defer wg.Done()
			debian.NewDebianCollector().Collect(*flagGenDot)
		}()
		go func() {
			defer wg.Done()
			deepin.NewDeepinCollector().Collect(*flagGenDot)
		}()
		go func() {
			defer wg.Done()
			ubuntu.NewUbuntuCollector().Collect(*flagGenDot)
		}()
		// go func() {
		// 	defer wg.Done()
		// 	nix.NewNixCollector().Collect(*workerCount, *batchSize, *flagGenDot)
		// }()
		go func() {
			defer wg.Done()
			homebrew.NewHomebrewCollector().Collect(*flagGenDot, *downloadDir)
		}()
		// go func() {
		// 	defer wg.Done()
		// 	gentoo.NewGentooCollector().Collect(*flagGenDot)
		// }()
		go func() {
			defer wg.Done()
			fedora.NewFedoraCollector().Collect(*flagGenDot)
		}()
		go func() {
			defer wg.Done()
			centos.NewCentosCollector().Collect(*flagGenDot)
		}()
		go func() {
			defer wg.Done()
			alpine.NewAlpineCollector().Collect(*flagGenDot)
		}()
		go func() {
			defer wg.Done()
			aur.NewAurCollector().Collect(*flagGenDot)
		}()

		wg.Wait()
	} else {
		switch *flagType {
		case "archlinux":
			archlinux.NewArchLinuxCollector().Collect(*flagGenDot)
		case "debian":
			debian.NewDebianCollector().Collect(*flagGenDot)
		case "deepin":
			deepin.NewDeepinCollector().Collect(*flagGenDot)
		case "ubuntu":
			ubuntu.NewUbuntuCollector().Collect(*flagGenDot)
		case "nix":
			nix.NewNixCollector().Collect(*workerCount, *batchSize, *flagGenDot)
		case "homebrew":
			homebrew.NewHomebrewCollector().Collect(*flagGenDot, *downloadDir)
		case "gentoo":
			gentoo.NewGentooCollector().Collect(*flagGenDot, *downloadDir)
		case "fedora":
			fedora.NewFedoraCollector().Collect(*flagGenDot)
		case "centos":
			centos.NewCentosCollector().Collect(*flagGenDot)
		case "alpine":
			alpine.NewAlpineCollector().Collect(*flagGenDot)
		case "aur":
			aur.NewAurCollector().Collect(*flagGenDot)
		}
	}
}
