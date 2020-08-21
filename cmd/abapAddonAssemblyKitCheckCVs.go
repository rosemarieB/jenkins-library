package cmd

import (
	"encoding/json"

	"github.com/SAP/jenkins-library/pkg/abaputils"
	"github.com/SAP/jenkins-library/pkg/command"
	piperhttp "github.com/SAP/jenkins-library/pkg/http"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/telemetry"
)

func abapAddonAssemblyKitCheckCVs(config abapAddonAssemblyKitCheckCVsOptions, telemetryData *telemetry.CustomData, cpe *abapAddonAssemblyKitCheckCVsCommonPipelineEnvironment) {
	// for command execution use Command
	c := command.Command{}
	// reroute command output to logging framework
	c.Stdout(log.Writer())
	c.Stderr(log.Writer())

	var autils = abaputils.AbapUtils{
		Exec: &c,
	}
	client := piperhttp.Client{}

	// error situations should stop execution through log.Entry().Fatal() call which leads to an os.Exit(1) in the end
	err := runAbapAddonAssemblyKitCheckCVs(&config, telemetryData, &autils, &client, cpe)
	if err != nil {
		log.Entry().WithError(err).Fatal("step execution failed")
	}
}

func runAbapAddonAssemblyKitCheckCVs(config *abapAddonAssemblyKitCheckCVsOptions, telemetryData *telemetry.CustomData, com abaputils.Communication, client piperhttp.Sender, cpe *abapAddonAssemblyKitCheckCVsCommonPipelineEnvironment) error {
	var addonDescriptorFromCPE abaputils.AddonDescriptor
	json.Unmarshal([]byte(config.AddonDescriptor), &addonDescriptorFromCPE)

	addonDescriptor, err := abaputils.ReadAddonDescriptor(config.AddonDescriptorFileName)
	addonDescriptor = transferProductFromCPE(addonDescriptor, addonDescriptorFromCPE)

	if err != nil {
		return nil
	}
	conn := new(connector)
	conn.initAAK(config.AbapAddonAssemblyKitEndpoint, config.Username, config.Password, &piperhttp.Client{})

	for i, repo := range addonDescriptor.Repositories {
		var c cv
		c.init(repo, *conn)
		err := c.validate()
		if err != nil {
			return err
		}
		addonDescriptor.Repositories[i] = c.addFields(addonDescriptor.Repositories[i])
	}
	toCPE, _ := json.Marshal(addonDescriptor)
	cpe.abap.addonDescriptor = string(toCPE)
	return nil
}

// TODO sobald abaputils geändert wurde muss das hier auch geändert werden!
func transferProductFromCPE(addonDescriptor abaputils.AddonDescriptor, addonDescriptorFromCPE abaputils.AddonDescriptor) abaputils.AddonDescriptor {
	addonDescriptor.AddonProduct = addonDescriptorFromCPE.AddonProduct
	addonDescriptor.AddonVersion = addonDescriptorFromCPE.AddonVersion
	addonDescriptor.AddonUniqueID = addonDescriptorFromCPE.AddonUniqueID
	addonDescriptor.CustomerID = addonDescriptorFromCPE.CustomerID
	addonDescriptor.AddonSpsLevel = addonDescriptorFromCPE.AddonSpsLevel
	addonDescriptor.AddonPatchLevel = addonDescriptorFromCPE.AddonPatchLevel
	addonDescriptor.TargetVectorID = addonDescriptorFromCPE.TargetVectorID
	return addonDescriptor
}

func (c *cv) init(repo abaputils.Repository, conn connector) {
	c.connector = conn
	c.Name = repo.Name
	c.VersionYAML = repo.VersionYAML
}

func (c *cv) addFields(initialRepo abaputils.Repository) abaputils.Repository {
	var repo abaputils.Repository
	repo = initialRepo
	repo.Version = c.Version
	repo.SpsLevel = c.SpsLevel
	repo.PatchLevel = c.PatchLevel
	return repo
}

func (c *cv) validate() error {
	appendum := "/odata/aas_ocs_package/ValidateComponentVersion?Name='" + c.Name + "'&Version='" + c.VersionYAML + "'"
	body, err := c.connector.get(appendum)
	if err != nil {
		return err
	}
	var jCV jsonCV
	json.Unmarshal(body, &jCV)
	c.Name = jCV.CV.Name
	c.Version = jCV.CV.Version
	c.SpsLevel = jCV.CV.SpsLevel
	c.PatchLevel = jCV.CV.PatchLevel
	return nil
}

type jsonCV struct {
	CV *cv `json:"d"`
}

type cv struct {
	connector
	Name        string `json:"Name"`
	VersionYAML string
	Version     string `json:"Version"`
	SpsLevel    string `json:"SpLevel"`
	PatchLevel  string `json:"PatchLevel"`
}
