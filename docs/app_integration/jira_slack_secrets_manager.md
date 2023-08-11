# JIRA and Slack Secrets Manager Configuration

The JIRA and Slack configuration needs to be manually created in the **Secrets Manager** as an "**Other type of secret**". You have two options for storing the configuration: a key/value pair or plaintext.

When entering the configuration, make sure to follow the JIRA and Slack configuration format. Once the configuration is created, make sure to ***copy the Secret ARN***, which uniquely identifies the secret. You can find the ARN in the Secrets Manager console.

Next, open the `env.go` file located in the `internal/app/config` directory. Look for the placeholder value '`JIRA-SECRET-ARN`' and '`SLACK-SECRET-ARN`' and replace it with the copied Secret ARN. This step links the application to the correct JIRA and Slack configuration stored in the Secrets Manager.

### JIRA Configuration
<table>
  <tr>
    <th>Field</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>
      <code>token</code>
    </td>
    <td>The JIRA API Token for authentication.</td>
  </tr>
  <tr>
    <td>
      <code>username</code>
    </td>
    <td>The username to be used for Basic Authentication.</td>
  </tr>
  <tr>
    <td>
      <code>endpoint</code>
    </td>
    <td>The Atlassian Domain of your JIRA.</td>
  </tr>
  <tr>
    <td>
      <code>issue_path</code>
    </td>
    <td>The REST API resource path for the JIRA issue.</td>
  </tr>
  <tr>
    <td>
      <code>project_path</code>
    </td>
    <td>The REST API resource path for the JIRA project.</td>
  </tr>
  <tr>
    <td>
      <code>priority_path</code>
    </td>
    <td>The REST API resource path for the JIRA priority.</td>
  </tr>
  <tr>
    <td>
      <code>users_search_path</code>
    </td>
    <td>The REST API resource path for the JIRA users.</td>
  </tr>
</table>

#### Sample JIRA Configuration
```json
{
  "token": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "username": "username@email.com",
  "endpoint": "https://domain_name.atlassian.net",
  "issue_path": "/rest/api/2/issue/",
  "project_path": "/rest/api/2/project/",
  "priority_path": "/rest/api/2/priority/",
  "users_search_path":"/rest/api/2/users/search/"
}
```

#### API Reference
* [Rest API V2](https://developer.atlassian.com/cloud/jira/platform/rest/v2/intro/)
* [Rest API V3](https://developer.atlassian.com/cloud/jira/platform/rest/v3/intro/)

### Slack Configuration
<table>
  <tr>
    <th>Field</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>
      <code>enabled</code>
    </td>
    <td>If the Slack integration is enabled or not. Should be "<code>true</code>" or "<code>false</code>" only.</td>
  </tr>
  <tr>
    <td>
      <code>token</code>
    </td>
    <td>The authentication token bearing required scopes.</td>
  </tr>
  <tr>
    <td>
      <code>channel</code>
    </td>
    <td>The public/private channel ID.</td>
  </tr>
  <tr>
    <td>
      <code>chat_endpoint</code>
    </td>
    <td>The Slack API endpoint of chat.postMessage.</td>
  </tr>
</table>

#### Sample Slack Configuration
```json
{
  "token": "xoxb-xxxxxxxxxxxx-xxxxxxxxxxxx-xxxxxxxxxxxxxxx",
  "channel": "xxxxxxxx",
  "enabled": "true",
  "chat_endpoint": "https://slack.com/api/chat.postMessage"
}
```

#### API Reference
* [Web API methods](https://api.slack.com/methods)
* [API object types](https://api.slack.com/types?ref=apis)
* [chat.postMessage](https://api.slack.com/methods/chat.postMessage)
* [Using the Slack Web API](https://api.slack.com/web#ssl)