package cmd

import (
	"context"

	"brave_signer/config"

	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	return &cobra.Command{
		Use:          "brave_signer",
		Short:        "Bravely generate key pairs, sign files, and check signatures.",
		Long:         `A collection of tools to generate key pairs in PEM files, sign files, and verify signatures.`,
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initializeConfig(cmd)
		},
	}
}

func initializeConfig(cmd *cobra.Command) error {
	localViper, err := config.LoadYamlConfig()
	if err != nil {
		return err
	}

	if err := config.BindFlags(cmd, localViper); err != nil {
		return err
	}

	ctx := context.WithValue(cmd.Context(), config.ViperKey, localViper)
	cmd.SetContext(ctx)
	return nil
}
