// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/piperenv"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type abapEnvironmentAssemblyOptions struct {
	CfAPIEndpoint          string `json:"cfApiEndpoint,omitempty"`
	CfOrg                  string `json:"cfOrg,omitempty"`
	CfSpace                string `json:"cfSpace,omitempty"`
	CfServiceInstance      string `json:"cfServiceInstance,omitempty"`
	CfServiceKeyName       string `json:"cfServiceKeyName,omitempty"`
	Host                   string `json:"host,omitempty"`
	Username               string `json:"username,omitempty"`
	Password               string `json:"password,omitempty"`
	PackageType            string `json:"PackageType,omitempty"`
	PackageName            string `json:"PackageName,omitempty"`
	SWC                    string `json:"SWC,omitempty"`
	SWCRelease             string `json:"SWCRelease,omitempty"`
	SpsLevel               string `json:"SpsLevel,omitempty"`
	Namespace              string `json:"Namespace,omitempty"`
	PreviousDeliveryCommit string `json:"PreviousDeliveryCommit,omitempty"`
	MaxRuntimeInMinutes    int    `json:"MaxRuntimeInMinutes,omitempty"`
}

type abapEnvironmentAssemblyCommonPipelineEnvironment struct {
	SAR_XML string
}

func (p *abapEnvironmentAssemblyCommonPipelineEnvironment) persist(path, resourceName string) {
	content := []struct {
		category string
		name     string
		value    string
	}{
		{category: "", name: "SAR_XML", value: p.SAR_XML},
	}

	errCount := 0
	for _, param := range content {
		err := piperenv.SetResourceParameter(path, resourceName, filepath.Join(param.category, param.name), param.value)
		if err != nil {
			log.Entry().WithError(err).Error("Error persisting piper environment.")
			errCount++
		}
	}
	if errCount > 0 {
		log.Entry().Fatal("failed to persist Piper environment")
	}
}

// AbapEnvironmentAssemblyCommand Assembly of installation, support package or patch in SAP Cloud Platform ABAP Environment system
func AbapEnvironmentAssemblyCommand() *cobra.Command {
	const STEP_NAME = "abapEnvironmentAssembly"

	metadata := abapEnvironmentAssemblyMetadata()
	var stepConfig abapEnvironmentAssemblyOptions
	var startTime time.Time
	var commonPipelineEnvironment abapEnvironmentAssemblyCommonPipelineEnvironment

	var createAbapEnvironmentAssemblyCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "Assembly of installation, support package or patch in SAP Cloud Platform ABAP Environment system",
		Long:  `Assembly of installation, support package or patch in SAP Cloud Platform ABAP Environment system`,
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
				commonPipelineEnvironment.persist(GeneralConfig.EnvRootPath, "commonPipelineEnvironment")
				telemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				telemetry.Send(&telemetryData)
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetry.Initialize(GeneralConfig.NoTelemetry, STEP_NAME)
			abapEnvironmentAssembly(stepConfig, &telemetryData, &commonPipelineEnvironment)
			telemetryData.ErrorCode = "0"
			log.Entry().Info("SUCCESS")
		},
	}

	addAbapEnvironmentAssemblyFlags(createAbapEnvironmentAssemblyCmd, &stepConfig)
	return createAbapEnvironmentAssemblyCmd
}

