package targets

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"text/template"
)

const chartTemplate = `serviceName: "{{ .Name }}"
serviceType: "{{ .Type }}"
serviceVersion: "0.0.1"

localRedis: 1

replicas: 3

pullPolicy: "Always"
podAutoscaling: 1
developmentVolumeMapping: 0

serviceIngress: 1
serviceGrpcIngress: 1

chartVersion: "v9.0.0"
`

var (
	chartTmpl = template.Must(template.New("chart").Parse(chartTemplate))

	// InfraDir represents the directory used for infrastructure files.
	InfraDir = "infra"

	// GCPStageChartFile configures the file path to the gcp chart for stage.
	GCPStageChartFile = path.Join(InfraDir, "stage.gcp.yaml")

	// GCPProdChartFile configures the file path to the gcp chart for prod.
	GCPProdChartFile = path.Join(InfraDir, "prod.gcp.yaml")
)

type chartData struct {
	Name string
	Type string
}

func writeInfraDir() error {
	return os.MkdirAll(InfraDir, os.ModePerm)
}

func writeStageChart() error {
	buf := bytes.NewBuffer(nil)
	if err := chartTmpl.Execute(buf, chartData{
		Name: ProjectName,
		Type: ProjectType,
	}); err != nil {
		return err
	}
	raw, err := ioutil.ReadAll(buf)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(GCPStageChartFile, raw, 0664)
}

func writeProdChart() error {
	buf := bytes.NewBuffer(nil)
	if err := chartTmpl.Execute(buf, chartData{
		Name: ProjectName,
		Type: ProjectType,
	}); err != nil {
		return err
	}
	raw, err := ioutil.ReadAll(buf)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(GCPProdChartFile, raw, 0664)
}

func writeDotEnv() error {
	raw := []byte{}
	return ioutil.WriteFile(path.Join(InfraDir, ".env"), raw, 0664)
}

func writeDotEnvDev() error {
	raw := []byte{}
	return ioutil.WriteFile(path.Join(InfraDir, "dev.env"), raw, 0664)
}

func writeDotEnvTest() error {
	raw := []byte{}
	return ioutil.WriteFile(path.Join(InfraDir, "test.env"), raw, 0664)
}

func writeDotEnvLocal() error {
	raw := []byte{}
	return ioutil.WriteFile(path.Join(InfraDir, "local.env"), raw, 0664)
}
