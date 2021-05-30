/*
import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"

	"encoding/json"
	"flag"
	"fmt"
	"os"
)
*/
package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"

	"fmt"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	myFunction := `function handler(event) {
		var request = event.request;
		var uri = request.uri;
	
		// Check whether the URI is missing a file name.
		if (uri.endsWith('/')) {
			request.uri += 'index.html';
		}
		// Check whether the URI is missing a file extension.
		else if (!uri.includes('.')) {
			request.uri += '/index.html';
		}
		return request;
	}`

	myCloudfront := cloudfront.New(sess)
	myFunctionName := "whyfunc"
	myRuntime := "cloudfront-js-1.0"
	myComment := "this is a function"

	myFunctionConfig := &cloudfront.FunctionConfig{
		Runtime: &myRuntime,
		Comment: &myComment,
	}
	myFunctionInput := &cloudfront.CreateFunctionInput{
		FunctionCode:   []byte(myFunction),
		FunctionConfig: myFunctionConfig,
		Name:           &myFunctionName,
	}
	result, err := myCloudfront.CreateFunction(myFunctionInput)
	if err != nil {
		fmt.Println("Cannot create function: " + err.Error())
	} else {
		fmt.Println(result)
	}
}
