package logger

import (
	"fmt"
	"os"
	"path"
	"sync"
	"time"
)

type timeRollingFileWriter struct {
	curDate    int
	filename   string
	outputFile *os.File
	mutex      sync.Mutex
}

func (p *timeRollingFileWriter) Write(data []byte) (int, error) {
	if len(data) == 0 {
		return 0, nil
	}

	_, _ = os.Stdout.Write(data)
	output, err := p.getOutputFile()
	if err != nil {
		return 0, err
	}

	return output.Write(data)
}

func (p *timeRollingFileWriter) getOutputFile() (*os.File, error) {
	yearDay := time.Now().YearDay()
	if yearDay == p.curDate && p.outputFile != nil {
		return p.outputFile, nil
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	if yearDay == p.curDate && p.outputFile != nil {
		return p.outputFile, nil
	}

	if err := os.MkdirAll(path.Dir(p.filename), os.ModePerm); err != nil {
		return nil, err
	}

	newFilename := fmt.Sprintf("%s.log.%s", p.filename, time.Now().Format("20060102"))
	file, err := os.OpenFile(newFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	if p.outputFile != nil {
		p.outputFile.Close()
	}

	p.outputFile = file

	return p.outputFile, nil
}
