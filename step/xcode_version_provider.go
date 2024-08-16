package step

import (
	"github.com/teamlapse/go-xcode/models"
	"github.com/teamlapse/go-xcode/utility"
)

type XcodeVersionProvider interface {
	GetXcodeVersion() (models.XcodebuildVersionModel, error)
}

type xcodebuildXcodeVersionProvider struct {
}

func NewXcodebuildXcodeVersionProvider() XcodeVersionProvider {
	return xcodebuildXcodeVersionProvider{}
}

// GetXcodeVersion ...
func (p xcodebuildXcodeVersionProvider) GetXcodeVersion() (models.XcodebuildVersionModel, error) {
	return utility.GetXcodeVersion()
}
