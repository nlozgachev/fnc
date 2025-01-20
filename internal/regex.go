package internal

import "regexp"

func ExtractBranchPrefix(branchName string) string {
	re := regexp.MustCompile(`([A-Z]+-\d+)_.*$`)
	matches := re.FindStringSubmatch(branchName)

	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}
