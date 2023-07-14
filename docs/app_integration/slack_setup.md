# Slack Setup and Authentication

### Create a Slack workspace
1. Go to https://slack.com/get-started#/createnew
2. Enter your email address and click **Continue**, or continue with **Apple** or **Google**
3. Check your email for a confirmation code
4. Enter your code, then click **Create a Workspace** and follow the prompts

### Installation and Permissions
1. Create a new [Slack App](https://api.slack.com/apps)

    a. Fill out the **App Name** and select a workspace where you'll develop your app

2. Select the scopes to add to your app by heading to the **OAuth & Permissions** in the sidebar

    a. Scroll down to the **Scopes** section and click to **Add an OAuth Scope** to your **Bot Token**
      * **`chat:write`**: Send messages
      * **`chat:write.public`**: Send messages to channels the bot isn't a member of
      * **`channels:read`**: View basic information about public channels in a workspace

3. Install the app to the workspace by selecting the **Install App** button on the sidebar

    a. After installation, you'll find an **access token** inside your app management page
      * Head to the **OAuth & Permissions** in the sidebar and see it under **OAuth Tokens for Your Workspace**

## Reference
* [Getting started](https://slack.com/help/articles/206845317-Create-a-Slack-workspace)
* [Authentication Basic app setup](https://api.slack.com/authentication/basics)