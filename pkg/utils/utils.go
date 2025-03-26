// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package utils

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// Returns Harbor v2 client for given clientConfig

func PrintPayloadInJSONFormat(payload any) {
	if payload == nil {
		return
	}

	jsonStr, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonStr))
}

func PrintPayloadInYAMLFormat(payload any) {
	if payload == nil {
		return
	}

	yamlStr, err := yaml.Marshal(payload)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(yamlStr))
}

func ParseProjectRepo(projectRepo string) (string, string) {
	split := strings.SplitN(projectRepo, "/", 2) //splits only at first slash
	if len(split) != 2 {
		log.Fatalf("invalid project/repository format: %s", projectRepo)
	}
	return split[0], split[1]
}

func ParseProjectRepoReference(projectRepoReference string) (string, string, string) {
	log.Infof("Parsing input: %s", projectRepoReference) // Debug log

	// Split project and repo
	parts := strings.SplitN(projectRepoReference, "/", 2)
	if len(parts) != 2 {
		log.Fatalf("Invalid format, expected <project>/<repository>:<tag> or <project>/<repository>@<digest>, got: %s", projectRepoReference)
	}

	project := parts[0]
	repoWithRef := parts[1]

	var repo, ref string
	if strings.Contains(repoWithRef, ":") {
		subParts := strings.SplitN(repoWithRef, ":", 2)
		repo = subParts[0]
		ref = subParts[1]
	} else if strings.Contains(repoWithRef, "@") {
		subParts := strings.SplitN(repoWithRef, "@", 2)
		repo = subParts[0]
		ref = subParts[1]
	} else {
		log.Fatalf("Invalid reference format: %s", repoWithRef)
	}

	log.Infof("Extracted -> Project: %s, Repo: %s, Reference: %s", project, repo, ref)
	return project, repo, ref
}

func SanitizeServerAddress(server string) string {
	re := regexp.MustCompile(`^https?://`)
	server = re.ReplaceAllString(server, "")
	re = regexp.MustCompile(`[^a-zA-Z0-9]`)
	server = re.ReplaceAllString(server, "-")
	return server
}
