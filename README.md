<div align="center">

  <p>
   <a href="https://travis-ci.org/yamamushi/du-authbot"><img src="https://img.shields.io/travis/yamamushi/du-authbot.svg?style=for-the-badge" alt="TravisCI"></a><!--
   --><a href="https://GitHub.com/yamamushi/du-authbot/issues/"><img src="https://img.shields.io/github/issues/yamamushi/du-authbot.svg?style=for-the-badge" alt="TravisCI"></a><!--
--><a href="https://github.com/Favna/Ribbon/blob/master/LICENSE.md"><img src="https://img.shields.io/github/license/favna/ribbon.svg?style=for-the-badge" alt="License"></a><!--
-->
  </p>
  
  <p>
  <img src="https://forthebadge.com/images/badges/uses-badges.svg" alt="Uses Badges"><!--
  --><img src="https://forthebadge.com/images/badges/built-with-love.svg" alt="Built with Love"><!--
  --><img src="https://forthebadge.com/images/badges/made-with-go.svg" alt="Made with Go"><!--
  --><img src="https://forthebadge.com/images/badges/powered-by-electricity.svg" alt="Powered by Electricity">
  </p>
  
  <p>
  <!--
  --><a href="https://dualuniverse.chat"><img src="https://img.shields.io/discord/184691218184273920.svg?style=for-the-badge" alt="TravisCI"></a>
  </p>
 
