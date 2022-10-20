package GoMWF

const version = "1.0.0"

type GoMWF struct {
	AppName string
	DEBUG   bool
	Version string
}

func (c *GoMWF) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "middleware", "migrations", "views", "data", "public", "tmp", "logs"},
	}
	err := c.Init(pathConfig)
	if err != nil {
		return err
	}
	return nil
}

func (c *GoMWF) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		// Create Folders if They Don't Exist
		err := c.CreateDir(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}
