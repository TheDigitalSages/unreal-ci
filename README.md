# unreal-ci
[![License](https://img.shields.io/github/license/eddiebeazer/ue4-ci?color=blue)](https://opensource.org/licenses/Apache-2.0) 

[![Go Report Card](https://goreportcard.com/badge/github.com/eddiebeazer/ue4-ci)](https://goreportcard.com/report/github.com/eddiebeazer/ue4-ci)

## What Unreal CI does

This is a commandline tool that's acts as a companion to the [ue4cli](https://docs.adamrehn.com/ue4cli/overview/introduction-to-ue4cli)
tool.  I mainly created a separate tool because this tool was designed to help out with more niche workflows such as
updating your projects version, parsing lint reports from [LinterV2](https://www.unrealengine.com/marketplace/en-US/product/linter-v2?sessionInvalidated=true),
cleaning your archive directory when building and parsing JSON produced from the built-in uat test tool into JUnit format
for test parsing.  I use this in my pipelines for all of my games and tools.  As my workload increases, I'll add in any
functionally that feels useful for other people to use

## Installation

This CLI can be used with Go installed on your system or through a docker container. The tool will assume that you are using
it in your projects root directory(directory with your *.uproject file in it), if you aren't you need to specify paths
which will be listed below.

### Go install

```
go install https://github.com/TheDigitalSages/unreal-ci
unreal-ci help
```

### Docker

This image was mainly meant to be used as a CI image for GitHub actions, TeamCity, Jenkins, etc.