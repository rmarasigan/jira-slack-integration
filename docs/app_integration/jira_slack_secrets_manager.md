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
    <td>
      The REST API resource path for the JIRA project.
    </td>
  </tr>
  <tr>
    <td>
      <code>priority_path</code>
    </td>
    <td>The REST API resource path for the JIRA priority.</td>
  </tr>
</table>

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