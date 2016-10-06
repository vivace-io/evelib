package zkill

import "strconv"

/* zkill/util.go
 * Defines utility functions for the library
 */

func queryFromOpts(opts *Options) string {
	var result string
	// TODO - time options
	if opts == nil {
		return result
	}
	if opts.BeforeKillID != 0 {
		result += "/beforeKillID/" + strconv.Itoa(opts.BeforeKillID)
	}
	if opts.AfterKillID != 0 {
		result += "/afterKillID/" + strconv.Itoa(opts.AfterKillID)
	}
	if opts.Solo {
		result += "/solo"
	}
	if opts.Kills {
		result += "/kills"
	}
	if opts.Losses {
		result += "/losses"
	}
	if opts.WSpace {
		result += "/w-space"
	}

	// lastly, ensure trailing slash in place
	if result != "" {
		result += "/"
	}
	return result
}