</div>

 [![Go Report Card](https://goreportcard.com/badge/github.com/yamamushi/du-authbot)](https://goreportcard.com/report/github.com/yamamushi/du-authbot?style=plastic) 


# du-authbot

A Discord bot to assist with the verification of Dual Universe backers in private discords.

Table of Contents
=================

   * [du-authbot](#du-authbot)
      * [Outline](#outline)
      * [Usage](#usage)
         * [Setup](#setup)
         * [Configuration](#configuration)
         * [Commands](#commands) 
      * [FAQ](#faq)
      * [Discord](#discord)

## Outline

- Unless otherwise stated, this bot is covered under the [GPL V3 License](https://www.gnu.org/licenses/gpl.txt).
- It is up to the individual Discord administrators to create channels as they see fit for use by various backer levels (ie NDA channels).
- Roles are not auto-assigned, as that may override permissions in place on individual discords. Instead permissions will need to be assigned (as described below) to allow for users to either self-auth, or for moderators to auth them directly.
- This bot works in conjunction with the [Dual Universe Discord Bot](https://github.com/yamamushi/du-authbot) for the retrieval of user data.
- This bot does NOT store user data.
- This bot does NOT log chats (as that would be a violation of the Discord TOS)
- This bot does NOT have permission to view all channels by default, you must specify which channels it can read.
- Upon joining a Discord, this bot will auto-create backer roles, but will not assign permissions to the roles it creates (see below for more information).
- The bot will store a list of which discords it belongs to, in order to store individual server configurations. 
- To prevent accidental deletion of roles, this bot will not delete any roles, including those it creates. However you can re-run the role creation process to replace backer roles that may have been removed by your Discord staff accidentally.
- To clarify the above, this bot can remove backer roles from users.
- Statistics may be collected periodically by the bot, such as the number of backers in a discord server, but it will not collect who those backers are. 
- This bot is provided without warranty, and by using it you agree to not hold the developers/contributors/or Novaquark liable for issues that may occur. 
- In keeping with that spirit, contributions are welcome and much appreciated. 
- If you discover an issue, you are encouraged to create an issue here on github so that it can be addressed in a timely manner.  


## Usage

### Setup

**Note**
Users will first need to authenticate through the [Official Dual Universe Discord](https://dualuniverse.chat) before their backer status can be verified.

To install the bot into your Discord you can [Click Here](https://discordapp.com/oauth2/authorize?client_id=487875933014065163&scope=bot&permissions=335858752). An outline of the permissions granted is provided below:
   
- Manage Roles (Necessary for creating and granting roles)
- Change Nickname (Only allows for the bot to change its own nickname)
- Read Messages (Necessary for reading command input, messages are NOT logged)
- Embed Links 
- Use External Emojis (Custom Emoji's may be used in the future)
- Send Messages 
- Attach Files (Necessary for generating backer user graphs)
- Add Reactions 

Upon joining a Discord, the following user roles will be automatically created (but not assigned):

- Alpha Team Vanguard
- Kyrium Founder
- Diamond Founder
- Emerald Founder
- Ruby Founder
- Sapphire Founder
- Gold Founder
- Silver Founder
- Bronze Founder
- Iron Founder
- Contributor Supporter
- Sponsor Supporter
- Patron Supporter
- ATV Authorized
- Pre Alpha Authorized
- Forum Authorized

Once the bot has joined a Discord, it will need to be assigned the correct permissions for your Discord to be available in the channels you would like it to view. It does not have permission to view all channels by default. 

This is intentional, as it's understood that there may be paranoia around letting the bot read private channels that servers may have.


### Configuration

A full list of commands is provided in the next section, however this can be used as a first-time configuration guide for a quick setup.

**After you have invited the bot to your server, you will need to give it the necessary access to read from the channels you would like it to read from before following this guide**

**Note** The roles management will **NOT** work if the bot role is not placed above the backer roles in your Discord. 


Add a moderator role to the bot (See [FAQ](#faq) for further information):

`$addmodrole Moderator`

**(Optional)** Set the command prefix for the bot:

`$setcp !` 

**(Optional)** Automatically assign backer roles to all members in your Discord

`$authall` 

**(Optional)** Assign backer roles to a specified user

`$auth @username` 


### Commands 

| Command       | Description   | Example Usage  |
| ------------- | ------------- | ------------- |
| auth | Verify the backer status of a mentioned user | auth @user |
| authall | Verifies the backer status of all users in a server | authall |
| setcp | Sets the command prefix for the bot (default is $) | setcp $ |
| reset | Removes backer roles from a mentioned users | reset @user |
| resetall | Removes backer roles from all users | resetall |
| addmodrole | Adds a server role as a moderator for the bot (case sensitive!) | addmodrole <rolename> | 
| removemodrole | Rmoves a server role as a moderator for the bot | removemodrole <rolename | 
| allowusers | Enables or disables the ability for users to self-auth | allowusers true | 
| adduserrole | Adds a server role as a user role for the bot (use * to allow all user roles). Without this correctly configured, user self-auth will not work (case sensitive!) | adduserrole <rolename> |
| removeuserrole | Removes a server role as a user role for the bot | removeuserrole <rolename> |  
| listuserroles | List user roles | listuserroles | 
| listmodroles | List moderator roles | listmodroles | 
| rebuildroles | Recreates backer roles that may have been removed on accident | rebuildroles | 



## FAQ 


#### What do the bot role permissions provide access to?

User roles (when enabled) will grant access to the following commands:

- auth
- reset


Moderator roles under the bot will grant access to the following commands:

- auth
- reset
- allowusers
- adduserrole
- removeuserrole 
- listuserroles
- listmodroles

Only server owners can use the following commands:

- authall
- resetall
- setcp
- addmodrole
- removemodrole


#### What data is being collected by this bot?

The following information is stored per-Discord for configuration and statistics tracking purposes:

- GuildID (Server ID)
- Command Prefix
- Backer Count (only a count, it does not collect user information)
- PreAlpha Count (a count of the number of pre-alpha users on a discord, it does not collect user information)
- Moderator Roles (the names of the configured roles the bot interprets as moderators for a Discord)
- User Access (a boolean that determines whether individual users can self-auth)
- User Roles (the names of the configured roles the bot interprets as valid user roles for self-auth, without this being configured, self-auth will not work)


#### How do I know this bot is secure? 

Authorization is performed by "Dual Universe Bot", and the source code for the authorization process can be found [here](https://github.com/yamamushi/du-discordbot/blob/master/backer.go).

This authorization process is what is currently relied upon for the Dual Universe Community Discord, and has been approved for use by Novaquark as a valid authorization mechanism. 


#### Why should I use this instead of manually verifying users?

This bot was created to make it easier for Discords to provide a backer authorization process to their users. It relies upon a secure validation mechanism (see above) for providing this authorization, and can alleviate issues caused by human error during manual verification. 

While it is not 100% necessary to use this bot, it is highly encouraged to ensure that users are correctly synchronized between Dual Universe-related Discords and to maintain consistency as further Backer verification information becomes available from Novaquark. 


## Discord

Join us on Discord @ [http://dualuniverse.chat](http://dualuniverse.chat)