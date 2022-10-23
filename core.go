package GoMWF

import (
	"log"
	"os"
	"strconv"
)

const version = "1.0.0"

type GoMWF struct {
	AppName  string
	DEBUG    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	config   config
}

type config struct {
	port     string
	renderer string
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

	//create loggers
	infoLog, errorLog := c.startLoggers()
	c.InfoLog = infoLog
	c.ErrorLog = errorLog
	c.DEBUG, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	c.Version = version
	c.RootPath = rootPath
	c.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}

	return nil
}

func (c *GoMWF) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
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
