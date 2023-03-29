### Deploying to Localstack

#### gateway.yaml
* aws cloudformation create-stack --stack-name my-stack --capabilities CAPABILITY_NAMED_IAM --template-body file://gateway.yaml --endpoint-url=http://localhost:4566
* aws cloudformation describe-stacks   --stack-name my-stack     --output text --endpoint-url=http://localhost:4566
* aws apigateway get-rest-apis   --endpoint-url=http://localhost:4566
  ```
  {
      "items": [
      {
          "id": "etf7xsf1l3",
          "name": "HelloWorldApi",
          "createdDate": 1679996507.0,
          "apiKeySource": "HEADER",
          "endpointConfiguration": {
              "types": [
                  "EDGE"
              ]
          },
          "tags": {
              "aws:cloudformation:logical-id": "ApiGatewayRestApi",
              "aws:cloudformation:stack-name": "my-stack",
              "aws:cloudformation:stack-id": "arn:aws:cloudformation:us-east-1:000000000000:stack/my-stack/53de0f30"
          },
          "disableExecuteApiEndpoint": false
      }
      ]
  }
  ```
    grab the items.id value and that will be the rest-api-id
* aws apigateway get-resources --rest-api-id etf7xsf1l3 --endpoint-url=http://localhost:4566
  ```
  {
      "items": [
      {
          "id": "pn4j492gxu",
          "path": "/"
      },
      {
          "id": "qzwpmgj2vf",
          "parentId": "pn4j492gxu",
          "pathPart": "helloworld",
          "path": "/helloworld",
          "resourceMethods": {
              "GET": {
                  "httpMethod": "GET",
                  "authorizationType": "NONE",
                  "apiKeyRequired": false,
                  "methodResponses": {},
                  "methodIntegration": {
                      "type": "AWS_PROXY",
                      "httpMethod": "POST",
                      "uri": "arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/arn:aws:lambda:us-east-1:000000000000:function:my-stack-HelloWorldFunction-d9ad3232/invocations",
                      "requestParameters": {},
                      "cacheNamespace": "qzwpmgj2vf",
                      "cacheKeyParameters": []
                  }
              }
          }
      }
      ]
  }
  ```
grab the items.id (items[1].id) value and that will be the resource-id
* aws  apigateway test-invoke-method --rest-api-id etf7xsf1l3 --resource-id qzwpmgj2vf --http-method GET --path-with-query-string '/helloworld'  --endpoint-url=http://localhost:4566
  ```
  {
      "status": 200,
      "body": "\"Hello World!\"",
      "headers": {
      "content-type": "application/json",
      "Content-Length": "14"
  }
  ``` 

#### go_gateway.yaml

* Build the .go application and zip it
    ```
    go build .
    zip hellolambda.zip hellolambda
    ```
* Upload the .zip file to S3 bucket
    ```
    aws s3 mb s3://my-lambda-bucket/ --endpoint-url=http://localhost:4566
    aws s3 cp ~/GolandProjects/cf/aws/hellolambda.zip s3://my-lambda-bucket/ --endpoint-url=http://localhost:4566
    ```  
* Create the stack
    ```
    aws cloudformation create-stack --stack-name my-stack --capabilities CAPABILITY_NAMED_IAM --template-body file://go_gateway.yaml --endpoint-url=http://localhost:4566 
    ```
* Get the rest-api using the command
    ```
    aws apigateway get-rest-apis --endpoint-url=http://localhost:4566
    ```
    
    The rest-id is inside items.id field.
    
* Get the resource id using the command

  ```
  aws apigateway get-resources --rest-api-id xnh7ipomrn --endpoint-url=http://localhost:4566
  ```
    
  Replace the `--rest-api-id` with the items.id from the previous step. 
  The items[1].id will be used in the next command

* Use the following command to invoke
    ```
    aws apigateway test-invoke-method --rest-api-id 304ascefor  --resource-id glvc6z2rkr  --http-method GET --path-with-query-string '/helloworld' --endpoint-url=http://localhost:4566 
    ```
    Replace `--rest-api-id` and `--resource-id` with the id from the previous 2 steps.

* On successful run will see output:
    ```
    {
        "status": 200,
        "body": "XXXX",
        "headers": {
            "content-type": "application/json",
            "Content-Length": "4"
        }
    }
    ```
