package targets

// createReadme creates and initial README file with instructions on how to use Mage
func setupReadme() error {
	return createReadme()
}

// setupInfra installs the infrastructure dependencies
func setupInfra() error {
	if err := writeInfraDir(); err != nil {
		return err
	}
	if err := writeStageChart(); err != nil {
		return err
	}
	if err := writeProdChart(); err != nil {
		return err
	}
	if err := writeDotEnvFiles(); err != nil {
		return err
	}
	return nil
}

// setupDocker installs the docker dependencies
func setupDocker() error {
	if err := writeDockerfile(); err != nil {
		return err
	}
	if err := writeDockerIgnorefile(); err != nil {
		return err
	}
	if err := writeDockerDir(); err != nil {
		return err
	}
	if err := writeDockerComposeDev(); err != nil {
		return err
	}
	if err := writeDockerComposeTest(); err != nil {
		return err
	}
	return nil
}

// setupGit sets up git inside the project
func setupGit() error {
	if err := initGit(); err != nil {
		return err
	}
	return writeGitIgnoreFile()
}

// setupGitlab sets up the gitlab pipeline
func setupGitlab() error {
	return writeGitlabCIFile()
}

// setupService sets up the service folder
func setupService() error {
	return writeServiceDir()
}
