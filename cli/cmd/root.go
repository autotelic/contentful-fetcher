package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/autotelic/contentful-fetcher/lib/contentful"
	"github.com/autotelic/contentful-fetcher/lib/filesystem"
)

// NewDefaultRootCommand wires the default dependencies for the NewRootCommand.
func NewDefaultRootCommand(name string) *cobra.Command {
	return NewRootCommand(
		name,
		filesystem.NewDefaultOpener(),
		contentful.NewManifestParser(),
		contentful.NewDefaultQueryExecutor(),
	)
}

// NewRootCommand constructs the root command and its nested children.
func NewRootCommand(
	name string,
	fileOpener filesystem.Opener,
	manifestParser contentful.ManifestParser,
	queryExecutor contentful.QueryExecutor,
) *cobra.Command {
	var manifestPaths []string

	cmd := &cobra.Command{
		Use:   name,
		Short: "Interacts with the Contentful API",
		RunE: func(cmd *cobra.Command, args []string) error {
			viper.SetEnvPrefix("contentful")
			viper.BindEnv("access_token")
			viper.BindEnv("api_url")

			if !viper.IsSet("access_token") {
				return fmt.Errorf("missing contentful access_token")
			}
			accessToken := viper.Get("access_token").(string)

			if !viper.IsSet("api_url") {
				return fmt.Errorf("missing contentful api_url")
			}
			apiURL := viper.Get("api_url").(string)

			for _, manifestPath := range manifestPaths {
				f, err := fileOpener.Open(manifestPath)
				if err != nil {
					return err
				}
				defer f.Close()

				manifest, err := manifestParser.Parse(f)
				if err != nil {
					return err
				}

				for _, q := range manifest.Queries {
					err := queryExecutor.Execute(&q, apiURL, accessToken)
					if err != nil {
						return err
					}
				}
			}
			return nil
		},
	}

	cmd.Flags().StringArrayVarP(
		&manifestPaths,
		"manifest",
		"m",
		[]string{},
		"a manifest file",
	)

	return cmd
}
