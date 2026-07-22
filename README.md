Creating telegram bot on golang.  
Right now there is some version of it on t.me/*****_bot  
To use it first build binary with  
`go build -ldflags "-X"=github.com/ptakha/go_bot/cmd.appVersion=YOUR_VERSION`  
then create your bot in telegram using botfather and save token  
then using terminal run next commands:  
`read -s token`  
After this command write or paste your token in shell  
`export token`  
`./go_bot start`  