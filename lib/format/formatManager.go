package format

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type FormatManager struct {
	Formaters map[string]Formater
	fileName  string
}

func NewFormatManager(fileName string) *FormatManager {
	fm := FormatManager{
		Formaters: make(map[string]Formater),
		fileName:  fileName,
	}

	// Add formaters
	fm.addFormat(NewJpeg())
	fm.addFormat(NewPng())

	return &fm
}

func (fm *FormatManager) Build() (Formater, error) {
	bytearr, err := ioutil.ReadFile(fm.fileName)
	if err != nil {
		log.Println("Failed to read file: ", fm.fileName)
		os.Exit(1)
	}

	for magicNumber := range fm.Formaters {
		if strings.HasPrefix(string(bytearr), magicNumber) {
			return fm.Formaters[magicNumber], nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Unknown format for: %v", fm.fileName))
}

func (fm *FormatManager) addFormat(formater Formater) {
	fm.Formaters[formater.MagicNumber()] = formater
}
