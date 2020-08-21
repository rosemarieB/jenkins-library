// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type abapAddonAssemblyKitRegisterPackagesOptions struct {
	AbapAddonAssemblyKitEndpoint string `json:"AbapAddonAssemblyKitEndpoint,omitempty"`
	Username                     string `json:"username,omitempty"`
	Password                     string `json:"password,omitempty"`
	AddonDescriptor              string `json:"addonDescriptor,omitempty"`
}

// AbapAddonAssemblyKitRegisterPackagesCommand TODO
func AbapAddonAssemblyKitRegisterPackagesCommand() *cobra.Command {
	const STEP_NAME = "abapAddonAssemblyKitRegisterPackages"

	metadata := abapAddonAssemblyKitRegisterPackagesMetadata()
	var stepConfig abapAddonAssemblyKitRegisterPackagesOptions
	var startTime time.Time

	var createAbapAddonAssemblyKitRegisterPackagesCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "TODO",
		Long:  `TODO`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			startTime = time.Now()
			log.SetStepName(STEP_NAME)
			log.SetVerbose(GeneralConfig.Verbose)

			path, _ := os.Getwd()
			fatalHook := &log.FatalHook{CorrelationID: GeneralConfig.CorrelationID, Path: path}
			log.RegisterHook(fatalHook)

			err := PrepareConfig(cmd, &metadata, STEP_NAME, &stepConfig, config.OpenPiperFile)
			if err != nil {
				log.SetErrorCategory(log.ErrorConfiguration)
				return err
			}
			log.RegisterSecret(stepConfig.Username)
			log.RegisterSecret(stepConfig.Password)

			if len(GeneralConfig.HookConfig.SentryConfig.Dsn) > 0 {
				sentryHook := log.NewSentryHook(GeneralConfig.HookConfig.SentryConfig.Dsn, GeneralConfig.CorrelationID)
				log.RegisterHook(&sentryHook)
			}

			return nil
		},
		Run: func(_ *cobra.Command, _ []string) {
			telemetryData := telemetry.CustomData{}
			telemetryData.ErrorCode = "1"
			handler := func() {
				telemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				telemetry.Send(&telemetryData)
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetry.Initialize(GeneralConfig.NoTelemetry, STEP_NAME)
			abapAddonAssemblyKitRegisterPackages(stepConfig, &telemetryData)
			telemetryData.ErrorCode = "0"
			log.Entry().Info("SUCCESS")
		},
	}

	addAbapAddonAssemblyKitRegisterPackagesFlags(createAbapAddonAssemblyKitRegisterPackagesCmd, &stepConfig)
	return createAbapAddonAssemblyKitRegisterPackagesCmd
}

func addAbapAddonAssemblyKitRegisterPackagesFlags(cmd *cobra.Command, stepConfig *abapAddonAssemblyKitRegisterPackagesOptions) {
	cmd.Flags().StringVar(&stepConfig.AbapAddonAssemblyKitEndpoint, "AbapAddonAssemblyKitEndpoint", `https://w7q.dmzwdf.sap.corp`, "TODO")
	cmd.Flags().StringVar(&stepConfig.Username, "username", os.Getenv("PIPER_username"), "User")
	cmd.Flags().StringVar(&stepConfig.Password, "password", os.Getenv("PIPER_password"), "User Password")
	cmd.Flags().StringVar(&stepConfig.AddonDescriptor, "addonDescriptor", os.Getenv("PIPER_addonDescriptor"), "AddonDescriptor")

	cmd.MarkFlagRequired("AbapAddonAssemblyKitEndpoint")
	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("password")
	cmd.MarkFlagRequired("addonDescriptor")
}

// retrieve step metadata
func abapAddonAssemblyKitRegisterPackagesMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:    "abapAddonAssemblyKitRegisterPackages",
			Aliases: []config.Alias{},
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name:        "AbapAddonAssemblyKitEndpoint",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS", "GENERAL"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "username",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "password",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "addonDescriptor",
						ResourceRef: []config.ResourceReference{{Name: "commonPipelineEnvironment", Param: "abap/addonDescriptor"}},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
				},
			},
		},
	}
	return theMetaData
}
