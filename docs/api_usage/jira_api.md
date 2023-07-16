# JIRA Integration API Documentation
The JIRA Integration API allows you to interact with some of the JIRA REST API programmatically.
* [Get all JIRA Users](#get-all-jira-users)
* [Get JIRA Project Details](#get-jira-project-details)
* [Get all Issue Priorities](#get-all-issue-priorities)
* [Create a JIRA Issue](#create-a-jira-issue)

## API Usage and Specification
#### Headers
<table>
  <tr>
    <th>Key</th>
    <th>Value</th>
    <th>Required</th>
  </tr>
  <tr>
    <td>
      <code>Content-Type</code>
    </td>
    <td>
      <code>application/json</code>
    </td>
    <td>✅</td>
  </tr>
  <tr>
    <td>
      <code>X-Api-Key</code>
    </td>
    <td>
      <code>xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx</code>
    </td>
    <td>✅</td>
  </tr>
</table>

Setting to **`application/json`** is recommended.

#### HTTP Response Status Codes
<table>
  <tr>
    <th>Status Code</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>200</td>
    <td>OK</td>
  </tr>
  <tr>
    <td>400</td>
    <td>Bad Request</td>
  </tr>
  <tr>
    <td>500</td>
    <td>Internal Server Error</td>
  </tr>
</table>

## Get all JIRA Users
Returns a list of all users, including active users, inactive users, and previously deleted users that have an Atlassian account.

**Method**: `GET`

**Endpoint**: `https://{api_id}.execute-api.{region}.amazonaws.com/production/jira/users`

### Responses
<table>
  <tr>
    <th>Field Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>
      <code>accountId</code>
    </td>
    <td>The account ID of the user (e.g. 123456:12345678-5b10-ac8d-2e05-b22cc7d4eef5)</td>
  </tr>
  <tr>
    <td>
      <code>accountType</code>
    </td>
    <td>The user account type (valid values: atlassian, app, customer, unknown)</td>
  </tr>
  <tr>
    <td>
      <code>displayName</code>
    </td>
    <td>The display name of the user</td>
  </tr>
  <tr>
    <td>
      <code>active</code>
    </td>
    <td>Whether the user is active</td>
  </tr>
</table>

#### Account Type
* `atlassian`: regular Atlassian user account
* `customer`: Jira Service Desk account representing an external service desk
* `app`: system account used for Connect applications and OAuth to represent external systems

#### Sample Response
```json
[
  {
    "accountId": "xxxxxx:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
    "accountType": "atlassian",
    "displayName": "Emily Davis",
    "active": true
  },
  {
    "accountId": "xxxxxx:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
    "accountType": "app",
    "displayName": "Slack",
    "active": true
  }
]
```

## Get JIRA Project Details
Returns the project details for a project.

**Method**: `GET`

**Endpoint**: `https://{api_id}.execute-api.{region}.amazonaws.com/production/jira/project`

### Query Parameters
<table>
  <tr>
    <th>Parameter</th>
    <th>Type</th>
    <th>Description</th>
    <th>Required</th>
  </tr>
  <tr>
    <td>
      <code>key</code>
    </td>
    <td>string</td>
    <td>The project ID or project key (case sensitive).</td>
    <td>✅</td>
  </tr>
</table>

### Responses
<table>
  <tr>
    <th>Field Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>
      <code>id</code>
    </td>
    <td>The ID of the project.</td>
  </tr>
  <tr>
    <td>
      <code>key</code>
    </td>
    <td>The key of the project.</td>
  </tr>
  <tr>
    <td>
      <code>name</code>
    </td>
    <td>The name of the project.</td>
  </tr>
  <tr>
    <td>
      <code>issueTypes</code>
    </td>
    <td>List of the issue types available in the project.</td>
  </tr>
  <tr>
    <td>
      <code>projectTypeKey</code>
    </td>
    <td>The project type of the project (valid values: software, service_desk, business).</td>
  </tr>
</table>

#### Issue Types
<table>
  <tr>
    <th>Field Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>
      <code>id</code>
    </td>
    <td>The ID of the issue type.</td>
  </tr>
  <tr>
    <td>
      <code>name</code>
    </td>
    <td>The name of the issue type.</td>
  </tr>
  <tr>
    <td>
      <code>subtask</code>
    </td>
    <td>Whether the issue type is used to create subtasks.</td>
  </tr>
  <tr>
    <td>
      <code>description</code>
    </td>
    <td>The description of the issue type.</td>
  </tr>
</table>

#### Sample Response
```json
{
  "id": "10000",
  "key": "IN",
  "name": "Integrations",
  "issueTypes": [
    {
      "id": "10001",
      "name": "Task",
      "subtask": false,
      "description": "Tasks track small, distinct pieces of work."
    },
    {
      "id": "10002",
      "name": "Epic",
      "subtask": false,
      "description": "Epics track collections of related bugs, stories, and tasks."
    },
    {
      "id": "10003",
      "name": "Subtask",
      "subtask": true,
      "description": "Subtasks track small pieces of work that are part of a larger task."
    }
  ],
  "projectTypeKey": "software"
}
```

## Get all Issue Priorities
Returns the list of all issue priorities.

**Method**: `GET`

**Endpoint**: `https://{api_id}.execute-api.{region}.amazonaws.com/production/jira/priorities`

### Responses
<table>
  <tr>
    <th>Field Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>
      <code>id</code>
    </td>
    <td>The ID of the issue priority.</td>
  </tr>
  <tr>
    <td>
      <code>name</code>
    </td>
    <td>The name of the issue priority.</td>
  </tr>
  <tr>
    <td>
      <code>description</code>
    </td>
    <td>The description of the issue priority.</td>
  </tr>
</table>

#### Sample Response
```json
[
  {
    "id": "1",
    "name": "Highest",
    "description": "This problem will block progress."
  },
  {
    "id": "2",
    "name": "High",
    "description": "Serious problem that could block progress."
  },
  {
    "id": "3",
    "name": "Medium",
    "description": "Has the potential to affect progress."
  },
  {
    "id": "4",
    "name": "Low",
    "description": "Minor problem or easily worked around."
  },
  {
    "id": "5",
    "name": "Lowest",
    "description": "Trivial problem with little or no impact on progress."
  }
]
```

## Create a JIRA Issue
Creates an issue or, where the option to create subtask is enabled in Jira, a subtask.

**Method**: `POST`

**Endpoint**: `https://{api_id}.execute-api.{region}.amazonaws.com/production/jira/ticket`

### Payload
<table>
  <tr>
    <th>Parameter</th>
    <th>Type</th>
    <th>Description</th>
    <th>Required</th>
  </tr>
  <tr>
    <td>
      <code>title</code>
    </td>
    <td>string</td>
    <td>The issue title/summary.</td>
    <td>✅</td>
  </tr>
  <tr>
    <td>
      <code>description</code>
    </td>
    <td>string</td>
    <td>The issue description.</td>
    <td>✅</td>
  </tr>
  <tr>
    <td>
      <code>project_key</code>
    </td>
    <td>string</td>
    <td>The issue Project Key.</td>
    <td>✅</td>
  </tr>
  <tr>
    <td>
      <code>priority_id</code>
    </td>
    <td>string</td>
    <td>The issue priority ID.</td>
    <td>✅</td>
  </tr>
  <tr>
    <td>
      <code>reporter_id</code>
    </td>
    <td>string</td>
    <td>The issue reporter (user account ID).</td>
    <td>✅</td>
  </tr>
  <tr>
    <td>
      <code>issue_type_id</code>
    </td>
    <td>string</td>
    <td>The issue type ID. If set as Subtask, <code>parent_key</code> is required.</td>
    <td>❌</td>
  </tr>
  <tr>
    <td>
      <code>parent_key</code>
    </td>
    <td>string</td>
    <td>Contain the ID or key of the parent issue.</td>
    <td>❌</td>
  </tr>
  <tr>
    <td>
      <code>assignee_id</code>
    </td>
    <td>string</td>
    <td>The issue assignee (user account ID).</td>
    <td>❌</td>
  </tr>
  <tr>
    <td>
      <code>labels</code>
    </td>
    <td>[]string</td>
    <td>The issue label(s).</td>
    <td>❌</td>
  </tr>
</table>

### Sample Request
#### Payload

* ***Issue as Task***
    ```json
    { 
      "priority_id": "1",
      "project_key": "IN",
      "issue_type_id": "10001",
      "title": "[BOOKING - Fetch] DynamoDBError",
      "reporter_id": "123456:12345678-5b10-ac8d-2e05-b22cc7d4eef5",
      "description": "source: get_bookings\nerror_message: failed to create a filter query to fetch all bookings\ntable: bookings_table"
    }
    ```

* ***Issue as Subtask***
    ```json
    { 
      "priority_id": "1",
      "project_key": "IN",
      "issue_type_id": "10003",
      "parent_key": "IN-1",
      "title": "[BOOKING - Fetch] DynamoDBError",
      "reporter_id": "123456:12345678-5b10-ac8d-2e05-b22cc7d4eef5",
      "description": "source: get_bookings\nerror_message: failed to create a filter query to fetch all bookings\ntable: bookings_table"
    }
    ```