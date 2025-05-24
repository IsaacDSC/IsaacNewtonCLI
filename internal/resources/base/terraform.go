package base

import (
	"context"
	"os"
)

type Resource struct{}

const TerraformBasePath = "terraform"

func (r Resource) CreateTerraform(ctx context.Context) error {
	basePath := TerraformBasePath

	// Create the base directory if it doesn't exist

	// Check if the directory already exists
	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func (r Resource) createIngress(ctx context.Context) error {
	// Create the ingress.tf file
	ingressTfPath := TerraformBasePath + "/ingress.tf"
	if _, err := os.Create(ingressTfPath); err != nil {
		return err
	}

	// Create the ingress variables.tf file
	ingressVariablesTfPath := TerraformBasePath + "/ingress_variables.tf"
	if _, err := os.Create(ingressVariablesTfPath); err != nil {
		return err
	}

	return nil
}

func (r Resource) createTerraformEnv(ctx context.Context, environment string) error {
	fnMountPath := func(fileName string) string {
		return TerraformBasePath + "/" + environment + "/" + fileName
	}

	// Create the provider.tf file
	providerTfPath := fnMountPath("provider.tf")
	if _, err := os.Create(providerTfPath); err != nil {
		return err
	}

	// Create the main.tf file
	mainTfPath := fnMountPath("main.tf")
	if _, err := os.Create(mainTfPath); err != nil {
		return err
	}

	// Create the variables.tf file
	variablesTfPath := fnMountPath("variables.tf")
	if _, err := os.Create(variablesTfPath); err != nil {
		return err
	}

	// Create the outputs.tf file
	outputsTfPath := fnMountPath("outputs.tf")
	if _, err := os.Create(outputsTfPath); err != nil {
		return err
	}

	// Create the terraform.tfvars file
	terraformTfvarsPath := fnMountPath("terraform.tfvars")
	if _, err := os.Create(terraformTfvarsPath); err != nil {
		return err
	}

	return nil
}
