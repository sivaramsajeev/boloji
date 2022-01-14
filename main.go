package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"

	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var pollySession *polly.Polly

const (
  fileName = "sample.txt"
)

func init() {
	pollySession = polly.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})))
}

func RunPolly() {
	contents, err := ioutil.ReadFile(fileName)
	Must(err)
	s := string(contents[:])
	input := &polly.SynthesizeSpeechInput{
		OutputFormat: aws.String("mp3"),
		Text:         aws.String(s),
		VoiceId:      aws.String("Aditi"),
	}
	output, err := pollySession.SynthesizeSpeech(input)
	Must(err)
	outFile, err := os.Create(fmt.Sprintf("%s.mp3", strings.Split(fileName, ".")[0]))
	Must(err)
	defer outFile.Close()
	_, err = io.Copy(outFile, output.AudioStream)
	Must(err)
}

func main() {
	RunPolly()
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
