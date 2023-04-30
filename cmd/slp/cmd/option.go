package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tkuchiki/slp/helper"
	"github.com/tkuchiki/slp/options"
	"github.com/tkuchiki/slp/stats"
)

func createOptions(rootCmd *cobra.Command, sortOptions *stats.SortOptions) (*options.Options, error) {
	config, err := rootCmd.PersistentFlags().GetString("config")
	if err != nil {
		return nil, err
	}
	changedOptions := make([]options.Option, 0)
	unchangedOptions := make([]options.Option, 0)

	file, err := rootCmd.PersistentFlags().GetString("file")
	if err != nil {
		return nil, err
	}
	if rootCmd.PersistentFlags().Changed("file") {
		changedOptions = append(changedOptions, options.File(file))
	} else {
		unchangedOptions = append(unchangedOptions, options.File(file))
	}

	format, err := rootCmd.PersistentFlags().GetString("format")
	if err != nil {
		return nil, err
	}
	if rootCmd.PersistentFlags().Changed("format") {
		changedOptions = append(changedOptions, options.Format(format))
	} else {
		unchangedOptions = append(unchangedOptions, options.Format(format))
	}

	sort, err := rootCmd.PersistentFlags().GetString("sort")
	if err != nil {
		return nil, err
	}
	err = sortOptions.SetAndValidate(sort)
	if err != nil {
		return nil, err
	}
	if rootCmd.PersistentFlags().Changed("sort") {
		changedOptions = append(changedOptions, options.Sort(sort))
	} else {
		unchangedOptions = append(unchangedOptions, options.Sort(sort))
	}

	reverse, err := rootCmd.PersistentFlags().GetBool("reverse")
	if err != nil {
		return nil, err
	}
	if rootCmd.PersistentFlags().Changed("reverse") {
		changedOptions = append(changedOptions, options.Reverse(reverse))
	} else {
		unchangedOptions = append(unchangedOptions, options.Reverse(reverse))
	}

	noHeaders, err := rootCmd.PersistentFlags().GetBool("noheaders")
	if err != nil {
		return nil, err
	}
	if rootCmd.PersistentFlags().Changed("noheaders") {
		changedOptions = append(changedOptions, options.NoHeaders(noHeaders))
	} else {
		unchangedOptions = append(unchangedOptions, options.NoHeaders(noHeaders))
	}

	showFooters, err := rootCmd.PersistentFlags().GetBool("show-footers")
	if err != nil {
		return nil, err
	}
	if rootCmd.PersistentFlags().Changed("show-footers") {
		changedOptions = append(changedOptions, options.ShowFooters(showFooters))
	} else {
		unchangedOptions = append(unchangedOptions, options.ShowFooters(showFooters))
	}

	limit, err := rootCmd.PersistentFlags().GetInt("limit")
	if err != nil {
		return nil, err
	}
	if rootCmd.PersistentFlags().Changed("limit") {
		changedOptions = append(changedOptions, options.Limit(limit))
	} else {
		unchangedOptions = append(unchangedOptions, options.Limit(limit))
	}

	output, err := rootCmd.PersistentFlags().GetString("output")
	if err != nil {
		return nil, err
	}
	if rootCmd.PersistentFlags().Changed("output") {
		changedOptions = append(changedOptions, options.Output(output))
	} else {
		unchangedOptions = append(unchangedOptions, options.Output(output))
	}

	matchingGroups, err := rootCmd.PersistentFlags().GetString("matching-groups")
	if err != nil {
		return nil, err
	}
	if rootCmd.PersistentFlags().Changed("matching-groups") {
		changedOptions = append(changedOptions, options.CSVGroups(matchingGroups))
	} else {
		unchangedOptions = append(unchangedOptions, options.CSVGroups(matchingGroups))
	}

	filters, err := rootCmd.PersistentFlags().GetString("filters")
	if err != nil {
		return nil, err
	}
	if rootCmd.PersistentFlags().Changed("filters") {
		changedOptions = append(changedOptions, options.Filters(filters))
	} else {
		unchangedOptions = append(unchangedOptions, options.Filters(filters))
	}

	pos, err := rootCmd.PersistentFlags().GetString("pos")
	if err != nil {
		return nil, err
	}
	if rootCmd.PersistentFlags().Changed("pos") {
		changedOptions = append(changedOptions, options.PosFile(pos))
	} else {
		unchangedOptions = append(unchangedOptions, options.PosFile(pos))
	}

	noSavePos, err := rootCmd.PersistentFlags().GetBool("nosave-pos")
	if err != nil {
		return nil, err
	}
	if rootCmd.PersistentFlags().Changed("nosave-pos") {
		changedOptions = append(changedOptions, options.NoSavePos(noSavePos))
	} else {
		unchangedOptions = append(unchangedOptions, options.NoSavePos(noSavePos))
	}

	ps, err := rootCmd.PersistentFlags().GetString("percentiles")
	if err != nil {
		return nil, err
	}

	var percentiles []int
	if ps != "" {
		percentiles, err = helper.SplitCSVIntoInts(ps)
		if err != nil {
			return nil, err
		}

		if err = helper.ValidatePercentiles(percentiles); err != nil {
			return nil, err
		}
	}

	if rootCmd.PersistentFlags().Changed("percentiles") {
		changedOptions = append(changedOptions, options.Percentiles(percentiles))
	} else {
		unchangedOptions = append(unchangedOptions, options.Percentiles(percentiles))
	}

	bundleWhereIn, err := rootCmd.PersistentFlags().GetBool("bundle-where-in")
	if err != nil {
		return nil, err
	}
	if rootCmd.PersistentFlags().Changed("bundle-where-in") {
		changedOptions = append(changedOptions, options.BundleWhereIn(bundleWhereIn))
	} else {
		unchangedOptions = append(unchangedOptions, options.BundleWhereIn(bundleWhereIn))
	}

	bundleValues, err := rootCmd.PersistentFlags().GetBool("bundle-values")
	if err != nil {
		return nil, err
	}
	if rootCmd.PersistentFlags().Changed("bundle-values") {
		changedOptions = append(changedOptions, options.BundleValues(bundleValues))
	} else {
		unchangedOptions = append(unchangedOptions, options.BundleValues(bundleValues))
	}

	noAbstract, err := rootCmd.PersistentFlags().GetBool("noabstract")
	if err != nil {
		return nil, err
	}
	if rootCmd.PersistentFlags().Changed("noabstract") {
		changedOptions = append(changedOptions, options.NoAbstract(noAbstract))
	} else {
		unchangedOptions = append(unchangedOptions, options.NoAbstract(noAbstract))
	}

	paginationLimit, err := rootCmd.PersistentFlags().GetInt("page")
	if err != nil {
		return nil, err
	}
	if rootCmd.PersistentFlags().Changed("page") {
		changedOptions = append(changedOptions, options.PaginationLimit(paginationLimit))
	} else {
		unchangedOptions = append(unchangedOptions, options.PaginationLimit(paginationLimit))
	}

	opts := options.NewOptions(unchangedOptions...)
	if config != "" {
		cf, err := os.Open(config)
		if err != nil {
			return nil, err
		}
		defer cf.Close()

		opts, err = options.LoadOptionsFromReader(opts, cf)
		if err != nil {
			return nil, err
		}

		err = sortOptions.SetAndValidate(opts.Sort)
		if err != nil {
			return nil, err
		}

		percentiles = opts.Percentiles
	}

	return options.SetOptions(opts, changedOptions...), nil
}
