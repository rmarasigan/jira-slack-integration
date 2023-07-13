# JIRA + Slack Integration

A mini-project that uses AWS services to automate the creation of JIRA tickets/issues and send notifications to a designated Slack channel. The **API Gateway** acts as the entry point, receiving HTTP requests that trigger Lambda Functions. The **Lambda Functions** process the requests, creating JIRA tickets/issues based on the provided information. To ensure reliable and scalable processing, **Simple Queue Service(SQS)** is employed as a buffer between the API Gateway and the Lambda Functions. The processed tickets/issues are then published to the designated Slack channel, notifying the team members.

Please note that this mini-project serves as a demonstration of integrating AWS services such as API Gateway, Lambda Functions, and Simple Queue Service (SQS), for creating JIRA tickets/issues and sending notifications to a Slack channel. It is not intended for real-world production use and may lack certain features, optimizations, and security issues required for a production-ready application.

### Create the JIRA API Token
An **API Token** is required to authenticate a script or other process with an Atlassasian Cloud Product.
1. Log in to the Atlassian platform to access Jira Cloud or click [here](https://id.atlassian.com/manage/api-tokens) to directly open the page to create the API token.
2. Go to **Settings** → **Atlassian account settings** → **Security** → **API token**.
3. Click on "**Create API token**".
  
    a. Enter a distinctive and concise **Label** for your token in the window that display, then click **Create**.

4. Copy the token to your clipboard.

    > **NOTE**:
    >
    >  * For security reasons it isn't possible to view token after closing the creation dialog; if necessary, create a new token.
    >  * You should store the token securely, just as for any password.

### Create a JIRA Basic Auth Header
1. Create a string that has a format of *user_email:api_token*.
2. Encode the string using Base64 encoding.

    **For Linux/MacOS**
    ```bash
    echo -n user@example.com:api_token | base64
    ```

    **For Windows 7 and later, using Microsoft Powershell**
    ```shell
    $Text = ‘user@example.com:api_token_string’
    $Bytes = [System.Text.Encoding]::UTF8.GetBytes($Text)
    $EncodedText = [Convert]::ToBase64String($Bytes)
    $EncodedText
    ```

3. Supply an `Authorization` header with content `Basic` followed by the encoded string.
    ```bash
    curl -D- \
      -X GET \
      -H "Authorization: Basic ZnJlZDpmcmVk" \
      -H "Content-Type: application/json" \
      "https://your-domain.atlassian.net/rest/api/2/issue/issue-id"
    ```

## Using `Makefile` to install, bootstrap, and deploy the project

1. Install all the dependencies and bootstrap your project
    ```bash
    dev@dev:~:bus-ticketing$ make init
    ```

    To initialize the project with specific AWS profile, you can pass a parameter called `profile`.
    ```bash
    dev@dev:~:bus-ticketing$ make init profile=profile_name
    ```

2. Deploy the project.
    ```bash
    dev@dev:~:bus-ticketing$ make deploy
    # Deploying with specific AWS profile
    dev@dev:~:bus-ticketing$ make deploy profile=profile_name
    ```

## Useful commands

* `go test`         run unit tests
* `cdk deploy`      deploy this stack to your default AWS account/region
* `cdk diff`        compare deployed stack with current state
* `cdk synth`       emits the synthesized CloudFormation template
* `cdk bootstrap`   deployment of AWS CloudFormation template to a specific AWS environment (account and region)
* `cdk destroy`     destroy this stack from your default AWS account/region
