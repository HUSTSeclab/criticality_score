// This file is used to clone the git repo to memory, and
// print its metadata to the console.
package main

import (
	"flag"
	"os"
	"strings"
	"sync"

	collector "github.com/HUSTSecLab/criticality_score/pkg/gitfile/collector"
	git "github.com/HUSTSecLab/criticality_score/pkg/gitfile/parser/git"
	url "github.com/HUSTSecLab/criticality_score/pkg/gitfile/parser/url"
	"github.com/HUSTSecLab/criticality_score/pkg/logger"
	"github.com/bytedance/gopkg/util/gopool"
	gogit "github.com/go-git/go-git/v5"
)

func main() {
	logger.ConfigAsCommandLineTool()
	flag.Usage = func() {
		logger.Printf("This tool is used to collect git metadata in storage path, but not clone the repository.\n")
		logger.Printf("Usage: %s [url1] [url2] ...\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	inputs := []string{}
	for i := 0; i < flag.NArg(); i++ {
		inputs = append(inputs, flag.Arg(i))
	}

	var wg sync.WaitGroup
	wg.Add(len(inputs))

	repos := make([]*git.Repo, 0)

	for index, input := range inputs {
		gopool.Go(func() {
			defer wg.Done()
			logger.Infof("[%d] Collecting %s", index, input)

			r := &gogit.Repository{}
			var err error

			//* if the input is url, parse and clone the repo
			//* if not, open the repo
			if strings.Contains(input, "://") {
				u := url.ParseURL(input)
				r, err = collector.EzCollect(&u)
				if err != nil {
					logger.Panicf("[%d] Collecting %s Failed", index, input)
				}
			} else {
				r, err = collector.Open(input)
				if err != nil {
					logger.Panicf("[%d] Opening %s Failed", index, input)
				}
			}

			repo, err := git.ParseRepo(r)
			if err != nil {
				logger.Panicf("[%d] Parsing %s Failed", index, input)
			}

			repos = append(repos, repo)
			logger.Infof("[%d] %s Collected", index, repo.Name)
		})
	}

	wg.Wait()
	for _, repo := range repos {
		repo.Show()
	}
}
