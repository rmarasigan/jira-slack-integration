# CloudWatch Alert Secrets Manager Configuration

The CloudWatch Alert configuration needs to be manually created in the **Secrets Manager** as an "**Other type of secret**". You have two options for storing the configuration: a key/value pair or plaintext.

When entering the configuration, make sure to follow the configuration format. Once the configuration is created, make sure to **copy the Secret ARN**, which uniquely identifies the secret. You can find the ARN in the Secrets Manager console.

Next, open the `env.go` file located in the `internal/app/config` directory. Look for the placeholder value '`CLOUDWATCH-SECRET-ARN`' and replace it with the copied Secret ARN. This step links the application to the correct CloudWatch configuration stored in the Secrets Manager.

### CloudWatch Alert Configuration
<table>
  <tr>
    <th>Field</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>
      <code>priority_id</code>
    </td>
    <td>The JIRA Issue Priority ID.</td>
  </tr>
  <tr>
    <td>
      <code>project_key</code>
    </td>
    <td>The JIRA Issue Project Key.</td>
  </tr>
  <tr>
    <td>
      <code>issue_type_id</code>
    </td>
    <td>The JIRA Issue Type ID.</td>
  </tr>
  <tr>
    <td>
      <code>reporter_id</code>
    </td>
    <td>The JIRA Issue Reporter ID.</td>
  </tr>
  <tr>
    <td>
      <code>parent_key</code>
    </td>
    <td>The JIRA Issue Parent Key. This is optional if the created issue will be a 'Task' issue type.</td>
  </tr>
  <tr>
    <td>
      <code>api_key</code>
    </td>
    <td>The REST API Gateway key.</td>
  </tr>
</table>

#### Sample CloudWatch Alert Configuration
```json
{
  "priority_id": "3",
  "project_key": "IN",
  "parent_key":"IN-1",
  "issue_type_id": "10017",
  "api_key":"xxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "reporter_id": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
```