func addAbapEnvironmentAssemblyFlags(cmd *cobra.Command, stepConfig *abapEnvironmentAssemblyOptions) {
	cmd.Flags().StringVar(&stepConfig.CfAPIEndpoint, "cfApiEndpoint", os.Getenv("PIPER_cfApiEndpoint"), "Cloud Foundry API endpoint")
	cmd.Flags().StringVar(&stepConfig.CfOrg, "cfOrg", os.Getenv("PIPER_cfOrg"), "CF org")
	cmd.Flags().StringVar(&stepConfig.CfSpace, "cfSpace", os.Getenv("PIPER_cfSpace"), "CF Space")
	cmd.Flags().StringVar(&stepConfig.CfServiceInstance, "cfServiceInstance", os.Getenv("PIPER_cfServiceInstance"), "Parameter of ServiceInstance Name to delete CloudFoundry Service")
	cmd.Flags().StringVar(&stepConfig.CfServiceKeyName, "cfServiceKeyName", `SAP_COM_0582`, "Parameter of CloudFoundry Service Key to be created")
	cmd.Flags().StringVar(&stepConfig.Host, "host", os.Getenv("PIPER_host"), "Specifies the host address of the SAP Cloud Platform ABAP Environment system")
	cmd.Flags().StringVar(&stepConfig.Username, "username", os.Getenv("PIPER_username"), "User or E-Mail for CF")
	cmd.Flags().StringVar(&stepConfig.Password, "password", os.Getenv("PIPER_password"), "User Password for CF User")
	cmd.Flags().StringVar(&stepConfig.PackageType, "PackageType", os.Getenv("PIPER_PackageType"), "Type of the delivery package(AOI, CSP, CPK) as provided by AAKaaS")
	cmd.Flags().StringVar(&stepConfig.PackageName, "PackageName", os.Getenv("PIPER_PackageName"), "Name of delivery package as provided by AAKaaS")
	cmd.Flags().StringVar(&stepConfig.SWC, "SWC", os.Getenv("PIPER_SWC"), "Name of software component as provided by AAKaaS")
	cmd.Flags().StringVar(&stepConfig.SWCRelease, "SWCRelease", os.Getenv("PIPER_SWCRelease"), "Software component release as provided by AAKaaS")
	cmd.Flags().StringVar(&stepConfig.SpsLevel, "SpsLevel", os.Getenv("PIPER_SpsLevel"), "Support package level as provided by AAKaaS")
	cmd.Flags().StringVar(&stepConfig.Namespace, "Namespace", os.Getenv("PIPER_Namespace"), "Development namespace for software component")
	cmd.Flags().StringVar(&stepConfig.PreviousDeliveryCommit, "PreviousDeliveryCommit", os.Getenv("PIPER_PreviousDeliveryCommit"), "Commit ID for the previous delivery event")
	cmd.Flags().IntVar(&stepConfig.MaxRuntimeInMinutes, "MaxRuntimeInMinutes", 360, "maximal runtime of the step")

	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("password")
	cmd.MarkFlagRequired("PackageType")
	cmd.MarkFlagRequired("PackageName")
	cmd.MarkFlagRequired("SWC")
	cmd.MarkFlagRequired("SWCRelease")
	cmd.MarkFlagRequired("SpsLevel")
	cmd.MarkFlagRequired("Namespace")
	cmd.MarkFlagRequired("MaxRuntimeInMinutes")
}

// retrieve step metadata
func abapEnvironmentAssemblyMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:    "abapEnvironmentAssembly",
			Aliases: []config.Alias{},
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name:        "cfApiEndpoint",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS", "GENERAL"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "cloudFoundry/apiEndpoint"}},
					},
					{
						Name:        "cfOrg",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS", "GENERAL"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "cloudFoundry/org"}},
					},
					{
						Name:        "cfSpace",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS", "GENERAL"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "cloudFoundry/space"}},
					},
					{
						Name:        "cfServiceInstance",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS", "GENERAL"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "cloudFoundry/serviceInstance"}},
					},
					{
						Name:        "cfServiceKeyName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS", "GENERAL"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "cloudFoundry/serviceKeyName"}},
					},
					{
						Name:        "host",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
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
						Name:        "PackageType",
						ResourceRef: []config.ResourceReference{{Name: "commonPipelineEnvironment", Param: "PackageType"}},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "PackageName",
						ResourceRef: []config.ResourceReference{{Name: "commonPipelineEnvironment", Param: "PackageName"}},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "SWC",
						ResourceRef: []config.ResourceReference{{Name: "commonPipelineEnvironment", Param: "SWC"}},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "SWCRelease",
						ResourceRef: []config.ResourceReference{{Name: "commonPipelineEnvironment", Param: "SWCRelease"}},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "SpsLevel",
						ResourceRef: []config.ResourceReference{{Name: "commonPipelineEnvironment", Param: "SpsLevel"}},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "Namespace",
						ResourceRef: []config.ResourceReference{{Name: "commonPipelineEnvironment", Param: "Namespace"}},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "PreviousDeliveryCommit",
						ResourceRef: []config.ResourceReference{{Name: "commonPipelineEnvironment", Param: "PreviousDeliveryCommit"}},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "MaxRuntimeInMinutes",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "int",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
				},
			},
		},
	}
	return theMetaData
}