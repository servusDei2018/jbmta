# jbmta
A simple, command-line mail transfer tool.

## Features
 - Fast
 - Small (3.5Mb)
 - 100% configurable via command-line flags
 - TLS security support
 - No dependencies
 - No installation, works out-of-the-box

## Usage
Run `jmbta` and supply the necessary information as flags:
```
jbmta.exe -from me@example.com -to someone@this.com,someoneElse@that.com -server smtp.somesite.com -port 23 -pwd mypassword -msg path\to\message.txt
```
The message to send the recipients is contained in a RFC 822 message in a file. The path to that file is specified with the -msg flag.

## Example
For an example of how to use this to send a gmail message through BASIC,
 1. Put jbmta.exe in the same folder as your `.bas` file
 2. Create `my_message.txt` and put in a RFC 822 style message. (See `example_msg.txt` in the `examples` folder for an example)
 3. Code something like this, customizing where needed:
```basic
run "jbmta.exe -from me@gmail.com -to someone@this.com,someoneElse@that.com -server smtp.google.com -port 587 -pwd mypassword -msg my_message.txt"
```
 4. Run the `.bas` file, and it should send the mail. If any errors occur, a file called `mta.log` will be created and an error message put there.

## RFC 822 message format
The message file should contain a RFC 822 style message, like so:
```
To: recipient@example.net
Subject: This is a message!

Hello! This is a message that you are receiving.

Cheers,
ME
```
The top lines are the RFC headers, there is a blank line, and then the rest is the actual message content. The `To:` header tells the people we're mailing who else is being mailed, so if you don't include someone in that header, nobody will know that he was mailed.

For more information about the RFC 822 format, <a href="https://docs.microsoft.com/en-us/previous-versions/office/developer/exchange-server-2010/aa493927(v=exchg.140)">see here</a>.
