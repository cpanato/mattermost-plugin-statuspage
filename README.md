# Mattermost Statuspage Plugin [![CircleCI](https://circleci.com/gh/cpanato/mattermost-plugin-statuspage.svg?style=svg)](https://circleci.com/gh/cpanato/mattermost-plugin-statuspage)

This plugin sends webhook notifications from Statuspage to Mattermost. Use it to get notified of system outages and of partially degraded services.

To see the plugin in action, join the **Statuspage** channel in our communtiy server at https://community.mattermost.com/core/channels/statuspage.

**Supported Mattermost Server Versions: 5.4+**

## Installation

1. Go to the [releases page of this GitHub repository](https://github.com/cpanato/mattermost-plugin-statuspage/releases) and download the latest release for your Mattermost server.
2. Upload this file in the Mattermost **System Console > Plugins > Management** page to install the plugin, and enable it. To learn more about how to upload a plugin, [see the documentation](https://docs.mattermost.com/administration/plugins.html#plugin-uploads).

## Usage

To configure the plugin, follow these steps:

1. After you've uploaded the plugin in **System Console > Plugins > Management**, go to the plugin's settings page at **System Console > Plugins > Statuspage**.
2. Specify the team and channel to send messages to. For each, use the URL of the team or channel instead of their respective display names.
3. Select the username that this plugin is attached to. You may optionally create a new user account for your Statuspage plugin, which can act as a bot account posting Statuspage updates to a Mattermost channel.
4. Hit **Save**.
5. Next, copy the **Token** above the **Save** button, which is used to configure the plugin for your Statuspage account.
6. Go to your Statuspage account, paste the following webhook URL and specfiy the name of the service and the token you copied in step 5.

```
https://SITEURL/plugins/statuspage/webhook?service=SERVICENAME&token=TOKEN
```
